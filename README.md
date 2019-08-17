# go-lib-ssv
The aim of this repository is simplify the chaos of the SSV

## Verifiers added for

* [Admob](https://admob.google.com/home/)

## Examples

* [echo](https://echo.labstack.com)

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
