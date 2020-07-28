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
	Expect(i.SharedDriveCredentials.StorageIPAddress).To(Equal(""))

}

var _sharedDriveFixture1 = "{\"shared_drive_id\":5033,\"shared_drive_change_id\":16508,\"infrastructure_id\":26849,\"shared_drive_size_mbytes\":2048,\"shared_drive_created_timestamp\":\"2020-07-28T12:26:23Z\",\"shared_drive_updated_timestamp\":\"2020-07-28T12:26:23Z\",\"shared_drive_storage_type\":\"iscsi_ssd\",\"shared_drive_has_gfs\":false,\"shared_drive_service_status\":\"ordered\",\"shared_drive_label\":\"csivolumename\",\"shared_drive_subdomain\":\"csivolumename.test-kube-csi.7.bigstep.io\",\"shared_drive_gui_settings_json\":\"\",\"shared_drive_operation\":{\"shared_drive_change_id\":16508,\"shared_drive_id\":5033,\"shared_drive_size_mbytes\":2048,\"shared_drive_updated_timestamp\":\"2020-07-28T12:26:23Z\",\"shared_drive_storage_type\":\"iscsi_ssd\",\"shared_drive_has_gfs\":false,\"shared_drive_label\":\"csivolumename\",\"shared_drive_subdomain\":\"csivolumename.test-kube-csi.7.bigstep.io\",\"shared_drive_deploy_type\":\"create\",\"shared_drive_deploy_status\":\"not_started\",\"shared_drive_gui_settings_json\":\"{\\\"nRowIndex\\\":0,\\\"nColumnIndex\\\":1,\\\"bShowWidgetChildren\\\":false,\\\"randomInstanceID\\\":\\\"rand:0.3180614591259636\\\",\\\"userAgent\\\":\\\"Mozilla\\\\/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit\\\\/605.1.15 (KHTML, like Gecko) Version\\\\/13.1.2 Safari\\\\/605.1.15\\\"}\",\"type\":\"SharedDriveOperation\",\"shared_drive_attached_instance_arrays\":[37824],\"shared_drive_attached_container_arrays\":[]},\"type\":\"SharedDrive\",\"shared_drive_credentials\":{\"iscsi\":{\"target_iqn\":null,\"storage_ip_address\":null,\"storage_port\":null,\"lun_id\":null,\"type\":\"iSCSI\"}},\"shared_drive_attached_instance_arrays\":[37824],\"shared_drive_attached_container_arrays\":[]}"
