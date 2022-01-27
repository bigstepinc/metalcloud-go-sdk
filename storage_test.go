package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestStoragePoolUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var i StoragePool
	err := json.Unmarshal([]byte(_storage_fixture1), &i)
	Expect(err).To(BeNil())
	Expect(i).NotTo(BeNil())

	Expect(i.StoragePoolID).To(Equal(1))

}

const _storage_fixture1 = "{\"storage_pool_id\":1,\"user_id\":null,\"datacenter_name\":\"us-chi-qts01-dc\",\"storage_driver\":\"bigstep_storage\",\"storage_pool_name\":\"us-chg-qts01-bstor-hdd01\",\"storage_pool_iscsi_host\":\"100.98.0.6\",\"storage_pool_iscsi_port\":3260,\"storage_pool_endpoint\":\"http://100.98.0.10/api\",\"storage_pool_username\":\"root\",\"storage_pool_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"storage_type\":\"iscsi_hdd\",\"storage_pool_status\":\"active\",\"storage_pool_in_maintenance\":false,\"storage_pool_options_json\":\"{\\\"volume_name\\\":\\\"Pool-A/Data\\\"}\",\"storage_pool_target_iqn\":\"iqn.2013-01.com.bigstep:storage.6wtsr4l.7yk703m.11em6rv\",\"storage_pool_is_experimental\":false,\"storage_pool_created_timestamp\":\"2021-07-26T13:35:31Z\",\"storage_pool_updated_timestamp\":\"2022-01-27T09:20:41Z\",\"storage_pool_capacity_total_cached_real_mbytes\":7364608,\"storage_pool_capacity_usable_cached_real_mbytes\":7364608,\"storage_pool_capacity_free_cached_real_mbytes\":7359925,\"storage_pool_capacity_used_cached_virtual_mbytes\":714677,\"storage_pool_drive_priority\":50,\"storage_pool_shared_drive_priority\":50,\"isns_portal_group_tag\":0,\"storage_pool_port_group_allocation_order_json\":null,\"storage_pool_port_group_physical_ports_json\":null,\"storage_pool_alternate_san_ips\":[\"100.98.0.6\"],\"storage_pool_password\":\"Use bsidev.password_decrypt:sdfd\",\"storage_pool_tags\":[],\"type\":\"StoragePool\"}"
