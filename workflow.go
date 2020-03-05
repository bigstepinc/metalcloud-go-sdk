package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//Workflow struct defines a server type
type Workflow struct {
	WorkflowID               int    `json:"workflow_id,omitempty"`
	UserIDOwner              int    `json:"user_id_owner,omitempty"`
	UserIDAuthenticated      int    `json:"user_id_authenticated,omitempty"`
	WorkflowLabel            string `json:"workflow_label,omitempty"`
	WorkflowUsage            string `json:"workflow_usage,omitempty"`
	WorkflowTitle            string `json:"workflow_title,omitempty"`
	WorkflowDescription      string `json:"workflow_description,omitempty"`
	WorkflowIsDeprecated     bool   `json:"workflow_is_deprecated,omitempty"`
	IconAssetDataURI         string `json:"icon_asset_data_uri,omitempty"`
	WorkflowCreatedTimestamp string `json:"workflow_created_timestamp,omitempty"`
	WorkflowUpdatedTimestamp string `json:"workflow_updated_timestamp,omitempty"`
}

//WorkflowStageDefinitionReference defines where in a workflow a stage definition resides
type WorkflowStageDefinitionReference struct {
	WorkflowStageID             int `json:"workflow_stage_id,omitempty"`
	WorkflowID                  int `json:"workflow_id,omitempty"`
	StageDefinitionID           int `json:"stage_definition_id,omitempty"`
	WorkflowStageRunLevel       int `json:"workflow_stage_run_level,omitempty"`
	WorkflowStageExecOutputJSON int `json:"workflow_stage_exec_output_json,omitempty"`
}

//WorkflowCreate creates a workflow
func (c *Client) WorkflowCreate(workflow Workflow) (*Workflow, error) {
	var createdObject Workflow

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"workflow_create",
		userID,
		workflow)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//WorkflowDelete Permanently destroys a Workflow.
func (c *Client) WorkflowDelete(workflowID int) error {

	if err := checkID(workflowID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_delete", workflowID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//WorkflowUpdate This function allows updating the workflow_usage, workflow_label and workflow_base64 of a Workflow
func (c *Client) WorkflowUpdate(workflowID int, workflow Workflow) (*Workflow, error) {
	var createdObject Workflow

	if err := checkID(workflowID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"workflow_update",
		workflowID,
		workflow)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//WorkflowGet returns a Workflow specified by nWorkflowID. The workflow's protected value is never returned.
func (c *Client) WorkflowGet(workflowID int) (*Workflow, error) {

	var createdObject Workflow

	if err := checkID(workflowID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"workflow_get",
		workflowID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//Workflows retrieves a list of all the Workflow objects which a specified User is allowed to see through ownership or delegation.
func (c *Client) Workflows() (*map[string]Workflow, error) {
	return c.WorkflowsWithUsage("")
}

//WorkflowsWithUsage retrieves a list of all the Workflow objects which the current User is allowed to see through ownership or delegation with a specific usage.
func (c *Client) WorkflowsWithUsage(usage string) (*map[string]Workflow, error) {

	userID := c.GetUserID()

	var res *jsonrpc.RPCResponse
	var err error
	if usage != "" {
		res, err = c.rpcClient.Call(
			"workflows",
			userID,
			usage)
	} else {
		res, err = c.rpcClient.Call(
			"workflows",
			userID)
	}

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Workflow{}
		return &m, nil
	}

	var createdObject map[string]Workflow

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//WorkflowStages retrieves a list of all the StageDefinitions objects in this workflow
func (c *Client) WorkflowStages(workflowID int) (*[]WorkflowStageDefinitionReference, error) {

	var res *jsonrpc.RPCResponse

	res, err := c.rpcClient.Call(
		"workflow_stages",
		workflowID)

	if err != nil {
		return nil, err
	}

	var createdObject []WorkflowStageDefinitionReference

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//WorkflowStageGet returns a StageDefinition specified by workflowStageID.
func (c *Client) WorkflowStageGet(workflowStageID int) (*WorkflowStageDefinitionReference, error) {

	var createdObject WorkflowStageDefinitionReference

	if err := checkID(workflowStageID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"workflow_stage_get",
		workflowStageID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//WorkflowStageAddAsNewRunLevel adds a new stage in this workflow
func (c *Client) WorkflowStageAddAsNewRunLevel(workflowID int, stageDefinitionID int, destinationRunLevel int) error {

	if err := checkID(workflowID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_stage_add_as_new_runlevel", workflowID, stageDefinitionID, destinationRunLevel)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//WorkflowStageAddIntoRunLevel adds a new stage in this workflow
func (c *Client) WorkflowStageAddIntoRunLevel(workflowID int, stageDefinitionID int, destinationRunLevel int) error {

	if err := checkID(workflowID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_stage_add_into_runlevel", workflowID, stageDefinitionID, destinationRunLevel)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//WorkflowDeleteFromRunLevel removes a  stage in this workflow from a runlevel
func (c *Client) WorkflowDeleteFromRunLevel(workflowID int, stageDefinitionID int, destinationRunLevel int) error {

	if err := checkID(workflowID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_stage_delete_from_runlevel", workflowID, stageDefinitionID, destinationRunLevel)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//WorkflowMoveAsNewRunLevel moves a stage in this workflow from a runlevel to another
func (c *Client) WorkflowMoveAsNewRunLevel(workflowID int, stageDefinitionID int, sourceRunLevel int, destinationRunLevel int) error {

	if err := checkID(workflowID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_stage_move_as_new_runlevel", workflowID, stageDefinitionID, sourceRunLevel, destinationRunLevel)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//WorkflowMoveIntoRunLevel moves a stage in this workflow from a runlevel to another
func (c *Client) WorkflowMoveIntoRunLevel(workflowID int, stageDefinitionID int, sourceRunLevel int, destinationRunLevel int) error {

	if err := checkID(workflowID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_stage_move_into_runlevel", workflowID, stageDefinitionID, sourceRunLevel, destinationRunLevel)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
