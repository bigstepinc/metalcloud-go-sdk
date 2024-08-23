package metalcloud

//go:generate go run helper/gen_exports.go

import "fmt"

// DriveArray represents a collection of identical drives
type DriveArray struct {
	DriveArrayID                      int                  `json:"drive_array_id,omitempty" yaml:"id,omitempty"`
	DriveArrayLabel                   string               `json:"drive_array_label,omitempty" yaml:"label,omitempty"`
	VolumeTemplateID                  int                  `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	DriveArrayStorageType             string               `json:"drive_array_storage_type,omitempty" yaml:"storageType,omitempty"`
	DriveSizeMBytesDefault            int                  `json:"drive_size_mbytes_default,omitempty" yaml:"sizeMBytesDefault,omitempty"`
	InstanceArrayID                   int                  `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	InfrastructureID                  int                  `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	DriveArrayServiceStatus           string               `json:"drive_array_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	DriveArrayCount                   int                  `json:"drive_array_count,omitempty" yaml:"count,omitempty"`
	DriveArrayExpandWithInstanceArray bool                 `json:"drive_array_expand_with_instance_array" yaml:"expandWithInstanceArray"`
	DriveArrayOperation               *DriveArrayOperation `json:"drive_array_operation,omitempty" yaml:"operation,omitempty"`
	DriveArrayIOLimitPolicy           string               `json:"drive_array_io_limit_policy,omitempty" yaml:"ioLimit,omitempty"`
	StoragePoolID                     int                  `json:"storage_pool_id,omitempty" yaml:"storagePoolID,omitempty"`
	DriveArrayAllocationAffinity      string               `json:"drive_array_allocation_affinity,omitempty" yaml:"affinity,omitempty"`
}

// DriveArrayOperation defines changes to be applied to a DriveArray
type DriveArrayOperation struct {
	DriveArrayID                      int         `json:"drive_array_id,omitempty" yaml:"id,omitempty"`
	DriveArrayLabel                   string      `json:"drive_array_label,omitempty" yaml:"label,omitempty"`
	VolumeTemplateID                  int         `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	DriveArrayStorageType             string      `json:"drive_array_storage_type,omitempty" yaml:"storageType,omitempty"`
	DriveSizeMBytesDefault            int         `json:"drive_size_mbytes_default,omitempty" yaml:"sizeMBytes,omitempty"`
	InstanceArrayID                   interface{} `json:"instance_array_id" yaml:"instanceArrayID"`
	InfrastructureID                  int         `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	DriveArrayCount                   int         `json:"drive_array_count,omitempty" yaml:"count,omitempty"`
	DriveArrayExpandWithInstanceArray bool        `json:"drive_array_expand_with_instance_array" yaml:"expandWithInstanceArray"`
	DriveArrayDeployType              string      `json:"drive_array_deploy_type,omitempty" yaml:"deployType,omitempty"`
	DriveArrayDeployStatus            string      `json:"drive_array_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	DriveArrayChangeID                int         `json:"drive_array_change_id,omitempty" yaml:"changeID,omitempty"`
	DriveArrayIOLimitPolicy           string      `json:"drive_array_io_limit_policy,omitempty" yaml:"ioLimit,omitempty"`
}

// DriveArrays retrieves the list of drives arrays of an infrastructure
func (c *Client) DriveArrays(infrastructureID int) (*map[string]DriveArray, error) {
	return c.driveArrays(infrastructureID)
}

// DriveArraysByLabel retrieves the list of drives arrays of an infrastructure
func (c *Client) DriveArraysByLabel(infrastructureLabel string) (*map[string]DriveArray, error) {
	return c.driveArrays(infrastructureLabel)
}

func (c *Client) driveArrays(infrastructureID id) (*map[string]DriveArray, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"drive_arrays",
		infrastructureID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})

	if ok {
		var m = map[string]DriveArray{}
		return &m, nil
	}

	var createdObject map[string]DriveArray

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// DriveArrayGet retrieves a DriveArray object with specified ids
func (c *Client) DriveArrayGet(driveArrayID int) (*DriveArray, error) {
	return c.driveArrayGet(driveArrayID)
}

// DriveArrayGetByLabel retrieves a DriveArray object with specified ids
func (c *Client) DriveArrayGetByLabel(driveArrayLabel string) (*DriveArray, error) {
	return c.driveArrayGet(driveArrayLabel)
}

func (c *Client) driveArrayGet(driveArrayID id) (*DriveArray, error) {

	var createdObject DriveArray

	if err := checkID(driveArrayID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_array_get",
		driveArrayID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// driveArrayCreate creates a drive array. Requires deploy.
func (c *Client) driveArrayCreate(infrastructureID id, driveArray DriveArray) (*DriveArray, error) {
	var createdObject DriveArray

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_array_create",
		infrastructureID,
		driveArray)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// driveArrayEdit alters a deployed drive array. Requires deploy.
func (c *Client) driveArrayEdit(driveArrayID id, driveArrayOperation DriveArrayOperation) (*DriveArray, error) {
	var createdObject DriveArray

	if err := checkID(driveArrayID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_array_edit",
		driveArrayID,
		driveArrayOperation)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// driveArrayDelete deletes a Drive Array with specified id
func (c *Client) driveArrayDelete(driveArrayID id) error {

	if err := checkID(driveArrayID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"drive_array_delete",
		driveArrayID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// driveArrayDrives returns the drives of a drive array
func (c *Client) driveArrayDrives(driveArray id) (*map[string]Drive, error) {

	if err := checkID(driveArray); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"drive_array_drives",
		driveArray,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Drive{}
		return &m, nil
	}

	var createdObject map[string]Drive

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (da *DriveArray) instanceToOperation(op *DriveArrayOperation) {
	operation := da.DriveArrayOperation
	operation.DriveArrayID = da.DriveArrayID
	operation.DriveArrayLabel = da.DriveArrayLabel
	operation.VolumeTemplateID = da.VolumeTemplateID
	operation.DriveArrayStorageType = da.DriveArrayStorageType
	operation.DriveSizeMBytesDefault = da.DriveSizeMBytesDefault
	operation.InstanceArrayID = da.InstanceArrayID
	operation.DriveArrayCount = da.DriveArrayCount
	operation.DriveArrayExpandWithInstanceArray = da.DriveArrayExpandWithInstanceArray
	operation.DriveArrayChangeID = op.DriveArrayChangeID
}

// CreateOrUpdate implements interface Applier
func (da DriveArray) CreateOrUpdate(client MetalCloudClient) error {
	var result *DriveArray
	var err error
	err = da.Validate()

	if err != nil {
		return err
	}

	if da.DriveArrayID != 0 {
		result, err = client.DriveArrayGet(da.DriveArrayID)
	} else {
		result, err = client.DriveArrayGetByLabel(da.DriveArrayLabel)
	}
	if err != nil {
		_, err = client.DriveArrayCreate(da.InfrastructureID, da)

		if err != nil {
			return err
		}
	} else {
		da.instanceToOperation(result.DriveArrayOperation)
		_, err = client.DriveArrayEdit(result.DriveArrayID, *da.DriveArrayOperation)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (da DriveArray) Delete(client MetalCloudClient) error {
	err := da.Validate()
	var result *DriveArray
	var id int

	if err != nil {
		return err
	}

	if da.DriveArrayLabel != "" {
		result, err = client.DriveArrayGetByLabel(da.DriveArrayLabel)
		if err != nil {
			return err
		}
		id = result.DriveArrayID
	} else {
		id = da.DriveArrayID
	}

	err = client.DriveArrayDelete(id)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (da DriveArray) Validate() error {
	if da.DriveArrayID == 0 && da.DriveArrayLabel == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
