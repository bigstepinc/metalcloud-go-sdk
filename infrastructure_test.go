package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestInfrastructures(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": {"test":` + _infrastructureFixture + `},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.Infrastructures()
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	Expect((*ret)["test"].InfrastructureID).To(Equal(4103))

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

}

func TestInfrastructureGet(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": ` + _infrastructureFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.InfrastructureGet(100)
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

}

func TestInfrastructureGetWithLabel(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": ` + _infrastructureFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.InfrastructureGet("my-test")
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	params := (m["params"].([]interface{}))

	Expect(params[0]).To(Equal("my-test"))

}

func TestInfrastructureGetWithWrongLabel(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": ` + _infrastructureFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.InfrastructureGet("my_test")
	Expect(err).NotTo(BeNil())
	Expect(err.Error()).To(ContainSubstring("label"))
	Expect(ret).To(BeNil())

	//	(<-requestChan)

}

func TestInfrastructureDeploy(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	opts := ShutdownOptions{
		HardShutdownAfterTimeout:   true,
		SoftShutdownTimeoutSeconds: 181,
		AttemptSoftShutdown:        false,
	}
	err = mc.InfrastructureDeploy(100, opts, false, true)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	param := m["params"].([]interface{})

	Expect(param[0].(float64)).To(Equal(100.0))

	shutDownOpts := param[1].(map[string]interface{})

	Expect(shutDownOpts["soft_shutdown_timeout_seconds"]).To(Equal(float64(opts.SoftShutdownTimeoutSeconds)))
	Expect(shutDownOpts["hard_shutdown_after_timeout"]).To(Equal(opts.HardShutdownAfterTimeout))
	Expect(shutDownOpts["attempt_soft_shutdown"]).To(Equal(opts.AttemptSoftShutdown))

	Expect(param[2]).To(BeNil())
	Expect(param[3].(bool)).To(BeFalse())
	Expect(param[4].(bool)).To(BeTrue())
}

func TestInfrastructureDelete(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	err = mc.InfrastructureDelete(100)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_delete"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(100.0))

}

func TestInfrastructureRevert(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	err = mc.InfrastructureOperationCancel("test-asdasd")
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_operation_cancel"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("test-asdasd"))

}

const _infrastructureFixture = "{\"infrastructure_id\":4103,\"datacenter_name\":\"us-santaclara\",\"user_id_owner\":2,\"infrastructure_label\":\"demo\",\"infrastructure_created_timestamp\":\"2019-11-12T20:44:04Z\",\"infrastructure_subdomain\":\"demo.2.poc.metalcloud.io\",\"infrastructure_change_id\":8805,\"infrastructure_service_status\":\"active\",\"infrastructure_touch_unixtime\":\"1573829237.9229\",\"infrastructure_updated_timestamp\":\"2019-11-12T20:44:04Z\",\"infrastructure_gui_settings_json\":\"\",\"infrastructure_private_datacenters_json\":null,\"infrastructure_deploy_id\":10420,\"infrastructure_design_is_locked\":false,\"infrastructure_operation\":{\"infrastructure_change_id\":8805,\"infrastructure_id\":4103,\"datacenter_name\":\"us-santaclara\",\"user_id_owner\":2,\"infrastructure_label\":\"demo\",\"infrastructure_subdomain\":\"demo.2.poc.metalcloud.io\",\"infrastructure_deploy_type\":\"create\",\"infrastructure_deploy_status\":\"finished\",\"infrastructure_updated_timestamp\":\"2019-11-12T20:44:04Z\",\"infrastructure_gui_settings_json\":\"\",\"infrastructure_private_datacenters_json\":null,\"infrastructure_deploy_id\":10420,\"type\":\"InfrastructureOperation\",\"subnet_pool_lan\":null,\"infrastructure_reserved_lan_ip_ranges\":[]},\"type\":\"Infrastructure\",\"subnet_pool_lan\":null,\"infrastructure_reserved_lan_ip_ranges\":[],\"user_email_owner\":\"alex.bordei@bigstep.com\"}"
