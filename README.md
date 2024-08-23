# Metal Cloud Go SDK

[![Build Status](https://travis-ci.com/metalsoft-io/metal-cloud-sdk-go.svg?branch=master)](https://travis-ci.com/metalsoft-io/metal-cloud-sdk-go)

This SDK allows control of the `Metalsoft Cloud` from Go.

Generated GoDoc documentation available <https://godoc.org/github.com/bigstepinc/metal-cloud-sdk-go>

## Setup

1. Make sure you have Go installed, version 1.20 or higher.

2. Clone the repository:

    ```sh
    git clone git@github.com:metalsoft-io/metal-cloud-sdk-go.git 
    cd metal-cloud-sdk-go
    ```

3. Install dependencies, run the following command to check if there are any missing dependencies and download them:

    ```sh
    go mod tidy
    ```

4. Set environment variables:
   You need to set the following environment variables:

   - `METALCLOUD_USER`: Your Metalsoft Cloud username.
   - `METALCLOUD_API_KEY`: Your Metalsoft Cloud API key.
   - `METALCLOUD_ENDPOINT`: The API endpoint.

    These are available in the MetalCloud UI that you are connecting to and will be used later on in the metalcloud-cli.

   Example:

   ```sh
   export METALCLOUD_USER="your_username"
   export METALCLOUD_API_KEY="your_api_key"
   export METALCLOUD_ENDPOINT="https://api.metalcloud.com"
   ```

## Development

All changes are done through pull requests against the `master` branch.  
Once the PR is merged a new release can be create manually in order to have the changes available in the SDK.

Make sure you also run the following before submitting a PR:

1. Format the code

```sh
go fmt .
```

2. Run the linter

```sh
go vet .
golangci-lint run 
```

3. Run the tests

```sh
go test -v
```

## Testing

To run the tests, run the following command:

```sh
go test -v
```

## Release

Releases are created manually by creating a new tag and release in the GitHub UI.

See the [Releases](https://github.com/metalsoft-io/metal-cloud-sdk-go/releases) page for more information.

## Sample Code

A sample code to get you started using the SDK:

```go
package main
import "github.com/metal-cloud-sdk-go"
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

ProxyFromEnvironment returns the URL of the proxy to use for a given request, as indicated by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions thereof).

Requests use the proxy from the environment variable matching their scheme, unless excluded by NO_PROXY.
