package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//Workflow struct defines a server type
type Workflow struct {
	WorkflowID               int    `json:"workflow_id,omitempty" yaml:"id,omitempty"`
	UserIDOwner              int    `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserIDAuthenticated      int    `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	WorkflowLabel            string `json:"workflow_label,omitempty" yaml:"label,omitempty"`
	WorkflowUsage            string `json:"workflow_usage,omitempty" yaml:"usage,omitempty"`
	WorkflowTitle            string `json:"workflow_title,omitempty" yaml:"title,omitempty"`
	WorkflowDescription      string `json:"workflow_description,omitempty" yaml:"description,omitempty"`
	WorkflowIsDeprecated     bool   `json:"workflow_is_deprecated,omitempty" yaml:"isDeprecated,omitempty"`
	IconAssetDataURI         string `json:"icon_asset_data_uri,omitempty" yaml:"assetDataURI,omitempty"`
	WorkflowCreatedTimestamp string `json:"workflow_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	WorkflowUpdatedTimestamp string `json:"workflow_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

//WorkflowStageDefinitionReference defines where in a workflow a stage definition resides
type WorkflowStageDefinitionReference struct {
	WorkflowStageID             int `json:"workflow_stage_id,omitempty"`
	WorkflowID                  int `json:"workflow_id,omitempty"`
	StageDefinitionID           int `json:"stage_definition_id,omitempty"`
	WorkflowStageRunLevel       int `json:"workflow_stage_run_level,omitempty"`
	WorkflowStageExecOutputJSON int `json:"workflow_stage_exec_output_json,omitempty"`
}

//WorkflowStageAssociation associations
type WorkflowStageAssociation struct {
	InfrastructureDeployCustomStageID             int    `json:"infrastructure_deploy_custom_stage_id,omitempty"`
	InfrastructureID                              int    `json:"infrastructure_id"`
	StageDefinitionID                             int    `json:"stage_definition_id,omitempty"`
	InfrastructureDeployCustomStageType           string `json:"infrastructure_deploy_custom_stage_type,omitempty"`
	InfrastructureDeployCustomStageRunLevel       int    `json:"infrastructure_deploy_custom_stage_run_level,omitempty"`
	InfrastructureDeployCustomStageExecOutputJSON string `json:"infrastructure_deploy_custom_stage_exec_output_json,omitempty"`
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

	var err error
	var createdObject map[string]Workflow
	var resp *jsonrpc.RPCResponse

	if usage != "" {
		resp, err = c.rpcClient.Call(
			"workflows",
			userID,
			usage,
		)
	} else {
		resp, err = c.rpcClient.Call(
			"workflows",
			userID,
		)
	}

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Workflow{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//WorkflowStages retrieves a list of all the StageDefinitions objects in this workflow
func (c *Client) WorkflowStages(workflowID int) (*[]WorkflowStageDefinitionReference, error) {
	var createdObject []WorkflowStageDefinitionReference

	err := c.rpcClient.CallFor(
		&createdObject,
		"workflow_stages",
		workflowID)

	if err != nil {
		return nil, err
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

//WorkflowStageDelete deletes a stage from a workflow entirelly
func (c *Client) WorkflowStageDelete(workflowStageID int) error {

	if err := checkID(workflowStageID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("workflow_stage_delete", workflowStageID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//InfrastructureDeployCustomStageAddIntoRunlevel adds a stage into a runlevel
func (c *Client) InfrastructureDeployCustomStageAddIntoRunlevel(infraID int, stageID int, runLevel int, stageRunMoment string) error {

	resp, err := c.rpcClient.Call("infrastructure_deploy_custom_stage_add_into_runlevel", infraID, stageID, runLevel, stageRunMoment)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//InfrastructureDeployCustomStageDeleteIntoRunlevel delete a stage into a runlevel
func (c *Client) InfrastructureDeployCustomStageDeleteIntoRunlevel(infraID int, stageID int, runLevel int, stageRunMoment string) error {

	resp, err := c.rpcClient.Call("infrastructure_deploy_custom_stage_delete_into_runlevel", infraID, stageID, runLevel, stageRunMoment)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//InfrastructureDeployCustomStages retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
func (c *Client) InfrastructureDeployCustomStages(infraID int, stageDefinitionType string) (*[]WorkflowStageAssociation, error) {
	var createdObject []WorkflowStageAssociation

	err := c.rpcClient.CallFor(
		&createdObject,
		"infrastructure_deploy_custom_stages",
		infraID,
		stageDefinitionType)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//CreateOrUpdate implements interface Applier
func (w Workflow) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var result *Workflow
	err = w.Validate()

	if err != nil {
		return err
	}

	if w.WorkflowID != 0 {
		result, err = client.WorkflowGet(w.WorkflowID)
	} else {
		wflows, err := client.Workflows()
		if err != nil {
			return err
		}
		for _, wflow := range *wflows {
			if wflow.WorkflowLabel == w.WorkflowLabel {
				result = &wflow
			}
		}
	}

	if err != nil {
		_, err = client.WorkflowCreate(w)

		if err != nil {
			return err
		}
	} else {
		_, err = client.WorkflowUpdate(result.WorkflowID, w)
		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (w Workflow) Delete(client MetalCloudClient) error {
	var result *Workflow
	var id int
	err := w.Validate()

	if err != nil {
		return err
	}

	if w.WorkflowID != 0 {
		id = w.WorkflowID
	} else {
		wflows, err := client.Workflows()
		if err != nil {
			return err
		}
		for _, wflow := range *wflows {
			if wflow.WorkflowLabel == w.WorkflowLabel {
				result = &wflow
			}
		}

		id = result.WorkflowID
	}

	err = client.WorkflowDelete(id)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (w Workflow) Validate() error {
	if w.WorkflowID == 0 && w.WorkflowLabel == "" {
		return fmt.Errorf("id is required")
	}
	if w.WorkflowUsage == "" {
		return fmt.Errorf("usage is required")
	}
	return nil
}
