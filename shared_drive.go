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
}

//SharedDriveCredentials iscsi or other forms of connection details
type SharedDriveCredentials struct {
	ISCSI ISCSI `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`
}

//SharedDriveOperation represents an ongoing or new operation on a shared drive
type SharedDriveOperation struct {
	SharedDriveDeployStatus           string `json:"shared_drive_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	SharedDriveDepoloyType            string `json:"shared_drive_deploy_type,omitempty" yaml:"deployType,omitempty"`
	SharedDriveLabel                  string `json:"shared_drive_label,omitempty" yaml:"label,omitempty"`
	SharedDriveSubdomain              string `json:"shared_drive_subdomain,omitempty" yaml:"subdomain,omitempty"`
	SharedDriveID                     int    `json:"shared_drive_id,omitempty" yaml:"id,omitempty"`
	SharedDriveSizeMbytes             int    `json:"shared_drive_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	SharedDriveStorageType            string `json:"shared_drive_storage_type,omitempty" yaml:"storageType,omitempty"`
	SharedDriveHasGFS                 bool   `json:"shared_drive_has_gfs,omitempty" yaml:"hasGFS,omitempty"`
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

func (c *Client) sharedDrives(infrastructureID id) (*map[string]SharedDrive, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	res, err := c.rpcClient.Call(
		"shared_drives",
		infrastructureID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]SharedDrive{}
		return &m, nil
	}

	var createdObject map[string]SharedDrive

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//CreateOrUpdate implements interface Applier
func (sd SharedDrive) CreateOrUpdate(c interface{}) error {
	client := c.(*Client)

	var result *SharedDrive
	var err error

	if sd.SharedDriveID != 0 {
		result, err = client.sharedDriveGet(sd.SharedDriveID)
	} else if sd.SharedDriveLabel != "" {
		result, err = client.sharedDriveGet(sd.SharedDriveLabel)
	} else {
		return fmt.Errorf("id is required")
	}

	if err != nil {
		_, err = client.sharedDriveCreate(sd.InfrastructureID, sd)

		if err != nil {
			return err
		}
	} else {
		sd.SharedDriveOperation.SharedDriveChangeID = result.SharedDriveOperation.SharedDriveChangeID
		sd.SharedDriveOperation.SharedDriveChangeID = result.SharedDriveOperation.SharedDriveChangeID
		_, err = client.sharedDriveEdit(sd.SharedDriveID, sd.SharedDriveOperation)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (sd SharedDrive) Delete(c interface{}) error {
	client := c.(*Client)

	err := client.sharedDriveDelete(sd.SharedDriveID)

	if err != nil {
		return err
	}

	return nil
}
