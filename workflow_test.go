package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestWorkflowStageDefinitionReferenceUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	var list []WorkflowStageDefinitionReference
	err := json.Unmarshal([]byte(_workflowStagesFixture1), &list)
	Expect(err).To(BeNil())

	obj := list[0]
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.WorkflowStageID).To(Equal(44000))
	Expect(obj.StageDefinitionID).To(Equal(322))
	Expect(obj.WorkflowID).To(Equal(9999))
}

func TestWorkflowStages(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result":` + _workflowStagesFixture1 + `},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.WorkflowStages(10)

	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	Expect((*ret)[0].WorkflowStageID).To(Equal(44000))

	(<-requestChan) //this clears the channel

}
func TestWorkflowCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _workflowFixture1 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Workflow{
		WorkflowID:    100,
		WorkflowUsage: "test-usage",
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("workflow_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("workflow_update"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

	responseBody = `{"error": {"message": "Workflow not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("workflow_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("workflow_create"))

	params = (m["params"].([]interface{}))

	Expect(params[1].(map[string]interface{})["workflow_id"].(float64)).To(Equal(100.0))
}

func TestWorkflowDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Workflow{
		WorkflowID:    100,
		WorkflowUsage: "test-usage",
	}
	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("workflow_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

const _workflowStagesFixture1 = "[\n        {\n            \"workflow_stage_id\": 44000,\n            \"workflow_id\": 9999,\n            \"stage_definition_id\": 322,\n            \"workflow_stage_run_level\": 10,\n            \"workflow_stage_exec_output_json\": null\n        }\n ]"
const _workflowFixture1 = "{\"workflow_id\": 100, \"workflow_usage\": \"test-usage\"}"
