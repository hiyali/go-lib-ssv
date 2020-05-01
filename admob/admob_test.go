package admob_test

import (
	"net/url"
	"testing"

	"github.com/hiyali/go-lib-ssv/admob"
)

const (
	testCBWithHashDataAndEncodedSpace = "https://api.example.com/callback?ad_network=4970775877303683148&ad_unit=3543424263&reward_amount=1&reward_item=Key%20Doubler&timestamp=1584428655496&transaction_id=0280088a3d615a1a28929ba7c00861d4&user_id=KK1nqvkZ4tQDon92LrStOXPJbx93&signature=MEUCIQChVBUhVph0ymVFqHzdld9PWZitruPy2Q-OELQZH9g35wIgO3cVUChofhVrlSXZodlQFQM5vIvf5dGLqBAGXz6AmaE&key_id=3335741209"
	testCBWithHashData                = "https://api.example.com/callback?ad_network=4970775877303683148&ad_unit=3543424263&reward_amount=1&reward_item=Key Doubler&timestamp=1584428655496&transaction_id=0280088a3d615a1a28929ba7c00861d4&user_id=KK1nqvkZ4tQDon92LrStOXPJbx93&signature=MEUCIQChVBUhVph0ymVFqHzdld9PWZitruPy2Q-OELQZH9g35wIgO3cVUChofhVrlSXZodlQFQM5vIvf5dGLqBAGXz6AmaE&key_id=3335741209"
)

func TestAdmobVerifyWithEscapedSpace(t *testing.T) {
	testUrl, err := url.Parse(testCBWithHashDataAndEncodedSpace)
	if err != nil {
		t.Errorf("url.Parse(%s) - err: %v", testCBWithHashDataAndEncodedSpace, err)
	}

	admob.LogEnabled = true

	if err = admob.Verify(testUrl); err != nil {
		t.Errorf("admob.Verify(%v) - url.RawQuery: %s, err: %v", testUrl, testUrl.RawQuery, err)
	}
}

func TestAdmobVerifyWithoutEscapedSpace(t *testing.T) {
	testUrl, err := url.Parse(testCBWithHashData)
	if err != nil {
		t.Errorf("url.Parse(%s) - err: %v", testCBWithHashData, err)
	}

	admob.LogEnabled = true

	if err = admob.Verify(testUrl); err != nil {
		t.Errorf("admob.Verify(%v) - url.RawQuery: %s, err: %v", testUrl, testUrl.RawQuery, err)
	}
}
