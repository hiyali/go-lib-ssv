package admob

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	keyServerEndpoint     = "https://gstatic.com/admob/reward/verifier-keys.json"
	keyServerEndpointTest = "https://gstatic.com/admob/reward/verifier-keys-test.json"
)

/*
* https://github.com/google/tink/blob/master/apps/rewardedads/src/main/java/com/google/crypto/tink/apps/rewardedads/RewardedAdsVerifier.java
* https://thanethomson.com/2018/11/30/validating-ecdsa-signatures-golang/
 */

type (
	ecdsaSignature struct {
		R, S *big.Int
	}

	verifierKeyJson struct {
		Keys []verifierKey `json:"keys"`
	}

	verifierKey struct {
		KeyId  int    `json:"keyId"`
		Pem    string `json:"pem"`
		Base64 string `json:"base64"`
	}

	keyMap map[int]string
)

var (
	LogEnabled bool = false
)

func getJson(url string, target interface{}) error {
	var c = &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	// compute the SHA256 hash
	return h.Sum(nil)
}

func parsePublicKey(publicKey string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("Failed to decode PEM public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Failed to parse ECDSA public key")
	}

	switch pub := pub.(type) {
	case *ecdsa.PublicKey:
		return pub, nil
	}

	return nil, errors.New("Unsupported public key type")
}

func keysToMap(keys []verifierKey) (keyMap, error) {
	kMap := keyMap{}
	for _, k := range keys {
		kMap[k.KeyId] = k.Pem
	}
	return kMap, nil
}

// cbUrl is admob rewarded video ads server-side-verification callback url
func Verify(cbUrl *url.URL) (err error) {
	// -- get verifier keys json
	verifierKeyJson := &verifierKeyJson{}
	if err = getJson(keyServerEndpoint, verifierKeyJson); err != nil {
		return
	}

	// -- prepare pre-params
	keyIdStr := cbUrl.Query().Get("key_id")
	keyId, err := strconv.Atoi(keyIdStr)
	if err != nil {
		return err
	}

	signatureDerStr := cbUrl.Query().Get("signature")
	signatureDer, err := base64.RawURLEncoding.DecodeString(signatureDerStr)
	if err != nil {
		return err
	}

	rawQuery, err := url.QueryUnescape(cbUrl.RawQuery)
	if err != nil {
		return err
	}
	if LogEnabled {
		log.Printf("admob.Verify - url.RawQuery: %s", rawQuery)
	}

	sigIdx := strings.Index(rawQuery, "&signature=")
	if sigIdx == -1 {
		return errors.New("Can't find signature")
	}

	// -- prepare specific params
	messageData := rawQuery[:sigIdx]
	msgHash := hash([]byte(messageData))

	signature := &ecdsaSignature{}
	_, err = asn1.Unmarshal(signatureDer, signature)
	if err != nil {
		return
	}

	// public key
	km, err := keysToMap(verifierKeyJson.Keys)
	if err != nil {
		return err
	}
	pem, ok := km[keyId]
	if !ok {
		return errors.New("Can't find key_id from keys")
	}
	publicKey, err := parsePublicKey(pem)
	if err != nil {
		return err
	}

	// -- verify
	verified := ecdsa.Verify(publicKey, msgHash, signature.R, signature.S)
	if !verified {
		return errors.New("Signature not valid")
	}
	return
}
