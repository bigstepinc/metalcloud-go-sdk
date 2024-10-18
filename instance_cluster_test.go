package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestClusterUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj Cluster
	err := json.Unmarshal([]byte(_clusterFixture2), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	Expect(obj.ClusterCustom["cluster_vcsa_admin_username"]).To(Equal("root"))
	Expect(obj.ClusterCustomJSON).NotTo(BeEmpty())

	var m2 map[string]string
	err = json.Unmarshal([]byte(obj.ClusterCustomJSON), &m2)
	Expect(err).To(BeNil())
	Expect(m2).To(HaveKey("cluster_saas_admin_username"))

}

func TestClusterAppUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj AppVMWareVsphereWrapper
	err := json.Unmarshal([]byte(_clusterAppFixture3), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())

	Expect(obj.ClusterApp.InstanceVCenterServerManagement).To(Equal("https://192.168.150.202:5480"))

	var obj2 AppVMWareVCFWrapper
	err = json.Unmarshal([]byte(_clusterAppFixture4), &obj2)
	Expect(err).To(BeNil())
	Expect(obj2).NotTo(BeNil())

	Expect(obj2.ClusterApp.NSXManagerAdminUsername).To(Equal("admin"))
	Expect(len(obj2.ClusterApp.VSphereManagement)).To(Equal(4))
	Expect(obj2.ClusterApp.VSphereManagement["instance-157"].ESXIUsername).To(Equal("root"))
	Expect(len(obj2.ClusterApp.VSphereWorkload)).To(Equal(0))

}

func TestClusterGet(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": ` + _clusterFixture2 + `,"jsonrpc": "2.0","id": 0}`
	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.ClusterGet(1623)
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	r := *ret

	body := (<-requestChan).body

	var m Cluster
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	var m2 map[string]string
	err = json.Unmarshal([]byte(r.ClusterCustomJSON), &m2)
	Expect(err).To(BeNil())
	Expect(m2).To(HaveKey("cluster_saas_admin_username"))

}

const _clusterFixture2 = "{\"cluster_id\":1657,\"infrastructure_id\":868,\"cluster_service_status\":\"ordered\",\"cluster_change_id\":3194,\"cluster_type\":\"vmware_vsphere\",\"cluster_label\":\"testvmware\",\"cluster_subdomain\":\"testvmware.test-infra.7.us01.metalsoft.io\",\"cluster_subdomain_permanent\":\"cluster-1657.us01.metalsoft.io\",\"dns_subdomain_id\":null,\"dns_subdomain_permanent_id\":null,\"cluster_gui_settings_json\":\"\",\"cluster_ssh_key_pair_internal_json_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"cluster_software_version\":\"7.0.0\",\"cluster_automatic_management\":true,\"cluster_custom_json\":\"{\\\"cluster_saas_admin_username\\\":\\\"administrator@vsphere.local\\\",\\\"cluster_saas_initial_password_encrypted\\\":\\\"rqi|aes-cbc|xpKy9qfN5XoaCj39swANeGSJPC6qCrI61UA63J2k3d7OtgN49QEkTrOYnrTB+mKq\\\",\\\"cluster_vcsa_admin_username\\\":\\\"root\\\",\\\"cluster_vcsa_initial_password_encrypted\\\":\\\"rqi|aes-cbc|U875SokMFfHqmEiNGvJz12+6DEajTwdMDEd2Tayzz5e07qk0iIBSfstTaXLC4UMr\\\",\\\"cluster_saas_internal_admin_username\\\":\\\"metalcloudmgmt\\\",\\\"cluster_saas_internal_password_encrypted\\\":\\\"rqi|aes-cbc|MjnrWnft4qCF1vUNe\\\\/6GvDz8MFHi7mW4xzv6EyMFRlnoarA1bqqFau\\\\/3oQfKahxN\\\"}\",\"cluster_updated_timestamp\":\"2024-01-15T18:27:57Z\",\"cluster_is_api_private\":false,\"cluster_connections_json\":\"[]\",\"cluster_operation\":{\"cluster_change_id\":3194,\"cluster_id\":1657,\"cluster_label\":\"testvmware\",\"dns_subdomain_change_id\":67924,\"cluster_subdomain\":\"testvmware.test-infra.7.us01.metalsoft.io\",\"cluster_deploy_type\":\"create\",\"cluster_deploy_status\":\"not_started\",\"cluster_software_version\":\"7.0.0\",\"cluster_automatic_management\":true,\"cluster_gui_settings_json\":\"\",\"cluster_updated_timestamp\":\"2024-01-15T18:27:57Z\",\"cluster_connections_json\":\"[]\",\"infrastructure_deploy_id\":null,\"cluster_empty_edit\":false,\"type\":\"ClusterOperation\",\"cluster_connections\":[],\"cluster_service_assignment\":{\"vsphere_master\":{\"1881\":[\"master\"],\"1883\":[\"master\"]},\"vsphere_worker\":{\"1884\":[\"worker\"],\"1885\":[\"worker\"],\"1886\":[\"worker\"]}}},\"type\":\"Cluster\",\"cluster_app\":null,\"cluster_custom\":{\"cluster_saas_admin_username\":\"administrator@vsphere.local\",\"cluster_saas_initial_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"cluster_vcsa_admin_username\":\"root\",\"cluster_vcsa_initial_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"cluster_saas_internal_admin_username\":\"metalcloudmgmt\",\"cluster_saas_internal_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\"},\"cluster_connections\":[],\"cluster_ssh_management_public_key\":\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCTC1/F9AokgWH1dAaX5MkgapAwHT3tfC7NyVVcioMJKvktmg07XKgIDHQerbHDBZGVdYug8AYVotzVV0RALARvuiBoAe9aBwHkHpHyNToUU+DPx29qIys/UJWB8f07iy9TCojRdSibuTxM6etzbhzMyHJi8NbEnXw3DpfUCnCMUlCJys+YoPhjZCEjS90G5+OpVN34Os7uej0SG/vYx2IXjTXi07wmaUATQfqPbhGh34u+ZplTZaxRZXTI5QAPUI0EC0D9DJnNwlGGZ5hfJVxVY5vldixhustDfDWhSC2AyphVOFaVhDQAv82s42x/q4qaDAURCk99txKjdi5iDpuZ bsi-rsa\",\"cluster_service_assignment\":{\"vsphere_master\":{\"1881\":[\"master\"],\"1883\":[\"master\"]},\"vsphere_worker\":{\"1884\":[\"worker\"],\"1885\":[\"worker\"],\"1886\":[\"worker\"]}}}"

// Unused fixture
// const _clusterFixture1 = "{\"cluster_id\":1651,\"infrastructure_id\":866,\"cluster_service_status\":\"ordered\",\"cluster_change_id\":3186,\"cluster_type\":\"vmware_vsphere\",\"cluster_label\":\"testvmware\",\"cluster_subdomain\":\"testvmware.test-infra.7.us01.metalsoft.io\",\"cluster_subdomain_permanent\":\"cluster-1651.us01.metalsoft.io\",\"dns_subdomain_id\":null,\"dns_subdomain_permanent_id\":null,\"cluster_gui_settings_json\":\"\",\"cluster_ssh_key_pair_internal_json_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"cluster_software_version\":\"7.0.0\",\"cluster_automatic_management\":true,\"cluster_custom_json\":\"{\\\"cluster_saas_admin_username\\\":\\\"administrator@vsphere.local\\\",\\\"cluster_saas_initial_password_encrypted\\\":\\\"rqi|aes-cbc|6iBzCrlUabrxqCrNx2JalmVDliEecqCC4OetO43p2Gjthmn0clPgmyV+yhTmBGI7\\\",\\\"cluster_vcsa_admin_username\\\":\\\"root\\\",\\\"cluster_vcsa_initial_password_encrypted\\\":\\\"rqi|aes-cbc|39p6n4Xniiz2Suhs\\\\/GIHIrZykTaWL66Ur17dCcoYdNHtQ5sa9VSro1P7osyCzVN+\\\",\\\"cluster_saas_internal_admin_username\\\":\\\"metalcloudmgmt\\\",\\\"cluster_saas_internal_password_encrypted\\\":\\\"rqi|aes-cbc|loyokDd4XhtX5pDNgLE2JC\\\\/bv0ZRrHUtRw0GDWx5PtgW540Sv0o7gDewS6KkG5ds\\\"}\",\"cluster_updated_timestamp\":\"2024-01-15T18:09:28Z\",\"cluster_is_api_private\":false,\"cluster_connections_json\":\"[]\",\"cluster_operation\":{\"cluster_change_id\":3186,\"cluster_id\":1651,\"cluster_label\":\"testvmware\",\"dns_subdomain_change_id\":67768,\"cluster_subdomain\":\"testvmware.test-infra.7.us01.metalsoft.io\",\"cluster_deploy_type\":\"create\",\"cluster_deploy_status\":\"not_started\",\"cluster_software_version\":\"7.0.0\",\"cluster_automatic_management\":true,\"cluster_gui_settings_json\":\"\",\"cluster_updated_timestamp\":\"2024-01-15T18:09:28Z\",\"cluster_connections_json\":\"[]\",\"infrastructure_deploy_id\":null,\"cluster_empty_edit\":false,\"type\":\"ClusterOperation\",\"cluster_connections\":[],\"cluster_service_assignment\":{\"vsphere_master\":{\"1869\":[\"master\"],\"1871\":[\"master\"]},\"vsphere_worker\":{\"1872\":[\"worker\"],\"1873\":[\"worker\"],\"1874\":[\"worker\"]}}},\"type\":\"Cluster\",\"cluster_app\":null,\"cluster_custom\":{\"cluster_saas_admin_username\":\"administrator@vsphere.local\",\"cluster_saas_initial_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"cluster_vcsa_admin_username\":\"root\",\"cluster_vcsa_initial_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\",\"cluster_saas_internal_admin_username\":\"metalcloudmgmt\",\"cluster_saas_internal_password_encrypted\":\"BSI\\\\JSONRPC\\\\Server\\\\Security\\\\Authorization\\\\DeveloperAuthorization: Not leaking database encrypted values for extra security.\"},\"cluster_connections\":[],\"cluster_ssh_management_public_key\":\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCGnUCuXxmH+Oi9XYtGeZslKYqGNm3plBda0/WtgscBvr7iNBMFhvCjST7it8cOMGonvsn7Vbe0EP6/EzAfnaNshordn4tUoPm6YtqYIcir1ye4QIPwWzCSJGGaUy59nfiI1o2pBiH9t68ohU7ZILXmhmCh5OotdrkpXnFb1dKkfFIelDQtrCzm0v2Mri9NHGwPmQ6HvR+jOTFcH8C/IcYsT6HK/VqC3E+34U+rAPTaxlMzc2Uo9Et/lTUTE4A+U81gU1XSOLnowH1yQXsJRRj8FsRxDk669pGvaOeTBquyK8z9VvxUscbJoObYGwrvjr7+Sui4dY7usfRh6SvRI0b3 bsi-rsa\",\"cluster_service_assignment\":{\"vsphere_master\":{\"1869\":[\"master\"],\"1871\":[\"master\"]},\"vsphere_worker\":{\"1872\":[\"worker\"],\"1873\":[\"worker\"],\"1874\":[\"worker\"]}}}"
const _clusterAppFixture3 = "{\"cluster_app\":{\"cluster_software_version\":null,\"cluster_software_available_versions\":[\"7.0.0\"],\"connectable_clusters\":{\"cluster\":[],\"container_cluster\":[]},\"vsphere_master\":{\"instance-87\":{\"instance_id\":87,\"instance_label\":\"instance-87\",\"instance_service_status\":\"active\",\"instance_hostname\":\"instance-87.us10.metalsoft.io\",\"instance_cluster_url\":\"unavailable\",\"instance_health\":\"unavailable\",\"type\":\"AppVMwarevSphereInstance\",\"esxi_username\":\"root\",\"esxi_password\":\"Use bsidev.password_decrypt:eyJycWkiOiJici5ZMURWNmRDVWVWV2xvSGp1czdINVc5ZktiU0JfZXVueGpocXBMSHVDTVFDQ01VRHdxVlI2bEhuWlhoVk5oREJzczAxS2stQ2hlekNydjB3eTlXeXdRLXQwR1h2ckFHdFlOanhDLUM0M0JNamVQb05pWWNjZ05Cby1LQVJlWGpNZVhqbTh1UnVNbnljY3R6WHVzc0VpZ1EiLCJ2IjoiTF9jN2lKY3JKNDlBLVhSWWw1SWlLUSJ9\"}},\"vsphere_worker\":{\"instance-89\":{\"instance_id\":89,\"instance_label\":\"instance-89\",\"instance_service_status\":\"active\",\"instance_hostname\":\"instance-89.us10.metalsoft.io\",\"instance_cluster_url\":\"unavailable\",\"instance_health\":\"unavailable\",\"type\":\"AppVMwarevSphereInstance\",\"esxi_username\":\"root\",\"esxi_password\":\"Use bsidev.password_decrypt:eyJycWkiOiJici5MdjhsWUNTeXItTmdkZlJLaEtJVGRRSTFTMm5YODZJcW13UnBEV09GQmxPNlRRU0pPMlIzd0Y5RXZrazRzcXg0LVFYZXJmS0s4YlR5dThuV0RsYWZ1R1NqSHExa20wYWJhS0E1dEIydGpJTS1Ua09HM1JCMHZtXzN3d1ZuQ2ZCMGxoUDV6YnVDQVh6eGVDbGZ3eXlIVFEiLCJ2Ijoia0d0a3Z3aHA1TkZMdXJzVi1XMGwzdyJ9\"},\"instance-90\":{\"instance_id\":90,\"instance_label\":\"instance-90\",\"instance_service_status\":\"active\",\"instance_hostname\":\"instance-90.us10.metalsoft.io\",\"instance_cluster_url\":\"unavailable\",\"instance_health\":\"unavailable\",\"type\":\"AppVMwarevSphereInstance\",\"esxi_username\":\"root\",\"esxi_password\":\"Use bsidev.password_decrypt:eyJycWkiOiJici5kMGdmMHpHclBiUkFBUlZTV3ZSQUxZTnVEbW1JUjZEamt4ZWJ2Q3RBMVNpT2Q0VlZoVEJLSTZGOVNaYUpkekNPRnJvMmtSYlhSOVlTMDdHMWw2ZHJOcmZwbnRSZ3lxU0Q0U1FQMm9UMTR1UW1VeDRGUlJBN0RkRWZ6Sy1ST3ltS1RfRzRhOXJKel9Cam5QcHNLRW5uR0EiLCJ2IjoiUkVtZ2ZjSGg0MzZaN3NwbUtEdXFzdyJ9\"}},\"admin_username\":\"administrator@vsphere.local\",\"admin_initial_password\":\"Use bsidev.password_decrypt:eyJycWkiOiJici5zZWtLN1otdlZIRVAwMURwZmxCX3RyZTlMS3RsbUJNNTF2bHo4eGdMNU92b3NET3dKQ3doLXpBdmprVjdTVi1ENC0tT29LUmRFbm8zLW9CU3dCNEIxMDNuTEdZLW8xc3FOdWJJaF9idXR3bFhPVzIyX2RZVUFtakx1ckVCdU1CQkk5UkpXSVVKREpGYllfSmhNM2liSmciLCJ2IjoiYkd2RE1zYUtpYkxyNWZKNGZNcWx0ZyJ9\",\"vcsa_username\":\"root\",\"vcsa_initial_password\":\"Use bsidev.password_decrypt:eyJycWkiOiJici43WkYyM285ZkYwR2lHWTVBTnE1Z25jNFg1MFhVd3lnSXB3RkN4V2c0bWw1RTRIQWZpUGxIeGZUeUdYQ29HSjcyelVxbXhXZmd5ZUxlaloza0VIcFNqZ3FOcC15N0FmSG5qc1lFWHA2ekUtYl9SZkdEUTEybDdFZTVNRFFXeGRwRmtqOXRVOU9wNjJDTnJoQ0YzWlI2SXciLCJ2IjoiMW1mSkJLUl9VTDhrX1M4ZGtVRmRZZyJ9\",\"instance_vcenter_server_management\":\"https://192.168.150.202:5480\",\"instance_vcenter_web_client\":\"https://192.168.150.202:443\",\"type\":\"AppVMwarevSphere\"},\"type\":\"ClusterApp\"}"
const _clusterAppFixture4 = "{\"cluster_app\": {\r\n            \"cluster_software_version\": null,\r\n            \"cluster_software_available_versions\": [\r\n                \"5.1.0\"\r\n            ],\r\n            \"connectable_clusters\": {\r\n                \"cluster\": [],\r\n                \"container_cluster\": []\r\n            },\r\n            \"vcf_management\": {\r\n                \"instance-155\": {\r\n                    \"instance_id\": 155,\r\n                    \"instance_label\": \"instance-155\",\r\n                    \"instance_service_status\": \"ordered\",\r\n                    \"instance_hostname\": \"instance-155\",\r\n                    \"instance_cluster_url\": \"unavailable\",\r\n                    \"instance_health\": \"unavailable\",\r\n                    \"type\": \"AppVMwareVCFInstance\",\r\n                    \"esxi_username\": \"root\",\r\n                    \"esxi_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici53eEFUbTl0UjJmRXNWc1N2aUpSYlBFdlJIX2Q5b1FEUEF2RjVuN3BGbHp2R0ZFam8wTDgxSVBFSlVxY1BzN2dINHVfUlZsazVWajZIUW9tNHowQU44dThiZjhUZG42RlZxaDhCdGpEN3U1MmtPYkx6RXNyWkl4U19xeW9ZQm1heFZfbTJEekhtd0tDMzQtQm1TbHJ1cFEiLCJ2IjoiWFIwSFZLV2Y0eURoeXYxd1g5TjlZQSJ9\"\r\n                },\r\n                \"instance-156\": {\r\n                    \"instance_id\": 156,\r\n                    \"instance_label\": \"instance-156\",\r\n                    \"instance_service_status\": \"ordered\",\r\n                    \"instance_hostname\": \"instance-156\",\r\n                    \"instance_cluster_url\": \"unavailable\",\r\n                    \"instance_health\": \"unavailable\",\r\n                    \"type\": \"AppVMwareVCFInstance\",\r\n                    \"esxi_username\": \"root\",\r\n                    \"esxi_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5YU09jWmVLNXhtNDNJS1lfTDBrQ3NvR0Z3M3NhSHBQeGQtb2VaWGZmcnpDTzB2SHd5MHRoaXFUT0N3dG9IX0d2Q1BaLWpRdmVJTVFXYk1EemZrR25TWktSUkpYT0ptMFNaQmZYR3FCRzUzV193MGljYlVwN0JWdVl6TmVIZmNja25Id0ZpT205UEd2ZXRzcmF3cHNjM3ciLCJ2IjoiUTlyTThocjFBVWlqMkw0ZkZpUm4tZyJ9\"\r\n                },\r\n                \"instance-157\": {\r\n                    \"instance_id\": 157,\r\n                    \"instance_label\": \"instance-157\",\r\n                    \"instance_service_status\": \"ordered\",\r\n                    \"instance_hostname\": \"instance-157\",\r\n                    \"instance_cluster_url\": \"unavailable\",\r\n                    \"instance_health\": \"unavailable\",\r\n                    \"type\": \"AppVMwareVCFInstance\",\r\n                    \"esxi_username\": \"root\",\r\n                    \"esxi_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5vaHk4a2dWVF9qUUcxMmctTHByeXIxYW5wQ3R1TEJoSEhXRXJiWFhpU3JEUlgyNEhQZkVHMWtrWWJhc2R1dTRZRVV4YTg5ZTZaV1BoNm1kN3FpaEg3blRoa0ExUU53aFdFMVN3M3B3bi05R09WWERjaXlqNjJScDZvMVRVYnlOamRQTHYwdFZhb191dFh5al9qM245X1EiLCJ2IjoiR2dVa3d1U2FBNHdrZ2g1VjBTT3B5ZyJ9\"\r\n                },\r\n                \"instance-158\": {\r\n                    \"instance_id\": 158,\r\n                    \"instance_label\": \"instance-158\",\r\n                    \"instance_service_status\": \"ordered\",\r\n                    \"instance_hostname\": \"instance-158\",\r\n                    \"instance_cluster_url\": \"unavailable\",\r\n                    \"instance_health\": \"unavailable\",\r\n                    \"type\": \"AppVMwareVCFInstance\",\r\n                    \"esxi_username\": \"root\",\r\n                    \"esxi_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5hQTYzUDBQa29abVNUcGowamlBeE1oY09nQXp3TWRzS1g3VW81aUlXajI2M3BNaEczWWp3cElpTzcyV19mc29BUnV0c0dVZ3V4TlgyRUs1aVhseGZQQmlNVUdSSFpnV2xwOFRsbmZ0eG1nZHZEaGZCanI4VGM5NnNaa29ITV9HbVlLRE5XMG5jbGVnMVB0NlAxdmJiZUZRVE1EOGJBTXVseUpLQzR1V1lISkkiLCJ2IjoiUG1oRVZyMERTWHdYTlByV294TldSQSJ9\"\r\n                }\r\n            },\r\n            \"vcf_workload\": [],\r\n            \"app_cluster_id\": 67,\r\n            \"cba_admin_username\": \"admin\",\r\n            \"cba_admin_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5XTm12aDVmQXBzSXZxWndMVFJYcnBZSFpiSkxKVldyV2IycWltMkNidnFtZ0pUSk85bTlBb1UtdXBhcG02NjRsemZmOG9xWUpfYlJDNElORklnTGVGeE5Kam9uYXprZFpTaGhibm5KdVhBLS1GUmpRSXhNNjdJR1M5QjFiUU4zQSIsInYiOiItRmljdzJESmpvNkU5ZldUdGJlaXBnIn0\",\r\n            \"cba_root_username\": \"root\",\r\n            \"cba_root_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5xcmNPWktyZG03UG9zWkFtUVFoVFVJajFGajRkOWFiUFQyYnphTTFVaHhNOXJtdUVhTUxJWmllNmdGaFVWRVRDZFBBM2tjZXh5c1FrWDUtbzIwdHRJZ2p1aTVTbTBGVFdTVEJ4TjRlWXZpMnVrTTV0WDAybTJoMzdydHlQd2FqYkN1SEphZHM1X19kQlVfdldhWFoxdVEiLCJ2IjoiOGE4b0M3WVpYa0ZTMkxCV1VPUjBfUSJ9\",\r\n            \"sddc_root_username\": \"root\",\r\n            \"sddc_root_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5Wa3htTElsdkt1UERSLVZac1ZwenItQTF6S20wVGdPZXlZMnVnZHhSMXY3NHJ6eHNLVUNYZ2JnNUdBSzQwQVRPQnh1aEFkdi1zbXI1dTQ4d05iMGVwUXpiN1JhVnNRMjZtc0EwZWRwY3duUVo5S0RGWHJKeXp3d1pIWVR2cDBQdDl3aW94VWRlc0FxRHpHSTZfMlN2cUEiLCJ2IjoiWjRCZVZOMkgzZ19nUnJ6NHBKeHVTQSJ9\",\r\n            \"sddc_local_admin_username\": \"admin\",\r\n            \"sddc_local_admin_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici4zdjkxVUZYSTl3RzVpYnlCaFFIRFozVUdKaEZCWjJrN3F5dnJsaWFxSjVyU1lqUnM0QWpCU19XdHNlRV9sTG9pYU9VdWpJcFpsbmRoemE3Rk5FVHd4NUhtTGNUNGNoc3gyV3NvT0pNenRWNWVva3dYT0V1NDBMem9yczdQQUFNYm5TU3NKR0VPYUZHeWdBV0pfSEVBQmciLCJ2IjoiWGlRVGs0dlIwamRnX3BrV2dMU1IzdyJ9\",\r\n            \"sddc_second_username\": \"vcf\",\r\n            \"sddc_second_user_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici44LVVaT3VyTVRYMGpQajBwVHo5U2NoQXRnUU1TcTk3am5UUEFnQWljekdRLUp5Rlk5bmhoZmtlQnB3WnhTNFVabjRVN1c0eFN1eHFHSlpSemJTSUc4TkNTcHdxeElneUc5MFNDMmt6d0E4eHlNdDdxaDhtYnF1VTgxYktYMGtzSFJXMklYZUhzT0FEQWdmMHBXQzNqQ0EiLCJ2IjoiRVRqZzV2WXJrWHMyZWRoU0Q2Sk5WQSJ9\",\r\n            \"nsx_manager_root_username\": \"root\",\r\n            \"nsx_manager_root_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5qeldpUk9oOERvS3R2TjRmYjR4Z2NLR3pDLW1PY1IxY3c5OVRJMW90RHlRalN5UHBKMVdsak1Fb1N2MzBGcDk3aV9iUFBERzJ0LUN3dG1KajkxN0tBbWhaUjRfd0MzZnE5VUk0dHhiSk15LVZNalFEX3RuTW5JcUg0QzlDX2l4dDdKLVlPWlR2MGROR3JyX0Q1SkN6dFEiLCJ2IjoiWWMtc3luVEVhMnNGMmlrVHk4TERKUSJ9\",\r\n            \"nsx_manager_admin_username\": \"admin\",\r\n            \"nsx_manager_admin_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5HZDVwQ2hna1ZVRWloWnh2cTlwQlc2R1M3TU9sT21RRmVzTEdLbnF5ekxSc1UzM0xVNzcxOVRVQnk3QTlqZVlPRG52bkFPeUxId25RUE9pdHp4THZNTTV0b2xZZTRRWDg3eUZfdm9zNk4yamkzdmpBOTkzNHNvWjk1R2didlE0Qnl4cGtna2dSc2VqZnItZ1g0OHl0T3ciLCJ2IjoiTG1yM2s3Q3cwaTZaR3JVamFZX2lSdyJ9\",\r\n            \"nsx_manager_audit_username\": \"audit\",\r\n            \"nsx_manager_audit_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5GNDMwZEV6bzVrZ1QxMGl3b01qbDJuLWl4SlUwOHlXMGx6bHYxY1QzWElFU0ZJNGJyU0tscGFvN1h3b1dxbEZLSm1IcWhRak1VUTNDbUl6NmxFd09KMkwyN1pNbVBZR3NhWVA1ODd6LVBHSnVhdktGUWpTME9ibjJ1NWRrN1RzbG1ubGYxMnpHT2tSb2FHTkpXNmxONWciLCJ2IjoieDUxdEthaUh6b2RPencteWFEelpBQSJ9\",\r\n            \"vcenter_root_username\": \"root\",\r\n            \"vcenter_root_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5HdXc1WGU1WTlvand5Q1lTQm1LTWQzNUZILXV1b1NIMnl0cHh0dGQwY2JpcENGRHktRFhpVi1kaTNDeHlSRklUeG81ZEIxaVQxNkdMSHlCZ3NmYV9oVm9xdnM1VzA5VFc5czBEdW1EcUFpYVJFTnNwOVFSTTBwNngyX2JuQWY0anFhMjNFVk5SNVdDNU5ZdEpjOExXTnciLCJ2IjoicXlrcnpKX3VmRzRUZGVWZkVGZEhUZyJ9\",\r\n            \"vcenter_sso_admin_username\": \"administrator@vsphere.local\",\r\n            \"vcenter_sso_admin_password\": \"Use bsidev.password_decrypt:eyJycWkiOiJici5FUUpNb0RCNzhZVm45QlhzVXBCZmpNMHlMTnJhbGFrbXd4UVNDVkg4OU1BZW81Z0trc1k1MURZNU9hVEMwa3haV3Z3dDA2OVFPNmw3T1NlSjVzUU0zbzlaVkhadnNFaVgtaThyTkhubGhhdzNDbzlQdnZvdW9YVnJ6b0lsaF9YZE5aVkxXYTl1em53YXVDOUtFNS1OU1EiLCJ2IjoiUTlRNGxZOWZieEhORnRlWjdUbFZLZyJ9\",\r\n            \"type\": \"AppVMwareVCF\",\r\n            \"cba_ip_url\": null,\r\n            \"cba_url\": null,\r\n            \"sddc_ip_url\": null,\r\n            \"sddc_url\": null,\r\n            \"m-vcs1_ip_url\": null,\r\n            \"m-vcs1_url\": null,\r\n            \"m-nsx1_ip_url\": null,\r\n            \"m-nsx1_url\": null,\r\n            \"m-nsx1a_ip_url\": null,\r\n            \"m-nsx1a_url\": null,\r\n            \"m-nsx1b_ip_url\": null,\r\n            \"m-nsx1b_url\": null,\r\n            \"m-nsx1c_ip_url\": null,\r\n            \"m-nsx1c_url\": null\r\n        },\r\n        \"type\": \"ClusterApp\"\r\n    }"
