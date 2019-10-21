package metalcloud

import "log"

type DriveArray struct {
	DriveArrayID float64 `json:"drive_array_id,omitempty"`
	DriveArrayLabel  string  `json:"drive_array_label,omitempty"`
	VolumeTemplateID  float64  `json:"volume_template_id,omitempty"`
	DriveArrayStorageType string  `json:"drive_array_storage_type,omitempty"`
	DriveSizeMBytesDefault float64  `json:"drive_size_mbytes_default,omitempty"`
	InstanceArrayID float64  `json:"instance_array_id,omitempty"`	
}

func (c *MetalCloudClient) DriveArrays(instanceArrayID float64) (*map[string]DriveArray, error) {
	var created_object map[string]DriveArray

	err := c.rpcClient.CallFor(
		&created_object,
		"drive_arrays",
		instanceArrayID)

	if err != nil {
		return nil, err
	}

	return &created_object, nil
}

func (c *MetalCloudClient) DriveArrayGet(driveArrayID float64) (*DriveArray, error) {
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

func (c *MetalCloudClient) DriveArrayCreate(infrastructureID float64, driveArray DriveArray) (*DriveArray, error) {
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