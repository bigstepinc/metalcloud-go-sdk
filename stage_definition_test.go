package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestStageDefinitionUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj StageDefinition
	err := json.Unmarshal([]byte(_stageDefinitionsFixture), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.StageDefinition).ToNot((BeNil()))
	Expect(obj.StageDefinitionID).To(Equal(402))
	Expect(obj.UserIDOwner).To(Equal(8030))

	req := obj.StageDefinition.(HTTPRequest)
	Expect(req.URL).To(Equal("https://raw.githubusercontent.com/bigstepinc/terraform-provider-metalcloud/master/README.md"))
}

func TestStageDefinitionMarshalTest(t *testing.T) {
	RegisterTestingT(t)

	//get an object

	ab := AnsibleBundle{
		AnsibleBundleArchiveFilename: "test",
	}

	obj := StageDefinition{
		StageDefinitionID:    100,
		StageDefinitionLabel: "test",
		StageDefinition:      ab,
		StageDefinitionType:  "AnsibleBundle",
	}

	//we get bytes
	b, err := json.Marshal(obj)
	Expect(err).To(BeNil())
	Expect(b).NotTo(BeNil())

	//unmarshal
	var obj2 StageDefinition
	err = json.Unmarshal(b, &obj2)
	Expect(err).To(BeNil())

	//check that we have ansible bundle
	ab2 := obj2.StageDefinition.(AnsibleBundle)
	Expect(ab2.AnsibleBundleArchiveFilename).To(Equal(ab.AnsibleBundleArchiveFilename))

}

func TestStageDefinitionCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _stageDefinitionsFixture1 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := StageDefinition{
		StageDefinitionID:    100,
		StageDefinitionType:  "sd-test-type",
		StageDefinitionTitle: "sd-test-title",
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("stage_definition_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("stage_definition_update"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

	responseBody = `{"error": {"message": "Stage definition not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("stage_definition_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("stage_definition_create"))

	params = (m["params"].([]interface{}))

	Expect(params[1].(map[string]interface{})["stage_definition_title"].(string)).To(Equal("sd-test-title"))
}

func TestStageDefinitionDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := StageDefinition{
		StageDefinitionID:    100,
		StageDefinitionType:  "sd-test-type",
		StageDefinitionTitle: "sd-test-title",
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("stage_definition_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

const _stageDefinitionsFixture = "{\"stage_definition_id\":402,\"user_id_owner\":8030,\"user_id_authenticated\":8030,\"stage_definition_is_deprecated\":false,\"stage_definition_label\":\"my-test-stage-3\",\"stage_definition_label_unique\":\"my-test-stage-3@8030\",\"stage_definition_title\":\"my title\",\"stage_definition_description\":\"\",\"stage_definition_type\":\"HTTPRequest\",\"icon_asset_data_uri\":null,\"stage_definition_created_timestamp\":\"2020-01-28T17:20:03Z\",\"stage_definition_updated_timestamp\":\"2020-01-28T17:20:03Z\",\"stage_definition_variable_names_required\":[],\"stage_definition\":{\"url\":\"https://raw.githubusercontent.com/bigstepinc/terraform-provider-metalcloud/master/README.md\",\"type\":\"HTTPRequest\",\"options\":{\"body\":null,\"size\":67108864,\"follow\":40,\"method\":\"GET\",\"headers\":{\"Pragma\":\"no-cache\",\"Connection\":\"close\",\"User-Agent\":\"MetalCloud/1.0 (WebFetchAPI)\",\"Cache-Control\":\"no-cache\"},\"timeout\":240000,\"compress\":true,\"redirect\":\"follow\",\"bodyBufferBase64\":null}}}"
const _stageDefinitionsFixture1 = "{\"stage_definition_id\": 100, \"stage_definition_title\": \"sd-test-title\", \"stage_definition_type\": \"sd-test-type\"}"
