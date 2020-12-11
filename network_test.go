package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestNetworkCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _networkFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := Network{
		NetworkID:                 101,
		NetworkLabel:              "net-test",
		InfrastructureID:          1,
		NetworkSubdomain:          "sub-test.test",
		NetworkType:               "test-net-type",
		NetworkLANAutoAllocateIPs: false,
		NetworkOperation: &NetworkOperation{
			NetworkID:        101,
			NetworkLabel:     "net-test",
			InfrastructureID: 1,
			NetworkChangeID:  3,
		},
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("network_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(101.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("network_edit"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(101.0))

	responseBody = `{"error": {"message": "Network not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("network_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(101.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("network_create"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(1.0))
}

func TestNetworkDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := Network{
		NetworkID:                 101,
		NetworkLabel:              "net-test",
		InfrastructureID:          1,
		NetworkSubdomain:          "sub-test.test",
		NetworkType:               "test-net-type",
		NetworkLANAutoAllocateIPs: false,
		NetworkOperation: &NetworkOperation{
			NetworkID:        101,
			NetworkLabel:     "net-test",
			InfrastructureID: 1,
			NetworkChangeID:  3,
		},
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("network_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(101.0))

}

const _networkFixture = "{\"network_id\": 101, \"network_label\": \"net-test\", \"infrastructure_id\": 1, \"network_operation\": {\"network_id\": 101, \"network_label\": \"net-test\", \"network_change_id\": 3, \"infrastructure_id\": 1}}"
