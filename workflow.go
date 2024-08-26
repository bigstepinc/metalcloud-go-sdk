//go:generate go run helper/docgen.go - $GOFILE ./ Workflow,WorkflowStageDefinitionReference,WorkflowStageAssociation Workflow
package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

// Workflow struct defines a server type
type Workflow struct {
	//description: Id of the object
	WorkflowID int `json:"workflow_id,omitempty" yaml:"id,omitempty"`
	//description: Id of the owner
	UserIDOwner int `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	//description: Internal.
	UserIDAuthenticated int `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	//description: Label
	WorkflowLabel string `json:"workflow_label,omitempty" yaml:"label,omitempty"`
	//description: Usage of this workflow
	//values:
	// - infrastructure
	// - switch_device
	// - server
	// - free_standing
	// - storage_pool
	// - user
	// - os_template
	// - global
	WorkflowUsage string `json:"workflow_usage,omitempty" yaml:"usage,omitempty"`
	//description: Title (name) of this workflow
	WorkflowTitle string `json:"workflow_title,omitempty" yaml:"title,omitempty"`
	//description: Description of this workflow
	WorkflowDescription string `json:"workflow_description,omitempty" yaml:"description,omitempty"`
	//description: Set to true if this workflow is deprecated
	WorkflowIsDeprecated bool `json:"workflow_is_deprecated,omitempty" yaml:"isDeprecated,omitempty"`
	//description: Internal
	IconAssetDataURI string `json:"icon_asset_data_uri,omitempty" yaml:"assetDataURI,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the workflow record was created. Readonly.
	//example: 2013-11-29T13:00:01Z
	WorkflowCreatedTimestamp string `json:"workflow_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the workflow record was updated. Readonly.
	//example: 2013-11-29T13:00:01Z
	WorkflowUpdatedTimestamp string `json:"workflow_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

// WorkflowStageDefinitionReference defines where in a workflow a stage definition resides
type WorkflowStageDefinitionReference struct {
	//description: Workflow stage id
	WorkflowStageID int `json:"workflow_stage_id,omitempty" yaml:"workflowStageID,omitempty"`
	//description: Workflow id
	WorkflowID int `json:"workflow_id,omitempty" yaml:"workflowID,omitempty"`
	//description: Stage definition id
	StageDefinitionID int `json:"stage_definition_id,omitempty" yaml:"stageDefinitionID,omitempty"`
	//description: Run level in the workflow
	WorkflowStageRunLevel int `json:"workflow_stage_run_level,omitempty" yaml:"runLevel,omitempty"`
}

// WorkflowStageAssociation associations
type WorkflowStageAssociation struct {
	//description: Infrastructure stage id
	InfrastructureDeployCustomStageID int `json:"infrastructure_deploy_custom_stage_id,omitempty" yaml:"infrastructureDeployStageID,omitempty"`
	//description: Infrastructure id
	InfrastructureID int `json:"infrastructure_id" yaml:"infrastructureID,omitempty"`
	//description: Stage definition id
	StageDefinitionID int `json:"stage_definition_id,omitempty" yaml:"definitionID,omitempty"`
	//description: Type of position in the infrastructure:
	//values:
	// - pre_deploy
	// - post_deploy
	InfrastructureDeployCustomStageType string `json:"infrastructure_deploy_custom_stage_type,omitempty" yaml:"type,omitempty"`
	//description: Run level (the depth in the tree where this stage resides).
	InfrastructureDeployCustomStageRunLevel int `json:"infrastructure_deploy_custom_stage_run_level,omitempty" yaml:"runLevel,omitempty"`
	//description: The output of the last execution of this task
	InfrastructureDeployCustomStageExecOutputJSON string `json:"infrastructure_deploy_custom_stage_exec_output_json,omitempty" yaml:"lastOutput,omitempty"`
}

// WorkflowCreate creates a workflow
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

// WorkflowDelete Permanently destroys a Workflow.
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

// WorkflowUpdate This function allows updating the workflow_usage, workflow_label and workflow_base64 of a Workflow
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

// WorkflowGet returns a Workflow specified by nWorkflowID. The workflow's protected value is never returned.
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

// Workflows retrieves a list of all the Workflow objects which a specified User is allowed to see through ownership or delegation.
func (c *Client) Workflows() (*map[string]Workflow, error) {
	return c.WorkflowsWithUsage("")
}

// WorkflowsWithUsage retrieves a list of all the Workflow objects which the current User is allowed to see through ownership or delegation with a specific usage.
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

// WorkflowStages retrieves a list of all the StageDefinitions objects in this workflow
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

// WorkflowStageGet returns a StageDefinition specified by workflowStageID.
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

// WorkflowStageAddAsNewRunLevel adds a new stage in this workflow
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

// WorkflowStageAddIntoRunLevel adds a new stage in this workflow
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

// WorkflowMoveAsNewRunLevel moves a stage in this workflow from a runlevel to another
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

// WorkflowMoveIntoRunLevel moves a stage in this workflow from a runlevel to another
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

// WorkflowStageDelete deletes a stage from a workflow entirelly
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

// InfrastructureDeployCustomStageAddIntoRunlevel adds a stage into a runlevel
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

// InfrastructureDeployCustomStageDelete delete a stage
func (c *Client) InfrastructureDeployCustomStageDelete(nInfrastructureCustomDeployStageID int) error {

	resp, err := c.rpcClient.Call(
		"infrastructure_deploy_custom_stage_delete",
		nInfrastructureCustomDeployStageID)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// InfrastructureDeployCustomStages retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
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

// CreateOrUpdate implements interface Applier
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

// Delete implements interface Applier
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

// Validate implements interface Applier
func (w Workflow) Validate() error {
	if w.WorkflowID == 0 && w.WorkflowLabel == "" {
		return fmt.Errorf("id is required")
	}
	if w.WorkflowUsage == "" {
		return fmt.Errorf("usage is required")
	}
	return nil
}
