//go:generate go run helper/docgen.go - $GOFILE ./ SwitchDevice SwitchDevice
package metalcloud

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

// SwitchDevice Represents a switch installed in a datacenter.
type SwitchDevice struct {
	//description: id of the network equipment
	NetworkEquipmentID int `json:"network_equipment_id,omitempty" yaml:"id,omitempty"`
	//description: Hostname (or unique label) of the network equipment
	NetworkEquipmentIdentifierString string `json:"network_equipment_identifier_string,omitempty" yaml:"identifierString,omitempty"`
	//description: Datacenter
	DatacenterName string `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
	//description: Type of provisioner. Read only
	//values:
	// - vpls
	// - vxlan
	// - vlan
	// - lan
	// - sdn
	// - evpnvxlanl2
	NetworkEquipmentProvisionerType string `json:"network_equipment_provisioner_type,omitempty" yaml:"provisionerType,omitempty"`
	//description: Role of this equipment in the provisioner
	//values:
	// - tor
	// - north
	// - leaf
	// - spine
	// - other
	NetworkEquipmentProvisionerPosition string `json:"network_equipment_position,omitempty" yaml:"provisionerPosition,omitempty"`
	//description: Driver to use. Note that this list may change frequently.
	// values:
	// - hp5800
	// - hp5900
	// - hp5950
	// - dell_s4000
	// - dell_s4048
	// - dell_s6010
	// - os_10
	// - cumulus42
	// - nexus9000
	// - cisco_aci51
	// - junos
	// - junos18
	// - sonic_enterprise
	// - dummy
	NetworkEquipmentDriver string `json:"network_equipment_driver,omitempty" yaml:"driver,omitempty"`
	//description: Username
	NetworkEquipmentManagementUsername string `json:"network_equipment_management_username,omitempty" yaml:"managementUsername,omitempty"`
	//description: Password
	NetworkEquipmentManagementPassword string `json:"network_equipment_management_password,omitempty" yaml:"managementPassword,omitempty"`
	//description: Address
	NetworkEquipmentManagementAddress string `json:"network_equipment_management_address,omitempty" yaml:"managementAddress,omitempty"`
	//description: Port
	NetworkEquipmentManagementPort int `json:"network_equipment_management_port,omitempty" yaml:"managementPort,omitempty"`
	//description: Deprecated.
	NetworkEquipmentManagementProtocol string `json:"network_equipment_management_protocol,omitempty" yaml:"managementProtocol,omitempty"`
	//description: Netmask of the management address
	NetworkEquipmentManagementAddressMask string `json:"network_equipment_management_address_mask,omitempty" yaml:"managementAddressMask,omitempty"`
	//description: Gateway of the management address
	NetworkEquipmentManagementAddressGateway string `json:"network_equipment_management_address_gateway,omitempty" yaml:"managementAddressGateway,omitempty"`
	//description: MAC address of the management address
	NetworkEquipmentManagementMACAddress string `json:"network_equipment_management_mac_address,omitempty" yaml:"managementMACAddress,omitempty"`
	//description: When set it will automatically create an IPv4 subnet pool for WAN. Deprecated
	NetworkEquipmentPrimaryWANIPv4SubnetPool string `json:"network_equipment_primary_wan_ipv4_subnet_pool,omitempty" yaml:"primaryWANIPv4SubnetPool,omitempty"`
	//description: Size of the subnet pool to automatically create. Deprecated
	//example:
	//  - 24
	NetworkEquipmentPrimaryWANIPv4SubnetPrefixSize int `json:"network_equipment_primary_wan_ipv4_subnet_prefix_size,omitempty" yaml:"primaryWANIPv4SubnetPrefixSize,omitempty"`
	//description: Label of an IPv6 subnet to use when creating the switch.
	NetworkEquipmentPrimaryWANIPv6SubnetPool string `json:"network_equipment_primary_wan_ipv6_subnet_pool,omitempty" yaml:"primaryWANIPv6SubnetPool,omitempty"`
	//description: ID of an IPv6 subnet to use when creating the switch.
	NetworkEquipmentPrimaryWANIPv6SubnetPoolID int `json:"network_equipment_primary_wan_ipv6_subnet_pool_id,omitempty" yaml:"primaryWANIPv6SubnetPoolID,omitempty"`
	//description: CIDR of the subnet to create a subnet pool with
	NetworkEquipmentPrimaryWANIPv6SubnetCIDR string `json:"network_equipment_primary_wan_ipv6_subnet_cidr,omitempty" yaml:"primaryWANIPv6SubnetCIDR,omitempty"`
	//description: Size of the subnet
	//example:
	//  - 64
	NetworkEquipmentPrimaryWANIPv6SubnetPrefixSize int `json:"network_equipment_primary_wan_ipv6_subnet_prefix_size,omitempty" yaml:"primaryWANIPv6SubnetPrefixSize,omitempty"`
	//description: Label of the SAN subnet to use
	NetworkEquipmentPrimarySANSubnetPool string `json:"network_equipment_primary_san_subnet_pool,omitempty" yaml:"primarySANSubnetPool,omitempty"`
	//description: Size of the san subnet to use
	//example:
	//  - 24
	NetworkEquipmentPrimarySANSubnetPrefixSize int `json:"network_equipment_primary_san_subnet_prefix_size,omitempty" yaml:"primarySANSubnetPrefixSize,omitempty"`
	//description: Only used for legacy operation. Start of the quarantine subnet to use during provisioning.
	NetworkEquipmentQuarantineSubnetStart string `json:"network_equipment_quarantine_subnet_start,omitempty" yaml:"quarantineSubnetStart,omitempty"`
	//description: Only used for legacy operation. End of the quarantine subnet to use during provisioning.
	NetworkEquipmentQuarantineSubnetEnd string `json:"network_equipment_quarantine_subnet_end,omitempty" yaml:"quarantineSubnetEnd,omitempty"`
	//description: Only used for legacy operation. End of the quarantine subnet to use during provisioning.
	NetworkEquipmentQuarantineSubnetPrefixSize int `json:"network_equipment_quarantine_subnet_prefix_size,omitempty" yaml:"quarantineSubnetPrefixSize,omitempty"`
	//description: Only used for legacy operation. Gateway to use during provisioning.
	NetworkEquipmentQuarantineSubnetGateway string `json:"network_equipment_quarantine_subnet_gateway,omitempty" yaml:"quarantineSubnetGateway,omitempty"`
	//description: Description of the network equipment
	NetworkEquipmentDescription string `json:"network_equipment_description,omitempty" yaml:"description,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentCountry string `json:"network_equipment_country,omitempty" yaml:"country,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentCity string `json:"network_equipment_city,omitempty" yaml:"city,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentDatacenter string `json:"network_equipment_datacenter,omitempty" yaml:"netDatacenter,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentDatacenterRoom string `json:"network_equipment_datacenter_room,omitempty" yaml:"datacenterRoom,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentDatacenterRack string `json:"network_equipment_datacenter_rack,omitempty" yaml:"datacenterRack,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentRackPositionUpperUnit int `json:"network_equipment_rack_position_upper_unit,omitempty" yaml:"rackPositionUpperUnit,omitempty"`
	//description: Location of the network equipment
	NetworkEquipmentRackPositionLowerUnit int `json:"network_equipment_rack_position_lower_unit,omitempty" yaml:"rackPositionLowerUnit,omitempty"`
	//description: Serial number of the network equipment
	NetworkEquipmentSerialNumber string `json:"network_equipment_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	//description: If applicable the chassis id
	ChassisRackID int `json:"chassis_rack_id,omitempty" yaml:"chassisRackID,omitempty"`
	//description: If applicable, the id of the linked equipment
	NetworkEquipmentTORLinkedID int `json:"network_equipment_tor_linked_id,omitempty"  yaml:"TORLinkedID,omitempty"`
	//description: Tags
	NetworkEquipmentTags []string `json:"network_equipment_tags,omitempty" yaml:"tags,omitempty"`
	//description: If set to true the OS will be reinstalled upon reboot
	NetworkEquipmentRequiresOSInstall bool `json:"network_equipment_requires_os_install" yaml:"requiresOSInstall"`
	//description: If set to true this network equipment can be used for external connections
	NetworkEquipmentIsBorderDevice bool `json:"network_equipment_is_border_device" yaml:"isBorderDevice"`
	//description: If set to true this network equipment can be used for SAN fabrics
	NetworkEquipmentIsStorageSwitch bool `json:"network_equipment_is_storage_switch" yaml:"isStorageSwitch"`
	//description: In certain situations (such as for the VPLS provisioner) a switch can be used as a router
	NetworkEquipmentIsGateway bool `json:"network_equipment_is_gateway" yaml:"isGateway"`
	//description: Types of networks allowed on this switch
	//values:
	// - wan
	// - lan
	// - san
	NetworkEquipmentNetworkTypesAllowed []string `json:"network_equipment_network_types_allowed,omitempty" yaml:"networkTypesAllowed,omitempty"`
	//description: Template id of the NOS to be installed
	VolumeTemplateID int `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	//description: Address of a loopback interface (if applicable)
	NetworkEquipmentLoopbackAddress string `json:"network_equipment_loopback_address,omitempty" yaml:"LoopbackAddress,omitempty"`
	//description: Address of a VTEP
	NetworkEquipmentVTEPAddress string `json:"network_equipment_vtep_address,omitempty" yaml:"VTEPAddress,omitempty"`
	//description: Address of an ASN
	NetworkEquipmentASN int `json:"network_equipment_asn,omitempty" yaml:"ASN,omitempty"`
	//description: In the SDN provisioner the ID of the controller device
	NetworkEquipmentControllerID int `json:"network_equipment_controller_id,omitempty" yaml:"controllerID,omitempty"`
}

// UnmarshalJSON to handle the shitty boolean being returned as 0 and 1 and true and false in different environments
func (s *SwitchDevice) UnmarshalJSON(data []byte) error {

	//SwitchDevice Represents a switch installed in a datacenter.
	var v struct {
		NetworkEquipmentID                             int      `json:"network_equipment_id,omitempty" yaml:"id,omitempty"`
		NetworkEquipmentIdentifierString               string   `json:"network_equipment_identifier_string,omitempty" yaml:"identifierString,omitempty"`
		DatacenterName                                 string   `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
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
		NetworkEquipmentNetworkTypesAllowed            []string `json:"network_equipment_network_types_allowed" yaml:"networkTypesAllowed"`
		//this is the culprit.
		NetworkEquipmentRequiresOSInstall interface{} `json:"network_equipment_requires_os_install" yaml:"requiresOSInstall"`
		NetworkEquipmentIsBorderDevice    interface{} `json:"network_equipment_is_border_device" yaml:"isBorderDevice"`
		NetworkEquipmentIsStorageSwitch   interface{} `json:"network_equipment_is_storage_switch" yaml:"isStorageSwitch"`
		NetworkEquipmentIsGateway         interface{} `json:"network_equipment_is_gateway" yaml:"isGateway"`
		VolumeTemplateID                  int         `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
		NetworkEquipmentLoopbackAddress   string      `json:"network_equipment_loopback_address,omitempty" yaml:"LoopbackAddress,omitempty"`
		NetworkEquipmentVTEPAddress       string      `json:"network_equipment_vtep_address,omitempty" yaml:"VTEPAddress,omitempty"`
		NetworkEquipmentASN               int         `json:"network_equipment_asn,omitempty" yaml:"ASN,omitempty"`
		NetworkEquipmentControllerID      int         `json:"network_equipment_controller_id,omitempty" yaml:"controllerID,omitempty"`
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
	switch v.NetworkEquipmentIsBorderDevice.(type) {
	case float64:
		if v.NetworkEquipmentIsBorderDevice.(float64) == 0.0 {
			v.NetworkEquipmentIsBorderDevice = false
		} else {
			v.NetworkEquipmentIsBorderDevice = true
		}
	}
	switch v.NetworkEquipmentIsStorageSwitch.(type) {
	case float64:
		if v.NetworkEquipmentIsStorageSwitch.(float64) == 0.0 {
			v.NetworkEquipmentIsStorageSwitch = false
		} else {
			v.NetworkEquipmentIsStorageSwitch = true
		}
	}
	switch v.NetworkEquipmentIsGateway.(type) {
	case float64:
		if v.NetworkEquipmentIsGateway.(float64) == 0.0 {
			v.NetworkEquipmentIsGateway = false
		} else {
			v.NetworkEquipmentIsGateway = true
		}
	}
	copier.Copy(&s, &v)

	return nil
}

// SwitchDeviceGet Retrieves information regarding a specified SwitchDevice.
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

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
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
			}
		}
	} else {
		createdObject.NetworkEquipmentManagementPassword = ""
	}

	return &createdObject, nil
}

// SwitchDeviceGetByIdentifierString Retrieves information regarding a specified SwitchDevice by identifier string.
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

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
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
			}
		}
	} else {
		createdObject.NetworkEquipmentManagementPassword = ""
	}

	return &createdObject, nil
}

// SwitchDeviceCreate Creates a record for a new SwitchDevice.
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

// SwitchDeviceDelete deletes a specified switch device and its registered interfaces.
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

// SwitchDevices retrieves all switch devices registered in the database.
func (c *Client) SwitchDevices(datacenter string, switchType string) (*map[string]SwitchDevice, error) {

	var dc *string
	if datacenter != "" {
		dc = &datacenter
	}
	var st *string
	if switchType != "" {
		st = &switchType
	}
	var createdObject map[string]SwitchDevice

	resp, err := c.rpcClient.Call(
		"switch_devices",
		dc,
		st,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]SwitchDevice{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SwitchDevicesInDatacenter retrieves all switch devices in a datacenter
func (c *Client) SwitchDevicesInDatacenter(datacenter string) (*map[string]SwitchDevice, error) {
	return c.SwitchDevices(datacenter, "")
}

// SwitchDeviceUpdate updates an existing switch configuration
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

// CreateOrUpdate implements interface Applier
func (s SwitchDevice) CreateOrUpdate(client MetalCloudClient) error {
	var err error

	err = s.Validate()
	if err != nil {
		return err
	}

	switches, err := client.SwitchDevices(s.DatacenterName, "")
	if err != nil {
		return err
	}

	found := false
	var foundSwitch *SwitchDevice
	for _, d := range *switches {
		if d.NetworkEquipmentIdentifierString == s.NetworkEquipmentIdentifierString {
			found = true
			foundSwitch = &d
			break
		}
	}

	if !found {
		_, err := client.SwitchDeviceCreate(s, false)

		if err != nil {
			return err
		}
	} else {
		s.NetworkEquipmentID = foundSwitch.NetworkEquipmentID
		_, err := client.SwitchDeviceUpdate(foundSwitch.NetworkEquipmentID, s, false)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (s SwitchDevice) Delete(client MetalCloudClient) error {
	err := s.Validate()
	var switchDevice *SwitchDevice
	var id int

	if err != nil {
		return err
	}

	if s.NetworkEquipmentIdentifierString != "" {
		switchDevice, err = client.SwitchDeviceGetByIdentifierString(s.NetworkEquipmentIdentifierString, false)
		if err != nil {
			return err
		}

		id = switchDevice.NetworkEquipmentID
	} else {
		id = s.NetworkEquipmentID
	}

	err = client.SwitchDeviceDelete(id)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (s SwitchDevice) Validate() error {
	if s.NetworkEquipmentID == 0 && s.NetworkEquipmentIdentifierString == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
