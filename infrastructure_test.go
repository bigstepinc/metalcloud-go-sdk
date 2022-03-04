package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestInfrastructures(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": {"test":` + _infrastructureFixture + `},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
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

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
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

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.InfrastructureGetByLabel("my-test")
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

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.InfrastructureGetByLabel("my_test")
	Expect(err).NotTo(BeNil())
	Expect(err.Error()).To(ContainSubstring("label"))
	Expect(ret).To(BeNil())

	//	(<-requestChan)

}

func TestInfrastructureDeploy(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
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

func TestInfrastructureDeployWithOptions(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	shOpts := ShutdownOptions{
		HardShutdownAfterTimeout:   true,
		SoftShutdownTimeoutSeconds: 181,
		AttemptSoftShutdown:        false,
	}

	dpOpts := DeployOptions{
		InstanceArrayMapping: map[int]map[string]DeployOptionsServerTypeMappingObject{
			1023: {
				"22": {
					ServerCount: 1,
					ServerIDs: []int{
						1,
						20,
					},
				},
			},
		},
	}

	err = mc.InfrastructureDeployWithOptions(100, shOpts, &dpOpts, false, true)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	param := m["params"].([]interface{})

	Expect(param[0].(float64)).To(Equal(100.0))

	shutDownOpts := param[1].(map[string]interface{})

	Expect(shutDownOpts["soft_shutdown_timeout_seconds"]).To(Equal(float64(shOpts.SoftShutdownTimeoutSeconds)))
	Expect(shutDownOpts["hard_shutdown_after_timeout"]).To(Equal(shOpts.HardShutdownAfterTimeout))
	Expect(shutDownOpts["attempt_soft_shutdown"]).To(Equal(shOpts.AttemptSoftShutdown))

	deployOpts := param[2].(map[string]interface{})

	iaOpts := deployOpts["instance_array"].(map[string]interface{})
	iaOpts2 := iaOpts["1023"].(map[string]interface{})
	stOpts := iaOpts2["22"].(map[string]interface{})

	Expect(stOpts["server_count"]).To(Equal(1.0))

	stiOpts := stOpts["server_ids"].([]interface{})

	Expect(stiOpts[0]).To(Equal(1.0))
	Expect(stiOpts[1]).To(Equal(20.0))

	Expect(param[3].(bool)).To(BeFalse())
	Expect(param[4].(bool)).To(BeTrue())
}

func TestInfrastructureDelete(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": [],"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
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

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	err = mc.InfrastructureOperationCancelByLabel("test-asdasd")
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

func TestInfrastructureCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _infrastructureFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Infrastructure{
		InfrastructureID:               4103,
		DatacenterName:                 "us-santaclara",
		UserIDowner:                    2,
		InfrastructureLabel:            "demo",
		InfrastructureCreatedTimestamp: "2019-11-12T20:44:04Z",
		InfrastructureSubdomain:        "demo.2.poc.metalcloud.io",
		InfrastructureChangeID:         8805,
		InfrastructureServiceStatus:    "active",
		InfrastructureTouchUnixtime:    "1573829237.9229",
		InfrastructureUpdatedTimestamp: "2019-11-12T20:44:04Z",
		InfrastructureDeployID:         10420,
		InfrastructureDesignIsLocked:   false,
		InfrastructureOperation: InfrastructureOperation{
			InfrastructureChangeID:         8805,
			InfrastructureID:               4103,
			DatacenterName:                 "us-santaclara",
			UserIDOwner:                    2,
			InfrastructureLabel:            "demo",
			InfrastructureSubdomain:        "demo.2.poc.metalcloud.io",
			InfrastructureDeployType:       "create",
			InfrastructureDeployStatus:     "finished",
			InfrastructureUpdatedTimestamp: "2019-11-12T20:44:04Z",
		},
	}
	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(4103.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_edit"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(4103.0))

	responseBody = `{"error": {"message": "Infrastructure not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(4103.0))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_create"))

	params = (m["params"].([]interface{}))

	Expect(params[1].(map[string]interface{})["datacenter_name"].(string)).To(Equal("us-santaclara"))
}

func TestInfrastructureDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _infrastructureFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Infrastructure{
		InfrastructureID:               4103,
		DatacenterName:                 "us-santaclara",
		UserIDowner:                    2,
		InfrastructureLabel:            "demo",
		InfrastructureCreatedTimestamp: "2019-11-12T20:44:04Z",
		InfrastructureSubdomain:        "demo.2.poc.metalcloud.io",
		InfrastructureChangeID:         8805,
		InfrastructureServiceStatus:    "active",
		InfrastructureTouchUnixtime:    "1573829237.9229",
		InfrastructureUpdatedTimestamp: "2019-11-12T20:44:04Z",
		InfrastructureDeployID:         10420,
		InfrastructureDesignIsLocked:   false,
		InfrastructureOperation: InfrastructureOperation{
			InfrastructureChangeID:         8805,
			InfrastructureID:               4103,
			DatacenterName:                 "us-santaclara",
			UserIDOwner:                    2,
			InfrastructureLabel:            "demo",
			InfrastructureSubdomain:        "demo.2.poc.metalcloud.io",
			InfrastructureDeployType:       "create",
			InfrastructureDeployStatus:     "finished",
			InfrastructureUpdatedTimestamp: "2019-11-12T20:44:04Z",
		},
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	var m map[string]interface{}
	body := (<-requestChan).body

	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("demo"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("infrastructure_delete"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(4103.0))
}

func TestInfrastructureSearch(t *testing.T) {
	RegisterTestingT(t)

	responseBody = _infrastructureSearchFixture1

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.InfrastructureSearch("*")
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	r := *ret
	Expect(r[0].InfrastructureID).To(Equal(346))
	Expect(r[1].InfrastructureID).To(Equal(341))

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

}

const _infrastructureFixture = "{\"infrastructure_id\":4103,\"datacenter_name\":\"us-santaclara\",\"user_id_owner\":2,\"infrastructure_label\":\"demo\",\"infrastructure_created_timestamp\":\"2019-11-12T20:44:04Z\",\"infrastructure_subdomain\":\"demo.2.poc.metalcloud.io\",\"infrastructure_change_id\":8805,\"infrastructure_service_status\":\"active\",\"infrastructure_touch_unixtime\":\"1573829237.9229\",\"infrastructure_updated_timestamp\":\"2019-11-12T20:44:04Z\",\"infrastructure_gui_settings_json\":\"\",\"infrastructure_private_datacenters_json\":null,\"infrastructure_deploy_id\":10420,\"infrastructure_design_is_locked\":false,\"infrastructure_operation\":{\"infrastructure_change_id\":8805,\"infrastructure_id\":4103,\"datacenter_name\":\"us-santaclara\",\"user_id_owner\":2,\"infrastructure_label\":\"demo\",\"infrastructure_subdomain\":\"demo.2.poc.metalcloud.io\",\"infrastructure_deploy_type\":\"create\",\"infrastructure_deploy_status\":\"finished\",\"infrastructure_updated_timestamp\":\"2019-11-12T20:44:04Z\",\"infrastructure_gui_settings_json\":\"\",\"infrastructure_private_datacenters_json\":null,\"infrastructure_deploy_id\":10420,\"type\":\"InfrastructureOperation\",\"subnet_pool_lan\":null,\"infrastructure_reserved_lan_ip_ranges\":[]},\"type\":\"Infrastructure\",\"subnet_pool_lan\":null,\"infrastructure_reserved_lan_ip_ranges\":[],\"user_email_owner\":\"ap.com\"}"
const _infrastructureSearchFixture1 = "{\"result\":{\"_user_infrastructures_extended\":{\"duration_milliseconds\":0.5091230869293213,\"rows\":[{\"infrastructure_id\":346,\"infrastructure_subdomain\":\"my-infrastructure.57.us01.metalsoft.io\",\"infrastructure_service_status\":\"active\",\"infrastructure_deploy_status\":\"finished\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-04T08:39:18Z\",\"infrastructure_updated_timestamp\":\"2022-03-04T08:43:00Z\",\"user_email\":[\"viktor.dddd.com\"],\"user_id_owner\":57,\"infrastructure_deploy_id\":342,\"afc_group_created_timestamp\":\"2022-03-04T09:21:11Z\",\"afc_group_finished_timestamp\":\"2022-03-04T09:24:35Z\",\"thrownError\":0,\"executedSuccess\":9,\"total\":9,\"infrastructure_change_id\":656,\"infrastructure_user_id\":[283]},{\"infrastructure_id\":341,\"infrastructure_subdomain\":\"my-infrastructure.6.us01.metalsoft.io\",\"infrastructure_service_status\":\"deleted\",\"infrastructure_deploy_status\":\"finished\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-02T08:35:15Z\",\"infrastructure_updated_timestamp\":\"2022-03-02T09:34:54Z\",\"user_email\":[\"aure.com\"],\"user_id_owner\":6,\"infrastructure_deploy_id\":334,\"afc_group_created_timestamp\":\"2022-03-02T09:35:02Z\",\"afc_group_finished_timestamp\":\"2022-03-02T09:37:54Z\",\"thrownError\":0,\"executedSuccess\":20,\"total\":20,\"infrastructure_change_id\":649,\"infrastructure_user_id\":[278]},{\"infrastructure_id\":345,\"infrastructure_subdomain\":\"my-infrastructure.55.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-04T07:58:35Z\",\"infrastructure_updated_timestamp\":\"2022-03-04T07:58:35Z\",\"user_email\":[\"kchu.com\"],\"user_id_owner\":55,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":653,\"infrastructure_user_id\":[282]},{\"infrastructure_id\":344,\"infrastructure_subdomain\":\"my-infrastructure.3.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-03T07:51:36Z\",\"infrastructure_updated_timestamp\":\"2022-03-03T07:51:36Z\",\"user_email\":[\"adrian..com\"],\"user_id_owner\":3,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":652,\"infrastructure_user_id\":[281]},{\"infrastructure_id\":343,\"infrastructure_subdomain\":\"my-infrastructure.59.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-02T12:21:59Z\",\"infrastructure_updated_timestamp\":\"2022-03-02T12:21:59Z\",\"user_email\":[\"darhoo.com\"],\"user_id_owner\":59,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":651,\"infrastructure_user_id\":[280]},{\"infrastructure_id\":342,\"infrastructure_subdomain\":\"my-infrastructure.6.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-02T09:35:26Z\",\"infrastructure_updated_timestamp\":\"2022-03-02T09:35:26Z\",\"user_email\":[\"aurelian.agstep.com\"],\"user_id_owner\":6,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":650,\"infrastructure_user_id\":[279]},{\"infrastructure_id\":340,\"infrastructure_subdomain\":\"my-infrastructure.2.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-03-01T08:30:03Z\",\"infrastructure_updated_timestamp\":\"2022-03-01T08:30:03Z\",\"user_email\":[\"matep.com\"],\"user_id_owner\":2,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":646,\"infrastructure_user_id\":[277]},{\"infrastructure_id\":331,\"infrastructure_subdomain\":\"my-infrastructure.51.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-02-18T01:32:41Z\",\"infrastructure_updated_timestamp\":\"2022-02-18T01:32:41Z\",\"user_email\":[\"mikealsoft.io\"],\"user_id_owner\":51,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":636,\"infrastructure_user_id\":[268]},{\"infrastructure_id\":329,\"infrastructure_subdomain\":\"my-infrastructure.7.us01.metalsoft.io\",\"infrastructure_service_status\":\"ordered\",\"infrastructure_deploy_status\":\"not_started\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-02-02T15:36:27Z\",\"infrastructure_updated_timestamp\":\"2022-02-02T15:36:27Z\",\"user_email\":[\"alex.bsoft.io\"],\"user_id_owner\":7,\"infrastructure_deploy_id\":null,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":632,\"infrastructure_user_id\":[266]},{\"infrastructure_id\":332,\"infrastructure_subdomain\":\"my-infrastructure.46.us01.metalsoft.io\",\"infrastructure_service_status\":\"active\",\"infrastructure_deploy_status\":\"finished\",\"datacenter_name\":\"us-chi-qts01-dc\",\"infrastructure_created_timestamp\":\"2022-02-22T03:42:35Z\",\"infrastructure_updated_timestamp\":\"2022-02-22T03:42:35Z\",\"user_email\":[\"johlsoft.io\"],\"user_id_owner\":46,\"infrastructure_deploy_id\":328,\"afc_group_created_timestamp\":null,\"afc_group_finished_timestamp\":null,\"thrownError\":null,\"executedSuccess\":null,\"total\":null,\"infrastructure_change_id\":637,\"infrastructure_user_id\":[269]}],\"rows_total\":12,\"rows_order\":[[\"thrownError\",\"DESC\"],[\"infrastructure_service_status\",\"ASC\"],[\"infrastructure_deploy_status\",\"DESC\"],[\"infrastructure_updated_timestamp\",\"DESC\"],[\"infrastructure_change_id\",\"ASC\"],[\"infrastructure_user_id\",\"ASC\"]],\"xls_for_pivot_tables_download_url\":\"https://api.us01.metalsoft.io/api/url?rqi=br.GF9yjZjqPIAyxF-2R2Yutw_RdgFn4kdNJut90RGhUHq3nYh4SqXqc3YoyfXuImHgPjTTXnBBclxgCxFsVuDVfvFaGssd49BylQu0EhW3m3-zcv4RCczGZd498AW5kd4am8ol7_ycbSO7iEivAtj1V8sgIYvfJYOq6sDApe4m6OBT-kEzfItet3lMDHOYxRPlCMH_khfNtzxEWipwqQJnR_0521GJuiAJJsENAXfiBl8GvTis-DnzobWV5BEaC9TNBseORiWAQL6g-EFWEV3FwZTQRQKGSbI0vG6rLY6xZV0rEwtn4LXN8IpBsQBFJv4ZfYsPKaU9naCVYLw48754y_rNYKQiXukKIk_KSFmHw46qd_py410vyaqiho609DqQQDKE_7qLNs28trLIWrIZ5Io3XstVBRx3-OUkcb3A7IjpnpTuEqI2xOp75256ApM2vxmUCn_aEcCpNQnub06tFmvpAvwzY990P7yDkM89_1Js-2owC662BiWiboVf-aG8oMjJpHJ_VhoT0Ypp6U8wiUHGYFEjh175C2_oi65Js3r5H-_-wYEmd6tzROPZMjJYnL2rrx0_tbDNdO3C8_p4hTMGDR2fCncQemZWMedzlchxKV_vd3NRLJ-hcTvdBo-miDhn5XqVU4Vdoyg_hdRFHaqrF2ML0zXoTu5KFDo2rZ-9l5dKx1zJSwprFDLuBz0NYdL_gWL6NAGMmnmnB73ELrUspJHj2STfKKqjFNqnw3uPFQM_Zp7Yrgij5R3cFxu-L6RCfodlOtyDvEVnxFyiapyPCblRDv5yPo6QOPnlM3kfZsN9gB0G4mDm8gqYs8z16ZLq-2bqF6f6Mhwc6fu9HXqjqXbeWQ8xb02u6bcM80SXoGUlAW0JrGAnY9mw7XePxcqJ1IvhJgFmEpPyxawaDjwqoQ6VjKgcwTBLJdzeolaAw8LdnIGA-0CZkvnlZhDyrD_ehImlDmmLOWKJVR8ZMJde5gJdDoa0OrEpzUGOb6z-s1uEMcMmvbqdo2Peeq0osRupvIkmUnzmZeEW3CRz9369Qx6XeF6KGB4dzMcC24mVX3OP27eS8CScuHrOjP4l5MWsVQ-mTFnZDr6FtQ1WC-OWrpXw_RrX8n7MDGfqq971JM4jOqYkdFHDQqbFNLChYS5VlxL0YwCYfrbNjVrfQlOENSJiLLCwzjASgSObPSPC6jfvQ49T1tqQff0rpv2qmRvposNE5i-d1QORg8ULUg&v=wGDPniMGuxdzpqMETrX4sg\"}},\"jsonrpc\":\"2.0\",\"id\":19}"
