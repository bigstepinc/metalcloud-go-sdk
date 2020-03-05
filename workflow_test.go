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

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.WorkflowStages(10)

	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	Expect((*ret)[0].WorkflowStageID).To(Equal(44000))

	(<-requestChan) //this clears the channel

}

const _workflowStagesFixture1 = "[\n        {\n            \"workflow_stage_id\": 44000,\n            \"workflow_id\": 9999,\n            \"stage_definition_id\": 322,\n            \"workflow_stage_run_level\": 10,\n            \"workflow_stage_exec_output_json\": null\n        }\n ]"
