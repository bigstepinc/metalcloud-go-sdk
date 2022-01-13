package metalcloud

import (
	"fmt"
)

type SwitchDeviceLink struct {
	NetworkEquipmentLinkID   int    `json:"network_equipment_link_id,omitempty" yaml:"id,omitempty"`
	NetworkEquipmentID1      int    `json:"network_equipment_id_1,omitempty" yaml:"switchID1,omitempty"`
	NetworkEquipmentID2      int    `json:"network_equipment_id_2,omitempty" yaml:"switchID2,omitempty"`
	NetworkEquipmentLinkType string `json:"network_equipment_link_type,omitempty" yaml:"type,omitempty"`
}

//SwitchDeviceLinks Returns all the switch device links found in the database.
func (c *Client) SwitchDeviceLinks() (*map[int]SwitchDeviceLink, error) {

	var createdObject map[int]SwitchDeviceLink

	resp, err := c.rpcClient.Call(
		"switch_device_links",
	)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("No response from call switch_device_links")
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[int]SwitchDeviceLink{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//SwitchDeviceLinkCreate Creates a record for a new SwitchDevice.
func (c *Client) SwitchDeviceLinkCreate(networkEquipmentID1 int, networkEquipmentID2 int, networkEquipmentLinkType string) (*SwitchDeviceLink, error) {
	var createdObject SwitchDeviceLink
	emptyObj := map[string]string{}
	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_link_create",
		networkEquipmentID1,
		networkEquipmentID2,
		networkEquipmentLinkType,
		emptyObj,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//SwitchDeviceLinkGet Retrieves information regarding a specified switch device link
func (c *Client) SwitchDeviceLinkGet(networkEquipmentID1 int, networkEquipmentID2 int, linkType string) (*SwitchDeviceLink, error) {

	var createdObject SwitchDeviceLink

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_link_get",
		networkEquipmentID1, networkEquipmentID2, linkType)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//SwitchDeviceLinkDelete deletes a specified switch device and its registered interfaces.
func (c *Client) SwitchDeviceLinkDelete(networkEquipmentID1 int, networkEquipmentID2 int, linkType string) error {

	resp, err := c.rpcClient.Call("switch_device_link_delete", networkEquipmentID1, networkEquipmentID2, linkType)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
