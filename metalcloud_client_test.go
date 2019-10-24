package metalcloud

import (
	. "github.com/onsi/gomega"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"strings"
	"encoding/json"
)


// needed to retrieve requests that arrived at httpServer for further investigation
var requestChan = make(chan *RequestData, 1)

// the request datastructure that can be retrieved for test assertions
type RequestData struct {
	request *http.Request
	body    string
}

// set the response body the httpServer should return for the next request
var responseBody = ""

var httpServer *httptest.Server

// start the testhttp server and stop it when tests are finished
func TestMain(m *testing.M) {
	httpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		
		defer r.Body.Close()
		// put request and body to channel for the client to investigate them
		requestChan <- &RequestData{r, string(data)}

		fmt.Fprintf(w, responseBody)
	}))
	defer httpServer.Close()

	os.Exit(m.Run())
}


func TestSignature(t *testing.T){
	RegisterTestingT(t)

	content := strings.NewReader("asldklk234mlk234asd")
	
	request,err := http.NewRequest("GET", httpServer.URL, content)
	Expect(request).NotTo(BeNil())
	Expect(err).To(BeNil())

	transport := &SignatureAdderRoundTripper{
		APIKey:   "asdasdasd",
		LoggingEnabled: false,
		DryRun: false,
	}

	_, err = transport.RoundTrip(request)
	Expect(err).To(BeNil())

	requestWithSignature := (<-requestChan).request
	
	gotSignature	  := (requestWithSignature.URL).Query().Get("verify")
	expectedSignature := "b8287028c41c7dee8e6f0479ff05dd71"


	Expect(gotSignature).To(Equal(expectedSignature))
}


func TestEmptyListReply(t *testing.T){

	RegisterTestingT(t)
	
	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user","APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret1,err1 := mc.InstanceArrays(100)
	Expect(err1).To(BeNil())
	Expect(ret1).NotTo(BeNil())

	<-requestChan

	ret2,err2 := mc.Infrastructures()
	Expect(err2).To(BeNil())
	Expect(ret2).NotTo(BeNil())



}


func TestInstanceArrayCreateOmitEmpty(t *testing.T){

	RegisterTestingT(t)

	responseBody = `{"result":{"instance_array_label":"test"},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user","APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())
	Expect(mc).NotTo(BeNil())

	instanceArray := InstanceArray{
		InstanceArrayLabel: "test",
	}

	<-requestChan

	ret1,err1 := mc.InstanceArrayCreate(100, instanceArray)
	Expect(err1).To(BeNil())
	Expect(ret1).NotTo(BeNil())

	body := (<-requestChan).body
	
	//fmt.Printf("body=%s\n", body)

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	param := m["params"].([]interface{})[1].(map[string]interface{})
	
	Expect(param["instance_array_label"]).To(Equal("test"))
	Expect(param["volume_template_id"]).To(BeNil())

}


