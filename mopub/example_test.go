package mopub_test

import (
	"fmt"
	"net/url"

	"github.com/hiyali/go-lib-ssv/mopub"
)

func ExampleVerify() {
	testSecretKey := "7dbcfd2a42134f47bfb72daa02f85ec9"
	testCallbackWithHashData := "https://api.example.com/callback?customer_id=3453523454&id=70bae1905f7844a3a012a5f4173021db&hash=28f3b28b09b2578db06ee371990b5a02882523eba954d5a1b57afe2c7e7d3f10&value=20&type=Coins"

	testUrl, err := url.Parse(testCallbackWithHashData)
	if err != nil {
		fmt.Errorf("url.Parse(%s) - err: %v", testCallbackWithHashData, err)
	}

	mopub.LogEnabled = true
	if err = mopub.Verify(testUrl, testSecretKey, "hash"); err != nil {
		fmt.Errorf("mopub.Verify(%v, %s, %s) - url.RawQuery: %s, err: %v", testUrl, testSecretKey, "hash", testUrl.RawQuery, err)
	}

	fmt.Println("PASSED")
	// Output: PASSED
}
