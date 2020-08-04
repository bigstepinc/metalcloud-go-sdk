package metalcloud

import "fmt"

//SwitchDevice Represents a switch installed in a datacenter.
type SwitchDevice struct {
	DatacenterName                                 string   `json:"datacenter_name,omitempty"`
	NetworkEquipmentProvisionerType                string   `json:"network_equipment_provisioner_type,omitempty"`
	NetworkEquipmentProvisionerPosition            string   `json:"network_equipment_position,omitempty"`
	NetworkEquipmentDriver                         string   `json:"network_equipment_driver,omitempty"`
	NetworkEquipmentManagementUsername             string   `json:"network_equipment_management_username,omitempty"`
	NetworkEquipmentManagementPassword             string   `json:"network_equipment_management_password,omitempty"`
	NetworkEquipmentManagementAddress              string   `json:"network_equipment_management_address,omitempty"`
	NetworkEquipmentManagementPort                 int      `json:"network_equipment_management_port,omitempty"`
	NetworkEquipmentManagementProtocol             string   `json:"network_equipment_management_protocol,omitempty"`
	NetworkEquipmentManagementRequiresOSInstall    int      `json:"network_equipment_requires_os_install,omitempty"`
	VolumeTemplateID                               int      `json:"volume_template_id,omitempty"`
	NetworkEquipmentID                             int      `json:"network_equipment_id,omitempty"`
	NetworkEquipmentIdentifierString               string   `json:"network_equipment_identifier_string,omitempty"`
	NetworkEquipmentPrimaryWANIPv4SubnetPool       string   `json:"network_equipment_primary_wan_ipv4_subnet_pool,omitempty"`
	NetworkEquipmentPrimaryWANIPv4SubnetPrefixSize int      `json:"network_equipment_primary_wan_ipv4_subnet_prefix_size,omitempty"`
	NetworkEquipmentPrimaryWANIPv6SubnetPool       string   `json:"network_equipment_primary_wan_ipv6_subnet_pool,omitempty"`
	NetworkEquipmentPrimaryWANIPv6SubnetPrefixSize int      `json:"network_equipment_primary_wan_ipv6_subnet_prefix_size,omitempty"`
	NetworkEquipmentPrimarySANSubnetPool           string   `json:"network_equipment_primary_san_subnet_pool,omitempty"`
	NetworkEquipmentPrimarySANSubnetPrefixSize     int      `json:"network_equipment_primary_san_subnet_prefix_size,omitempty"`
	NetworkEquipmentQuarantineSubnetStart          string   `json:"network_equipment_quarantine_subnet_start,omitempty"`
	NetworkEquipmentQuarantineSubnetEnd            string   `json:"network_equipment_quarantine_subnet_end,omitempty"`
	NetworkEquipmentQuarantineSubnetPrefixSize     int      `json:"network_equipment_quarantine_subnet_prefix_size,omitempty"`
	NetworkEquipmentQuarantineSubnetGateway        string   `json:"network_equipment_quarantine_subnet_gateway,omitempty"`
	NetworkEquipmentManagementAddressMask          string   `json:"network_equipment_management_address_mask,omitempty"`
	NetworkEquipmentManagementAddressGateway       string   `json:"network_equipment_management_address_gateway,omitempty"`
	NetworkEquipmentManagementMACAddress           string   `json:"network_equipment_management_mac_address,omitempty"`
	NetworkEquipmentDescription                    string   `json:"network_equipment_description,omitempty"`
	NetworkEquipmentCountry                        string   `json:"network_equipment_country,omitempty"`
	NetworkEquipmentCity                           string   `json:"network_equipment_city,omitempty"`
	NetworkEquipmentDatacenter                     string   `json:"network_equipment_datacenter,omitempty"`
	NetworkEquipmentDatacenterRoom                 string   `json:"network_equipment_datacenter_room,omitempty"`
	NetworkEquipmentDatacenterRack                 string   `json:"network_equipment_datacenter_rack,omitempty"`
	NetworkEquipmentRackPositionUpperUnit          int      `json:"network_equipment_rack_position_upper_unit,omitempty"`
	NetworkEquipmentRackPositionLowerUnit          int      `json:"network_equipment_rack_position_lower_unit,omitempty"`
	NetworkEquipmentSerialNumber                   string   `json:"network_equipment_serial_number,omitempty"`
	ChassisRackID                                  int      `json:"chassis_rack_id,omitempty"`
	NetworkEquipmentTORLinkedID                    int      `json:"network_equipment_tor_linked_id,omitempty"`
	NetworkEquipmentTags                           []string `json:"network_equipment_tags,omitempty"`
}

//SwitchDeviceGet Retrieves information regarding a specified SwitchDevice.
func (c *Client) SwitchDeviceGet(networkEquipmentID int) (*SwitchDevice, error) {

	var createdObject SwitchDevice

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_get",
		networkEquipmentID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//SwitchDeviceGetByIdentifierString Retrieves information regarding a specified SwitchDevice by identifier string.
func (c *Client) SwitchDeviceGetByIdentifierString(networkEquipmentIdentifierString string) (*SwitchDevice, error) {

	var createdObject SwitchDevice

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_get",
		networkEquipmentIdentifierString)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//SwitchDeviceCreate Creates a record for a new SwitchDevice.
func (c *Client) SwitchDeviceCreate(switchDevice SwitchDevice, bOverwriteWithHostnameFromFetchedSwitch bool) (*SwitchDevice, error) {
	var createdObject SwitchDevice

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_create",
		switchDevice,
		bOverwriteWithHostnameFromFetchedSwitch)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//SwitchDeviceDelete deletes a specified switch device and its registered interfaces.
func (c *Client) SwitchDeviceDelete(networkEquipmentID int) error {

	resp, err := c.rpcClient.Call("switch_device_delete", networkEquipmentID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//SwitchDevices retrieves all switch devices registered in the database.
func (c *Client) SwitchDevices(datacenter string, switchType string) (*map[string]SwitchDevice, error) {

	var dc *string
	if datacenter != "" {
		dc = &datacenter
	}
	var st *string
	if switchType != "" {
		st = &switchType
	}

	res, err := c.rpcClient.Call(
		"switch_devices",
		dc,
		st)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]SwitchDevice{}
		return &m, nil
	}

	var createdObject map[string]SwitchDevice

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//SwitchDevicesInDatacenter retrieves all switch devices in a datacenter
func (c *Client) SwitchDevicesInDatacenter(datacenter string) (*map[string]SwitchDevice, error) {
	return c.SwitchDevices(datacenter, "")
}

//SwitchDeviceUpdate updates an existing switch configuration
func (c *Client) SwitchDeviceUpdate(networkEquipmentID int, switchDevice SwitchDevice, bOverwriteWithHostnameFromFetchedSwitch bool) (*SwitchDevice, error) {
	var createdObject SwitchDevice

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_update",
		networkEquipmentID,
		switchDevice,
		bOverwriteWithHostnameFromFetchedSwitch)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}
