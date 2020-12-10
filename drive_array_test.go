package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestDriveArrays(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": {"test":` + _driveArrayFixture + `},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.DriveArrays(100)
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	Expect((*ret)["test"].InfrastructureID).To(Equal(25524))

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

}

func TestDriveArraylUnmarshalTest(t *testing.T) {

	RegisterTestingT(t)

	var da DriveArray

	err2 := json.Unmarshal([]byte(_driveArrayFixture), &da)

	Expect(err2).To(BeNil())
	Expect(da).NotTo(BeNil())
	Expect(da.DriveArrayID).To(Equal(45928))
	Expect(da.DriveSizeMBytesDefault).To(Equal(40960))
	Expect(da.VolumeTemplateID).To(Equal(78))
	Expect(da.InstanceArrayID).To(Equal(35516))
	Expect(da.DriveArrayCount).To(Equal(2))
}

func TestDriveArrayGet(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _driveArrayFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	ret, err := mc.DriveArrayGet(100)
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("drive_array_get"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

	//make sure we return what we need to
	Expect(ret.DriveArrayID).To(Equal(45928))

}

func TestDriveArrayEdit(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _driveArrayFixture + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := DriveArrayOperation{
		DriveArrayCount: 10,
	}

	ret, err := mc.DriveArrayEdit(100, obj)
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("drive_array_edit"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))
	//make sure we send the object
	Expect(params[1].(map[string]interface{})["drive_array_count"]).To(Equal(10.0))
	//make sure we return what we need to
	Expect(ret.DriveArrayID).To(Equal(45928))

}

func TestDriveArrayDelete(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	err = mc.DriveArrayDelete(100)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("drive_array_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(100.0))

}

func TestDriveArrayCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _driveArrayFixture1 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := DriveArray{
		VolumeTemplateID:        78,
		DriveArrayStorageType:   "iscsi_ssd",
		InfrastructureID:        25524,
		DriveArrayServiceStatus: "active",
		DriveArrayCount:         2,
		DriveArrayLabel:         "drive-array-45928",
		DriveArrayOperation: &DriveArrayOperation{
			DriveArrayID:                      45928,
			DriveArrayChangeID:                215701,
			VolumeTemplateID:                  78,
			DriveArrayLabel:                   "drive-array-45928",
			DriveArrayStorageType:             "iscsi_ssd",
			DriveArrayCount:                   2,
			DriveSizeMBytesDefault:            40960,
			InstanceArrayID:                   35516,
			DriveArrayExpandWithInstanceArray: true,
		},
	}
	err = obj.CreateOrUpdate(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("drive_array_get"))

	params := (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("drive-array-45928"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("drive_array_edit"))

	params = (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(45928.0))

	responseBody = `{"error": {"message": "DriveArray not found.","code": 103}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("drive_array_get"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("drive-array-45928"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("drive_array_create"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(float64)).To(Equal(25524.0))
}

func TestDriveArrayDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": [] ,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("user", "APIKey", httpServer.URL, false)
	Expect(err).To(BeNil())

	obj := DriveArray{
		VolumeTemplateID:        78,
		DriveArrayStorageType:   "iscsi_ssd",
		InfrastructureID:        25524,
		DriveArrayServiceStatus: "active",
		DriveArrayCount:         2,
		DriveArrayLabel:         "drive-array-45928",
		DriveArrayID:            1,
	}

	err = obj.Delete(mc)
	Expect(err).To(BeNil())

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	//make sure we use the proper method
	Expect(m["method"].(string)).To(Equal("drive_array_delete"))

	params := (m["params"].([]interface{}))

	//make sure we ask for the proper ID
	Expect(params[0].(float64)).To(Equal(1.0))

}

const _driveArrayFixture = "{\"drive_array_id\":45928,\"drive_array_change_id\":215701,\"volume_template_id\":78,\"drive_array_storage_type\":\"iscsi_ssd\",\"infrastructure_id\":25524,\"drive_array_service_status\":\"active\",\"drive_array_count\":2,\"drive_array_label\":\"drive-array-45928\",\"drive_array_subdomain\":\"drive-array-45928.vanilla.complex-demo.7.bigstep.io\",\"drive_size_mbytes_default\":40960,\"instance_array_id\":35516,\"container_array_id\":null,\"drive_array_expand_with_instance_array\":true,\"drive_array_gui_settings_json\":\"{\\\"bUnsetGridPosition\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.43156788473635765\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/12.0.3 Safari\\\\/605.1.15\\\"}\",\"drive_array_updated_timestamp\":\"2019-03-28T15:24:05Z\",\"drive_array_created_timestamp\":\"2019-03-28T15:23:19Z\",\"cluster_id\":40559,\"container_cluster_id\":null,\"cluster_role_group\":\"none\",\"license_utilization_type\":\"subscription\",\"drive_array_operation\":{\"drive_array_change_id\":215701,\"drive_array_id\":45928,\"volume_template_id\":78,\"drive_array_storage_type\":\"iscsi_ssd\",\"infrastructure_id\":25524,\"drive_array_count\":2,\"drive_array_deploy_type\":\"create\",\"drive_array_deploy_status\":\"finished\",\"drive_array_label\":\"drive-array-45928\",\"drive_array_subdomain\":\"drive-array-45928.vanilla.complex-demo.7.bigstep.io\",\"drive_size_mbytes_default\":40960,\"instance_array_id\":35516,\"container_array_id\":null,\"drive_array_expand_with_instance_array\":true,\"drive_array_gui_settings_json\":\"{\\\"bUnsetGridPosition\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.43156788473635765\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/12.0.3 Safari\\\\/605.1.15\\\"}\",\"drive_array_updated_timestamp\":\"2019-03-28T15:24:05Z\",\"license_utilization_type\":\"subscription\",\"type\":\"DriveArrayOperation\",\"drive_array_filesystem\":{\"type\":\"DriveArrayFilesystem\",\"drive_array_filesystem_type_default\":\"none\",\"drive_array_filesystem_block_size_bytes_default\":4096}},\"type\":\"DriveArray\",\"drive_array_filesystem\":{\"type\":\"DriveArrayFilesystem\",\"drive_array_filesystem_type_default\":\"none\",\"drive_array_filesystem_block_size_bytes_default\":4096}}"
const _driveArrayFixture1 = "{\"drive_array_id\":45928,\"volume_template_id\":78,\"drive_array_storage_type\":\"iscsi_ssd\",\"infrastructure_id\":25524,\"drive_array_service_status\":\"active\",\"drive_array_count\":2,\"drive_array_label\":\"drive-array-45928\", \"drive_array_operation\":{\"drive_array_change_id\":215701,\"drive_array_id\":45928,\"volume_template_id\":78,\"drive_array_storage_type\":\"iscsi_ssd\",\"infrastructure_id\":25524,\"drive_array_count\":2,\"drive_array_deploy_type\":\"create\",\"drive_array_deploy_status\":\"finished\",\"drive_array_label\":\"drive-array-45928\",\"drive_array_subdomain\":\"drive-array-45928.vanilla.complex-demo.7.bigstep.io\",\"drive_size_mbytes_default\":40960,\"instance_array_id\":35516,\"container_array_id\":null,\"drive_array_expand_with_instance_array\":true,\"drive_array_gui_settings_json\":\"{\\\"bUnsetGridPosition\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.43156788473635765\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/12.0.3 Safari\\\\/605.1.15\\\"}\",\"drive_array_updated_timestamp\":\"2019-03-28T15:24:05Z\",\"license_utilization_type\":\"subscription\",\"type\":\"DriveArrayOperation\",\"drive_array_filesystem\":{\"type\":\"DriveArrayFilesystem\",\"drive_array_filesystem_type_default\":\"none\",\"drive_array_filesystem_block_size_bytes_default\":4096}},\"type\":\"DriveArray\",\"drive_array_filesystem\":{\"type\":\"DriveArrayFilesystem\",\"drive_array_filesystem_type_default\":\"none\",\"drive_array_filesystem_block_size_bytes_default\":4096}}"
