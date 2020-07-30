package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSharedDriveUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var i SharedDrive
	err := json.Unmarshal([]byte(_sharedDriveFixture1), &i)
	Expect(err).To(BeNil())
	Expect(i).NotTo(BeNil())

	Expect(i.SharedDriveID).To(Equal(5033))

}

func TestProvisionedSharedDriveUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var i SharedDrive
	err := json.Unmarshal([]byte(_sharedDriveFixture2), &i)
	Expect(err).To(BeNil())
	Expect(i).NotTo(BeNil())
	Expect(i.SharedDriveLabel).To(Equal("csivolumename3"))
	Expect(i.SharedDriveID).To(Equal(5039))
	Expect(i.SharedDriveCredentials.ISCSI.StorageIPAddress).To(Equal("100.96.0.192"))
	Expect(i.SharedDriveCredentials.ISCSI.StoragePort).To(Equal(3260))
	Expect(i.SharedDriveCredentials.ISCSI.TargetIQN).To(Equal("iqn.2013-01.com.bigstep:storage.91nqwgd.4lys6bq.iel1v3k"))
}

var _sharedDriveFixture1 = "{\"shared_drive_id\":5033,\"shared_drive_change_id\":16508,\"infrastructure_id\":26849,\"shared_drive_size_mbytes\":2048,\"shared_drive_created_timestamp\":\"2020-07-28T12:26:23Z\",\"shared_drive_updated_timestamp\":\"2020-07-28T12:26:23Z\",\"shared_drive_storage_type\":\"iscsi_ssd\",\"shared_drive_has_gfs\":false,\"shared_drive_service_status\":\"ordered\",\"shared_drive_label\":\"csivolumename\",\"shared_drive_subdomain\":\"csivolumename.test-kube-csi.7.bigstep.io\",\"shared_drive_gui_settings_json\":\"\",\"shared_drive_operation\":{\"shared_drive_change_id\":16508,\"shared_drive_id\":5033,\"shared_drive_size_mbytes\":2048,\"shared_drive_updated_timestamp\":\"2020-07-28T12:26:23Z\",\"shared_drive_storage_type\":\"iscsi_ssd\",\"shared_drive_has_gfs\":false,\"shared_drive_label\":\"csivolumename\",\"shared_drive_subdomain\":\"csivolumename.test-kube-csi.7.bigstep.io\",\"shared_drive_deploy_type\":\"create\",\"shared_drive_deploy_status\":\"not_started\",\"shared_drive_gui_settings_json\":\"{\\\"nRowIndex\\\":0,\\\"nColumnIndex\\\":1,\\\"bShowWidgetChildren\\\":false,\\\"randomInstanceID\\\":\\\"rand:0.3180614591259636\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/13.1.2 Safari\\\\/605.1.15\\\"}\",\"type\":\"SharedDriveOperation\",\"shared_drive_attached_instance_arrays\":[37824],\"shared_drive_attached_container_arrays\":[]},\"type\":\"SharedDrive\",\"shared_drive_credentials\":{\"iscsi\":{\"target_iqn\":null,\"storage_ip_address\":null,\"storage_port\":null,\"lun_id\":null,\"type\":\"iSCSI\"}},\"shared_drive_attached_instance_arrays\":[37824],\"shared_drive_attached_container_arrays\":[]}"
var _sharedDriveFixture2 = "{\"shared_drive_id\":5039,\"shared_drive_change_id\":16514,\"infrastructure_id\":26849,\"shared_drive_size_mbytes\":2048,\"shared_drive_created_timestamp\":\"2020-07-28T15:40:39Z\",\"shared_drive_updated_timestamp\":\"2020-07-30T13:47:35Z\",\"shared_drive_storage_type\":\"iscsi_ssd\",\"shared_drive_has_gfs\":false,\"shared_drive_service_status\":\"active\",\"shared_drive_label\":\"csivolumename3\",\"shared_drive_subdomain\":\"csivolumename3.test-kube-csi.7.bigstep.io\",\"shared_drive_gui_settings_json\":\"{\\\"nRowIndex\\\":0,\\\"nColumnIndex\\\":1,\\\"bShowWidgetChildren\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.17130147893495928\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/13.1.2 Safari\\\\/605.1.15\\\"}\",\"shared_drive_operation\":{\"shared_drive_change_id\":16514,\"shared_drive_id\":5039,\"shared_drive_size_mbytes\":2048,\"shared_drive_updated_timestamp\":\"2020-07-30T13:46:38Z\",\"shared_drive_storage_type\":\"iscsi_ssd\",\"shared_drive_has_gfs\":false,\"shared_drive_label\":\"csivolumename3\",\"shared_drive_subdomain\":\"csivolumename3.test-kube-csi.7.bigstep.io\",\"shared_drive_deploy_type\":\"edit\",\"shared_drive_deploy_status\":\"finished\",\"shared_drive_gui_settings_json\":\"{\\\"nRowIndex\\\":0,\\\"nColumnIndex\\\":1,\\\"bShowWidgetChildren\\\":true,\\\"randomInstanceID\\\":\\\"rand:0.17130147893495928\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/13.1.2 Safari\\\\/605.1.15\\\"}\",\"type\":\"SharedDriveOperation\",\"shared_drive_attached_instance_arrays\":[37824],\"shared_drive_attached_container_arrays\":[]},\"type\":\"SharedDrive\",\"shared_drive_credentials\":{\"iscsi\":{\"target_iqn\":\"iqn.2013-01.com.bigstep:storage.91nqwgd.4lys6bq.iel1v3k\",\"storage_ip_address\":\"100.96.0.192\",\"storage_port\":3260,\"lun_id\":null,\"type\":\"iSCSI\"}},\"shared_drive_attached_instance_arrays\":[37824],\"shared_drive_attached_container_arrays\":[]}"
