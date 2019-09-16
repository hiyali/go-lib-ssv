# go-lib-ssv
The aim of this repository is simplify the chaos of the SSV

[Admob rewarded-video-ssv docs](https://developers.google.com/admob/android/rewarded-video-ssv)

## Verifiers added for

* Admob ([admob.google.com](https://admob.google.com/home/))

## Usage

```golang
import "github.com/hiyali/go-lib-ssv/admob"

func main() {
  admob.LogRawQuery = true // enable log query raw, default is: false
  if err := admob.Verify(adRewardUrl); err != nil {
    // Verification failed
  }

  // Verified
}
```

adRewardUrl look like:
```
https://www.yourdomain.com/path?ad_network=5450213213286189855&ad_unit=12345678&reward_amount=10&reward_item=coins×tamp=1507770365237823&transaction_id=1234567890ABCDEF1234567890ABCDEF&user_id=1234567&signature=MEUCIQDGx44BZgQU6TU4iYEo1nyzh3NgDEvqNAUXlax-XPBQ5AIgCXSdjgKZvs_6QNYad29NJRqwGIhGb7GfuI914MDDZ1c&key_id=1268222887
```

## Examples

* echo ([echo.labstack.com](https://echo.labstack.com))

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
