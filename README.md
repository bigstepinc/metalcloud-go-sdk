# Metal Cloud Go SDK

[![GoDoc](https://godoc.org/github.com/bigstepinc/metal-cloud-sdk-go?status.svg)](https://godoc.org/github.com/bigstepinc/metal-cloud-sdk-go) 
[![Build Status](https://travis-ci.com/metalsoft-io/metal-cloud-sdk-go.svg?branch=master)](https://travis-ci.com/metalsoft-io/metal-cloud-sdk-go)

This SDK allows control of the [Bigstep Metal Cloud](https://bigstep.com) from Go.

GoDoc documentation available [here](https://godoc.org/github.com/bigstepinc/metal-cloud-sdk-go)

## Getting started

```go
package main
import "github.com/bigstepinc/metal-cloud-sdk-go"
import "os"
import "log"

func main(){
  user := os.Getenv("METALCLOUD_USER")
  apiKey := os.Getenv("METALCLOUD_API_KEY")
  endpoint := os.Getenv("METALCLOUD_ENDPOINT")

  if(user=="" || apiKey=="" || endpoint==""){
    log.Fatal("METALCLOUD_USER, METALCLOUD_API_KEY, METALCLOUD_ENDPOINT environment variables must be set")
  }

  client, err := metalcloud.GetMetalcloudClient(user, apiKey, endpoint)
  if err != nil {
    log.Fatal("Error initiating client: %s", err)
  }

  infras,err :=client.Infrastructures()
  if err != nil {
    log.Fatal("Error retrieving a list of infrastructures: %s", err)
  }

  for _,infra := range *infras{
    log.Printf("%s(%d)",infra.InfrastructureLabel, infra.InfrastructureID)
  }
}
```

### Configuring a proxy:

ProxyFromEnvironment returns the URL of the proxy to use for a given request, as indicated by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions thereof). Requests use the proxy from the environment variable matching their scheme, unless excluded by NO_PROXY.


## Building the SDK

You should always rebuild the automatically generated components of some of the objects as well as documentation by running `go generate`. Before doing so install the following:
```
go install github.com/vburenin/ifacemaker@latest
```

Then run go generate
```
go generate
```

Note that you may need to run the generation process two times for the documentation generation to update as the first pass generates the doc functions for the objects that need to be documented and the other generates the actual markdown documentation that can be found here:

## A note on apply-able objects documentation

Certain objects implement the Applier interface:
1. datacenter
2. infrastructure
3. asset
4. template
5. secret
6. shared_drive
7. subnet_oob
8. subnet_pool
9. switch_device
10. variable
11. workflow

The documentation for the object's fields is available here: [Applier types YAML documentation](docs/types.md)