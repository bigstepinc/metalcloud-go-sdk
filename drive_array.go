package metalcloud

import "log"

//DriveArray represents a collection of identical drives
type DriveArray struct {
	DriveArrayID            int    `json:"drive_array_id,omitempty"`
	DriveArrayLabel         string `json:"drive_array_label,omitempty"`
	VolumeTemplateID        int    `json:"volume_template_id,omitempty"`
	DriveArrayStorageType   string `json:"drive_array_storage_type,omitempty"`
	DriveSizeMBytesDefault  int    `json:"drive_size_mbytes_default,omitempty"`
	InstanceArrayID         int    `json:"instance_array_id,omitempty"`
	InfrastructureID        int    `json:"infrastructure_id,omitempty"`
	DriveArrayServiceStatus string `json:"drive_array_service_status,omitempty"`

	DriveArrayOperation *DriveArrayOperation `json:"drive_array_operation,omitempty"`
}

//DriveArrayOperation defines changes to be applied to a DriveArray
type DriveArrayOperation struct {
	DriveArrayID           int    `json:"drive_array_id,omitempty"`
	DriveArrayLabel        string `json:"drive_array_label,omitempty"`
	VolumeTemplateID       int    `json:"volume_template_id,omitempty"`
	DriveArrayStorageType  string `json:"drive_array_storage_type,omitempty"`
	DriveSizeMBytesDefault int    `json:"drive_size_mbytes_default,omitempty"`
	InstanceArrayID        int    `json:"instance_array_id,omitempty"`
	InfrastructureID       int    `json:"infrastructure_id,omitempty"`

	DriveArrayDeployType string `json:"drive_array_deploy_type,omitempty"`
	DriveArrayChangeID   int    `json:"drive_array_change_id,omitempty"`
}

//DriveArrays retrieves the list of drives arrays of an infrastructure
func (c *Client) DriveArrays(infrastructureID int) (*map[string]DriveArray, error) {
	res, err := c.rpcClient.Call(
		"drive_arrays",
		infrastructureID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]DriveArray{}
		return &m, nil
	}

	var createdObject map[string]DriveArray

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//DriveArrayGet retrieves a DriveArray object with specified ids
func (c *Client) DriveArrayGet(driveArrayID int) (*DriveArray, error) {
	var createdObject DriveArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_array_get",
		driveArrayID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &createdObject, nil
}

//DriveArrayCreate creates a drive array. Requires deploy.
func (c *Client) DriveArrayCreate(infrastructureID int, driveArray DriveArray) (*DriveArray, error) {
	var createdObject DriveArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_array_create",
		infrastructureID,
		driveArray)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &createdObject, nil
}

//DriveArrayEdit alters a deployed drive array. Requires deploy.
func (c *Client) DriveArrayEdit(driveArrayID int, driveArrayOperation DriveArrayOperation) (*DriveArray, error) {
	var createdObject DriveArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_array_edit",
		driveArrayID,
		driveArrayOperation)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &createdObject, nil
}

//DriveArrayDelete deletes a Drive Array with specified id
func (c *Client) DriveArrayDelete(driveArrayID int) error {

	_, err := c.rpcClient.Call(
		"drive_array_delete",
		driveArrayID)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}
