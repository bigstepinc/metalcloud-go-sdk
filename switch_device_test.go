package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSwitchDeviceUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj SwitchDevice
	err := json.Unmarshal([]byte(_switchDeviceFixture1), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.NetworkEquipmentRequiresOSInstall).To(BeTrue())
	Expect(obj.NetworkEquipmentIsBorderDevice).To(BeTrue())
	Expect(obj.NetworkEquipmentIsStorageSwitch).To(BeTrue())
	Expect(obj.NetworkEquipmentIdentifierString).To(Equal("UK_RDG_EVR01_00_0001_00A9_01"))
	Expect(obj.NetworkEquipmentProvisionerPosition).To(Equal("other"))
	Expect(obj.NetworkEquipmentQuarantineSubnetGateway).To(Equal("11.16.0.1"))
	Expect(obj.NetworkEquipmentNetworkTypesAllowed).To(Equal([]string{"wan", "san", "lan", "quarantine"}))

	err = json.Unmarshal([]byte(_switchDeviceFixture2), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.NetworkEquipmentRequiresOSInstall).To(BeFalse())
	Expect(obj.NetworkEquipmentIsBorderDevice).To(BeFalse())
	Expect(obj.NetworkEquipmentIsStorageSwitch).To(BeFalse())
	Expect(obj.NetworkEquipmentIdentifierString).To(Equal("UK_RDG_EVR01_00_0001_00A9_01"))
	Expect(obj.NetworkEquipmentQuarantineSubnetGateway).To(Equal("11.16.0.1"))
	Expect(obj.NetworkEquipmentNetworkTypesAllowed).To(Equal([]string{"wan", "san", "lan", "quarantine"}))

	err = json.Unmarshal([]byte(_switchDeviceFixture3), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.NetworkEquipmentRequiresOSInstall).To(BeFalse())
	Expect(obj.NetworkEquipmentIsBorderDevice).To(BeFalse())
	Expect(obj.NetworkEquipmentIsStorageSwitch).To(BeFalse())
	Expect(obj.NetworkEquipmentIdentifierString).To(Equal("UK_RDG_EVR01_00_0001_00A9_01"))
	Expect(obj.NetworkEquipmentQuarantineSubnetGateway).To(Equal("11.16.0.1"))
	Expect(obj.NetworkEquipmentNetworkTypesAllowed).To(Equal([]string{"wan", "san", "lan", "quarantine"}))

	err = json.Unmarshal([]byte(_switchDeviceFixture4), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.NetworkEquipmentRequiresOSInstall).To(BeTrue())
	Expect(obj.NetworkEquipmentIsBorderDevice).To(BeTrue())
	Expect(obj.NetworkEquipmentIsStorageSwitch).To(BeTrue())
	Expect(obj.NetworkEquipmentIdentifierString).To(Equal("UK_RDG_EVR01_00_0001_00A9_01"))
	Expect(obj.NetworkEquipmentQuarantineSubnetGateway).To(Equal("11.16.0.1"))
	Expect(obj.NetworkEquipmentNetworkTypesAllowed).To(Equal([]string{"wan", "san", "lan", "quarantine"}))

}

func TestSwitchDeviceCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	///// UPDATE TEST
	obj := SwitchDevice{
		NetworkEquipmentIdentifierString: "sw1-env1", //this is among the ones in the _switchesFixture1 list
	}

	//this will include switch with id 4
	responseBody = `{"result": ` + _switchesFixture1 + `,"jsonrpc": "2.0","id": 0}`

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("switch_devices"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())
	//since the switch is in the list
	Expect(m["method"].(string)).To(Equal("switch_device_update"))

	var params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(4.0))

	///// CREATE TEST

	responseBody = `{"result": ` + _switchesFixture1 + `,"jsonrpc": "2.0","id": 0}`
	obj = SwitchDevice{
		NetworkEquipmentIdentifierString: "sw1-env100", //this is NOT among the ones in the _switchesFixture1 list
	}
	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("switch_devices"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())
	//we should now call this methid as the switch does not exist
	Expect(m["method"].(string)).To(Equal("switch_device_create"))
	params = (m["params"].([]interface{}))
	Expect(params[0].(map[string]interface{})["network_equipment_identifier_string"].(string)).To(Equal("sw1-env100"))
}

func TestSwitchDeviceDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _switchDeviceFixture5 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := SwitchDevice{
		NetworkEquipmentID:               100,
		NetworkEquipmentIdentifierString: "UK_RDG_EVR01_00_0001_00A9_01",
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}

	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("switch_device_get"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(string)).To(Equal("UK_RDG_EVR01_00_0001_00A9_01"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("switch_device_delete"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

const _switchDeviceFixture1 = "{\"network_equipment_id\":1,\"datacenter_name\":\"uk-reading\",\"network_equipment_driver\":\"hp5900\",\"network_equipment_position\":\"other\",\"network_equipment_provisioner_type\":\"vpls\",\"network_equipment_identifier_string\":\"UK_RDG_EVR01_00_0001_00A9_01\",\"network_equipment_description\":\"HP Comware Software, Version 7.1.045, Release 2311P06\",\"network_equipment_management_address\":\"10.0.0.0\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"sad\",\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"11.16.0.1\",\"network_equipment_quarantine_subnet_end\":\"11.16.0.00\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"11.16.0.1\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"11.24.0.2\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"100.64.0.0\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":1,\"network_equipment_primary_wan_ipv6_subnet_cidr\":\"2A02:0CB8:0000:0000:0000:0000:0000:0000/53\",\"network_equipment_cached_updated_timestamp\":\"2020-08-04T20:11:49Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":\"\",\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_management_address_mask\":null,\"network_equipment_management_address_gateway\":null,\"network_equipment_requires_os_install\":true,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_numbers\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":1,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_tags\":[],\"network_equipment_management_password\":\"zk3enQ4VXROZyJ9\",\"network_equipment_is_border_device\":true,\"network_equipment_is_storage_switch\":true,\"network_equipment_is_gateway\":true,\"network_equipment_network_types_allowed\": [\"wan\",\"san\",\"lan\",\"quarantine\"]}"
const _switchDeviceFixture2 = "{\"network_equipment_id\":1,\"datacenter_name\":\"uk-reading\",\"network_equipment_driver\":\"hp5900\",\"network_equipment_position\":\"test1\",\"network_equipment_provisioner_type\":\"vpls\",\"network_equipment_identifier_string\":\"UK_RDG_EVR01_00_0001_00A9_01\",\"network_equipment_description\":\"HP Comware Software, Version 7.1.045, Release 2311P06\",\"network_equipment_management_address\":\"10.0.0.0\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"sad\",\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"11.16.0.1\",\"network_equipment_quarantine_subnet_end\":\"11.16.0.00\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"11.16.0.1\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"11.24.0.2\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"100.64.0.0\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":1,\"network_equipment_primary_wan_ipv6_subnet_cidr\":\"2A02:0CB8:0000:0000:0000:0000:0000:0000/53\",\"network_equipment_cached_updated_timestamp\":\"2020-08-04T20:11:49Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":\"\",\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_management_address_mask\":null,\"network_equipment_management_address_gateway\":null,\"network_equipment_requires_os_install\":false,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_numbers\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":1,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_tags\":[],\"network_equipment_management_password\":\"zk3enQ4VXROZyJ9\",\"network_equipment_is_border_device\":false,\"network_equipment_is_storage_switch\":false,\"network_equipment_is_gateway\":true,\"network_equipment_network_types_allowed\": [\"wan\",\"san\",\"lan\",\"quarantine\"]}"
const _switchDeviceFixture3 = "{\"network_equipment_id\":1,\"datacenter_name\":\"uk-reading\",\"network_equipment_driver\":\"hp5900\",\"network_equipment_position\":\"test1\",\"network_equipment_provisioner_type\":\"vpls\",\"network_equipment_identifier_string\":\"UK_RDG_EVR01_00_0001_00A9_01\",\"network_equipment_description\":\"HP Comware Software, Version 7.1.045, Release 2311P06\",\"network_equipment_management_address\":\"10.0.0.0\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"sad\",\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"11.16.0.1\",\"network_equipment_quarantine_subnet_end\":\"11.16.0.00\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"11.16.0.1\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"11.24.0.2\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"100.64.0.0\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":1,\"network_equipment_primary_wan_ipv6_subnet_cidr\":\"2A02:0CB8:0000:0000:0000:0000:0000:0000/53\",\"network_equipment_cached_updated_timestamp\":\"2020-08-04T20:11:49Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":\"\",\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_management_address_mask\":null,\"network_equipment_management_address_gateway\":null,\"network_equipment_requires_os_install\":0,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_numbers\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":1,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_tags\":[],\"network_equipment_management_password\":\"zk3enQ4VXROZyJ9\",\"network_equipment_is_border_device\":0,\"network_equipment_is_storage_switch\":0,\"network_equipment_is_gateway\":true,\"network_equipment_network_types_allowed\": [\"wan\",\"san\",\"lan\",\"quarantine\"]}"
const _switchDeviceFixture4 = "{\"network_equipment_id\":1,\"datacenter_name\":\"uk-reading\",\"network_equipment_driver\":\"hp5900\",\"network_equipment_position\":\"test1\",\"network_equipment_provisioner_type\":\"vpls\",\"network_equipment_identifier_string\":\"UK_RDG_EVR01_00_0001_00A9_01\",\"network_equipment_description\":\"HP Comware Software, Version 7.1.045, Release 2311P06\",\"network_equipment_management_address\":\"10.0.0.0\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"sad\",\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"11.16.0.1\",\"network_equipment_quarantine_subnet_end\":\"11.16.0.00\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"11.16.0.1\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"11.24.0.2\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"100.64.0.0\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":1,\"network_equipment_primary_wan_ipv6_subnet_cidr\":\"2A02:0CB8:0000:0000:0000:0000:0000:0000/53\",\"network_equipment_cached_updated_timestamp\":\"2020-08-04T20:11:49Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":\"\",\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_management_address_mask\":null,\"network_equipment_management_address_gateway\":null,\"network_equipment_requires_os_install\":1,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_numbers\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":1,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_tags\":[],\"network_equipment_management_password\":\"zk3enQ4VXROZyJ9\",\"network_equipment_is_border_device\":1,\"network_equipment_is_storage_switch\":1,\"network_equipment_is_gateway\":true,\"network_equipment_network_types_allowed\": [\"wan\",\"san\",\"lan\",\"quarantine\"]}"
const _switchDeviceFixture5 = "{\"network_equipment_id\": 100}"
const _switchesFixture1 = "{\"4\":{\"network_equipment_id\":4,\"network_equipment_status\":\"active\",\"datacenter_name\":\"us-chi-qts01-dc\",\"network_equipment_driver\":\"os_10\",\"network_equipment_position\":\"leaf\",\"network_equipment_provisioner_type\":\"evpnvxlanl2\",\"network_equipment_identifier_string\":\"sw1-env1\",\"network_equipment_description\":\"\",\"network_equipment_management_address\":\"10.0.5.4\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"admin\",\"network_equipment_loopback_address\":null,\"network_equipment_loopback_address_ipv6\":null,\"network_equipment_vtep_address\":null,\"network_equipment_vtep_address_ipv6\":null,\"network_equipment_asn\":null,\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"192.168.254.0\",\"network_equipment_quarantine_subnet_end\":\"192.168.254.255\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"192.168.254.1\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":2,\"network_equipment_primary_wan_ipv6_subnet_cidr\":null,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_cached_updated_timestamp\":\"0000-00-00T00:00:00Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":null,\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_tags_json\":null,\"network_equipment_management_address_mask\":\"\",\"network_equipment_management_address_gateway\":\"\",\"network_equipment_requires_os_install\":false,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_number\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":true,\"network_equipment_interfaces_blacklist_json\":null,\"network_equipment_controller_id\":null,\"network_equipment_is_border_device\":true,\"network_equipment_is_storage_switch\":true,\"network_equipment_network_types_allowed_json\":\"[\\\"wan\\\", \\\"quarantine\\\", \\\"san\\\"]\",\"network_equipment_order_index\":10,\"network_equipment_chassis_identifier\":\"88:6f:d4:b6:67:fd\",\"subnet_oob_id\":null,\"subnet_oob_index\":0,\"network_equipment_variables_materialized_for_os_assets_json\":null,\"network_equipment_secrets_materialized_for_os_assets_json\":null,\"network_equipment_ready_for_initial_configuration\":false,\"network_equipment_bootstrap_readiness_check_in_progress\":false,\"network_equipment_bootstrap_readiness_check_result_json\":null,\"network_equipment_external_id\":null,\"network_equipment_sdn_pod_id\":null,\"network_equipment_mlag_system_mac\":null,\"network_equipment_mlag_domain_id\":null,\"network_equipment_bootstrap_skip_initial_configuration\":false,\"network_equipment_bootstrap_expected_partner_hostname\":null,\"network_equipment_is_gateway\":true,\"network_equipment_management_password\":\"Use bsidev.password_decrypt:asd\"},\"5\":{\"network_equipment_id\":5,\"network_equipment_status\":\"active\",\"datacenter_name\":\"us-chi-qts01-dc\",\"network_equipment_driver\":\"os_10\",\"network_equipment_position\":\"leaf\",\"network_equipment_provisioner_type\":\"evpnvxlanl2\",\"network_equipment_identifier_string\":\"sw2-env1\",\"network_equipment_description\":\"\",\"network_equipment_management_address\":\"10.0.5.5\",\"network_equipment_management_port\":22,\"network_equipment_management_username\":\"admin\",\"network_equipment_loopback_address\":null,\"network_equipment_loopback_address_ipv6\":null,\"network_equipment_vtep_address\":null,\"network_equipment_vtep_address_ipv6\":null,\"network_equipment_asn\":null,\"network_equipment_quarantine_vlan\":5,\"network_equipment_quarantine_subnet_start\":\"192.168.254.0\",\"network_equipment_quarantine_subnet_end\":\"192.168.254.255\",\"network_equipment_quarantine_subnet_prefix_size\":24,\"network_equipment_quarantine_subnet_gateway\":\"192.168.254.2\",\"network_equipment_primary_wan_ipv4_subnet_pool\":\"\",\"network_equipment_primary_wan_ipv4_subnet_prefix_size\":22,\"network_equipment_primary_san_subnet_pool\":\"\",\"network_equipment_primary_san_subnet_prefix_size\":21,\"network_equipment_primary_wan_ipv6_subnet_pool_id\":2,\"network_equipment_primary_wan_ipv6_subnet_cidr\":null,\"network_equipment_driver_dump_cached_json\":null,\"network_equipment_cached_updated_timestamp\":\"0000-00-00T00:00:00Z\",\"network_equipment_management_protocol\":\"ssh\",\"chassis_rack_id\":null,\"network_equipment_cache_wrapper_json\":null,\"network_equipment_cache_wrapper_phpserialize\":null,\"network_equipment_tor_linked_id\":null,\"network_equipment_uplink_ip_addresses_json\":null,\"network_equipment_tags_json\":null,\"network_equipment_management_address_mask\":\"\",\"network_equipment_management_address_gateway\":\"\",\"network_equipment_requires_os_install\":false,\"network_equipment_management_mac_address\":\"00:00:00:00:00:00\",\"volume_template_id\":null,\"network_equipment_country\":null,\"network_equipment_city\":null,\"network_equipment_datacenter\":null,\"network_equipment_datacenter_room\":null,\"network_equipment_datacenter_rack\":null,\"network_equipment_rack_position_upper_unit\":null,\"network_equipment_rack_position_lower_unit\":null,\"network_equipment_serial_number\":null,\"network_equipment_info_json\":null,\"network_equipment_management_subnet\":null,\"network_equipment_management_subnet_prefix_size\":null,\"network_equipment_management_subnet_start\":null,\"network_equipment_management_subnet_end\":null,\"network_equipment_management_subnet_gateway\":null,\"datacenter_id_parent\":null,\"network_equipment_dhcp_packet_sniffing_is_enabled\":true,\"network_equipment_interfaces_blacklist_json\":null,\"network_equipment_controller_id\":null,\"network_equipment_is_border_device\":false,\"network_equipment_is_storage_switch\":true,\"network_equipment_network_types_allowed_json\":\"[\\\"wan\\\", \\\"quarantine\\\", \\\"san\\\"]\",\"network_equipment_order_index\":20,\"network_equipment_chassis_identifier\":\"f0:d4:e2:0f:d5:ce\",\"subnet_oob_id\":null,\"subnet_oob_index\":0,\"network_equipment_variables_materialized_for_os_assets_json\":null,\"network_equipment_secrets_materialized_for_os_assets_json\":null,\"network_equipment_ready_for_initial_configuration\":false,\"network_equipment_bootstrap_readiness_check_in_progress\":false,\"network_equipment_bootstrap_readiness_check_result_json\":null,\"network_equipment_external_id\":null,\"network_equipment_sdn_pod_id\":null,\"network_equipment_mlag_system_mac\":null,\"network_equipment_mlag_domain_id\":null,\"network_equipment_bootstrap_skip_initial_configuration\":false,\"network_equipment_bootstrap_expected_partner_hostname\":null,\"network_equipment_is_gateway\":true,\"network_equipment_management_password\":\"Use bsidev.password_decrypt:sdsd\"}}"
