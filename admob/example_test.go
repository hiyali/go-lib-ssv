package admob_test

import (
	"fmt"
	"net/url"

	"github.com/hiyali/go-lib-ssv/admob"
)

func ExampleVerify() {
	rawUrl := "https://api.example.com/callback?ad_network=4970775877303683148&ad_unit=3543424263&reward_amount=1&reward_item=Key%20Doubler&timestamp=1584428655496&transaction_id=0280088a3d615a1a28929ba7c00861d4&user_id=KK1nqvkZ4tQDon92LrStOXPJbx93&signature=MEUCIQChVBUhVph0ymVFqHzdld9PWZitruPy2Q-OELQZH9g35wIgO3cVUChofhVrlSXZodlQFQM5vIvf5dGLqBAGXz6AmaE&key_id=3335741209"
	testUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Errorf("url.Parse(%s) - err: %v", rawUrl, err)
	}

	admob.LogEnabled = true

	if err = admob.Verify(testUrl); err != nil {
		fmt.Errorf("admob.Verify(%v) - url.RawQuery: %s, err: %v", testUrl, testUrl.RawQuery, err)
	}
	fmt.Println("PASSED")
	// Output: PASSED
}
