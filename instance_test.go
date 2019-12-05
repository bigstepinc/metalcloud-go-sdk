package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestInstanceArrayUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var i Instance
	err := json.Unmarshal([]byte(_instanceFixture1), &i)
	Expect(err).To(BeNil())
	Expect(i).NotTo(BeNil())

	Expect(i.InstanceID).To(Equal(20639))
	Expect(i.InstanceCredentials.SSH.InitialPassword).To(Equal("asdasd"))
	Expect(i.InstanceInterfaces[0].InstanceInterfaceLabel).To(Equal("if0"))

	found := false
	for _, i := range i.InstanceInterfaces {
		for _, ip := range i.InstanceInterfaceIPs {
			if ip.IPHumanReadable == "172.17.106.22" {
				found = true
			}
		}
	}

	Expect(found).To(BeTrue())

}

func TestInstanceArrayInstances(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": {"test":` + _instanceFixture1 + `},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.instanceArrayInstances("test")
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())

	i := (*ret)["test"]

	Expect(i.InstanceID).To(Equal(20639))
	Expect(i.InstanceCredentials.SSH.InitialPassword).To(Equal("asdasd"))

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	params := (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("test"))

}

const _instanceFixture1 = "{\"instance_id\":20639,\"instance_array_id\":23739,\"server_id\":1,\"server_type_id\":1,\"instance_change_id\":48471,\"instance_service_status\":\"active\",\"drive_id_bootable\":12146,\"instance_label\":\"instance-20639\",\"instance_subdomain\":\"instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_subdomain_permanent\":\"instance-20639.poc.metalcloud.io\",\"instance_updated_timestamp\":\"2019-11-27T14:38:26Z\",\"instance_created_timestamp\":\"2019-11-27T14:38:25Z\",\"template_id_origin\":null,\"instance_operation\":{\"instance_change_id\":48471,\"instance_id\":20639,\"instance_array_id\":23739,\"server_id\":1,\"server_type_id\":1,\"instance_deploy_type\":\"create\",\"instance_deploy_status\":\"finished\",\"instance_label\":\"instance-20639\",\"instance_subdomain\":\"instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_updated_timestamp\":\"2019-11-27T14:38:26Z\",\"drive_id_bootable\":12146,\"template_id_origin\":null,\"type\":\"InstanceOperation\"},\"type\":\"Instance\",\"instance_interfaces\":[{\"instance_interface_id\":82706,\"network_id\":9013,\"instance_interface_index\":0,\"instance_id\":20639,\"instance_interface_change_id\":203393,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if0\",\"instance_interface_subdomain\":\"if0.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203393,\"instance_interface_id\":82706,\"network_id\":9013,\"instance_interface_index\":0,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if0.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if0\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"24:6e:96:6a:37:9a\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":82707,\"network_id\":9012,\"instance_interface_index\":1,\"instance_id\":20639,\"instance_interface_change_id\":203392,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if1\",\"instance_interface_subdomain\":\"if1.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203392,\"instance_interface_id\":82707,\"network_id\":9012,\"instance_interface_index\":1,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if1.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if1\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"a0:36:9f:ee:b2:f0\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[{\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91240,\"subnet_netmask_human_readable\":\"ffff:ffff:ffff:ffff:0000:0000:0000:0000\",\"subnet_gateway_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0001\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91240,\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59395\",\"ip_subdomain\":\"ip-59395.subnet-20504.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"},{\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91241,\"subnet_netmask_human_readable\":\"255.255.255.252\",\"subnet_gateway_human_readable\":\"172.17.106.21\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91241,\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59396\",\"ip_subdomain\":\"ip-59396.subnet-20505.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"}]},{\"instance_interface_id\":82708,\"network_id\":null,\"instance_interface_index\":2,\"instance_id\":20639,\"instance_interface_change_id\":203390,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if2\",\"instance_interface_subdomain\":\"if2.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203390,\"instance_interface_id\":82708,\"network_id\":null,\"instance_interface_index\":2,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if2.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if2\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"a0:36:9f:f0:ee:3a\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":82709,\"network_id\":null,\"instance_interface_index\":3,\"instance_id\":20639,\"instance_interface_change_id\":203391,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if3\",\"instance_interface_subdomain\":\"if3.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203391,\"instance_interface_id\":82709,\"network_id\":null,\"instance_interface_index\":3,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if3.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if3\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"a0:36:9f:f0:ee:6e\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]}],\"instance_credentials\":{\"ipmi\":{\"username\":null,\"initial_password\":null,\"ip_address\":null,\"version\":null,\"type\":\"IPMI\"},\"ilo\":{\"control_panel_url\":null,\"username\":null,\"initial_password\":null,\"type\":\"iLO\"},\"idrac\":{\"control_panel_url\":null,\"username\":null,\"initial_password\":null,\"type\":\"iDRAC\"},\"rdp\":{\"port\":null,\"username\":null,\"initial_password\":null,\"type\":\"RDP\"},\"ssh\":{\"port\":22,\"username\":\"root\",\"initial_password\":\"asdasd\",\"initial_ssh_keys\":{\"marius.boeru@bigstep.com\":[{\"user_ssh_key_id\":38,\"user_id\":1,\"user_ssh_key\":\"ssh-rsa AAAasdasdasd\",\"user_ssh_key_created_timestamp\":\"2019-10-17T14:14:52Z\",\"user_ssh_key_status\":\"active\",\"type\":\"SSHKey\"}]},\"type\":\"SSH\"},\"remote_console\":{\"remote_protocol\":\"ssh\",\"remote_control_panel_url\":\"?product=instance&id=20639\",\"tunnel_path_url\":\"https://us-santaclara-api.poc.metalcloud.io/remote-console/instance-tunnel\",\"type\":\"RemoteConsole\"},\"telnet\":null,\"iscsi\":{\"username\":\"asdasd\",\"password\":\"asdad\",\"initiator_iqn\":\"iqn.2019-11.com.bigstep.storage:instance-20639\",\"gateway\":\"100.64.0.1\",\"netmask\":\"255.255.255.248\",\"initiator_ip_address\":\"100.64.0.6\",\"type\":\"iSCSIInitiator\"},\"shared_drives\":[],\"ip_addresses_public\":[{\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91240,\"subnet_netmask_human_readable\":\"ffff:ffff:ffff:ffff:0000:0000:0000:0000\",\"subnet_gateway_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0001\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91240,\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59395\",\"ip_subdomain\":\"ip-59395.subnet-20504.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"},{\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91241,\"subnet_netmask_human_readable\":\"255.255.255.252\",\"subnet_gateway_human_readable\":\"172.17.106.21\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91241,\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59396\",\"ip_subdomain\":\"ip-59396.subnet-20505.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"}],\"ip_addresses_private\":[],\"type\":\"InstanceCredentials\"}}"
