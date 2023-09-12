package metalcloud

import (
	"fmt"
)

// SwitchDeviceDefaults represents a set of switch device defaults for a datacenter.
type SwitchDeviceDefaults struct {
	NetworkEquipmentDefaultsID                int    `json:"network_equipment_defaults_id,omitempty" yaml:"id,omitempty"`
	DatacenterName                            string `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
	NetworkEquipmentSerialNumber              string `json:"network_equipment_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	NetworkEquipmentManagementMacAddress      string `json:"network_equipment_management_mac_address,omitempty" yaml:"managementMacAddress,omitempty"`
	NetworkEquipmentPosition                  string `json:"network_equipment_position,omitempty" yaml:"position,omitempty"`
	NetworkEquipmentIdentifierString          string `json:"network_equipment_identifier_string,omitempty" yaml:"identifierString,omitempty"`
	NetworkEquipmentAsn                       int    `json:"network_equipment_asn,omitempty" yaml:"asn,omitempty"`
	NetworkEquipmentPartOfMlagPair            bool   `json:"network_equipment_part_of_mlag_pair,omitempty" yaml:"partOfMlagPair,omitempty"`
	NetworkEquipmentMlagSystemMac             string `json:"network_equipment_mlag_system_mac,omitempty" yaml:"mlagSystemMac,omitempty"`
	NetworkEquipmentMlagDomainId              int    `json:"network_equipment_mlag_domain_id,omitempty" yaml:"mlagDomainId,omitempty"`
	NetworkEquipmentMlagPeerLinkPortChannelId int    `json:"network_equipment_mlag_peer_link_port_channel_id,omitempty" yaml:"mlagPeerLinkPortChannelId,omitempty"`
	NetworkEquipmentMlagPartnerVlanId         int    `json:"network_equipment_mlag_partner_vlan_id,omitempty" yaml:"mlagPartnerVlanId,omitempty"`
	NetworkEquipmentMlagPartnerHostname       string `json:"network_equipment_mlag_partner_hostname,omitempty" yaml:"mlagPartnerHostname,omitempty"`
	NetworkEquipmentLoopbackAddressIpv4       string `json:"network_equipment_loopback_address_ipv4,omitempty" yaml:"loopbackAddressIpv4,omitempty"`
	NetworkEquipmentLoopbackAddressIpv6       string `json:"network_equipment_loopback_address_ipv6,omitempty" yaml:"loopbackAddressIpv6,omitempty"`
	NetworkEquipmentVtepAddressIpv4           string `json:"network_equipment_vtep_address_ipv4,omitempty" yaml:"vtepAddressIpv4,omitempty"`
	NetworkEquipmentVtepAddressIpv6           string `json:"network_equipment_vtep_address_ipv6,omitempty" yaml:"vtepAddressIpv6,omitempty"`
	NetworkEquipmentSkipInitialConfiguration  bool   `json:"network_equipment_skip_initial_configuration,omitempty" yaml:"skipInitialConfiguration,omitempty"`
	VolumeTemplateID                          int    `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	NetworkEquipmentCustomVariablesJson       string `json:"network_equipment_custom_variables_json,omitempty" yaml:"customVariablesJson,omitempty"`
}

// SwitchDeviceDefaults retrieves all switch defaults registered in the database for a datacenter.
func (c *Client) SwitchDeviceDefaults(datacenter string) (*map[int]SwitchDeviceDefaults, error) {
	var createdObject map[int]SwitchDeviceDefaults

	resp, err := c.rpcClient.Call(
		"switch_device_defaults",
		datacenter,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[int]SwitchDeviceDefaults{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SwitchDeviceDefaultsAdd adds records for a new set of switch device defaults.
func (c *Client) SwitchDeviceDefaultsAdd(switchDeviceDefaultsArray []SwitchDeviceDefaults) error {
	resp, err := c.rpcClient.Call(
		"switch_device_defaults_add",
		switchDeviceDefaultsArray,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// SwitchDeviceDefaultsRemove removes records for the specified switch device defaults.
func (c *Client) SwitchDeviceDefaultsRemove(switchDeviceDefaultsIDs []int) error {
	resp, err := c.rpcClient.Call(
		"switch_device_defaults_remove",
		switchDeviceDefaultsIDs,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
