# go-lib-ssv
The aim of this repository is simplify the chaos of the SSV

## Verifiers added for

* [Admob](https://admob.google.com/home/) ([doc](https://developers.google.com/admob/android/rewarded-video-ssv))
* [MoPub](https://app.mopub.com/) ([doc](https://developers.mopub.com/publishers/android/rewarded-video/#4-configure-the-callback-server))

## Quick look

```golang
import "github.com/hiyali/go-lib-ssv/admob"

func main() {
  admob.LogEnabled = true // enable log query raw, default is: false
  if err := admob.Verify(adRewardUrl); err != nil {
    // Verification failed
  }

  // Verified
}
```

admob server callback url look like:
```
https://www.yourdomain.com/path?ad_network=5450213213286189855&ad_unit=12345678&reward_amount=10&reward_item=coins×tamp=1507770365237823&transaction_id=1234567890ABCDEF1234567890ABCDEF&user_id=1234567&signature=MEUCIQDGx44BZgQU6TU4iYEo1nyzh3NgDEvqNAUXlax-XPBQ5AIgCXSdjgKZvs_6QNYad29NJRqwGIhGb7GfuI914MDDZ1c&key_id=1268222887
```

## Lib / Method

> All libs have `LogEnabled` property

| lib.Method | Description |
| --- | --- |
| `admob.Verify(url *url.Url) error` | |
| `mopub.Verify(url *url.Url, secret, verifierKey string) error` | verifierKey usually is `hash`, you'll find secret key in `Rewarded video` tab in `https://app.mopub.com/account` page |

## Test
```
go test ./...
```

## Example with [echo](https://echo.labstack.com) framework

```golang
import "github.com/hiyali/go-lib-ssv/admob"

func ApiGetHandler(c echo.Context) (err error) {
  ...
  if err := admob.Verify(c.Request().URL); err != nil {
    log.Errorf("Verification failed - err: %v", err)
    return err
  }

  // do something after verified
  ...
}
```

## Contribution
> Feel free
