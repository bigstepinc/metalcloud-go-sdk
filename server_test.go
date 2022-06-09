package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestServerUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj Server
	err := json.Unmarshal([]byte(_serverFixture1), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	Expect(obj.ServerID).To(Equal(310))

	err = json.Unmarshal([]byte(_serverFixture2), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	Expect(obj.ServerID).To(Equal(685))
	Expect(obj.ServerDisks[0].ServerDiskSizeGB).To(Equal(279))

	Expect(*obj.ServerRackName).To(Equal("Rack Name"))
	Expect(*obj.ServerRackPositionLowerUnit).To(Equal("L-2004"))
	Expect(*obj.ServerRackPositionUpperUnit).To(Equal("U-2404"))
	Expect(*obj.ServerInventoryId).To(Equal("id-20040424"))
}

func TestServerUnmarshalWithNICDetailsTest(t *testing.T) {
	RegisterTestingT(t)

	var obj Server
	err := json.Unmarshal([]byte(_serverFixture4), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	Expect(obj.NICDetails["00:15:17:c0:50:24"].SwitchHostname).To(Equal("RO_LAB_BUH01_00_0001_DC1_Rack01"))
}

func TestServerCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _serverFixture3 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Server{
		ServerID: 100,
	}

	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("server_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("server_edit"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

	responseBody = `{"error": {"message": "Server not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("server_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("server_create"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(map[string]interface{})["server_id"].(float64)).To(Equal(100.0))
}

func TestServerDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Server{
		ServerID: 100,
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("server_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))
}

func TestServerRegister(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _serverFixture3 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	serverCreateAndRegister := ServerCreateAndRegister{
		DatacenterName: "datacenter",
		ServerVendor: "hp",
		ServerManagementAddress: "127.0.0.1",
		ServerManagementUser: "root",
		ServerManagementPassword: "calvin",
	}

	_, err = mc.ServerCreateAndRegister(serverCreateAndRegister)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err = json.Unmarshal([]byte(body), &m)
	Expect(err).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("server_create_and_register"))

	params := (m["params"].([]interface{}))
	param := (params[0].(map[string]interface{}))
	Expect(param["datacenter_name"].(string)).To(Equal(serverCreateAndRegister.DatacenterName))
	Expect(param["server_vendor"].(string)).To(Equal(serverCreateAndRegister.ServerVendor))
	Expect(param["server_ipmi_host"].(string)).To(Equal(serverCreateAndRegister.ServerManagementAddress))
	Expect(param["server_ipmi_user"].(string)).To(Equal(serverCreateAndRegister.ServerManagementUser))
	Expect(param["server_ipmi_password"].(string)).To(Equal(serverCreateAndRegister.ServerManagementPassword))
}

func TestServerCheckForMissingProperties(t *testing.T) {
	RegisterTestingT(t)

	//this test checks if there are missing properties in the process of
	//umarshaling to the Server object and re-marshaling

	//we put properties in a Server object
	var obj Server
	err := json.Unmarshal([]byte(_serverFixture5), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	//we serialize the server object again into json
	jsonFromServerObj, err := json.Marshal(obj)
	Expect(err).To(BeNil())

	//we build a map of properties from the json we created
	//so we can compare it with the original one
	var mapFromServerObj map[string]interface{}
	err = json.Unmarshal([]byte(jsonFromServerObj), &mapFromServerObj)
	Expect(err).To(BeNil())

	//we use this to determine the properties
	//in the json
	var mapFromOriginalJSON map[string]interface{}
	err = json.Unmarshal([]byte(_serverFixture5), &mapFromOriginalJSON)
	Expect(err).To(BeNil())

	//we take each property from the original json and look for it in the

	for k, v := range mapFromOriginalJSON {

		//skip some properties
		switch k {
		case
			"type":
			continue

		}

		if v != nil && v != false {
			Expect(mapFromServerObj).To(HaveKey(k))

			switch k {
			case
				"server_interfaces",
				"server_disks":

				//Expect(reflect.DeepEqual(mapFromServerObj[k], v)).To(BeTrue(), fmt.Sprintf("%s is not equal in both sides", k))
				continue
			default:
				Expect(mapFromServerObj[k]).To(Equal(v))
			}
		}

	}

	//we take each property from the server object
	//to see if we have extra properties
	for k, _ := range mapFromServerObj {
		if k == "server_rack_id" {
			continue
		}
		Expect(mapFromOriginalJSON).To(HaveKey(k))
	}

}

const _serverFixture1 = "{\"server_id\":310,\"agent_id\":44,\"datacenter_name\":\"es-madrid\",\"server_uuid\":\"44454C4C-5900-1033-8032-B9C04F434631\",\"server_serial_number\":\"9Y32CF1\",\"server_product_name\":\"PowerEdge 1950\",\"server_vendor\":\"Dell Inc.\",\"server_vendor_sku_id\":\"0\",\"server_ipmi_host\":\"10.255.237.28\",\"server_ipmi_internal_username\":\"bsii3Cu\",\"server_ipmi_internal_password_encrypted\":\"\",\"server_ipmi_version\":\"2\",\"server_ram_gbytes\":8,\"server_processor_count\":2,\"server_processor_core_mhz\":2333,\"server_processor_core_count\":4,\"server_processor_name\":\"Intel(R) Xeon(R) CPU           E5345  @ 2.33GHz\",\"server_processor_cpu_mark\":0,\"server_processor_threads\":1,\"server_type_id\":14,\"server_status\":\"available\",\"server_comments\":\"a\",\"server_details_xml\":null,\"server_network_total_capacity_mbps\":4000,\"server_ipmi_channel\":0,\"server_power_status\":\"off\",\"server_power_status_last_update_timestamp\":\"2020-08-19T13:05:04Z\",\"server_ilo_reset_timestamp\":\"0000-00-00T00:00:00Z\",\"server_boot_last_update_timestamp\":null,\"server_bdk_debug\":false,\"server_dhcp_status\":\"deny_requests\",\"server_bios_info_json\":\"{\\\"server_bios_vendor\\\":\\\"Dell Inc.\\\",\\\"server_bios_version\\\":\\\"2.7.0\\\"}\",\"server_vendor_info_json\":\"{\\\"management\\\":\\\"iDRAC\\\",\\\"version\\\":\\\"2.75.75\\\"}\",\"server_class\":\"bigdata\",\"server_created_timestamp\":\"2019-07-02T07:57:19Z\",\"subnet_oob_id\":2,\"subnet_oob_index\":28,\"server_boot_type\":\"classic\",\"server_disk_wipe\":true,\"server_disk_count\":0,\"server_disk_size_mbytes\":0,\"server_disk_type\":\"none\",\"server_requires_manual_cleaning\":false,\"chassis_rack_id\":null,\"server_custom_json\":\"\",\"server_instance_custom_json\":null,\"server_last_cleanup_start\":\"2020-08-12T14:26:47Z\",\"server_allocation_timestamp\":null,\"server_dhcp_packet_sniffing_is_enabled\":true,\"snmp_community_password_dcencrypted\":null,\"server_mgmt_snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_port\":161,\"server_mgmt_snmp_version\":2,\"server_dhcp_relay_security_is_enabled\":true,\"server_keys_json\":\"\",\"server_info_json\":null,\"server_ipmi_credentials_need_update\":false,\"server_gpu_count\":0,\"server_gpu_vendor\":\"\",\"server_gpu_model\":\"\",\"server_bmc_mac_address\":null,\"server_metrics_metadata_json\":null,\"server_interfaces\":[{\"server_interface_mac_address\":\"00:1d:09:64:f0:2b\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:1d:09:64:f0:2d\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:15:17:c0:4c:e6\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:15:17:c0:4c:e7\",\"type\":\"ServerInterface\"}],\"server_disks\":[],\"server_tags\":[],\"type\":\"Server\",\"server_rack_name\":\"Rack Name\",\"server_rack_position_lower_unit\":\"L-2004\",\"server_rack_position_upper_unit\":\"U-2404\",\"server_inventory_id\":\"id-20040424\"}"
const _serverFixture2 = "{\"server_id\":685,\"agent_id\":137,\"datacenter_name\":\"es-madrid\",\"server_uuid\":\"44454C4C-5800-1039-8052-C2C04F373632\",\"server_serial_number\":\"BX9R762\",\"server_product_name\":\"PowerEdge R730 (SKU=NotProvided;ModelName=PowerEdge R730)\",\"server_vendor\":\"Dell Inc.\",\"server_vendor_sku_id\":\"SKU=NotProvided;ModelName=PowerEdge R730\",\"server_ipmi_host\":\"10.255.237.40\",\"server_ipmi_internal_username\":\"\",\"server_ipmi_internal_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_ipmi_version\":\"2\",\"server_ram_gbytes\":32,\"server_processor_count\":2,\"server_processor_core_mhz\":1295,\"server_processor_core_count\":8,\"server_processor_name\":\"Intel(R) Xeon(R) CPU E5-2630 v3 @ 2.40GHz\",\"server_processor_cpu_mark\":0,\"server_processor_threads\":2,\"server_type_id\":33,\"server_status\":\"available\",\"server_comments\":null,\"server_details_xml\":null,\"server_network_total_capacity_mbps\":4000,\"server_ipmi_channel\":0,\"server_power_status\":\"off\",\"server_power_status_last_update_timestamp\":\"2020-08-19T13:05:04Z\",\"server_ilo_reset_timestamp\":\"2020-07-22T08:26:38Z\",\"server_boot_last_update_timestamp\":null,\"server_bdk_debug\":false,\"server_dhcp_status\":\"deny_requests\",\"server_bios_info_json\":\"\",\"server_vendor_info_json\":\"\",\"server_class\":\"bigdata\",\"server_created_timestamp\":\"2019-10-08T10:30:48Z\",\"subnet_oob_id\":2,\"subnet_oob_index\":21,\"server_boot_type\":\"uefi\",\"server_disk_wipe\":true,\"server_disk_count\":1,\"server_disk_size_mbytes\":285696,\"server_disk_type\":\"HDD\",\"server_requires_manual_cleaning\":false,\"chassis_rack_id\":null,\"server_custom_json\":\"\",\"server_instance_custom_json\":null,\"server_last_cleanup_start\":\"2020-07-22T08:28:45Z\",\"server_allocation_timestamp\":null,\"server_dhcp_packet_sniffing_is_enabled\":true,\"snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_port\":161,\"server_mgmt_snmp_version\":2,\"server_dhcp_relay_security_is_enabled\":true,\"server_keys_json\":\"\",\"server_info_json\":null,\"server_ipmi_credentials_need_update\":false,\"server_gpu_count\":0,\"server_gpu_vendor\":\"\",\"server_gpu_model\":\"\",\"server_bmc_mac_address\":null,\"server_metrics_metadata_json\":\"{\\\"Fans\\\": [{\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan1\\\", \\\"Label\\\": \\\"fan.systemboard.1\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 1, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan2\\\", \\\"Label\\\": \\\"fan.systemboard.2\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 2, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan3\\\", \\\"Label\\\": \\\"fan.systemboard.3\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 3, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan4\\\", \\\"Label\\\": \\\"fan.systemboard.4\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 4, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan5\\\", \\\"Label\\\": \\\"fan.systemboard.5\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 5, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan6\\\", \\\"Label\\\": \\\"fan.systemboard.6\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 6, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}], \\\"Temperatures\\\": [{\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Inlet Temp\\\", \\\"Label\\\": \\\"temperature.intake.4\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 4, \\\"PhysicalContext\\\": \\\"Intake\\\", \\\"UpperThresholdFatal\\\": 47, \\\"UpperThresholdCritical\\\": 47}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Exhaust Temp\\\", \\\"Label\\\": \\\"temperature.exhaust.1\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 1, \\\"PhysicalContext\\\": \\\"Exhaust\\\", \\\"UpperThresholdFatal\\\": 75, \\\"UpperThresholdCritical\\\": 75}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"CPU1 Temp\\\", \\\"Label\\\": \\\"temperature.cpu.14\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 14, \\\"PhysicalContext\\\": \\\"CPU\\\", \\\"UpperThresholdFatal\\\": 87, \\\"UpperThresholdCritical\\\": 87}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"CPU2 Temp\\\", \\\"Label\\\": \\\"temperature.cpu.15\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 15, \\\"PhysicalContext\\\": \\\"CPU\\\", \\\"UpperThresholdFatal\\\": 87, \\\"UpperThresholdCritical\\\": 87}]}\",\"server_interfaces\":[{\"server_interface_mac_address\":\"44:a8:42:24:93:4e\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"44:a8:42:24:93:50\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:1b:21:72:f2:74\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:1b:21:72:f2:75\",\"type\":\"ServerInterface\"}],\"server_disks\":[{\"server_disk_id\":226,\"server_disk_model\":\"PERC H730 Mini\",\"server_disk_size_gb\":279,\"server_id\":685,\"server_disk_serial\":\"644a84203323c500265e82e96590c769\",\"server_disk_vendor\":\"DELL\",\"server_disk_status\":\"installed\",\"server_disk_type\":\"HDD\",\"type\":\"ServerDisk\"}],\"server_tags\":[\"0\",\"1\",\"2\"],\"type\":\"Server\"}"
const _serverFixture3 = "{\"server_id\": 100}"
const _serverFixture4 = "{\r\n   \"server_uuid\":\"44454C4C-5300-1047-804E-B2C04F5A4331\",\r\n   \"server_product_name\":\"PowerEdge 1950\",\r\n   \"server_vendor\":\"Dell Inc.\",\r\n   \"server_serial_number\":\"2SGNZC1\",\r\n   \"slot\":null,\r\n   \"server_vendor_sku_id\":\"0\",\r\n   \"server_ipmi_channel\":null,\r\n   \"server_bios_vendor\":\"Dell Inc.\",\r\n   \"server_bios_version\":\"2.7.0\",\r\n   \"enabled_efi_variables\":null,\r\n   \"server_processor_count\":2,\r\n   \"server_processor_threads\":1,\r\n   \"server_processor_core_count\":4,\r\n   \"server_processor_enabled_core_count\":4,\r\n   \"server_processor_name\":\"Intel(R) Xeon(R) CPU           E5345  @ 2.33GHz\",\r\n   \"server_processor_core_mhz\":2333,\r\n   \"server_processor_cpu_mark\":null,\r\n   \"server_ram_gbytes\":8,\r\n   \"diskCollection\":[\r\n      \r\n   ],\r\n   \"server_disk_count\":0,\r\n   \"server_disk_size_gb\":0,\r\n   \"server_disk_size_mbytes\":0,\r\n   \"server_disk_type\":\"none\",\r\n   \"server_network_total_capacity_mbps\":4000,\r\n   \"nic_details\":{\r\n      \"00:15:c5:f0:62:47\":{\r\n         \"network_equipment_interface_lldp_information\":\"Chassis ID TLV MAC: 44:31:92:6b:5e:fc Port ID TLV Ifname: GigabitEthernet2 0 13 Time to Live TLV 120 Port Description TLV QUARANTINE System Name TLV RO_LAB_BUH01_00_0001_DC1_Rack01 System Description TLV HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P. System Capabilities TLV System capabilities: Bridge, Router Enabled capabilities: Bridge, Router Management Address TLV IPv4: 172.16.0.1 Ifindex: 975 Port VLAN ID TLV PVID: 5 Link Aggregation TLV Aggregation capable Currently not aggregated Aggregated Port ID: 0 MAC\\/PHY Configuration Status TLV Auto-negotiation supported and enabled PMD auto-negotiation capabilities: 0x0000 MAU type: 1000 BaseTFD Power via MDI TLV Port class PSE PSE MDI power not supported PSE pairs not controllable PSE Power pair: signal Power class 2 Maximum Frame Size TLV 10000 End of LLDPDU TLV\",\r\n         \"network_equipment_interface_mac_address\":\"44:31:92:6b:5e:fc\",\r\n         \"switch_port_id\":\"GigabitEthernet2\\/0\\/13\",\r\n         \"switch_port_description\":\"QUARANTINE\",\r\n         \"switch_hostname\":\"RO_LAB_BUH01_00_0001_DC1_Rack01\",\r\n         \"network_equipment_description\":\"HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P.\",\r\n         \"switch_vlan_id\":\"5\",\r\n         \"server_interface_index\":1,\r\n         \"server_interface_mac_address\":\"00:15:c5:f0:62:47\",\r\n         \"server_interface_capacity_mbps\":1000\r\n      },\r\n      \"00:15:c5:f0:62:49\":{\r\n         \"network_equipment_interface_lldp_information\":\"Chassis ID TLV MAC: 44:31:92:6b:5e:fc Port ID TLV Ifname: GigabitEthernet2\\/0\\/15 Time to Live TLV 120 Port Description TLV QUARANTINE System Name TLV RO_LAB_BUH01_00_0001_DC1_Rack01 System Description TLV HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P. System Capabilities TLV System capabilities: Bridge, Router Enabled capabilities: Bridge, Router Management Address TLV IPv4: 172.16.0.1 Ifindex: 975 Port VLAN ID TLV PVID: 5 Link Aggregation TLV Aggregation capable Currently not aggregated Aggregated Port ID: 0 MAC\\/PHY Configuration Status TLV Auto-negotiation supported and enabled PMD auto-negotiation capabilities: 0x0000 MAU type: 1000 BaseTFD Power via MDI TLV Port class PSE PSE MDI power not supported PSE pairs not controllable PSE Power pair: signal Power class 2 Maximum Frame Size TLV 10000 End of LLDPDU TLV\",\r\n         \"network_equipment_interface_mac_address\":\"44:31:92:6b:5e:fc\",\r\n         \"switch_port_id\":\"GigabitEthernet2\\/0\\/15\",\r\n         \"switch_port_description\":\"QUARANTINE\",\r\n         \"switch_hostname\":\"RO_LAB_BUH01_00_0001_DC1_Rack01\",\r\n         \"network_equipment_description\":\"HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P.\",\r\n         \"switch_vlan_id\":\"5\",\r\n         \"server_interface_index\":3,\r\n         \"server_interface_mac_address\":\"00:15:c5:f0:62:49\",\r\n         \"server_interface_capacity_mbps\":1000\r\n      },\r\n      \"00:15:17:c0:50:24\":{\r\n         \"network_equipment_interface_lldp_information\":\"Chassis ID TLV MAC: 44:31:92:6b:5e:fc Port ID TLV Ifname: GigabitEthernet1\\/0\\/13 Time to Live TLV 120 Port Description TLV QUARANTINE System Name TLV RO_LAB_BUH01_00_0001_DC1_Rack01 System Description TLV HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P. System Capabilities TLV System capabilities: Bridge, Router Enabled capabilities: Bridge, Router Management Address TLV IPv4: 172.16.0.1 Ifindex: 975 Port VLAN ID TLV PVID: 5 Link Aggregation TLV Aggregation capable Currently not aggregated Aggregated Port ID: 0 MAC\\/PHY Configuration Status TLV Auto-negotiation supported and enabled PMD auto-negotiation capabilities: 0x0000 MAU type: 1000 BaseTFD Power via MDI TLV Port class PSE PSE MDI power not supported PSE pairs not controllable PSE Power pair: signal Power class 2 Maximum Frame Size TLV 10000 End of LLDPDU TLV\",\r\n         \"network_equipment_interface_mac_address\":\"44:31:92:6b:5e:fc\",\r\n         \"switch_port_id\":\"GigabitEthernet1\\/0\\/13\",\r\n         \"switch_port_description\":\"QUARANTINE\",\r\n         \"switch_hostname\":\"RO_LAB_BUH01_00_0001_DC1_Rack01\",\r\n         \"network_equipment_description\":\"HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P.\",\r\n         \"switch_vlan_id\":\"5\",\r\n         \"server_interface_index\":0,\r\n         \"server_interface_mac_address\":\"00:15:17:c0:50:24\",\r\n         \"server_interface_capacity_mbps\":1000\r\n      },\r\n      \"00:15:17:c0:50:25\":{\r\n         \"network_equipment_interface_lldp_information\":\"Chassis ID TLV MAC: 44:31:92:6b:5e:fc Port ID TLV Ifname: GigabitEthernet1\\/0\\/15 Time to Live TLV 120 Port Description TLV QUARANTINE System Name TLV RO_LAB_BUH01_00_0001_DC1_Rack01 System Description TLV HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P. System Capabilities TLV System capabilities: Bridge, Router Enabled capabilities: Bridge, Router Management Address TLV IPv4: 172.16.0.1 Ifindex: 975 Port VLAN ID TLV PVID: 5 Link Aggregation TLV Aggregation capable Currently not aggregated Aggregated Port ID: 0 MAC\\/PHY Configuration Status TLV Auto-negotiation supported and enabled PMD auto-negotiation capabilities: 0x0000 MAU type: 1000 BaseTFD Power via MDI TLV Port class PSE PSE MDI power not supported PSE pairs not controllable PSE Power pair: signal Power class 2 Maximum Frame Size TLV 10000 End of LLDPDU TLV\",\r\n         \"network_equipment_interface_mac_address\":\"44:31:92:6b:5e:fc\",\r\n         \"switch_port_id\":\"GigabitEthernet1\\/0\\/15\",\r\n         \"switch_port_description\":\"QUARANTINE\",\r\n         \"switch_hostname\":\"RO_LAB_BUH01_00_0001_DC1_Rack01\",\r\n         \"network_equipment_description\":\"HP Comware Platform Software, Software Version 7.1.045, Release 2311P06 HP 5900AF-48G-4XG-2QSFP+ Switch Copyright (c) 2010-2015 Hewlett-Packard Development Company, L.P.\",\r\n         \"switch_vlan_id\":\"5\",\r\n         \"server_interface_index\":2,\r\n         \"server_interface_mac_address\":\"00:15:17:c0:50:25\",\r\n         \"server_interface_capacity_mbps\":1000\r\n      }\r\n   },\r\n   \"server_gpu_count\":0,\r\n   \"server_gpu_list\":[\r\n      \r\n   ]\r\n}"
const _serverFixture5 = "{\"server_id\":1749,\"agent_id\":396,\"datacenter_name\":\"ke-nairobi\",\"server_uuid\":\"30373637-3233-5A43-4A35-303630433251\",\"server_serial_number\":\"CZJ5060C2Q\",\"server_product_name\":\"ProLiant DL380 Gen9\",\"server_vendor\":\"HP\",\"server_vendor_sku_id\":\"767032-B21\",\"server_ipmi_host\":\"10.255.238.13\",\"server_ipmi_internal_username\":\"bsiUdYk\",\"server_ipmi_internal_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_ipmi_version\":\"2\",\"server_ram_gbytes\":32,\"server_processor_count\":1,\"server_processor_core_mhz\":2300,\"server_processor_core_count\":18,\"server_processor_name\":\"6.63.2\",\"server_processor_cpu_mark\":0,\"server_processor_threads\":2,\"server_type_id\":13134,\"server_status\":\"cleaning\",\"server_comments\":null,\"server_details_xml\":null,\"server_network_total_capacity_mbps\":2000,\"server_ipmi_channel\":2,\"server_power_status\":\"on\",\"server_power_status_last_update_timestamp\":\"2022-02-25T06:23:19Z\",\"server_ilo_reset_timestamp\":\"2022-02-11T08:08:46Z\",\"server_boot_last_update_timestamp\":null,\"server_bdk_debug\":false,\"server_dhcp_status\":\"quarantine\",\"server_bios_info_json\":\"{\\\"server_bios_vendor\\\":\\\"HP\\\",\\\"server_bios_version\\\":\\\"P89\\\"}\",\"server_vendor_info_json\":\"{\\\"management\\\":\\\"iLO\\\",\\\"version\\\":\\\"2.76\\\"}\",\"server_class\":\"bigdata\",\"server_created_timestamp\":\"2021-12-13T12:59:03Z\",\"subnet_oob_id\":5,\"subnet_oob_index\":13,\"server_boot_type\":\"classic\",\"server_disk_wipe\":true,\"server_disk_count\":2,\"server_disk_size_mbytes\":953344,\"server_disk_type\":\"HDD\",\"server_requires_manual_cleaning\":false,\"chassis_rack_id\":null,\"server_custom_json\":\"{\\\"previous_ipmi_username\\\":\\\"bsizyBs\\\",\\\"previous_ipmi_password_encrypted\\\":\\\"rqi|aes-cbc|nLPFIIg8DAAKkawrADW5DvwyxHS3WbDADlpW2rAWMXHD5sepNXKMe8ORFcrHsl\\\\/C7bR\\\\/igaFZVbe\\\\/hjWycMM8w==\\\"}\",\"server_instance_custom_json\":null,\"server_last_cleanup_start\":\"2022-02-14T19:37:29Z\",\"server_allocation_timestamp\":\"2022-02-11T09:19:30Z\",\"server_dhcp_packet_sniffing_is_enabled\":true,\"snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_port\":161,\"server_mgmt_snmp_version\":2,\"server_dhcp_relay_security_is_enabled\":true,\"server_keys_json\":\"{\\\"keys\\\": {\\\"r1\\\": {\\\"created\\\": \\\"2021-12-13T13:02:40Z\\\", \\\"salt_encrypted\\\": \\\"rqi|aes-cbc|Aso/ESc/hcipjclSo52EhVundxRxCTgLbzYkEpwVkL3pucFOMU8ri36oSXa/9rIl1TyTb+lAfhQP3SXf0sArng==\\\", \\\"aes_key_encrypted\\\": \\\"rqi|aes-cbc|ebuH5QNF7cKgSNuvFo2kqtbNAER6I1NMTw8pYlS5GCag0LuY5K6HHufOZVsvZUcadneZIhTVz7XxZXpGavYwN2P7B7fyNTDf5pJJdK5lkrA=\\\"}}, \\\"active_index\\\": \\\"r1\\\", \\\"keys_partition\\\": \\\"server_id_1749\\\"}\",\"server_info_json\":null,\"server_ipmi_credentials_need_update\":false,\"server_gpu_count\":0,\"server_gpu_vendor\":\"\",\"server_gpu_model\":\"\",\"server_bmc_mac_address\":null,\"server_metrics_metadata_json\":\"{\\\"Fans\\\": [{\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpServerFan.1.0.0\\\", \\\"Location\\\": \\\"System\\\", \\\"@odata.type\\\": \\\"#HpServerFan.1.0.0.HpServerFan\\\"}}, \\\"Name\\\": \\\"Fan 1\\\", \\\"Label\\\": \\\"fan.system.1\\\", \\\"Units\\\": \\\"Percent\\\", \\\"Number\\\": 1, \\\"PhysicalContext\\\": \\\"System\\\"}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpServerFan.1.0.0\\\", \\\"Location\\\": \\\"System\\\", \\\"@odata.type\\\": \\\"#HpServerFan.1.0.0.HpServerFan\\\"}}, \\\"Name\\\": \\\"Fan 2\\\", \\\"Label\\\": \\\"fan.system.2\\\", \\\"Units\\\": \\\"Percent\\\", \\\"Number\\\": 2, \\\"PhysicalContext\\\": \\\"System\\\"}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpServerFan.1.0.0\\\", \\\"Location\\\": \\\"System\\\", \\\"@odata.type\\\": \\\"#HpServerFan.1.0.0.HpServerFan\\\"}}, \\\"Name\\\": \\\"Fan 3\\\", \\\"Label\\\": \\\"fan.system.3\\\", \\\"Units\\\": \\\"Percent\\\", \\\"Number\\\": 3, \\\"PhysicalContext\\\": \\\"System\\\"}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpServerFan.1.0.0\\\", \\\"Location\\\": \\\"System\\\", \\\"@odata.type\\\": \\\"#HpServerFan.1.0.0.HpServerFan\\\"}}, \\\"Name\\\": \\\"Fan 4\\\", \\\"Label\\\": \\\"fan.system.4\\\", \\\"Units\\\": \\\"Percent\\\", \\\"Number\\\": 4, \\\"PhysicalContext\\\": \\\"System\\\"}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpServerFan.1.0.0\\\", \\\"Location\\\": \\\"System\\\", \\\"@odata.type\\\": \\\"#HpServerFan.1.0.0.HpServerFan\\\"}}, \\\"Name\\\": \\\"Fan 5\\\", \\\"Label\\\": \\\"fan.system.5\\\", \\\"Units\\\": \\\"Percent\\\", \\\"Number\\\": 5, \\\"PhysicalContext\\\": \\\"System\\\"}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpServerFan.1.0.0\\\", \\\"Location\\\": \\\"System\\\", \\\"@odata.type\\\": \\\"#HpServerFan.1.0.0.HpServerFan\\\"}}, \\\"Name\\\": \\\"Fan 6\\\", \\\"Label\\\": \\\"fan.system.6\\\", \\\"Units\\\": \\\"Percent\\\", \\\"Number\\\": 6, \\\"PhysicalContext\\\": \\\"System\\\"}], \\\"Temperatures\\\": [{\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 1, \\\"LocationYmm\\\": 1}}, \\\"Name\\\": \\\"01-Inlet Ambient\\\", \\\"Label\\\": \\\"temperature.intake.1\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 1, \\\"PhysicalContext\\\": \\\"Intake\\\", \\\"UpperThresholdFatal\\\": 46, \\\"UpperThresholdCritical\\\": 42}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 5}}, \\\"Name\\\": \\\"02-CPU 1\\\", \\\"Label\\\": \\\"temperature.cpu.2\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 2, \\\"PhysicalContext\\\": \\\"CPU\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 70}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 14, \\\"LocationYmm\\\": 5}}, \\\"Name\\\": \\\"04-P1 DIMM 1-6\\\", \\\"Label\\\": \\\"temperature.systemboard.3\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 3, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 9, \\\"LocationYmm\\\": 5}}, \\\"Name\\\": \\\"05-P1 DIMM 7-12\\\", \\\"Label\\\": \\\"temperature.systemboard.4\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 4, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 87}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 6, \\\"LocationYmm\\\": 5}}, \\\"Name\\\": \\\"06-P2 DIMM 1-6\\\", \\\"Label\\\": \\\"temperature.systemboard.5\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 5, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 1, \\\"LocationYmm\\\": 5}}, \\\"Name\\\": \\\"07-P2 DIMM 7-12\\\", \\\"Label\\\": \\\"temperature.systemboard.6\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 6, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 2, \\\"LocationYmm\\\": 3}}, \\\"Name\\\": \\\"08-HD Max\\\", \\\"Label\\\": \\\"temperature.systemboard.7\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 7, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 60}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 13, \\\"LocationYmm\\\": 10}}, \\\"Name\\\": \\\"10-Chipset\\\", \\\"Label\\\": \\\"temperature.systemboard.8\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 8, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 105}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 1, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"11-PS 1 Inlet\\\", \\\"Label\\\": \\\"temperature.systemboard.9\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 9, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 4, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"12-PS 2 Inlet\\\", \\\"Label\\\": \\\"temperature.systemboard.10\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 10, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 10, \\\"LocationYmm\\\": 1}}, \\\"Name\\\": \\\"13-VR P1\\\", \\\"Label\\\": \\\"temperature.systemboard.11\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 11, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 120, \\\"UpperThresholdCritical\\\": 115}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 13, \\\"LocationYmm\\\": 1}}, \\\"Name\\\": \\\"15-VR P1 Mem\\\", \\\"Label\\\": \\\"temperature.systemboard.12\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 12, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 120, \\\"UpperThresholdCritical\\\": 115}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 9, \\\"LocationYmm\\\": 1}}, \\\"Name\\\": \\\"16-VR P1 Mem\\\", \\\"Label\\\": \\\"temperature.systemboard.13\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 13, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 120, \\\"UpperThresholdCritical\\\": 115}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 8, \\\"LocationYmm\\\": 1}}, \\\"Name\\\": \\\"19-PS 1 Internal\\\", \\\"Label\\\": \\\"temperature.powersupply.14\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 14, \\\"PhysicalContext\\\": \\\"PowerSupply\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 1, \\\"LocationYmm\\\": 8}}, \\\"Name\\\": \\\"20-PS 2 Internal\\\", \\\"Label\\\": \\\"temperature.powersupply.15\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 15, \\\"PhysicalContext\\\": \\\"PowerSupply\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 5, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"21-PCI 1\\\", \\\"Label\\\": \\\"temperature.systemboard.16\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 16, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"22-PCI 2\\\", \\\"Label\\\": \\\"temperature.systemboard.17\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 17, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 13, \\\"LocationYmm\\\": 13}}, \\\"Name\\\": \\\"23-PCI 3\\\", \\\"Label\\\": \\\"temperature.systemboard.18\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 18, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"24-PCI 4\\\", \\\"Label\\\": \\\"temperature.systemboard.19\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 19, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 12, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"25-PCI 5\\\", \\\"Label\\\": \\\"temperature.systemboard.20\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 20, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 12, \\\"LocationYmm\\\": 13}}, \\\"Name\\\": \\\"26-PCI 6\\\", \\\"Label\\\": \\\"temperature.systemboard.21\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 21, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 8, \\\"LocationYmm\\\": 8}}, \\\"Name\\\": \\\"27-HD Controller\\\", \\\"Label\\\": \\\"temperature.systemboard.22\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 22, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 100}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 14, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"28-LOM Card\\\", \\\"Label\\\": \\\"temperature.systemboard.23\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 23, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 7, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"29-LOM\\\", \\\"Label\\\": \\\"temperature.systemboard.24\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 24, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 9, \\\"LocationYmm\\\": 0}}, \\\"Name\\\": \\\"30-Front Ambient\\\", \\\"Label\\\": \\\"temperature.intake.25\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 25, \\\"PhysicalContext\\\": \\\"Intake\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 65}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 5, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"31-PCI 1 Zone.\\\", \\\"Label\\\": \\\"temperature.systemboard.26\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 26, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 75, \\\"UpperThresholdCritical\\\": 70}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"32-PCI 2 Zone.\\\", \\\"Label\\\": \\\"temperature.systemboard.27\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 27, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 75, \\\"UpperThresholdCritical\\\": 70}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 13, \\\"LocationYmm\\\": 13}}, \\\"Name\\\": \\\"33-PCI 3 Zone.\\\", \\\"Label\\\": \\\"temperature.systemboard.28\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 28, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 75, \\\"UpperThresholdCritical\\\": 70}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 12}}, \\\"Name\\\": \\\"34-PCI 4 Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.29\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 29, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"35-PCI 5 Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.30\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 30, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 15}}, \\\"Name\\\": \\\"36-PCI 6 Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.31\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 31, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 11, \\\"LocationYmm\\\": 7}}, \\\"Name\\\": \\\"37-HD Cntlr Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.32\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 32, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 75}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 14, \\\"LocationYmm\\\": 11}}, \\\"Name\\\": \\\"38-I/O Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.33\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 33, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 80, \\\"UpperThresholdCritical\\\": 75}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 3, \\\"LocationYmm\\\": 7}}, \\\"Name\\\": \\\"39-P/S 2 Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.34\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 34, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 70}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 7, \\\"LocationYmm\\\": 10}}, \\\"Name\\\": \\\"40-Battery Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.35\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 35, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 80, \\\"UpperThresholdCritical\\\": 75}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 9, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"41-iLO Zone\\\", \\\"Label\\\": \\\"temperature.systemboard.36\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 36, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 95, \\\"UpperThresholdCritical\\\": 90}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 9, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"42-Rear HD Max\\\", \\\"Label\\\": \\\"temperature.systemboard.37\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 37, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 5, \\\"LocationYmm\\\": 1}}, \\\"Name\\\": \\\"43-Storage Batt\\\", \\\"Label\\\": \\\"temperature.systemboard.38\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 38, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 60}, {\\\"Oem\\\": {\\\"Hp\\\": {\\\"Type\\\": \\\"HpSeaOfSensors.1.0.0\\\", \\\"@odata.type\\\": \\\"#HpSeaOfSensors.1.0.0.HpSeaOfSensors\\\", \\\"LocationXmm\\\": 3, \\\"LocationYmm\\\": 14}}, \\\"Name\\\": \\\"44-Fuse\\\", \\\"Label\\\": \\\"temperature.systemboard.39\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 39, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": 0, \\\"UpperThresholdCritical\\\": 0}]}\",\"server_secure_boot_is_enabled\":false,\"server_chipset_name\":\"\",\"server_requires_reregister\":false,\"server_rack_name\":null,\"server_rack_position_upper_unit\":null,\"server_rack_position_lower_unit\":null,\"server_inventory_id\":null,\"server_is_in_diagnostics\":false,\"server_supports_sol\":false,\"server_supports_virtual_media\":false,\"server_supports_oob_provisioning\":false,\"server_cleanup_in_progress\":false,\"server_interfaces\":[{\"server_interface_mac_address\":\"38:63:bb:3b:ec:84\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"38:63:bb:3b:ec:85\",\"type\":\"ServerInterface\"}],\"server_disks\":[{\"server_disk_id\":751,\"server_disk_model\":\"LOGICAL VOLUME\",\"server_disk_size_gb\":931,\"server_id\":1749,\"server_disk_serial\":\"600508b1001cb3a2319d3b2df27e17a4\",\"server_disk_vendor\":\"HP\",\"server_disk_status\":\"installed\",\"server_disk_type\":\"HDD\",\"type\":\"ServerDisk\"},{\"server_disk_id\":752,\"server_disk_model\":\"LOGICAL VOLUME\",\"server_disk_size_gb\":931,\"server_id\":1749,\"server_disk_serial\":\"600508b1001c37682776483ee7137a02\",\"server_disk_vendor\":\"HP\",\"server_disk_status\":\"installed\",\"server_disk_type\":\"HDD\",\"type\":\"ServerDisk\"}],\"server_tags\":[],\"type\":\"Server\", \"server_rack_id\": null}"
