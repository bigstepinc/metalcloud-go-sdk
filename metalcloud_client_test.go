package metalcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

// needed to retrieve requests that arrived at httpServer for further investigation
var requestChan = make(chan *RequestData, 1)

// the request datastructure that can be retrieved for test assertions
type RequestData struct {
	request *http.Request
	body    string
}

// set the response body the httpServer should return for the next request
var responseBody = ""

var httpServer *httptest.Server

// start the testhttp server and stop it when tests are finished
func TestMain(m *testing.M) {
	httpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)

		defer r.Body.Close()
		// put request and body to channel for the client to investigate them
		requestChan <- &RequestData{r, string(data)}

		fmt.Fprintf(w, responseBody)
	}))
	defer httpServer.Close()

	os.Exit(m.Run())
}

func TestSignature(t *testing.T) {
	RegisterTestingT(t)

	content := strings.NewReader("asldklk234mlk234asd")

	request, err := http.NewRequest("GET", httpServer.URL, content)
	Expect(request).NotTo(BeNil())
	Expect(err).To(BeNil())

	transport := &signatureAdderRoundTripper{
		APIKey:         "asdasdasd",
		LoggingEnabled: false,
		DryRun:         false,
	}

	_, err = transport.RoundTrip(request)
	Expect(err).To(BeNil())

	requestWithSignature := (<-requestChan).request

	gotSignature := (requestWithSignature.URL).Query().Get("verify")
	expectedSignature := "b8287028c41c7dee8e6f0479ff05dd71"

	Expect(gotSignature).To(Equal(expectedSignature))
}

func TestEmptyListReply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret1, err1 := mc.InstanceArrays(100)
	Expect(err1).To(BeNil())
	Expect(ret1).NotTo(BeNil())

	<-requestChan

	ret2, err2 := mc.Infrastructures()
	Expect(err2).To(BeNil())
	Expect(ret2).NotTo(BeNil())

}

func TestInstanceArrayCreateOmitEmpty(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result":{"instance_array_label":"test"},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())
	Expect(mc).NotTo(BeNil())

	instanceArray := InstanceArray{
		InstanceArrayLabel: "test",
	}

	<-requestChan

	ret1, err1 := mc.InstanceArrayCreate(100, instanceArray)
	Expect(err1).To(BeNil())
	Expect(ret1).NotTo(BeNil())

	body := (<-requestChan).body

	//fmt.Printf("body=%s\n", body)

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	param := m["params"].([]interface{})[1].(map[string]interface{})

	Expect(param["instance_array_label"]).To(Equal("test"))
	Expect(param["volume_template_id"]).To(BeNil())

}

func testEmptyObjMarshalingToBeEmpty(obj interface{}) {
	s, err := json.MarshalIndent(obj, "", "\t")
	Expect(err).To(BeNil())
	Expect(string(s)).To(Equal("{}"))
}

func TestResourcesOmitEmptyMarshaling(t *testing.T) {
	RegisterTestingT(t)

	obj := FirewallRule{}
	testEmptyObjMarshalingToBeEmpty(obj)

	obj2 := InstanceArray{}
	testEmptyObjMarshalingToBeEmpty(obj2)

	obj3 := InstanceArrayOperation{}
	testEmptyObjMarshalingToBeEmpty(obj3)

	obj4 := DriveArray{}
	testEmptyObjMarshalingToBeEmpty(obj4)

	obj5 := DriveArrayOperation{}
	testEmptyObjMarshalingToBeEmpty(obj5)

	obj6 := ServerType{}
	testEmptyObjMarshalingToBeEmpty(obj6)

}

func TestInstanceArraylUnmarshalTest(t *testing.T) {

	RegisterTestingT(t)

	iaStr := "{\"instance_array_id\":35516,\"instance_array_instance_count\":2,\"instance_array_ipv4_subnet_create_auto\":true,\"instance_array_ip_allocate_auto\":true,\"instance_array_ram_gbytes\":1,\"instance_array_processor_count\":1,\"instance_array_processor_core_mhz\":1000,\"instance_array_processor_core_count\":1,\"infrastructure_id\":25524,\"instance_array_service_status\":\"active\",\"instance_array_change_id\":215807,\"instance_array_label\":\"workers\",\"instance_array_subdomain\":\"workers.vanilla.complex-demo.7.bigstep.io\",\"drive_array_id_boot\":45928,\"instance_array_gui_settings_json\":\"{\\\"nRowIndex\\\":0,\\\"nColumnIndex\\\":3,\\\"bShowWidgetChildren\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.6337124950169671\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/13.0.2 Safari\\\\/605.1.15\\\"}\",\"instance_array_updated_timestamp\":\"2019-10-11T07:18:57Z\",\"instance_array_created_timestamp\":\"2019-03-28T15:23:18Z\",\"cluster_id\":40559,\"cluster_role_group\":\"none\",\"instance_array_disk_count\":0,\"instance_array_disk_size_mbytes\":0,\"instance_array_firewall_managed\":true,\"volume_template_id\":null,\"instance_array_boot_method\":\"pxe_iscsi\",\"instance_array_virtual_interfaces_enabled\":false,\"instance_array_operation\":{\"instance_array_change_id\":215807,\"instance_array_id\":35516,\"instance_array_instance_count\":2,\"instance_array_ipv4_subnet_create_auto\":true,\"instance_array_ip_allocate_auto\":true,\"instance_array_ram_gbytes\":1,\"instance_array_processor_count\":1,\"instance_array_processor_core_mhz\":1000,\"instance_array_processor_core_count\":1,\"instance_array_deploy_type\":\"edit\",\"instance_array_deploy_status\":\"finished\",\"instance_array_label\":\"workers\",\"instance_array_subdomain\":\"workers.vanilla.complex-demo.7.bigstep.io\",\"drive_array_id_boot\":45928,\"instance_array_gui_settings_json\":\"{\\\"nRowIndex\\\":0,\\\"nColumnIndex\\\":3,\\\"bShowWidgetChildren\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.6337124950169671\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/13.0.2 Safari\\\\/605.1.15\\\"}\",\"instance_array_updated_timestamp\":\"2019-10-11T07:18:57Z\",\"instance_array_disk_count\":0,\"instance_array_disk_size_mbytes\":0,\"instance_array_firewall_managed\":true,\"volume_template_id\":null,\"instance_array_boot_method\":\"pxe_iscsi\",\"instance_array_virtual_interfaces_enabled\":false,\"type\":\"InstanceArrayOperation\",\"instance_array_disk_types\":[],\"instance_array_firewall_rules\":[{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Allow SSH traffic.\",\"firewall_rule_port_range_end\":22,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":22,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Allow RDP traffic.\",\"firewall_rule_port_range_end\":3389,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":3389,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"udp\",\"firewall_rule_description\":\"Allow RDP traffic.\",\"firewall_rule_port_range_end\":3389,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":3389,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"all\",\"firewall_rule_description\":\"Allow traffic on 10.0.0.0/8.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":\"10.255.255.255\",\"firewall_rule_destination_ip_address_range_start\":\"10.0.0.0\"},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"all\",\"firewall_rule_description\":\"Allow traffic on 172.16.0.0/12.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":\"172.31.255.255\",\"firewall_rule_destination_ip_address_range_start\":\"172.16.0.0\"},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"all\",\"firewall_rule_description\":\"Allow traffic on 192.168.0.0/16.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":\"192.168.255.255\",\"firewall_rule_destination_ip_address_range_start\":\"192.168.0.0\"},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"icmp\",\"firewall_rule_description\":\"Allow IPv4 ICMP traffic.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"icmp\",\"firewall_rule_description\":\"Allow IPv6 ICMP traffic.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv6\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"udp\",\"firewall_rule_description\":\"Allow SNMP traffic.\",\"firewall_rule_port_range_end\":161,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":161,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Rule description\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":\"89.36.24.2\",\"firewall_rule_source_ip_address_range_start\":\"89.36.24.2\",\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Rule description\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":\"5.13.86.205\",\"firewall_rule_source_ip_address_range_start\":\"5.13.86.205\",\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null}],\"instance_array_interfaces\":[{\"instance_array_interface_change_id\":633555,\"instance_array_interface_id\":139705,\"network_id\":58438,\"instance_array_interface_index\":0,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:19Z\",\"instance_array_interface_subdomain\":\"if0.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if0\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},{\"instance_array_interface_change_id\":633554,\"instance_array_interface_id\":139706,\"network_id\":58437,\"instance_array_interface_index\":1,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_subdomain\":\"if1.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if1\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},{\"instance_array_interface_change_id\":633552,\"instance_array_interface_id\":139707,\"network_id\":null,\"instance_array_interface_index\":2,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_subdomain\":\"if2.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if2\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},{\"instance_array_interface_change_id\":633562,\"instance_array_interface_id\":139708,\"network_id\":58439,\"instance_array_interface_index\":3,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:25:10Z\",\"instance_array_interface_subdomain\":\"if3.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if3\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]}]},\"type\":\"InstanceArray\",\"instance_array_disk_types\":[],\"instance_array_firewall_rules\":[{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Allow SSH traffic.\",\"firewall_rule_port_range_end\":22,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":22,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Allow RDP traffic.\",\"firewall_rule_port_range_end\":3389,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":3389,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"udp\",\"firewall_rule_description\":\"Allow RDP traffic.\",\"firewall_rule_port_range_end\":3389,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":3389,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"all\",\"firewall_rule_description\":\"Allow traffic on 10.0.0.0/8.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":\"10.255.255.255\",\"firewall_rule_destination_ip_address_range_start\":\"10.0.0.0\"},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"all\",\"firewall_rule_description\":\"Allow traffic on 172.16.0.0/12.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":\"172.31.255.255\",\"firewall_rule_destination_ip_address_range_start\":\"172.16.0.0\"},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"all\",\"firewall_rule_description\":\"Allow traffic on 192.168.0.0/16.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":\"192.168.255.255\",\"firewall_rule_destination_ip_address_range_start\":\"192.168.0.0\"},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"icmp\",\"firewall_rule_description\":\"Allow IPv4 ICMP traffic.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"icmp\",\"firewall_rule_description\":\"Allow IPv6 ICMP traffic.\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv6\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":false,\"firewall_rule_protocol\":\"udp\",\"firewall_rule_description\":\"Allow SNMP traffic.\",\"firewall_rule_port_range_end\":161,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":161,\"firewall_rule_source_ip_address_range_end\":null,\"firewall_rule_source_ip_address_range_start\":null,\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Rule description\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":\"89.36.24.2\",\"firewall_rule_source_ip_address_range_start\":\"89.36.24.2\",\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null},{\"type\":\"FirewallRule\",\"firewall_rule_enabled\":true,\"firewall_rule_protocol\":\"tcp\",\"firewall_rule_description\":\"Rule description\",\"firewall_rule_port_range_end\":null,\"firewall_rule_ip_address_type\":\"ipv4\",\"firewall_rule_port_range_start\":null,\"firewall_rule_source_ip_address_range_end\":\"5.13.86.205\",\"firewall_rule_source_ip_address_range_start\":\"5.13.86.205\",\"firewall_rule_destination_ip_address_range_end\":null,\"firewall_rule_destination_ip_address_range_start\":null}],\"instance_array_interfaces\":[{\"instance_array_interface_id\":139705,\"network_id\":58438,\"instance_array_interface_index\":0,\"instance_array_id\":35516,\"instance_array_interface_change_id\":633555,\"instance_array_interface_service_status\":\"active\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:19Z\",\"instance_array_interface_created_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_label\":\"if0\",\"instance_array_interface_subdomain\":\"if0.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_operation\":{\"instance_array_interface_change_id\":633555,\"instance_array_interface_id\":139705,\"network_id\":58438,\"instance_array_interface_index\":0,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:19Z\",\"instance_array_interface_subdomain\":\"if0.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if0\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},\"type\":\"InstanceArrayInterface\",\"instance_array_interface_lagg_indexes\":[]},{\"instance_array_interface_id\":139706,\"network_id\":58437,\"instance_array_interface_index\":1,\"instance_array_id\":35516,\"instance_array_interface_change_id\":633554,\"instance_array_interface_service_status\":\"active\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_created_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_label\":\"if1\",\"instance_array_interface_subdomain\":\"if1.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_operation\":{\"instance_array_interface_change_id\":633554,\"instance_array_interface_id\":139706,\"network_id\":58437,\"instance_array_interface_index\":1,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_subdomain\":\"if1.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if1\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},\"type\":\"InstanceArrayInterface\",\"instance_array_interface_lagg_indexes\":[]},{\"instance_array_interface_id\":139707,\"network_id\":null,\"instance_array_interface_index\":2,\"instance_array_id\":35516,\"instance_array_interface_change_id\":633552,\"instance_array_interface_service_status\":\"active\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_created_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_label\":\"if2\",\"instance_array_interface_subdomain\":\"if2.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_operation\":{\"instance_array_interface_change_id\":633552,\"instance_array_interface_id\":139707,\"network_id\":null,\"instance_array_interface_index\":2,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_subdomain\":\"if2.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if2\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},\"type\":\"InstanceArrayInterface\",\"instance_array_interface_lagg_indexes\":[]},{\"instance_array_interface_id\":139708,\"network_id\":58439,\"instance_array_interface_index\":3,\"instance_array_id\":35516,\"instance_array_interface_change_id\":633562,\"instance_array_interface_service_status\":\"active\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:25:10Z\",\"instance_array_interface_created_timestamp\":\"2019-03-28T15:23:18Z\",\"instance_array_interface_label\":\"if3\",\"instance_array_interface_subdomain\":\"if3.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_operation\":{\"instance_array_interface_change_id\":633562,\"instance_array_interface_id\":139708,\"network_id\":58439,\"instance_array_interface_index\":3,\"instance_array_id\":35516,\"instance_array_interface_deploy_type\":\"create\",\"instance_array_interface_deploy_status\":\"finished\",\"instance_array_interface_updated_timestamp\":\"2019-03-28T15:25:10Z\",\"instance_array_interface_subdomain\":\"if3.frontend.vanilla.complex-demo.7.bigstep.io\",\"instance_array_interface_label\":\"if3\",\"type\":\"InstanceArrayInterfaceOperation\",\"instance_array_interface_lagg_indexes\":[]},\"type\":\"InstanceArrayInterface\",\"instance_array_interface_lagg_indexes\":[]}]}"

	var ia InstanceArray

	err2 := json.Unmarshal([]byte(iaStr), &ia)

	Expect(err2).To(BeNil())
	Expect(ia).NotTo(BeNil())
	Expect(ia.InstanceArrayID).To(Equal(35516))
	Expect(ia.InstanceArrayLabel).To(Equal("workers"))
	Expect(ia.VolumeTemplateID).To(Equal(0))
	Expect(ia.InstanceArrayInstanceCount).To(Equal(2))
	Expect(ia.InstanceArrayProcessorCount).To(Equal(1))
}
