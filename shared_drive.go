package metalcloud

import "fmt"

//go:generate go run helper/gen_exports.go

//SharedDrive represents a drive that can be shared between instances
type SharedDrive struct {
	SharedDriveLabel                  string                 `json:"shared_drive_label,omitempty" yaml:"label,omitempty"`
	SharedDriveSubdomain              string                 `json:"shared_drive_subdomain,omitempty" yaml:"subdomain,omitempty"`
	SharedDriveID                     int                    `json:"shared_drive_id,omitempty" yaml:"id,omitempty"`
	SharedDriveStorageType            string                 `json:"shared_drive_storage_type,omitempty" yaml:"storageType,omitempty"`
	SharedDriveHasGFS                 bool                   `json:"shared_drive_has_gfs,omitempty" yaml:"hasGFS,omitempty"`
	InfrastructureID                  int                    `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	SharedDriveServiceStatus          string                 `json:"shared_drive_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	SharedDriveCreatedTimestamp       string                 `json:"shared_drive_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	SharedDriveUpdatedTimestamp       string                 `json:"shared_drive_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	SharedDriveSizeMbytes             int                    `json:"shared_drive_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	SharedDriveAttachedInstanceArrays []int                  `json:"shared_drive_attached_instance_arrays,omitempty" yaml:"attachedInstaceArrays,omitempty"`
	SharedDriveOperation              SharedDriveOperation   `json:"shared_drive_operation,omitempty" yaml:"operation,omitempty"`
	SharedDriveCredentials            SharedDriveCredentials `json:"shared_drive_credentials,omitempty" yaml:"credentials,omitempty"`
	SharedDriveChangeID               int                    `json:"shared_drive_change_id,omitempty" yaml:"changeID,omitempty"`
	SharedDriveTargetsJSON            string                 `json:"shared_drive_targets_json,omitempty" yaml:"targetsJSON,omitempty"`
	SharedDriveIOLimitPolicy          string                 `json:"shared_drive_io_limit_policy,omitempty" yaml:"ioLimit,omitempty"`
	SharedDriveWWN                    string                 `json:"shared_drive_wwn,omitempty" yaml:"wwn,omitempty"`
	StoragePoolID                     int                    `json:"storage_pool_id,omitempty" yaml:"storagePoolID,omitempty"`
}

//SharedDriveCredentials iscsi or other forms of connection details
type SharedDriveCredentials struct {
	ISCSI ISCSI `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`
}

//SharedDriveOperation represents an ongoing or new operation on a shared drive
type SharedDriveOperation struct {
	SharedDriveDeployStatus           string `json:"shared_drive_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	SharedDriveDeployType             string `json:"shared_drive_deploy_type,omitempty" yaml:"deployType,omitempty"`
	SharedDriveLabel                  string `json:"shared_drive_label,omitempty" yaml:"label,omitempty"`
	SharedDriveSubdomain              string `json:"shared_drive_subdomain,omitempty" yaml:"subdomain,omitempty"`
	SharedDriveID                     int    `json:"shared_drive_id,omitempty" yaml:"id,omitempty"`
	SharedDriveSizeMbytes             int    `json:"shared_drive_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	SharedDriveStorageType            string `json:"shared_drive_storage_type,omitempty" yaml:"storageType,omitempty"`
	SharedDriveHasGFS                 bool   `json:"shared_drive_has_gfs" yaml:"hasGFS"`
	InfrastructureID                  int    `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	SharedDriveServiceStatus          string `json:"shared_drive_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	SharedDriveAttachedInstanceArrays []int  `json:"shared_drive_attached_instance_arrays,omitempty" yaml:"attachedInstanceArrays,omitempty"`
	SharedDriveChangeID               int    `json:"shared_drive_change_id,omitempty" yaml:"changeID,omitempty"`
}

//sharedDriveCreate creates a shared drive array. Requires deploy.
func (c *Client) sharedDriveCreate(infrastructureID id, sharedDrive SharedDrive) (*SharedDrive, error) {
	var createdObject SharedDrive

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"shared_drive_create",
		infrastructureID,
		sharedDrive)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//sharedDriveGet Retrieves a shared drive
func (c *Client) sharedDriveGet(sharedDriveID id) (*SharedDrive, error) {

	var createdObject SharedDrive

	if err := checkID(sharedDriveID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"shared_drive_get",
		sharedDriveID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//sharedDriveEdit alters a deployed drive array. Requires deploy.
func (c *Client) sharedDriveEdit(sharedDriveID id, sharedDriveOperation SharedDriveOperation) (*SharedDrive, error) {
	var createdObject SharedDrive

	if err := checkID(sharedDriveID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"shared_drive_edit",
		sharedDriveID,
		sharedDriveOperation)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//sharedDriveDelete deletes a shared drive.
func (c *Client) sharedDriveDelete(sharedDriveID id) error {

	if err := checkID(sharedDriveID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"shared_drive_delete",
		sharedDriveID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (c *Client) SharedDriveAttachInstanceArray(sharedDriveID int, instanceArrayID int) (*SharedDrive, error) {
	var updatedObject SharedDrive

	if err := checkID(sharedDriveID); err != nil {
		return nil, err
	}

	if err := checkID(instanceArrayID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&updatedObject,
		"shared_drive_attach_instance_array",
		sharedDriveID,
		instanceArrayID,
	)
	if err != nil {
		return nil, err
	}

	return &updatedObject, nil
}

func (c *Client) SharedDriveDetachInstanceArray(sharedDriveID int, instanceArrayID int) (*SharedDrive, error) {
	var updatedObject SharedDrive

	if err := checkID(sharedDriveID); err != nil {
		return nil, err
	}

	if err := checkID(instanceArrayID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&updatedObject,
		"shared_drive_detach_instance_array",
		sharedDriveID,
		instanceArrayID,
	)
	if err != nil {
		return nil, err
	}

	return &updatedObject, nil
}

//SharedDrives retrieves the list of shared drives of an infrastructure
func (c *Client) SharedDrives(infrastructureID int) (*map[string]SharedDrive, error) {
	return c.sharedDrives(infrastructureID)
}

func (c *Client) sharedDrives(infrastructureID id) (*map[string]SharedDrive, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}
	var createdObject map[string]SharedDrive

	resp, err := c.rpcClient.Call(
		"shared_drives",
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
		var m = map[string]SharedDrive{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (sd *SharedDrive) instanceToOperation(op *SharedDriveOperation) {
	operation := &sd.SharedDriveOperation
	operation.SharedDriveID = sd.SharedDriveID
	operation.SharedDriveLabel = sd.SharedDriveLabel
	operation.SharedDriveSubdomain = sd.SharedDriveSubdomain
	operation.SharedDriveSizeMbytes = sd.SharedDriveSizeMbytes
	operation.SharedDriveStorageType = sd.SharedDriveStorageType
	operation.SharedDriveHasGFS = sd.SharedDriveHasGFS
	operation.SharedDriveAttachedInstanceArrays = sd.SharedDriveAttachedInstanceArrays
	operation.SharedDriveChangeID = op.SharedDriveChangeID
}

//CreateOrUpdate implements interface Applier
func (sd SharedDrive) CreateOrUpdate(client MetalCloudClient) error {
	var result *SharedDrive
	var err error
	err = sd.Validate()

	if err != nil {
		return err
	}

	if sd.SharedDriveID != 0 {
		result, err = client.SharedDriveGet(sd.SharedDriveID)
	} else {
		result, err = client.SharedDriveGetByLabel(sd.SharedDriveLabel)
	}

	if err != nil {
		_, err = client.SharedDriveCreate(sd.InfrastructureID, sd)

		if err != nil {
			return err
		}
	} else {
		sd.instanceToOperation(&result.SharedDriveOperation)
		// return fmt.Errorf("value is obj %+v", sd)
		_, err = client.SharedDriveEdit(result.SharedDriveID, sd.SharedDriveOperation)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (sd SharedDrive) Delete(client MetalCloudClient) error {
	var result *SharedDrive
	var id int
	err := sd.Validate()

	if err != nil {
		return err
	}

	if sd.SharedDriveLabel != "" {
		result, err = client.SharedDriveGetByLabel(sd.SharedDriveLabel)
		if err != nil {
			return err
		}
		id = result.SharedDriveID
	} else {
		id = sd.SharedDriveID
	}
	err = client.SharedDriveDelete(id)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (sd SharedDrive) Validate() error {
	if sd.SharedDriveID == 0 && sd.SharedDriveLabel == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
