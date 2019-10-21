package metalcloud

import "github.com/ybbus/jsonrpc"
import "net/http"
import "net/url"
import "crypto/hmac"
import "crypto/md5"
import "encoding/hex"
import "io/ioutil"
import "bytes"
import "log"
import "strings"
import "errors"

func DEFAULT_ENDPOINT() string {
	return "https://api.bigstep.com/metal-cloud"
}

type SignatureAdderRoundTripper struct {
	APIKey string
	http.RoundTripper
	LogReply bool
	DryRun bool
}

func (c *SignatureAdderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	components := strings.Split(c.APIKey, ":")

	var strKeyMetaData *string

	strKeyMetaData = nil

	if len(components) > 1 {
		strKeyMetaData = &components[0]
	}

	key := []byte(c.APIKey)

	// Read the content
	var message []byte
	if req.Body != nil {
		message, _ = ioutil.ReadAll(req.Body)
	}

	if c.LogReply {
		log.Println(string(message))
	}

	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(message))

	hmac := hmac.New(md5.New, key)
	hmac.Write(message)

	var signature = hex.EncodeToString(hmac.Sum(nil))

	values, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	if strKeyMetaData != nil {
		signature = *strKeyMetaData + ":" + signature
	}

	values.Add("verify", signature)

	url := req.URL

	url.RawQuery = values.Encode()

	req.URL = url

	var resp *http.Response = nil

	if !c.DryRun {
		resp, err = http.DefaultTransport.RoundTrip(req)
	}

	if c.LogReply {
		//log the reply
		if resp.Body != nil {
			message, _ = ioutil.ReadAll(resp.Body)
		}

		log.Println(string(message))

		// Restore the io.ReadCloser to its original state
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(message))
	}

	return resp, err
}

type MetalCloudClient struct {
	rpcClient jsonrpc.RPCClient
	user      string
	apiKey    string
	endpoint  string
}

func GetMetalcloudClient(user string, apiKey string, endpoint string) (*MetalCloudClient, error) {

	if user == "" {
		return nil, errors.New("user cannot be an empty string! It is typically in the form of user's email address.")
	}

	if apiKey == "" {
		return nil, errors.New("apiKey cannot be empty string!")
	}

	if endpoint == "" {
		return nil, errors.New("endpoint cannot be an empty string! It is typically in the form of user's email address.")
	}

	_, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return nil, err
	}

	transport := &SignatureAdderRoundTripper{
		APIKey:   apiKey,
		LogReply: false,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	rpcClient := jsonrpc.NewClientWithOpts(endpoint, &jsonrpc.RPCClientOpts{
		HTTPClient: httpClient,
	})

	return &MetalCloudClient{
		rpcClient: rpcClient,
		user:      user,
		apiKey:    apiKey,
		endpoint:  endpoint,
	}, nil

}
