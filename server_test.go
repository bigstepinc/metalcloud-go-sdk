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

}

func TestServerCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _serverFixture3 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
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

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
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

const _serverFixture1 = "{\"server_id\":310,\"agent_id\":44,\"datacenter_name\":\"es-madrid\",\"server_uuid\":\"44454C4C-5900-1033-8032-B9C04F434631\",\"server_serial_number\":\"9Y32CF1\",\"server_product_name\":\"PowerEdge 1950\",\"server_vendor\":\"Dell Inc.\",\"server_vendor_sku_id\":\"0\",\"server_ipmi_host\":\"10.255.237.28\",\"server_ipmi_internal_username\":\"bsii3Cu\",\"server_ipmi_internal_password_encrypted\":\"\",\"server_ipmi_version\":\"2\",\"server_ram_gbytes\":8,\"server_processor_count\":2,\"server_processor_core_mhz\":2333,\"server_processor_core_count\":4,\"server_processor_name\":\"Intel(R) Xeon(R) CPU           E5345  @ 2.33GHz\",\"server_processor_cpu_mark\":0,\"server_processor_threads\":1,\"server_type_id\":14,\"server_status\":\"available\",\"server_comments\":\"a\",\"server_details_xml\":null,\"server_network_total_capacity_mbps\":4000,\"server_ipmi_channel\":0,\"server_power_status\":\"off\",\"server_power_status_last_update_timestamp\":\"2020-08-19T13:05:04Z\",\"server_ilo_reset_timestamp\":\"0000-00-00T00:00:00Z\",\"server_boot_last_update_timestamp\":null,\"server_bdk_debug\":false,\"server_dhcp_status\":\"deny_requests\",\"server_bios_info_json\":\"{\\\"server_bios_vendor\\\":\\\"Dell Inc.\\\",\\\"server_bios_version\\\":\\\"2.7.0\\\"}\",\"server_vendor_info_json\":\"{\\\"management\\\":\\\"iDRAC\\\",\\\"version\\\":\\\"er] rpcRoundRobinConnectedAgentsOfType() failed with error: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n FetchError: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n\\\\n    at ClientRequest.<anonymous> (\\\\/var\\\\/datacenter-agents-binary-compiled-temp\\\\/Power\\\\/Power.portable.js:8:469877)\\\\n    at ClientRequest.emit (events.js:209:13)\\\\n    at TLSSocket.socketErrorListener (_http_client.js:406:9)\\\\n    at TLSSocket.emit (events.js:209:13)\\\\n    at errorOrDestroy (internal\\\\/streams\\\\/destroy.js:107:12)\\\\n    at onwriteError (_stream_writable.js:449:5)\\\\n    at onwrite (_stream_writable.js:470:5)\\\\n    at internal\\\\/streams\\\\/destroy.js:49:7\\\\n    at TLSSocket.Socket._destroy (net.js:595:3)\\\\n    at TLSSocket.destroy (internal\\\\/streams\\\\/destroy.js:37:8) Exception: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n FetchError: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n\\\\n    at ClientRequest.<anonymous> (\\\\/var\\\\/datacenter-agents-binary-compiled-temp\\\\/Power\\\\/Power.portable.js:8:469877)\\\\n    at ClientRequest.emit (events.js:209:13)\\\\n    at TLSSocket.socketErrorListener (_http_client.js:406:9)\\\\n    at TLSSocket.emit (events.js:209:13)\\\\n    at errorOrDestroy (internal\\\\/streams\\\\/destroy.js:107:12)\\\\n    at onwriteError (_stream_writable.js:449:5)\\\\n    at onwrite (_stream_writable.js:470:5)\\\\n    at internal\\\\/streams\\\\/destroy.js:49:7\\\\n    at TLSSocket.Socket._destroy (net.js:595:3)\\\\n    at TLSSocket.destroy (internal\\\\/streams\\\\/destroy.js:37:8)\\\\n    at \\\\/var\\\\/vhosts\\\\/bsiintegration.bigstepcloud.com\\\\/BSIWebSocketServer\\\\/node_modules\\\\/jsonrpc-bidirectional\\\\/src\\\\/Client.js:331:37\\\\n    at runMicrotasks (<anonymous>)\\\\n    at processTicksAndRejections (internal\\\\/process\\\\/task_queues.js:97:5) Exception: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n FetchError: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n\\\\n    at ClientRequest.<anonymous> (\\\\/var\\\\/datacenter-agents-binary-compiled-temp\\\\/Power\\\\/Power.portable.js:8:469877)\\\\n    at ClientRequest.emit (events.js:209:13)\\\\n    at TLSSocket.socketErrorListener (_http_client.js:406:9)\\\\n    at TLSSocket.emit (events.js:209:13)\\\\n    at errorOrDestroy (internal\\\\/streams\\\\/destroy.js:107:12)\\\\n    at onwriteError (_stream_writable.js:449:5)\\\\n    at onwrite (_stream_writable.js:470:5)\\\\n    at internal\\\\/streams\\\\/destroy.js:49:7\\\\n    at TLSSocket.Socket._destroy (net.js:595:3)\\\\n    at TLSSocket.destroy (internal\\\\/streams\\\\/destroy.js:37:8) Exception: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n FetchError: request to https:\\\\/\\\\/10.255.237.28\\\\/cgi-bin\\\\/webcgi\\\\/about failed, reason: write EPROTO 38858976:error:1425F102:SSL routines:ssl_choose_client_version:unsupported protocol:..\\\\/deps\\\\/openssl\\\\/openssl\\\\/ssl\\\\/statem\\\\/statem_lib.c:1922:\\\\n\\\\n    at ClientRequest.<anonymous> (\\\\/var\\\\/datacenter-agents-binary-compiled-temp\\\\/Power\\\\/Power.portable.js:8:469877)\\\\n    at ClientRequest.emit (events.js:209:13)\\\\n    at TLSSocket.socketErrorListener (_http_client.js:406:9)\\\\n    at TLSSocket.emit (events.js:209:13)\\\\n    at errorOrDestroy (internal\\\\/streams\\\\/destroy.js:107:12)\\\\n    at onwriteError (_stream_writable.js:449:5)\\\\n    at onwrite (_stream_writable.js:470:5)\\\\n    at internal\\\\/streams\\\\/destroy.js:49:7\\\\n    at TLSSocket.Socket._destroy (net.js:595:3)\\\\n    at TLSSocket.destroy (internal\\\\/streams\\\\/destroy.js:37:8)\\\\n    at \\\\/var\\\\/vhosts\\\\/bsiintegration.bigstepcloud.com\\\\/BSIWebSocketServer\\\\/node_modules\\\\/jsonrpc-bidirectional\\\\/src\\\\/Client.js:331:37\\\\n    at runMicrotasks (<anonymous>)\\\\n    at processTicksAndRejections (internal\\\\/process\\\\/task_queues.js:97:5)\\\\n    at \\\\/var\\\\/vhosts\\\\/bsiintegration.bigstepcloud.com\\\\/BSIWebSocketServer\\\\/node_modules\\\\/jsonrpc-bidirectional\\\\/src\\\\/Client.js:331:37\\\\n    at runMicrotasks (<anonymous>)\\\\n    at processTicksAndRejections (internal\\\\/process\\\\/tas\\\"}\",\"server_class\":\"bigdata\",\"server_created_timestamp\":\"2019-07-02T07:57:19Z\",\"subnet_oob_id\":2,\"subnet_oob_index\":28,\"server_boot_type\":\"classic\",\"server_disk_wipe\":true,\"server_disk_count\":0,\"server_disk_size_mbytes\":0,\"server_disk_type\":\"none\",\"server_requires_manual_cleaning\":false,\"chassis_rack_id\":null,\"server_custom_json\":\"\",\"server_instance_custom_json\":null,\"server_last_cleanup_start\":\"2020-08-12T14:26:47Z\",\"server_allocation_timestamp\":null,\"server_dhcp_packet_sniffing_is_enabled\":true,\"snmp_community_password_dcencrypted\":null,\"server_mgmt_snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_port\":161,\"server_mgmt_snmp_version\":2,\"server_dhcp_relay_security_is_enabled\":true,\"server_keys_json\":\"\",\"server_info_json\":null,\"server_ipmi_credentials_need_update\":false,\"server_gpu_count\":0,\"server_gpu_vendor\":\"\",\"server_gpu_model\":\"\",\"server_bmc_mac_address\":null,\"server_metrics_metadata_json\":null,\"server_interfaces\":[{\"server_interface_mac_address\":\"00:1d:09:64:f0:2b\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:1d:09:64:f0:2d\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:15:17:c0:4c:e6\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:15:17:c0:4c:e7\",\"type\":\"ServerInterface\"}],\"server_disks\":[],\"server_tags\":[],\"type\":\"Server\"}"
const _serverFixture2 = "{\"server_id\":685,\"agent_id\":137,\"datacenter_name\":\"es-madrid\",\"server_uuid\":\"44454C4C-5800-1039-8052-C2C04F373632\",\"server_serial_number\":\"BX9R762\",\"server_product_name\":\"PowerEdge R730 (SKU=NotProvided;ModelName=PowerEdge R730)\",\"server_vendor\":\"Dell Inc.\",\"server_vendor_sku_id\":\"SKU=NotProvided;ModelName=PowerEdge R730\",\"server_ipmi_host\":\"10.255.237.40\",\"server_ipmi_internal_username\":\"\",\"server_ipmi_internal_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_ipmi_version\":\"2\",\"server_ram_gbytes\":32,\"server_processor_count\":2,\"server_processor_core_mhz\":1295,\"server_processor_core_count\":8,\"server_processor_name\":\"Intel(R) Xeon(R) CPU E5-2630 v3 @ 2.40GHz\",\"server_processor_cpu_mark\":0,\"server_processor_threads\":2,\"server_type_id\":33,\"server_status\":\"available\",\"server_comments\":null,\"server_details_xml\":null,\"server_network_total_capacity_mbps\":4000,\"server_ipmi_channel\":0,\"server_power_status\":\"off\",\"server_power_status_last_update_timestamp\":\"2020-08-19T13:05:04Z\",\"server_ilo_reset_timestamp\":\"2020-07-22T08:26:38Z\",\"server_boot_last_update_timestamp\":null,\"server_bdk_debug\":false,\"server_dhcp_status\":\"deny_requests\",\"server_bios_info_json\":\"\",\"server_vendor_info_json\":\"\",\"server_class\":\"bigdata\",\"server_created_timestamp\":\"2019-10-08T10:30:48Z\",\"subnet_oob_id\":2,\"subnet_oob_index\":21,\"server_boot_type\":\"uefi\",\"server_disk_wipe\":true,\"server_disk_count\":1,\"server_disk_size_mbytes\":285696,\"server_disk_type\":\"HDD\",\"server_requires_manual_cleaning\":false,\"chassis_rack_id\":null,\"server_custom_json\":\"\",\"server_instance_custom_json\":null,\"server_last_cleanup_start\":\"2020-07-22T08:28:45Z\",\"server_allocation_timestamp\":null,\"server_dhcp_packet_sniffing_is_enabled\":true,\"snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_community_password_dcencrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"server_mgmt_snmp_port\":161,\"server_mgmt_snmp_version\":2,\"server_dhcp_relay_security_is_enabled\":true,\"server_keys_json\":\"\",\"server_info_json\":null,\"server_ipmi_credentials_need_update\":false,\"server_gpu_count\":0,\"server_gpu_vendor\":\"\",\"server_gpu_model\":\"\",\"server_bmc_mac_address\":null,\"server_metrics_metadata_json\":\"{\\\"Fans\\\": [{\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan1\\\", \\\"Label\\\": \\\"fan.systemboard.1\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 1, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan2\\\", \\\"Label\\\": \\\"fan.systemboard.2\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 2, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan3\\\", \\\"Label\\\": \\\"fan.systemboard.3\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 3, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan4\\\", \\\"Label\\\": \\\"fan.systemboard.4\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 4, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan5\\\", \\\"Label\\\": \\\"fan.systemboard.5\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 5, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Fan6\\\", \\\"Label\\\": \\\"fan.systemboard.6\\\", \\\"Units\\\": \\\"RPM\\\", \\\"Number\\\": 6, \\\"PhysicalContext\\\": \\\"SystemBoard\\\", \\\"UpperThresholdFatal\\\": null, \\\"UpperThresholdCritical\\\": null}], \\\"Temperatures\\\": [{\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Inlet Temp\\\", \\\"Label\\\": \\\"temperature.intake.4\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 4, \\\"PhysicalContext\\\": \\\"Intake\\\", \\\"UpperThresholdFatal\\\": 47, \\\"UpperThresholdCritical\\\": 47}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"System Board Exhaust Temp\\\", \\\"Label\\\": \\\"temperature.exhaust.1\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 1, \\\"PhysicalContext\\\": \\\"Exhaust\\\", \\\"UpperThresholdFatal\\\": 75, \\\"UpperThresholdCritical\\\": 75}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"CPU1 Temp\\\", \\\"Label\\\": \\\"temperature.cpu.14\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 14, \\\"PhysicalContext\\\": \\\"CPU\\\", \\\"UpperThresholdFatal\\\": 87, \\\"UpperThresholdCritical\\\": 87}, {\\\"Oem\\\": null, \\\"Name\\\": \\\"CPU2 Temp\\\", \\\"Label\\\": \\\"temperature.cpu.15\\\", \\\"Units\\\": \\\"Celsius\\\", \\\"Number\\\": 15, \\\"PhysicalContext\\\": \\\"CPU\\\", \\\"UpperThresholdFatal\\\": 87, \\\"UpperThresholdCritical\\\": 87}]}\",\"server_interfaces\":[{\"server_interface_mac_address\":\"44:a8:42:24:93:4e\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"44:a8:42:24:93:50\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:1b:21:72:f2:74\",\"type\":\"ServerInterface\"},{\"server_interface_mac_address\":\"00:1b:21:72:f2:75\",\"type\":\"ServerInterface\"}],\"server_disks\":[{\"server_disk_id\":226,\"server_disk_model\":\"PERC H730 Mini\",\"server_disk_size_gb\":279,\"server_id\":685,\"server_disk_serial\":\"644a84203323c500265e82e96590c769\",\"server_disk_vendor\":\"DELL\",\"server_disk_status\":\"installed\",\"server_disk_type\":\"HDD\",\"type\":\"ServerDisk\"}],\"server_tags\":[\"0\",\"1\",\"2\"],\"type\":\"Server\"}"
const _serverFixture3 = "{\"server_id\": 100}"
