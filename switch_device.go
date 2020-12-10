package metalcloud

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

//SwitchDevice Represents a switch installed in a datacenter.
type SwitchDevice struct {
	NetworkEquipmentID                             int      `json:"network_equipment_id,omitempty" yaml:"id,omitempty"`
	NetworkEquipmentIdentifierString               string   `json:"network_equipment_identifier_string,omitempty" yaml:"identifierString,omitempty"`
	DatacenterName                                 string   `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	NetworkEquipmentProvisionerType                string   `json:"network_equipment_provisioner_type,omitempty" yaml:"provisionerType,omitempty"`
	NetworkEquipmentProvisionerPosition            string   `json:"network_equipment_position,omitempty" yaml:"provisionerPosition,omitempty"`
	NetworkEquipmentDriver                         string   `json:"network_equipment_driver,omitempty" yaml:"driver,omitempty"`
	NetworkEquipmentManagementUsername             string   `json:"network_equipment_management_username,omitempty" yaml:"managementUsername,omitempty"`
	NetworkEquipmentManagementPassword             string   `json:"network_equipment_management_password,omitempty" yaml:"managementPassword,omitempty"`
	NetworkEquipmentManagementAddress              string   `json:"network_equipment_management_address,omitempty" yaml:"managementAddress,omitempty"`
	NetworkEquipmentManagementPort                 int      `json:"network_equipment_management_port,omitempty" yaml:"managementPort,omitempty"`
	NetworkEquipmentManagementProtocol             string   `json:"network_equipment_management_protocol,omitempty" yaml:"managementProtocol,omitempty"`
	NetworkEquipmentManagementAddressMask          string   `json:"network_equipment_management_address_mask,omitempty" yaml:"managementAddressMask,omitempty"`
	NetworkEquipmentManagementAddressGateway       string   `json:"network_equipment_management_address_gateway,omitempty" yaml:"managementAddressGateway,omitempty"`
	NetworkEquipmentManagementMACAddress           string   `json:"network_equipment_management_mac_address,omitempty" yaml:"managementMACAddress,omitempty"`
	NetworkEquipmentPrimaryWANIPv4SubnetPool       string   `json:"network_equipment_primary_wan_ipv4_subnet_pool,omitempty" yaml:"primaryWANIPv4SubnetPool,omitempty"`
	NetworkEquipmentPrimaryWANIPv4SubnetPrefixSize int      `json:"network_equipment_primary_wan_ipv4_subnet_prefix_size,omitempty" yaml:"primaryWANIPv4SubnetPrefixSize,omitempty"`
	NetworkEquipmentPrimaryWANIPv6SubnetPool       string   `json:"network_equipment_primary_wan_ipv6_subnet_pool,omitempty" yaml:"primaryWANIPv6SubnetPool,omitempty"`
	NetworkEquipmentPrimaryWANIPv6SubnetPoolID     int      `json:"network_equipment_primary_wan_ipv6_subnet_pool_id,omitempty" yaml:"primaryWANIPv6SubnetPoolID,omitempty"`
	NetworkEquipmentPrimaryWANIPv6SubnetCIDR       string   `json:"network_equipment_primary_wan_ipv6_subnet_cidr,omitempty" yaml:"primaryWANIPv6SubnetCIDR,omitempty"`
	NetworkEquipmentPrimaryWANIPv6SubnetPrefixSize int      `json:"network_equipment_primary_wan_ipv6_subnet_prefix_size,omitempty" yaml:"primaryWANIPv6SubnetPrefixSize,omitempty"`
	NetworkEquipmentPrimarySANSubnetPool           string   `json:"network_equipment_primary_san_subnet_pool,omitempty" yaml:"primarySANSubnetPool,omitempty"`
	NetworkEquipmentPrimarySANSubnetPrefixSize     int      `json:"network_equipment_primary_san_subnet_prefix_size,omitempty" yaml:"primarySANSubnetPrefixSize,omitempty"`
	NetworkEquipmentQuarantineSubnetStart          string   `json:"network_equipment_quarantine_subnet_start,omitempty" yaml:"quarantineSubnetStart,omitempty"`
	NetworkEquipmentQuarantineSubnetEnd            string   `json:"network_equipment_quarantine_subnet_end,omitempty" yaml:"quarantineSubnetEnd,omitempty"`
	NetworkEquipmentQuarantineSubnetPrefixSize     int      `json:"network_equipment_quarantine_subnet_prefix_size,omitempty" yaml:"quarantineSubnetPrefixSize,omitempty"`
	NetworkEquipmentQuarantineSubnetGateway        string   `json:"network_equipment_quarantine_subnet_gateway,omitempty" yaml:"quarantineSubnetGateway,omitempty"`
	NetworkEquipmentDescription                    string   `json:"network_equipment_description,omitempty" yaml:"description,omitempty"`
	NetworkEquipmentCountry                        string   `json:"network_equipment_country,omitempty" yaml:"country,omitempty"`
	NetworkEquipmentCity                           string   `json:"network_equipment_city,omitempty" yaml:"city,omitempty"`
	NetworkEquipmentDatacenter                     string   `json:"network_equipment_datacenter,omitempty" yaml:"netDatacenter,omitempty"`
	NetworkEquipmentDatacenterRoom                 string   `json:"network_equipment_datacenter_room,omitempty" yaml:"datacenterRoom,omitempty"`
	NetworkEquipmentDatacenterRack                 string   `json:"network_equipment_datacenter_rack,omitempty" yaml:"datacenterRack,omitempty"`
	NetworkEquipmentRackPositionUpperUnit          int      `json:"network_equipment_rack_position_upper_unit,omitempty" yaml:"rackPositionUpperUnit,omitempty"`
	NetworkEquipmentRackPositionLowerUnit          int      `json:"network_equipment_rack_position_lower_unit,omitempty" yaml:"rackPositionLowerUnit,omitempty"`
	NetworkEquipmentSerialNumber                   string   `json:"network_equipment_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	ChassisRackID                                  int      `json:"chassis_rack_id,omitempty" yaml:"chassisRackID,omitempty"`
	NetworkEquipmentTORLinkedID                    int      `json:"network_equipment_tor_linked_id,omitempty"  yaml:"TORLinkedID,omitempty"`
	NetworkEquipmentTags                           []string `json:"network_equipment_tags,omitempty" yaml:"tags,omitempty"`
	NetworkEquipmentRequiresOSInstall              bool     `json:"network_equipment_requires_os_install" yaml:"requiresOSInstall"`
	VolumeTemplateID                               int      `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
}

//UnmarshalJSON to handle the shitty boolean being returned as 0 and 1 and true and false in different environments
func (s *SwitchDevice) UnmarshalJSON(data []byte) error {

	//SwitchDevice Represents a switch installed in a datacenter.
	var v struct {
		NetworkEquipmentID                             int      `json:"network_equipment_id,omitempty" yaml:"id,omitempty"`
		NetworkEquipmentIdentifierString               string   `json:"network_equipment_identifier_string,omitempty" yaml:"identifierString,omitempty"`
		DatacenterName                                 string   `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
		NetworkEquipmentProvisionerType                string   `json:"network_equipment_provisioner_type,omitempty" yaml:"provisionerType,omitempty"`
		NetworkEquipmentProvisionerPosition            string   `json:"network_equipment_position,omitempty" yaml:"provisionerPosition,omitempty"`
		NetworkEquipmentDriver                         string   `json:"network_equipment_driver,omitempty" yaml:"driver,omitempty"`
		NetworkEquipmentManagementUsername             string   `json:"network_equipment_management_username,omitempty" yaml:"managementUsername,omitempty"`
		NetworkEquipmentManagementPassword             string   `json:"network_equipment_management_password,omitempty" yaml:"managementPassword,omitempty"`
		NetworkEquipmentManagementAddress              string   `json:"network_equipment_management_address,omitempty" yaml:"managementAddress,omitempty"`
		NetworkEquipmentManagementPort                 int      `json:"network_equipment_management_port,omitempty" yaml:"managementPort,omitempty"`
		NetworkEquipmentManagementProtocol             string   `json:"network_equipment_management_protocol,omitempty" yaml:"managementProtocol,omitempty"`
		NetworkEquipmentManagementAddressMask          string   `json:"network_equipment_management_address_mask,omitempty" yaml:"managementAddressMask,omitempty"`
		NetworkEquipmentManagementAddressGateway       string   `json:"network_equipment_management_address_gateway,omitempty" yaml:"managementAddressGateway,omitempty"`
		NetworkEquipmentManagementMACAddress           string   `json:"network_equipment_management_mac_address,omitempty" yaml:"managementMACAddress,omitempty"`
		NetworkEquipmentPrimaryWANIPv4SubnetPool       string   `json:"network_equipment_primary_wan_ipv4_subnet_pool,omitempty" yaml:"primaryWANIPv4SubnetPool,omitempty"`
		NetworkEquipmentPrimaryWANIPv4SubnetPrefixSize int      `json:"network_equipment_primary_wan_ipv4_subnet_prefix_size,omitempty" yaml:"primaryWANIPv4SubnetPrefixSize,omitempty"`
		NetworkEquipmentPrimaryWANIPv6SubnetPool       string   `json:"network_equipment_primary_wan_ipv6_subnet_pool,omitempty" yaml:"primaryWANIPv6SubnetPool,omitempty"`
		NetworkEquipmentPrimaryWANIPv6SubnetPoolID     int      `json:"network_equipment_primary_wan_ipv6_subnet_pool_id,omitempty" yaml:"primaryWANIPv6SubnetPoolID,omitempty"`
		NetworkEquipmentPrimaryWANIPv6SubnetCIDR       string   `json:"network_equipment_primary_wan_ipv6_subnet_cidr,omitempty" yaml:"primaryWANIPv6SubnetCIDR,omitempty"`
		NetworkEquipmentPrimaryWANIPv6SubnetPrefixSize int      `json:"network_equipment_primary_wan_ipv6_subnet_prefix_size,omitempty" yaml:"primaryWANIPv6SubnetPrefixSize,omitempty"`
		NetworkEquipmentPrimarySANSubnetPool           string   `json:"network_equipment_primary_san_subnet_pool,omitempty" yaml:"primarySANSubnetPool,omitempty"`
		NetworkEquipmentPrimarySANSubnetPrefixSize     int      `json:"network_equipment_primary_san_subnet_prefix_size,omitempty" yaml:"primarySANSubnetPrefixSize,omitempty"`
		NetworkEquipmentQuarantineSubnetStart          string   `json:"network_equipment_quarantine_subnet_start,omitempty" yaml:"quarantineSubnetStart,omitempty"`
		NetworkEquipmentQuarantineSubnetEnd            string   `json:"network_equipment_quarantine_subnet_end,omitempty" yaml:"quarantineSubnetEnd,omitempty"`
		NetworkEquipmentQuarantineSubnetPrefixSize     int      `json:"network_equipment_quarantine_subnet_prefix_size,omitempty" yaml:"quarantineSubnetPrefixSize,omitempty"`
		NetworkEquipmentQuarantineSubnetGateway        string   `json:"network_equipment_quarantine_subnet_gateway,omitempty" yaml:"quarantineSubnetGateway,omitempty"`
		NetworkEquipmentDescription                    string   `json:"network_equipment_description,omitempty" yaml:"description,omitempty"`
		NetworkEquipmentCountry                        string   `json:"network_equipment_country,omitempty" yaml:"country,omitempty"`
		NetworkEquipmentCity                           string   `json:"network_equipment_city,omitempty" yaml:"city,omitempty"`
		NetworkEquipmentDatacenter                     string   `json:"network_equipment_datacenter,omitempty" yaml:"datacenter,omitempty"`
		NetworkEquipmentDatacenterRoom                 string   `json:"network_equipment_datacenter_room,omitempty" yaml:"datacenterRoom,omitempty"`
		NetworkEquipmentDatacenterRack                 string   `json:"network_equipment_datacenter_rack,omitempty" yaml:"datacenterRack,omitempty"`
		NetworkEquipmentRackPositionUpperUnit          int      `json:"network_equipment_rack_position_upper_unit,omitempty" yaml:"rackPositionUpperUnit,omitempty"`
		NetworkEquipmentRackPositionLowerUnit          int      `json:"network_equipment_rack_position_lower_unit,omitempty" yaml:"rackPositionLowerUnit,omitempty"`
		NetworkEquipmentSerialNumber                   string   `json:"network_equipment_serial_number,omitempty" yaml:"serialNumber,omitempty"`
		ChassisRackID                                  int      `json:"chassis_rack_id,omitempty" yaml:"chassisRackID,omitempty"`
		NetworkEquipmentTORLinkedID                    int      `json:"network_equipment_tor_linked_id,omitempty"  yaml:"TORLinkedID,omitempty"`
		NetworkEquipmentTags                           []string `json:"network_equipment_tags,omitempty" yaml:"tags,omitempty"`
		//this is the culprit.
		NetworkEquipmentRequiresOSInstall interface{} `json:"network_equipment_requires_os_install" yaml:"requiresOSInstall"`
		VolumeTemplateID                  int         `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	}

	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.NetworkEquipmentRequiresOSInstall.(type) {
	case float64:
		if v.NetworkEquipmentRequiresOSInstall.(float64) == 0.0 {
			v.NetworkEquipmentRequiresOSInstall = false
		} else {
			v.NetworkEquipmentRequiresOSInstall = true
		}
	}
	copier.Copy(&s, &v)

	return nil
}

//SwitchDeviceGet Retrieves information regarding a specified SwitchDevice.
func (c *Client) SwitchDeviceGet(networkEquipmentID int, decryptPasswd bool) (*SwitchDevice, error) {

	var createdObject SwitchDevice

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_get",
		networkEquipmentID)

	if err != nil {

		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.NetworkEquipmentManagementPassword, ":")
		if len(passwdComponents) != 2 {
			return nil, fmt.Errorf("Password not returned with proper components")
		}
		var passwd string
		err = c.rpcClient.CallFor(
			&passwd,
			"password_decrypt",
			passwdComponents[1],
		)
		if err != nil {
			return nil, err
		}
		createdObject.NetworkEquipmentManagementPassword = passwd
	} else {
		createdObject.NetworkEquipmentManagementPassword = ""
	}

	return &createdObject, nil
}

//SwitchDeviceGetByIdentifierString Retrieves information regarding a specified SwitchDevice by identifier string.
func (c *Client) SwitchDeviceGetByIdentifierString(networkEquipmentIdentifierString string, decryptPasswd bool) (*SwitchDevice, error) {

	var createdObject SwitchDevice

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_get",
		networkEquipmentIdentifierString)

	if err != nil {

		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.NetworkEquipmentManagementPassword, ":")
		if len(passwdComponents) != 2 {
			return nil, fmt.Errorf("Password not returned with proper components")
		}
		var passwd string
		err = c.rpcClient.CallFor(
			&passwd,
			"password_decrypt",
			passwdComponents[1],
		)
		if err != nil {
			return nil, err
		}
		createdObject.NetworkEquipmentManagementPassword = passwd
	} else {
		createdObject.NetworkEquipmentManagementPassword = ""
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

//CreateOrUpdate implements interface Applier
func (s SwitchDevice) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var switchDevice *SwitchDevice
	err = s.Validate()

	if err != nil {
		return err
	}
	if s.NetworkEquipmentIdentifierString != "" {
		switchDevice, err = client.SwitchDeviceGetByIdentifierString(s.NetworkEquipmentIdentifierString, false)
	} else {
		switchDevice, err = client.SwitchDeviceGet(s.NetworkEquipmentID, false)
	}

	if err != nil {
		_, err := client.SwitchDeviceCreate(s, false)

		if err != nil {
			return err
		}
	} else {
		_, err := client.SwitchDeviceUpdate(switchDevice.NetworkEquipmentID, s, false)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (s SwitchDevice) Delete(client MetalCloudClient) error {
	err := s.Validate()

	if err != nil {
		return err
	}
	err = client.SwitchDeviceDelete(s.NetworkEquipmentID)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (s SwitchDevice) Validate() error {
	if s.NetworkEquipmentID != 0 && s.NetworkEquipmentIdentifierString != "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
