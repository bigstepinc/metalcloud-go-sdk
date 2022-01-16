package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestAFCUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj map[string]searchResultResponseWrapperForAFC
	err := json.Unmarshal([]byte(_afc_fixture), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	r := obj["_afc_queue"].Rows[1]
	Expect(r.AFCID).To(Equal(5815980))

}

const _afc_fixture = `{
"_afc_queue": {
	"duration_milliseconds": 0.018967866897583008,
	"rows": [
		{
			"afc_type": "asynchronous",
			"afc_response_json": "null",
			"afc_exception_json": "{\"exception_type\":\"BSI_Exception\",\"message\":\"Some power agents from us-chi-qts01-dc could not reach some L3 quarantine gateways at IPs: \\nAgent 8 cannot reach 172.16.0.1 on equipmentId 1.\"}",
			"afc_ip_address_human_readable": "10.244.235.212",
			"afc_id_blocked": null,
			"afc_id_blocked_by": null,
			"infrastructure_id": null,
			"afc_group_id": 315,
			"datacenter_name": "us-chi-qts01-dc",
			"instance_id": null,
			"server_id": null,
			"afc_execute_engine": "BSI",
			"afc_id": 6147305,
			"afc_call_count": 3,
			"afc_retry_max": 2,
			"afc_retry_count": 2,
			"afc_start_timestamp": null,
			"afc_retry_min_sec": 60,
			"afc_duration_milliseconds": 12057,
			"afc_updated_timestamp": "2021-12-14T03:35:48Z",
			"afc_created_timestamp": "2021-12-14T03:33:11Z",
			"afc_call_memory_used_bytes": 2097152,
			"afc_api_error_code": 0,
			"afc_function_name": "datacenter_diagnostics_reach_l3_interface",
			"afc_params_json": "[\"us-chi-qts01-dc\"]",
			"afc_status": "thrown_error"
		},
		{
			"afc_type": "asynchronous",
			"afc_response_json": "null",
			"afc_exception_json": "{\"exception_type\":\"BSI_Exception\",\"message\":\"Some power agents from us-chi-qts01-dc could not reach some L3 quarantine gateways at IPs: \\nAgent 8 cannot reach 172.16.0.1 on equipmentId 1.\"}",
			"afc_ip_address_human_readable": "10.244.73.190",
			"afc_id_blocked": null,
			"afc_id_blocked_by": null,
			"infrastructure_id": null,
			"afc_group_id": 304,
			"datacenter_name": "us-chi-qts01-dc",
			"instance_id": null,
			"server_id": null,
			"afc_execute_engine": "BSI",
			"afc_id": 5815980,
			"afc_call_count": 3,
			"afc_retry_max": 2,
			"afc_retry_count": 2,
			"afc_start_timestamp": null,
			"afc_retry_min_sec": 60,
			"afc_duration_milliseconds": 12041,
			"afc_updated_timestamp": "2021-11-29T08:48:38Z",
			"afc_created_timestamp": "2021-11-29T08:45:58Z",
			"afc_call_memory_used_bytes": 2097152,
			"afc_api_error_code": 0,
			"afc_function_name": "datacenter_diagnostics_reach_l3_interface",
			"afc_params_json": "[\"us-chi-qts01-dc\"]",
			"afc_status": "thrown_error"
		},
		{
			"afc_type": "asynchronous",
			"afc_response_json": "null",
			"afc_exception_json": "{\"exception_type\":\"BSI_Exception\",\"message\":\"Some power agents from us-chi-qts01-dc could not reach some L3 quarantine gateways at IPs: \\nAgent 8 cannot reach 172.16.0.1 on equipmentId 1.\"}",
			"afc_ip_address_human_readable": "10.244.236.178",
			"afc_id_blocked": null,
			"afc_id_blocked_by": null,
			"infrastructure_id": null,
			"afc_group_id": 292,
			"datacenter_name": "us-chi-qts01-dc",
			"instance_id": null,
			"server_id": null,
			"afc_execute_engine": "BSI",
			"afc_id": 4727469,
			"afc_call_count": 6,
			"afc_retry_max": 5,
			"afc_retry_count": 5,
			"afc_start_timestamp": null,
			"afc_retry_min_sec": 60,
			"afc_duration_milliseconds": 12041,
			"afc_updated_timestamp": "2021-10-13T00:30:56Z",
			"afc_created_timestamp": "2021-10-13T00:24:41Z",
			"afc_call_memory_used_bytes": 2097152,
			"afc_api_error_code": 0,
			"afc_function_name": "datacenter_diagnostics_reach_l3_interface",
			"afc_params_json": "[\"us-chi-qts01-dc\"]",
			"afc_status": "thrown_error"
		},
		{
			"afc_type": "asynchronous",
			"afc_response_json": "",
			"afc_exception_json": "",
			"afc_ip_address_human_readable": "",
			"afc_id_blocked": null,
			"afc_id_blocked_by": null,
			"infrastructure_id": null,
			"afc_group_id": null,
			"datacenter_name": "master",
			"instance_id": null,
			"server_id": null,
			"afc_execute_engine": "BSI",
			"afc_id": 6485849,
			"afc_call_count": 0,
			"afc_retry_max": 2147483647,
			"afc_retry_count": 0,
			"afc_start_timestamp": null,
			"afc_retry_min_sec": 233,
			"afc_duration_milliseconds": null,
			"afc_updated_timestamp": "2021-12-29T05:45:49Z",
			"afc_created_timestamp": "2021-12-29T05:44:12Z",
			"afc_call_memory_used_bytes": 0,
			"afc_api_error_code": 0,
			"afc_function_name": "logs_archive",
			"afc_params_json": "[]",
			"afc_status": "running"
		},
		{
			"afc_type": "asynchronous",
			"afc_response_json": "",
			"afc_exception_json": "",
			"afc_ip_address_human_readable": "",
			"afc_id_blocked": null,
			"afc_id_blocked_by": null,
			"infrastructure_id": null,
			"afc_group_id": null,
			"datacenter_name": "master",
			"instance_id": null,
			"server_id": null,
			"afc_execute_engine": "BSI",
			"afc_id": 6485836,
			"afc_call_count": 0,
			"afc_retry_max": 2147483647,
			"afc_retry_count": 0,
			"afc_start_timestamp": null,
			"afc_retry_min_sec": 41,
			"afc_duration_milliseconds": null,
			"afc_updated_timestamp": "2021-12-29T05:47:12Z",
			"afc_created_timestamp": "2021-12-29T05:43:23Z",
			"afc_call_memory_used_bytes": 0,
			"afc_api_error_code": 0,
			"afc_function_name": "lock_garbage_files_delete",
			"afc_params_json": "[]",
			"afc_status": "running"
		},
		{
			"afc_type": "asynchronous",
			"afc_response_json": "",
			"afc_exception_json": "",
			"afc_ip_address_human_readable": "",
			"afc_id_blocked": null,
			"afc_id_blocked_by": null,
			"infrastructure_id": null,
			"afc_group_id": null,
			"datacenter_name": "master",
			"instance_id": null,
			"server_id": null,
			"afc_execute_engine": "BSI",
			"afc_id": 6485816,
			"afc_call_count": 0,
			"afc_retry_max": 2147483647,
			"afc_retry_count": 0,
			"afc_start_timestamp": null,
			"afc_retry_min_sec": 39,
			"afc_duration_milliseconds": null,
			"afc_updated_timestamp": "2021-12-29T05:47:30Z",
			"afc_created_timestamp": "2021-12-29T05:42:05Z",
			"afc_call_memory_used_bytes": 0,
			"afc_api_error_code": 0,
			"afc_function_name": "data_lake_hdfs_keytab_regenerate_if_not_exists_now",
			"afc_params_json": "[]",
			"afc_status": "running"
		}
	]
	}
}`
