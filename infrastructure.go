package metalcloud

//go:generate go run helper/gen_exports.go
//go:generate go run helper/docgen.go - $GOFILE ./ Infrastructure,InfrastructureOperation Infrastructure

import (
	"fmt"
)

// description - The infrastructure object. All other client-side objects (InstanceArrays, DriveArrays) fall under it in the object hierarchy
type Infrastructure struct {
	// description: The id of the object
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"id,omitempty"`
	// description: The label of the object
	InfrastructureLabel string `json:"infrastructure_label" yaml:"label"`
	// description: The datacenter label on which this infrastructure resides
	DatacenterName string `json:"datacenter_name" yaml:"datacenter"`
	// description: An automatically generated string that represents this infrastructure. Deprecated.
	InfrastructureSubdomain string `json:"infrastructure_subdomain,omitempty" yaml:"subdomain,omitempty"`
	// description: The ID of the user that is the owner of the infrastructure.
	UserIDowner int `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	// description: The email of the user that is the owner of the infrastructure. Readonly
	UserEmailOwner string `json:"user_email_owner,omitempty" yaml:"ownerEmail,omitempty"`
	// description: Internal.
	InfrastructureTouchUnixtime string `json:"infrastructure_touch_unixtime,omitempty" yaml:"touchUnixTime,omitempty"`
	// description: The service status.  Read only.
	// values:
	//  - ordered #designed but not yet deployed
	//  - active #deployed
	//  - suspended #deployed but suspended in a way that cannot be resumed by the user. (typically this means shutdown and disconnected from the internet).
	//  - stopped #deployed but suspended in a way that can be resumed by the user.
	//  - deleted #deleted permanently
	InfrastructureServiceStatus string `json:"infrastructure_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	// description: ISO 8601 timestamp which holds the date and time when the infrastructure was created.
	// example: 2013-11-29T13:00:01Z
	InfrastructureCreatedTimestamp string `json:"infrastructure_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	// description: ISO 8601 timestamp which holds the date and time when the infrastructure was last updated.
	// example: 2013-11-29T13:00:01Z
	InfrastructureUpdatedTimestamp string `json:"infrastructure_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	// description: The id of the ongoing change operation. readonly
	InfrastructureChangeID int `json:"infrastructure_change_id,omitempty" yaml:"changeID,omitempty"`
	// description: The ID of the current deploy operation (AFC group id). Readonly.
	InfrastructureDeployID int `json:"infrastructure_deploy_id,omitempty" yaml:"deployID,omitempty"`
	// description: Set if the infrastructure cannot be changed.
	InfrastructureDesignIsLocked bool `json:"infrastructure_design_is_locked,omitempty" yaml:"designIsLocked,omitempty"`
	// description: The operation object
	InfrastructureOperation InfrastructureOperation `json:"infrastructure_operation,omitempty" yaml:"operation,omitempty"`
	// description: Internal
	InfrastructureExperimentalPriority string `json:"infrastructure_experimental_priority,omitempty"`
	// description: Custom variables at the infrastructure level. They will be pushed to all objects (such as operating systems) and workflows.
	InfrastructureCustomVariables interface{} `json:"infrastructure_custom_variables,omitempty" yaml:"customVariables,omitempty"`
}

// InfrastructureOperation - object with alternations to be applied
type InfrastructureOperation struct {
	// description: The id of the object
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"id,omitempty"`
	// description: The label of the object
	InfrastructureLabel string `json:"infrastructure_label" yaml:"label"`
	// description: The datacenter label on which this infrastructure resides
	DatacenterName string `json:"datacenter_name" yaml:"datacenter"`
	// description: The status of the ongoing deployment operation (if any)
	// values:
	//     - not_started
	//     - ongoing
	//     - finished
	InfrastructureDeployStatus string `json:"infrastructure_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	// description: The type of the current deploy operation (if any)
	// values:
	//     - create
	// 	   - delete
	//     - edit
	// 	   - start
	//     - stop
	// 	   - suspend
	InfrastructureDeployType string `json:"infrastructure_deploy_type,omitempty" yaml:"deployType,omitempty"`
	// description: An automatically generated string that represents this infrastructure. Deprecated.
	InfrastructureSubdomain string `json:"infrastructure_subdomain,omitempty" yaml:"subdomain,omitempty"`
	// description: The ID of the user that is the owner of the infrastructure.
	UserIDOwner int `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	// description: ISO 8601 timestamp which holds the date and time when the infrastructure was last updated.
	// example: 2013-11-29T13:00:01Z
	InfrastructureUpdatedTimestamp string `json:"infrastructure_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	// description: The ID of the current change operation. Readonly.
	InfrastructureChangeID int `json:"infrastructure_change_id,omitempty" yaml:"changeID,omitempty"`
	// description: The ID of the current deploy operation (AFC group id). Readonly.
	InfrastructureDeployID int `json:"infrastructure_deploy_id,omitempty" yaml:"deployID,omitempty"`
	// description: Custom variables at the infrastructure level. They will be pushed to all objects (such as operating systems) and workflows.
	InfrastructureCustomVariables interface{} `json:"infrastructure_custom_variables,omitempty" yaml:"customVariables,omitempty"`
}

// ShutdownOptions controls how the deploy engine handles running instances
type ShutdownOptions struct {
	HardShutdownAfterTimeout   bool `json:"hard_shutdown_after_timeout"`
	AttemptSoftShutdown        bool `json:"attempt_soft_shutdown"`
	SoftShutdownTimeoutSeconds int  `json:"soft_shutdown_timeout_seconds"`
}

// DeployOptions controls server allocation
type DeployOptions struct {
	InstanceArrayMapping map[int]map[string]DeployOptionsServerTypeMappingObject `json:"instance_array"`
}

// DeployOptionsServerTypeMappingObject respresents one of the server type mappings
type DeployOptionsServerTypeMappingObject struct {
	ServerCount int   `json:"server_count"`
	ServerIDs   []int `json:"server_ids"`
}

type searchResultWrapperForInfrastructures struct {
	DurationMilliseconds int                           `json:"duration_millisecnds,omitempty"`
	Rows                 []InfrastructuresSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string                    `json:"rows_order,omitempty"`
	RowsTotal            int                           `json:"rows_total,omitempty"`
}

type InfrastructuresSearchResult struct {
	InfrastructureID               int      `json:"infrastructure_id,omitempty" yaml:"id,omitempty"`
	InfrastructureLabel            string   `json:"infrastructure_label" yaml:"label"`
	InfrastructureSubdomain        string   `json:"infrastructure_subdomain" yaml:"label"`
	InfrastructureServiceStatus    string   `json:"infrastructure_service_status" yaml:"label"`
	InfrastructureDeployStatus     string   `json:"infrastructure_deploy_status" yaml:"label"`
	DatacenterName                 string   `json:"datacenter_name" yaml:"datacenter"`
	InfrastructureCreatedTimestamp string   `json:"infrastructure_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	InfrastructureUpdatedTimestamp string   `json:"infrastructure_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	UserIDOwner                    int      `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserEmail                      []string `json:"user_email,omitempty" yaml:"userEmail,omitempty"`
	InfrastructureDeployID         int      `json:"infrastructure_deploy_id,omitempty" yaml:"deployID,omitempty"`
	AFCGroupCreatedTimestamp       string   `json:"afc_group_created_timestamp,omitempty" yaml:"afcGroupCreatedTimestamp,omitempty"`
	AFCGroupFinishedTimestamp      string   `json:"afc_group_finished_timestamp,omitempty" yaml:"afcGroupFinishedTimestamp,omitempty"`
	AFCThrownError                 int      `json:"thrownError,omitempty" yaml:"thrownError,omitempty"`
	AFCExecutedSuccess             int      `json:"executedSuccess,omitempty" yaml:"executedSuccess,omitempty"`
	AFCTotal                       int      `json:"total,omitempty" yaml:"total,omitempty"`
}

// InfrastructureCreate creates an infrastructure
func (c *Client) InfrastructureCreate(infrastructure Infrastructure) (*Infrastructure, error) {
	var createdObject Infrastructure

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"infrastructure_create",
		userID,
		infrastructure)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// infrastructureEdit alters an infrastructure
func (c *Client) infrastructureEdit(infrastructureID id, infrastructureOperation InfrastructureOperation) (*Infrastructure, error) {
	var createdObject Infrastructure

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"infrastructure_edit",
		infrastructureID,
		infrastructureOperation)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// infrastructureDelete deletes an infrastructure and all associated elements. Requires deploy
func (c *Client) infrastructureDelete(infrastructureID id) error {

	if err := checkID(infrastructureID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("infrastructure_delete", infrastructureID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// infrastructureOperationCancel reverts (undos) alterations done before deploy
func (c *Client) infrastructureOperationCancel(infrastructureID id) error {

	if err := checkID(infrastructureID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"infrastructure_operation_cancel",
		infrastructureID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// infrastructureDeploy initiates a deploy operation that will apply all registered changes for the respective infrastructure
func (c *Client) infrastructureDeploy(infrastructureID id, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error {
	return c.infrastructureDeployWithOptions(infrastructureID, shutdownOptions, nil, allowDataLoss, skipAnsible)
}

// infrastructureDeployWithOptions initiates a deploy operation that will apply all registered changes for the respective infrastructure. With options.
func (c *Client) infrastructureDeployWithOptions(infrastructureID id, shutdownOptions ShutdownOptions, deployOptions *DeployOptions, allowDataLoss bool, skipAnsible bool) error {

	if err := checkID(infrastructureID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"infrastructure_deploy",
		infrastructureID,
		shutdownOptions,
		deployOptions,
		allowDataLoss,
		skipAnsible,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// Infrastructures returns a list of infrastructures
func (c *Client) Infrastructures() (*map[string]Infrastructure, error) {
	userID := c.GetUserID()

	resp, err := c.rpcClient.Call(
		"infrastructures",
		userID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Infrastructure{}
		return &m, nil
	}

	var createdObject map[string]Infrastructure

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// infrastructureGet returns a specific infrastructure by id
func (c *Client) infrastructureGet(infrastructureID id) (*Infrastructure, error) {
	var infrastructure Infrastructure

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&infrastructure, "infrastructure_get", infrastructureID)

	if err != nil {
		return nil, err
	}

	return &infrastructure, nil
}

// infrastructureUserLimits returns user metadata
func (c *Client) infrastructureUserLimits(infrastructureID id) (*map[string]interface{}, error) {
	var userLimits map[string]interface{}

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&userLimits, "infrastructure_user_limits", infrastructureID)

	if err != nil {
		return nil, err
	}

	return &userLimits, nil
}

func (c *Client) infrastructureInstances(infrastructureID id) (*map[string]interface{}, error) {
	var instances interface{}

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&instances, "infrastructure_instances_info", infrastructureID)

	if err != nil {
		return nil, err
	}

	_, ok := instances.([]interface{})

	if ok {
		var m = make(map[string]interface{})
		return &m, nil
	}

	instancesIntf := instances.(map[string]interface{})

	return &instancesIntf, nil
}

func (i *Infrastructure) instanceToOperation(op *InfrastructureOperation) {
	operation := &i.InfrastructureOperation
	operation.InfrastructureID = i.InfrastructureID
	operation.InfrastructureLabel = i.InfrastructureLabel
	operation.DatacenterName = i.DatacenterName
	operation.InfrastructureSubdomain = i.InfrastructureSubdomain
	operation.InfrastructureUpdatedTimestamp = i.InfrastructureUpdatedTimestamp
	operation.InfrastructureChangeID = op.InfrastructureChangeID
}

// CreateOrUpdate implements interface Applier
func (i Infrastructure) CreateOrUpdate(client MetalCloudClient) error {
	var result *Infrastructure
	var err error

	err = i.Validate()

	if err != nil {
		return err
	}

	if i.InfrastructureID != 0 {
		result, err = client.InfrastructureGet(i.InfrastructureID)
	} else {
		result, err = client.InfrastructureGetByLabel(i.InfrastructureLabel)
	}

	if err != nil {
		_, err = client.InfrastructureCreate(i)

		if err != nil {
			return err
		}
	} else {
		i.instanceToOperation(&result.InfrastructureOperation)
		_, err = client.InfrastructureEdit(result.InfrastructureID, i.InfrastructureOperation)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (i Infrastructure) Delete(client MetalCloudClient) error {
	var result *Infrastructure
	var id int
	err := i.Validate()

	if err != nil {
		return err
	}

	if i.InfrastructureLabel != "" {
		result, err = client.InfrastructureGetByLabel(i.InfrastructureLabel)
		if err != nil {
			return err
		}
		id = result.InfrastructureID
	} else {
		id = i.InfrastructureID
	}
	err = client.InfrastructureDelete(id)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (i Infrastructure) Validate() error {
	if i.InfrastructureID == 0 && i.InfrastructureLabel == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}

// InfrastructureSearch searches for infrastructures with filtering support
func (c *Client) InfrastructureSearch(filter string) (*[]InfrastructuresSearchResult, error) {

	tables := []string{"_user_infrastructures_extended"}
	columns := map[string][]string{

		"_user_infrastructures_extended": {
			"infrastructure_id",
			"infrastructure_subdomain",
			"infrastructure_label",
			"infrastructure_service_status",
			"infrastructure_deploy_status",
			"datacenter_name",
			"infrastructure_created_timestamp",
			"infrastructure_updated_timestamp",
			"user_email",
			"user_id_owner",
			"infrastructure_deploy_id",
			"afc_group_created_timestamp",
			"afc_group_finished_timestamp",
			"thrownError",
			"executedSuccess",
			"total",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	sortBy := [][]string{
		{
			"thrownError",
			"DESC",
		},
		{
			"infrastructure_service_status",
			"ASC",
		},
		{
			"infrastructure_deploy_status",
			"DESC",
		},
		{
			"infrastructure_updated_timestamp",
			"DESC",
		},
	}

	var createdObject map[string]searchResultWrapperForInfrastructures

	resp, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType,
		sortBy,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]searchResultWrapperForInfrastructures{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []InfrastructuresSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil
}
