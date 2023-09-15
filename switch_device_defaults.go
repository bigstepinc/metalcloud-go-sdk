package metalcloud

import (
	"fmt"
)

// SwitchDeviceDefaults represents a set of switch device defaults for a datacenter.
type SwitchDeviceDefaults struct {
	NetworkEquipmentDefaultsID                int               `json:"network_equipment_defaults_id,omitempty" yaml:"id,omitempty"`
	DatacenterName                            string            `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
	NetworkEquipmentSerialNumber              string            `json:"network_equipment_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	NetworkEquipmentManagementMacAddress      string            `json:"network_equipment_management_mac_address,omitempty" yaml:"managementMacAddress,omitempty"`
	NetworkEquipmentPosition                  string            `json:"network_equipment_position,omitempty" yaml:"position,omitempty"`
	NetworkEquipmentIdentifierString          string            `json:"network_equipment_identifier_string,omitempty" yaml:"identifierString,omitempty"`
	NetworkEquipmentAsn                       int               `json:"network_equipment_asn,omitempty" yaml:"asn,omitempty"`
	NetworkEquipmentPartOfMlagPair            bool              `json:"network_equipment_part_of_mlag_pair,omitempty" yaml:"partOfMlagPair,omitempty"`
	NetworkEquipmentMlagSystemMac             string            `json:"network_equipment_mlag_system_mac,omitempty" yaml:"mlagSystemMac,omitempty"`
	NetworkEquipmentMlagDomainId              int               `json:"network_equipment_mlag_domain_id,omitempty" yaml:"mlagDomainId,omitempty"`
	NetworkEquipmentMlagPeerLinkPortChannelId int               `json:"network_equipment_mlag_peer_link_port_channel_id,omitempty" yaml:"mlagPeerLinkPortChannelId,omitempty"`
	NetworkEquipmentMlagPartnerVlanId         int               `json:"network_equipment_mlag_partner_vlan_id,omitempty" yaml:"mlagPartnerVlanId,omitempty"`
	NetworkEquipmentMlagPartnerHostname       string            `json:"network_equipment_mlag_partner_hostname,omitempty" yaml:"mlagPartnerHostname,omitempty"`
	NetworkEquipmentLoopbackAddressIpv4       string            `json:"network_equipment_loopback_address_ipv4,omitempty" yaml:"loopbackAddressIpv4,omitempty"`
	NetworkEquipmentLoopbackAddressIpv6       string            `json:"network_equipment_loopback_address_ipv6,omitempty" yaml:"loopbackAddressIpv6,omitempty"`
	NetworkEquipmentVtepAddressIpv4           string            `json:"network_equipment_vtep_address_ipv4,omitempty" yaml:"vtepAddressIpv4,omitempty"`
	NetworkEquipmentVtepAddressIpv6           string            `json:"network_equipment_vtep_address_ipv6,omitempty" yaml:"vtepAddressIpv6,omitempty"`
	NetworkEquipmentSkipInitialConfiguration  bool              `json:"network_equipment_skip_initial_configuration,omitempty" yaml:"skipInitialConfiguration,omitempty"`
	VolumeTemplateID                          int               `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	NetworkEquipmentCustomVariables           map[string]string `json:"network_equipment_custom_variables,omitempty" yaml:"customVariables,omitempty"`
}

// SwitchDeviceDefaults retrieves all switch defaults registered in the database for a datacenter.
func (c *Client) SwitchDeviceDefaults(datacenter string) (*[]SwitchDeviceDefaults, error) {
	var createdObject []SwitchDeviceDefaults

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_defaults",
		datacenter,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SwitchDeviceDefaultsCreate adds records for a new set of switch device defaults.
func (c *Client) SwitchDeviceDefaultsCreate(switchDeviceDefaultsArray []SwitchDeviceDefaults) error {
	// When making a call with a single object parameter, we have to put it into an array.
	resp, err := c.rpcClient.Call(
		"switch_device_defaults_add",
		[]interface{}{switchDeviceDefaultsArray},
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// SwitchDeviceDefaultsDelete removes records for the specified switch device defaults.
func (c *Client) SwitchDeviceDefaultsDelete(switchDeviceDefaultsIDs []int) error {
	// When making a call with a single object parameter, we have to put it into an array.
	resp, err := c.rpcClient.Call(
		"switch_device_defaults_remove",
		[]interface{}{switchDeviceDefaultsIDs},
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
