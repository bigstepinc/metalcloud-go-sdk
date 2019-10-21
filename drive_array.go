package metalcloud

import "log"

type DriveArray struct {
	DriveArrayID int64 `json:"drive_array_id,omitempty"`
	DriveArrayLabel  string  `json:"drive_array_label,omitempty"`
	VolumeTemplateID  int64  `json:"volume_template_id,omitempty"`
	DriveArrayStorageType string  `json:"drive_array_storage_type,omitempty"`
	DriveSizeMBytesDefault int64  `json:"drive_size_mbytes_default,omitempty"`
	InstanceArrayID int64  `json:"instance_array_id,omitempty"`	
}

func (c *MetalCloudClient) DriveArrays(infrastructureID int64) (*map[string]DriveArray, error) {
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

	var created_object map[string]DriveArray
	
	err2 := res.GetObject(&created_object)
	if err2 != nil {
			return nil, err2
	}

	return &created_object, nil
}

func (c *MetalCloudClient) DriveArrayGet(driveArrayID int64) (*DriveArray, error) {
	var created_object DriveArray

	err := c.rpcClient.CallFor(
		&created_object,
		"drive_array_get",
		driveArrayID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &created_object, nil
}

func (c *MetalCloudClient) DriveArrayCreate(infrastructureID int64, driveArray DriveArray) (*DriveArray, error) {
	var created_object DriveArray

	err := c.rpcClient.CallFor(
		&created_object,
		"drive_array_create",
		infrastructureID,
		driveArray)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &created_object, nil
}