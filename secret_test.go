package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSecretCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _secretFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := Secret{
		SecretID:   100,
		SecretName: "secret-test",
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("secret_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("secret_update"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

	responseBody = `{"error": {"message": "Secret not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("secret_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("secret_create"))

	params = (m["params"].([]interface{}))

	Expect(params[1].(map[string]interface{})["secret_id"].(float64)).To(Equal(100.0))
}

func TestSecretDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := Secret{
		SecretID:   100,
		SecretName: "secret-test",
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("secret_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

const _secretFixture = "{\"secret_id\": 100, \"secret_name\": \"test-secret\"}"
