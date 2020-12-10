package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestInstanceArrayCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _instanceArrayFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := InstanceArray{
		InstanceArrayID:    100,
		InstanceArrayLabel: "ia-test",
		InfrastructureID:   2,
		InstanceArrayOperation: &InstanceArrayOperation{
			InstanceArrayID:       100,
			InstanceArrayLabel:    "ia-test",
			InstanceArrayChangeID: 200,
		},
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("instance_array_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("instance_array_edit"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(string)).To(Equal("ia-test"))

	responseBody = `{"error": {"message": "Instance array not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("instance_array_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("instance_array_create"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(2.0))
}

func TestInstanceArrayDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := InstanceArray{
		InstanceArrayID:    100,
		InstanceArrayLabel: "ia-test",
		InfrastructureID:   2,
		InstanceArrayOperation: &InstanceArrayOperation{
			InstanceArrayID:       100,
			InstanceArrayLabel:    "ia-test",
			InstanceArrayChangeID: 200,
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
	Expect(m["method"].(string)).To(Equal("instance_array_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

const _instanceArrayFixture = "{\"instance_array_id\": 100, \"instance_array_label\": \"ia-test\", \"instance_array_operation\": {\"instance_array_change_id\": 200, \"instance_array_id\": 100, \"instance_array_label\": \"ia-test\"}}"
