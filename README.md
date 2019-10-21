# Metal Cloud Go SDK

This SDK allows control of the [Bigstep Metal Cloud](https://bigstep.com) from Go.

GoDoc documentation available [here](https://godoc.org/github.com/bigstepinc/metal-cloud-sdk-go)

## Getting started

```go
import "github.com/bigstepinc/metal-cloud-sdk"

client, err := metalcloud.GetMetalcloudClient(
    d.Get("user").(string),
    d.Get("api_key").(string),
    d.Get("endpoint").(string),
  )
  if err != nil {
    return nil, err
  }


```
