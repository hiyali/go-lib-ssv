package mopub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"net/url"
	"sort"
)

/*
* https://developers.mopub.com/publishers/ui/apps/manage-ad-units/#callback-macros
 */

var (
	LogEnabled bool = false
)

// mopub rewarded video ads server-side-verification callback url
func Verify(cbUrl *url.URL, secret, verifierKey string) (err error) {
	// -- 1. Parse the callback URL and retrieve the query parameters
	query := cbUrl.Query()
	if LogEnabled {
		log.Printf("mopub.Verify - url.RawQuery: %s, secret: %s, verifierKey: %s", cbUrl.RawQuery, secret, verifierKey)
	}

	// -- 2. Drop the verifier key and value from the array of parameters.
	verifierVal := query.Get(verifierKey)
	verifier, err := hex.DecodeString(verifierVal)
	if err != nil {
		return errors.New("Can not decode verifier")
	}
	delete(query, verifierKey)

	// -- 3. Sort the query parameters alphabetically
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// -- 4. Concatenate the values in that alphabetical order
	concatStr := ""
	for _, k := range keys {
		concatStr += query[k][0] // just use first one
	}

	// -- 5. Hash the concatenated string using a HMAC SHA256 hash algorithm and the Callback Secret Key.
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(concatStr))
	freshHash := mac.Sum(nil)
	// log.Printf("%x", freshHash)
	// log.Printf("%x", verifier)
	if !hmac.Equal(freshHash, verifier) {
		return errors.New("Not verified")
	}

	return
}
