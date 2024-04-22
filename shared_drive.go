package metalcloud

//go:generate go run helper/docgen.go - $GOFILE ./ SharedDrive,SharedDriveCredentials,SharedDriveOperation SharedDrive

import "fmt"

//go:generate go run helper/gen_exports.go

// SharedDrive represents a drive that can be shared between instances
type SharedDrive struct {
	//description: Label of the shared drive
	SharedDriveLabel string `json:"shared_drive_label,omitempty" yaml:"label,omitempty"`
	//description: Unique string representing the shared drive
	SharedDriveSubdomain string `json:"shared_drive_subdomain,omitempty" yaml:"subdomain,omitempty"`
	//description: Id of the shared drive
	SharedDriveID int `json:"shared_drive_id,omitempty" yaml:"id,omitempty"`
	//description: Type of the shared drive
	//values:
	// - iscsi_ssd
	// - iscsi_hdd
	SharedDriveStorageType string `json:"shared_drive_storage_type,omitempty" yaml:"storageType,omitempty"`
	//description: ID of the infrastructure
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	//description: Service status of the shared drive
	//values:
	// - ordered
	// - active
	// - stopped
	// - deleted
	SharedDriveServiceStatus string `json:"shared_drive_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the SharedDrive was created.
	//example: 2013-11-29T13:00:01Z
	SharedDriveCreatedTimestamp string `json:"shared_drive_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the SharedDrive was updated.
	//example: 2013-11-29T13:00:01Z
	SharedDriveUpdatedTimestamp string `json:"shared_drive_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	//description: Size of the drive
	SharedDriveSizeMbytes int `json:"shared_drive_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	//description: An array of the instance array ids attached to this drive
	SharedDriveAttachedInstanceArrays []int `json:"shared_drive_attached_instance_arrays,omitempty" yaml:"attachedInstaceArrays,omitempty"`
	//description: The operation object
	SharedDriveOperation SharedDriveOperation `json:"shared_drive_operation,omitempty" yaml:"operation,omitempty"`
	//description: Credentials of the shared drive
	SharedDriveCredentials SharedDriveCredentials `json:"shared_drive_credentials,omitempty" yaml:"credentials,omitempty"`
	//description: The operation id
	SharedDriveChangeID int `json:"shared_drive_change_id,omitempty" yaml:"changeID,omitempty"`
	//description: Details on the ISCSI or FC targets
	SharedDriveTargetsJSON string `json:"shared_drive_targets_json,omitempty" yaml:"targetsJSON,omitempty"`
	//description: Used by certain storage types
	SharedDriveIOLimitPolicy string `json:"shared_drive_io_limit_policy,omitempty" yaml:"ioLimit,omitempty"`
	//description: WWN of the drive as reported by the storage appliance
	SharedDriveWWN string `json:"shared_drive_wwn,omitempty" yaml:"wwn,omitempty"`
	//description: The storage pool id to use if set.
	StoragePoolID int `json:"storage_pool_id,omitempty" yaml:"storagePoolID,omitempty"`
	// description: Used to control if drives in this drive array should be allocated on the same storage pool or different storage pools
	// values:
	// - same_storage
	// - different_storage
	SharedDriveAllocationAffinity string `json:"shared_drive_allocation_affinity,omitempty" yaml:"affinity,omitempty"`
}

// description: SharedDriveCredentials iscsi or other forms of connection details
type SharedDriveCredentials struct {
	//description: details
	ISCSI ISCSI `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`
}

// SharedDriveOperation represents an ongoing or new operation on a shared drive
type SharedDriveOperation struct {
	//description: Deploy status
	//values:
	// - not_started
	// - ongoing
	// - finished
	SharedDriveDeployStatus string `json:"shared_drive_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	//description: Type of operation
	//values:
	// - create
	// - edit
	// - delete
	// - start
	// - stop
	SharedDriveDeployType string `json:"shared_drive_deploy_type,omitempty" yaml:"deployType,omitempty"`
	//description: Label
	SharedDriveLabel string `json:"shared_drive_label,omitempty" yaml:"label,omitempty"`
	//description: Unique string describing this shared drive
	SharedDriveSubdomain string `json:"shared_drive_subdomain,omitempty" yaml:"subdomain,omitempty"`
	//description: ID of the shared drive
	SharedDriveID int `json:"shared_drive_id,omitempty" yaml:"id,omitempty"`
	//description: Size of the drive
	SharedDriveSizeMbytes int `json:"shared_drive_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	//description: Type of the shared drive. Readonly.
	//values:
	// - iscsi_ssd
	// - iscsi_hdd
	SharedDriveStorageType string `json:"shared_drive_storage_type,omitempty" yaml:"storageType,omitempty"`
	//description: ID of the infrastructure
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	//description: Status of the service
	//values:
	// - active
	// - ordered
	// - deleted
	// - suspended
	// - stopped
	SharedDriveServiceStatus string `json:"shared_drive_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	//description: List of instance arrays to which this shared drive is attached
	SharedDriveAttachedInstanceArrays []int `json:"shared_drive_attached_instance_arrays,omitempty" yaml:"attachedInstanceArrays,omitempty"`
	//description: ID of the operation
	SharedDriveChangeID int `json:"shared_drive_change_id,omitempty" yaml:"changeID,omitempty"`
	//description: Used for certain SAN appliances
	SharedDriveIOLimitPolicy string `json:"shared_drive_io_limit_policy,omitempty" yaml:"ioLimit,omitempty"`
}

// sharedDriveCreate creates a shared drive array. Requires deploy.
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

// sharedDriveGet Retrieves a shared drive
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

// sharedDriveEdit alters a deployed drive array. Requires deploy.
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

// sharedDriveDelete deletes a shared drive.
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

// SharedDrives retrieves the list of shared drives of an infrastructure
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
	operation.SharedDriveAttachedInstanceArrays = sd.SharedDriveAttachedInstanceArrays
	operation.SharedDriveChangeID = op.SharedDriveChangeID
}

// CreateOrUpdate implements interface Applier
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

// Delete implements interface Applier
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

// Validate implements interface Applier
func (sd SharedDrive) Validate() error {
	if sd.SharedDriveID == 0 && sd.SharedDriveLabel == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
