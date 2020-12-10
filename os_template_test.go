package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestOSTemplateCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _osTemplateFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := OSTemplate{
		VolumeTemplateID:          100,
		VolumeTemplateDisplayName: "test-display-template",
		VolumeTemplateBootType:    "test-boot",
		VolumeTemplateOperatingSystem: &OperatingSystem{
			OperatingSystemType:         "os-type",
			OperatingSystemVersion:      "os-version",
			OperatingSystemArchitecture: "os-arch",
		},
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("os_template_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("os_template_update"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

	responseBody = `{"error": {"message": "OsTemplate not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("os_template_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("os_template_create"))

	params = (m["params"].([]interface{}))

	Expect(params[1].(map[string]interface{})["volume_template_id"].(float64)).To(Equal(100.0))
}

func TestOSTemplateDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := OSTemplate{
		VolumeTemplateID:          100,
		VolumeTemplateDisplayName: "test-display-template",
		VolumeTemplateBootType:    "test-boot",
		VolumeTemplateOperatingSystem: &OperatingSystem{
			OperatingSystemType:         "os-type",
			OperatingSystemVersion:      "os-version",
			OperatingSystemArchitecture: "os-arch",
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
	Expect(m["method"].(string)).To(Equal("os_template_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

const _osTemplateFixture = "{\"volume_template_id\": 100, \"volume_template_display_name\": \"test-display-template\", \"volume_template_boot_type\": \"test-boot\", \"volume_template_operating_system\": {\"operating_system_type\": \"os-type\", \"operating_system_version\": \"os-version\", \"operating_system_architecture\": \"os-arch\"}}"
