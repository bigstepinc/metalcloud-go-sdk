package metalcloud

import "github.com/ybbus/jsonrpc"

//go:generate go run helper/gen_exports.go

//WorkflowStageAssociation associations
type WorkflowStageAssociation struct {
	InfrastructureDeployCustomStageID             int    `json:"infrastructure_deploy_custom_stage_id,omitempty"`
	InfrastructureID                              int    `json:"infrastructure_id,omitempty"`
	StageDefinitionID                             int    `json:"stage_definition_id,omitempty"`
	InfrastructureDeployCustomStageType           string `json:"infrastructure_deploy_custom_stage_type,omitempty"`
	InfrastructureDeployCustomStageRunLevel       int    `json:"infrastructure_deploy_custom_stage_run_level,omitempty"`
	InfrastructureDeployCustomStageExecOutputJSON string `json:"infrastructure_deploy_custom_stage_exec_output_json,omitempty"`
}

//InfrastructureDeployCustomStageAddIntoRunlevel adds a stage into a runlevel
func (c *Client) InfrastructureDeployCustomStageAddIntoRunlevel(infraID int, stageID int, runLevel int, stageRunMoment string) error {

	_, err := c.rpcClient.Call("infrastructure_deploy_custom_stage_add_into_runlevel", infraID, stageID, runLevel, stageRunMoment)
	if err != nil {
		return err
	}

	return nil
}

//InfrastructureDeployCustomStageDeleteIntoRunlevel delete a stage into a runlevel
func (c *Client) InfrastructureDeployCustomStageDeleteIntoRunlevel(infraID int, stageID int, runLevel int, stageRunMoment string) error {

	_, err := c.rpcClient.Call("infrastructure_deploy_custom_stage_delete_into_runlevel", infraID, stageID, runLevel, stageRunMoment)
	if err != nil {
		return err
	}

	return nil
}

//InfrastructureDeployCustomStages retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
func (c *Client) InfrastructureDeployCustomStages(infraID int, stageDefinitionType string) (*[]WorkflowStageAssociation, error) {

	var res *jsonrpc.RPCResponse

	res, err := c.rpcClient.Call(
		"infrastructure_deploy_custom_stages",
		infraID,
		stageDefinitionType)

	if err != nil {
		return nil, err
	}

	var createdObject []WorkflowStageAssociation

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}
