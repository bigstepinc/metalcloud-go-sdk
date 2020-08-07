package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSubnetsUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	var obj map[string]searchResultForSubnets
	err := json.Unmarshal([]byte(_subnetsFixture1), &obj)
	Expect(err).To(BeNil())
	Expect(obj["_subnet_pools"]).NotTo(BeNil())
	Expect(len(obj["_subnet_pools"].Rows)).To(Equal(2))
	Expect(obj["_subnet_pools"].Rows[0].SubnetPoolPrefixHumanReadable).To(Equal("2a02:0cb8:0000:0000:0000:0000:0000:0000"))

}
func TestSubnetUtilizationUnmarshal(t *testing.T) {
	RegisterTestingT(t)
	var obj map[string]searchResultForSubnets
	err := json.Unmarshal([]byte(_subnetsFixture1), &obj)
	Expect(err).To(BeNil())

	j := obj["_subnet_pools"].Rows[0].SubnetPoolUtilizationCachedJSON
	Expect(j).NotTo(BeNil())

	var u SubnetPoolUtilization
	json.Unmarshal([]byte(j), &u)
	Expect(u.PrefixCountAllocated["64"]).To(Equal(20))
	Expect(u.PrefixCountFree["64"]).To(Equal(2028))
	Expect(u.IPAddressesUsableCountAllocated).To(Equal("368934881474191032320"))
	Expect(u.IPAddressesUsableCountFree).To(Equal("37409996981482970677248"))
	Expect(u.IPAddressesUsableFreePercentOptimistic).To(Equal("99"))

}

const _subnetsFixture1 = "{\"_subnet_pools\":{\"duration_milliseconds\":0.02988600730895996,\"rows\":[{\"subnet_pool_id\":1,\"subnet_pool_prefix_human_readable\":\"2a02:0cb8:0000:0000:0000:0000:0000:0000\",\"subnet_pool_prefix_hex\":\"2a020cb8000000000000000000000000\",\"subnet_pool_netmask_human_readable\":\"ffff:ffff:ffff:f800:0000:0000:0000:0000\",\"subnet_pool_netmask_hex\":\"fffffffffffff8000000000000000000\",\"subnet_pool_prefix_size\":53,\"subnet_pool_routable\":true,\"user_id\":null,\"subnet_pool_destination\":\"WAN\",\"datacenter_name\":\"uk-reading\",\"network_equipment_id\":1,\"subnet_pool_utilization_cached_json\":\"{\\\"prefix_count_free\\\": {\\\"64\\\": 2028}, \\\"prefix_count_allocated\\\": {\\\"64\\\": 20}, \\\"ip_addresses_usable_count_free\\\": \\\"37409996981482970677248\\\", \\\"ip_addresses_usable_count_allocated\\\": \\\"368934881474191032320\\\", \\\"ip_addresses_usable_free_percent_optimistic\\\": \\\"99\\\"}\",\"subnet_pool_cached_updated_timestamp\":\"2020-08-07T07:14:17Z\"},{\"subnet_pool_id\":2,\"subnet_pool_prefix_human_readable\":\"84.40.60.0\",\"subnet_pool_prefix_hex\":\"54283c00\",\"subnet_pool_netmask_human_readable\":\"255.255.252.0\",\"subnet_pool_netmask_hex\":\"fffffc00\",\"subnet_pool_prefix_size\":22,\"subnet_pool_routable\":true,\"user_id\":null,\"subnet_pool_destination\":\"WAN\",\"datacenter_name\":\"uk-reading\",\"network_equipment_id\":null,\"subnet_pool_utilization_cached_json\":\"{\\\"prefix_count_free\\\": {\\\"27\\\": 3, \\\"28\\\": 3, \\\"29\\\": 9, \\\"30\\\": 41}, \\\"prefix_count_allocated\\\": {\\\"27\\\": 2, \\\"28\\\": 8, \\\"29\\\": 32, \\\"30\\\": 49}, \\\"ip_addresses_usable_count_free\\\": \\\"212\\\", \\\"ip_addresses_usable_count_allocated\\\": \\\"371\\\", \\\"ip_addresses_usable_free_percent_optimistic\\\": \\\"36\\\"}\",\"subnet_pool_cached_updated_timestamp\":\"2020-08-07T07:14:17Z\"}],\"rows_total\":19,\"rows_order\":[[\"subnet_pool_id\",\"ASC\"]],\"xls_for_pivot_tables_download_url\":\"https://api.bigstep.com/api/url?rqj=br.ye2l7qGCphx_yNXl4DA6hP9BJvsLBYzVeJ_JR4GmqNW6AhEuIWFv53IQF2QnAzc4yTIEfBA0DYPI1ZvgZBW_AnXTp2C-YqqfG9yx_RQH_zGEEm86tDoGyD02wn76HAnaAXelNko5QP0RmXUYrrPq-hK9kFcQLiJ9ozGwscDhSkse7chGfm1BDXP92GCyBWZ-9mVDcwqL3ykfWaJHAGvuI-aiPbx9Vhz6ZgqmhLpy6wGhJWn-eOXTH-jhH9KLwjfwR42mjcg9KAnkjq90PeI8n8U8GmjFhbz6Tnc0cTEG98oVHqDnG4-jdNIOOwF55voDIkd0HhiLrIOCoxlFfN0Q-K6VDn0s1w07WKQJp0esMHp-Q1pktH3NH5nmDvC4CJBcKK5eNLoz6hCKikqL4vN2z3l5brcg8dfo2G5Jj50mq_OCjS4aIklBcSY7nbgM2n4iwbBCESt7xeCyEaeMAXEVlNubsmf6YygaivOv35vp_C3J-2RGEcd3JbVN4MrpJeY1sA9_px00xzMezJ9APBbMd7tEo56AANIvm_G23d5fu_mbECKtgnS0w4gUJhnN2n2PaVNThdMcLngr5z7QoqI8pwWsBLB32dibjGvPuHT9Ff8Tb1bR9oMfPnuHUqBoXJPlQYS5ie726-ogwa6eoz6IG3hqeRpr5LVEXliQMj-sU_4b6Mil80T7wIUgfEyQjPMYcQjFYy9aYSlFJAEC3I3DslPzBByOaWxHMBLabYBcXr3lM1ycXaZdSLE7_ZUQqlEpHoapTKOnnshzKbIiFyUiLg&v=ffI8M48lE-x3TuL4r9wwBA\"}}"
