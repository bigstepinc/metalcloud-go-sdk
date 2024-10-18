package metalcloud

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
)

//go:generate go run helper/gen_exports.go

type Cluster struct {
	ClusterID                  int    `json:"cluster_id,omitempty" yaml:"clusterId,omitempty"`
	ClusterLabel               string `json:"cluster_label,omitempty" yaml:"clusterLabel,omitempty"`
	ClusterSubdomain           string `json:"cluster_subdomain,omitempty" yaml:"clusterSubdomain,omitempty"`
	ClusterSubdomainPermanent  string `json:"cluster_subdomain_permanent,omitempty" yaml:"clusterSubdomainPermanent,omitempty"`
	InfrastructureID           int    `json:"infrastructure_id,omitempty" yaml:"infrastructureId,omitempty"`
	ClusterType                string `json:"cluster_type,omitempty" yaml:"clusterType,omitempty"`
	ClusterServiceStatus       string `json:"cluster_service_status,omitempty" yaml:"clusterServiceStatus,omitempty"`
	ClusterSoftwareVersion     string `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
	ClusterAutomaticManagement bool   `json:"cluster_automatic_management,omitempty" yaml:"clusterAutomaticManagement,omitempty"`
	//ClusterApp                 interface{} `json:"cluster_app,omitempty" yaml:"clusterApp,omitempty"`
	ClusterChangeId   int               `json:"cluster_change_id,omitempty" yaml:"clusterChangeId,omitempty"`
	ClusterOperation  ClusterOperation  `json:"cluster_operation,omitempty" yaml:"clusterOperation,omitempty"`
	ClusterCustom     map[string]string `json:"cluster_custom,omitempty" yaml:"clusterCustom,omitempty"`
	ClusterCustomJSON string            `json:"cluster_custom_json,omitempty" yaml:"clusterCustomJson,omitempty"`
}

type ClusterOperation struct {
	ClusterLabel               string `json:"cluster_label,omitempty" yaml:"clusterLabel,omitempty"`
	ClusterSubdomain           string `json:"cluster_subdomain,omitempty" yaml:"clusterSubdomain,omitempty"`
	ClusterSubdomainPermanent  string `json:"cluster_subdomain_permanent,omitempty" yaml:"clusterSubdomainPermanent,omitempty"`
	InfrastructureID           string `json:"infrastructure_id,omitempty" yaml:"infrastructureId,omitempty"`
	ClusterType                string `json:"cluster_type,omitempty" yaml:"clusterType,omitempty"`
	ClusterServiceStatus       string `json:"cluster_service_status,omitempty" yaml:"clusterServiceStatus,omitempty"`
	ClusterSoftwareVersion     string `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
	ClusterAutomaticManagement bool   `json:"cluster_automatic_management,omitempty" yaml:"clusterAutomaticManagement,omitempty"`
	ClusterDeployStatus        string `json:"cluster_deploy_status,omitempty" yaml:"clusterDeployStatus,omitempty"`
	ClusterDeployType          string `json:"cluster_deploy_type,omitempty" yaml:"clusterDeployType,omitempty"`
	ClusterChangeId            int    `json:"cluster_change_id,omitempty" yaml:"clusterChangeId,omitempty"`
}

type AppInstanceDetails struct {
	InstanceID         int    `json:"instance_id,omitempty" yaml:"instanceId,omitempty"`
	InstanceLabel      string `json:"instance_label,omitempty" yaml:"instanceLabel,omitempty"`
	InstanceHostname   string `json:"instance_hostname,omitempty" yaml:"instanceHostname,omitempty"`
	InstanceClusterUrl string `json:"instance_cluster_url,omitempty" yaml:"instanceClusterUrl,omitempty"`
	InstanceHealth     string `json:"instance_health,omitempty" yaml:"instanceHealth,omitempty"`
	Type               string `json:"type,omitempty" yaml:"type,omitempty"`
	ESXIUsername       string `json:"esxi_username,omitempty" yaml:"esxiUsername,omitempty"`
	ESXIPassword       string `json:"esxi_password,omitempty" yaml:"esxiUsername,omitempty"`
}

type AppVMWareVsphere struct {
	VSphereWorker                    map[string]AppInstanceDetails `json:"vsphere_worker,omitempty" yaml:"vsphereWorker,omitempty"`
	VSphereMaster                    map[string]AppInstanceDetails `json:"vsphere_master,omitempty" yaml:"vsphereWorker,omitempty"`
	AdminUsername                    string                        `json:"admin_username,omitempty" yaml:"adminUsername,omitempty"`
	AdminPassword                    string                        `json:"admin_password,omitempty" yaml:"adminPassword,omitempty"`
	ClusterSoftwareAvailableVersions []string                      `json:"cluster_software_available_versions,omitempty" yaml:"clusterSoftwareAvailableVersions,omitempty"`
	ClusterSoftwareVersion           string                        `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
	Type                             string                        `json:"type,omitempty" yaml:"type,omitempty"`
	InstanceVCenterServerManagement  string                        `json:"instance_vcenter_server_management,omitempty" yaml:"instanceVcenterServerManagement,omitempty"`
	InstanceVcenterWebClient         string                        `json:"instance_vcenter_web_client,omitempty" yaml:"instanceVcenterWebClient,omitempty"`
	VCSAUsername                     string                        `json:"vcsa_username,omitempty" yaml:"vcsaUsername,omitempty"`
	VCSAInitialPassword              string                        `json:"vcsa_initial_password,omitempty" yaml:"vcsaUsername,omitempty"`
}

type AppVMWareVsphereWrapper struct {
	ClusterApp AppVMWareVsphere `json:"cluster_app,omitempty" yaml:"clusterApp,omitempty"`
	Type       string           `json:"type,omitempty" yaml:"type,omitempty"`
}

// note if you change this structure change also the custom unmarshaling below that fixes the shitty [] instead of {} problem
type AppVMWareVCF struct {
	VSphereWorkload                  map[string]AppInstanceDetails `json:"vcf_workload,omitempty" yaml:"vcfWorkload,omitempty"`
	VSphereManagement                map[string]AppInstanceDetails `json:"vcf_management,omitempty" yaml:"vcfManagement,omitempty"`
	AdminUsername                    string                        `json:"admin_username,omitempty" yaml:"adminUsername,omitempty"`
	AdminPassword                    string                        `json:"admin_password,omitempty" yaml:"adminPassword,omitempty"`
	ClusterSoftwareAvailableVersions []string                      `json:"cluster_software_available_versions,omitempty" yaml:"clusterSoftwareAvailableVersions,omitempty"`
	ClusterSoftwareVersion           string                        `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
	Type                             string                        `json:"type,omitempty" yaml:"type,omitempty"`
	InstanceVCenterServerManagement  string                        `json:"instance_vcenter_server_management,omitempty" yaml:"instanceVcenterServerManagement,omitempty"`
	InstanceVcenterWebClient         string                        `json:"instance_vcenter_web_client,omitempty" yaml:"instanceVcenterWebClient,omitempty"`
	VCSAUsername                     string                        `json:"vcsa_username,omitempty" yaml:"vcsaUsername,omitempty"`
	VCSAInitialPassword              string                        `json:"vcsa_initial_password,omitempty" yaml:"vcsaUsername,omitempty"`
	CBAAdminUsername                 string                        `json:"cba_admin_username,omitempty" yaml:"cbaAdminUsername,omitempty"`
	CBAAdminPassword                 string                        `json:"cba_admin_password,omitempty" yaml:"cbaAdminPassword,omitempty"`
	CBARootUsername                  string                        `json:"cba_root_username,omitempty" yaml:"cbaRootUsername,omitempty"`
	CBARootPassword                  string                        `json:"cba_root_password,omitempty" yaml:"cbaRootPassword,omitempty"`
	SDDCRootUsername                 string                        `json:"sddc_root_username,omitempty" yaml:"sddcRootUsername,omitempty"`
	SDDCRootPassword                 string                        `json:"sdcc_root_password,omitempty" yaml:"sddcRootPassword,omitempty"`
	SDDCLocalAdminUsername           string                        `json:"sddc_local_admin_username,omitempty" yaml:"sddcLocalAdminUsername,omitempty"`
	SDDCLocalAdminPassword           string                        `json:"sdcc_local_admin_password,omitempty" yaml:"sddcLocalAdminPassword,omitempty"`
	SDDCSecondUsername               string                        `json:"sddc_second_username,omitempty" yaml:"sddcSecondUsername,omitempty"`
	SDDCSecondPassword               string                        `json:"sdcc_second_password,omitempty" yaml:"sddcSecondPassword,omitempty"`
	NSXManagerRootUsername           string                        `json:"nsx_manager_root_username,omitempty" yaml:"nsxManagerRootUsername,omitempty"`
	NSXManagerRootPassword           string                        `json:"nsx_manager_root_password,omitempty" yaml:"nsxManagerRootPassword,omitempty"`
	NSXManagerAdminUsername          string                        `json:"nsx_manager_admin_username,omitempty" yaml:"nsxManagerAdminUsername,omitempty"`
	NSXManagerAdminPassword          string                        `json:"nsx_manager_admin_password,omitempty" yaml:"nsxManagerAdminPassword,omitempty"`
	NSXManagerAuditUsername          string                        `json:"nsx_manager_audit_username,omitempty" yaml:"nsxManagerAuditUsername,omitempty"`
	NSXManagerAuditPassword          string                        `json:"nsx_manager_audit_password,omitempty" yaml:"nsxManagerAuditPassword,omitempty"`
	VCenterRootUsername              string                        `json:"vcenter_root_username,omitempty" yaml:"vcenterRootUsername,omitempty"`
	VCenterRootPassword              string                        `json:"vcenter_root_password,omitempty" yaml:"vcenterRootPassword,omitempty"`
	VCenterSSOAdminUsername          string                        `json:"vcenter_sso_admin_username,omitempty" yaml:"vcenterSSOAdminUsername,omitempty"`
	VCenterSSOAdminPassword          string                        `json:"vcenter_sso_admin_password,omitempty" yaml:"vcenterSSOAdminPassword,omitempty"`
	CBAIPURL                         string                        `json:"cba_ip_url,omitempty" yaml:"cbaIPUrl,omitempty"`
	CBAURL                           string                        `json:"cba_url,omitempty" yaml:"cbaURL,omitempty"`
	SDDCIPURL                        string                        `json:"sddc_ip_url,omitempty" yaml:"sddcIPUrl,omitempty"`
	MVCS1IPURL                       string                        `json:"m-vcs1_ip_url,omitempty" yaml:"mVCS1IPURL,omitempty"`
	MVCS1URL                         string                        `json:"m-vcs1_url,omitempty" yaml:"mVCS1URL,omitempty"`
	MNSX1IPURL                       string                        `json:"m-nsx1_ip_url,omitempty" yaml:"mNSX1IPURL,omitempty"`
	MNSX1URL                         string                        `json:"m-nsx1_url,omitempty" yaml:"mNSX1URL,omitempty"`
	MNSX1AIPURL                      string                        `json:"m-nsx1a_ip_url,omitempty" yaml:"mNSX1AIPURL,omitempty"`
	MNSX1AURL                        string                        `json:"m-nsx1a_url,omitempty" yaml:"mNSX1AURL,omitempty"`
	MNSX1BIPURL                      string                        `json:"m-nsx1b_ip_url,omitempty" yaml:"mNSX1BIPURL,omitempty"`
	MNSX1BURL                        string                        `json:"m-nsx1b_url,omitempty" yaml:"mNSX1BURL,omitempty"`
	MNSX1CIPURL                      string                        `json:"m-nsx1c_ip_url,omitempty" yaml:"mNSX1CIPURL,omitempty"`
	MNSX1CURL                        string                        `json:"m-nsx1c_url,omitempty" yaml:"mNSX1CURL,omitempty"`
}

// UnmarshalJSON to handle the shitty [] instead of {} or undefined returned by the serverside.
func (s *AppVMWareVCF) UnmarshalJSON(data []byte) error {

	var v struct {
		VSphereWorkload                  interface{} `json:"vcf_workload,omitempty" yaml:"vcfWorkload,omitempty"`
		VSphereManagement                interface{} `json:"vcf_management,omitempty" yaml:"vcfManagement,omitempty"`
		AdminUsername                    string      `json:"admin_username,omitempty" yaml:"adminUsername,omitempty"`
		AdminPassword                    string      `json:"admin_password,omitempty" yaml:"adminPassword,omitempty"`
		ClusterSoftwareAvailableVersions []string    `json:"cluster_software_available_versions,omitempty" yaml:"clusterSoftwareAvailableVersions,omitempty"`
		ClusterSoftwareVersion           string      `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
		Type                             string      `json:"type,omitempty" yaml:"type,omitempty"`
		InstanceVCenterServerManagement  string      `json:"instance_vcenter_server_management,omitempty" yaml:"instanceVcenterServerManagement,omitempty"`
		InstanceVcenterWebClient         string      `json:"instance_vcenter_web_client,omitempty" yaml:"instanceVcenterWebClient,omitempty"`
		VCSAUsername                     string      `json:"vcsa_username,omitempty" yaml:"vcsaUsername,omitempty"`
		VCSAInitialPassword              string      `json:"vcsa_initial_password,omitempty" yaml:"vcsaUsername,omitempty"`
		CBAAdminUsername                 string      `json:"cba_admin_username,omitempty" yaml:"cbaAdminUsername,omitempty"`
		CBAAdminPassword                 string      `json:"cba_admin_password,omitempty" yaml:"cbaAdminPassword,omitempty"`
		CBARootUsername                  string      `json:"cba_root_username,omitempty" yaml:"cbaRootUsername,omitempty"`
		CBARootPassword                  string      `json:"cba_root_password,omitempty" yaml:"cbaRootPassword,omitempty"`
		SDDCRootUsername                 string      `json:"sddc_root_username,omitempty" yaml:"sddcRootUsername,omitempty"`
		SDDCRootPassword                 string      `json:"sdcc_root_password,omitempty" yaml:"sddcRootPassword,omitempty"`
		SDDCLocalAdminUsername           string      `json:"sddc_local_admin_username,omitempty" yaml:"sddcLocalAdminUsername,omitempty"`
		SDDCLocalAdminPassword           string      `json:"sdcc_local_admin_password,omitempty" yaml:"sddcLocalAdminPassword,omitempty"`
		SDDCSecondUsername               string      `json:"sddc_second_username,omitempty" yaml:"sddcSecondUsername,omitempty"`
		SDDCSecondPassword               string      `json:"sdcc_second_password,omitempty" yaml:"sddcSecondPassword,omitempty"`
		NSXManagerRootUsername           string      `json:"nsx_manager_root_username,omitempty" yaml:"nsxManagerRootUsername,omitempty"`
		NSXManagerRootPassword           string      `json:"nsx_manager_root_password,omitempty" yaml:"nsxManagerRootPassword,omitempty"`
		NSXManagerAdminUsername          string      `json:"nsx_manager_admin_username,omitempty" yaml:"nsxManagerAdminUsername,omitempty"`
		NSXManagerAdminPassword          string      `json:"nsx_manager_admin_password,omitempty" yaml:"nsxManagerAdminPassword,omitempty"`
		NSXManagerAuditUsername          string      `json:"nsx_manager_audit_username,omitempty" yaml:"nsxManagerAuditUsername,omitempty"`
		NSXManagerAuditPassword          string      `json:"nsx_manager_audit_password,omitempty" yaml:"nsxManagerAuditPassword,omitempty"`
		VCenterRootUsername              string      `json:"vcenter_root_username,omitempty" yaml:"vcenterRootUsername,omitempty"`
		VCenterRootPassword              string      `json:"vcenter_root_password,omitempty" yaml:"vcenterRootPassword,omitempty"`
		VCenterSSOAdminUsername          string      `json:"vcenter_sso_admin_username,omitempty" yaml:"vcenterSSOAdminUsername,omitempty"`
		VCenterSSOAdminPassword          string      `json:"vcenter_sso_admin_password,omitempty" yaml:"vcenterSSOAdminPassword,omitempty"`
		CBAIPURL                         string      `json:"cba_ip_url,omitempty" yaml:"cbaIPUrl,omitempty"`
		CBAURL                           string      `json:"cba_url,omitempty" yaml:"cbaURL,omitempty"`
		SDDCIPURL                        string      `json:"sddc_ip_url,omitempty" yaml:"sddcIPUrl,omitempty"`
		MVCS1IPURL                       string      `json:"m-vcs1_ip_url,omitempty" yaml:"mVCS1IPURL,omitempty"`
		MVCS1URL                         string      `json:"m-vcs1_url,omitempty" yaml:"mVCS1URL,omitempty"`
		MNSX1IPURL                       string      `json:"m-nsx1_ip_url,omitempty" yaml:"mNSX1IPURL,omitempty"`
		MNSX1URL                         string      `json:"m-nsx1_url,omitempty" yaml:"mNSX1URL,omitempty"`
		MNSX1AIPURL                      string      `json:"m-nsx1a_ip_url,omitempty" yaml:"mNSX1AIPURL,omitempty"`
		MNSX1AURL                        string      `json:"m-nsx1a_url,omitempty" yaml:"mNSX1AURL,omitempty"`
		MNSX1BIPURL                      string      `json:"m-nsx1b_ip_url,omitempty" yaml:"mNSX1BIPURL,omitempty"`
		MNSX1BURL                        string      `json:"m-nsx1b_url,omitempty" yaml:"mNSX1BURL,omitempty"`
		MNSX1CIPURL                      string      `json:"m-nsx1c_ip_url,omitempty" yaml:"mNSX1CIPURL,omitempty"`
		MNSX1CURL                        string      `json:"m-nsx1c_url,omitempty" yaml:"mNSX1CURL,omitempty"`
	}

	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	copier.Copy(&s, &v)

	s.VSphereManagement = map[string]AppInstanceDetails{}
	if _, ok := v.VSphereManagement.([]interface{}); !ok {
		for i, x := range v.VSphereManagement.(map[string]interface{}) {
			var arr = x.(map[string]interface{})
			s.VSphereManagement[i] = AppInstanceDetails{
				InstanceID:         int(arr["instance_id"].(float64)),
				InstanceLabel:      arr["instance_label"].(string),
				InstanceHostname:   arr["instance_hostname"].(string),
				InstanceClusterUrl: arr["instance_cluster_url"].(string),
				InstanceHealth:     arr["instance_health"].(string),
				ESXIUsername:       arr["esxi_username"].(string),
				ESXIPassword:       arr["esxi_password"].(string),
				Type:               arr["type"].(string),
			}
		}
	}

	s.VSphereWorkload = map[string]AppInstanceDetails{}
	if _, ok := v.VSphereWorkload.([]interface{}); !ok {
		for i, x := range v.VSphereWorkload.(map[string]interface{}) {
			var arr = x.(map[string]interface{})
			s.VSphereWorkload[i] = AppInstanceDetails{
				InstanceID:         int(arr["instance_id"].(float64)),
				InstanceLabel:      arr["instance_label"].(string),
				InstanceHostname:   arr["instance_hostname"].(string),
				InstanceClusterUrl: arr["instance_cluster_url"].(string),
				InstanceHealth:     arr["instance_health"].(string),
				ESXIUsername:       arr["esxi_username"].(string),
				ESXIPassword:       arr["esxi_password"].(string),
				Type:               arr["type"].(string),
			}
		}
	}

	return nil
}

type AppVMWareVCFWrapper struct {
	ClusterApp AppVMWareVCF `json:"cluster_app,omitempty" yaml:"clusterApp,omitempty"`
	Type       string       `json:"type,omitempty" yaml:"type,omitempty"`
}

type AppKubernetes struct {
	KubernetesWorker                 map[string]AppInstanceDetails `json:"kubernetes_worker,omitempty" yaml:"kubernetesWorker,omitempty"`
	KubernetesMaster                 map[string]AppInstanceDetails `json:"kubernetes_master,omitempty" yaml:"kubernetesMaster,omitempty"`
	AdminUsername                    string                        `json:"admin_username,omitempty" yaml:"adminUsername,omitempty"`
	AdminPassword                    string                        `json:"admin_password,omitempty" yaml:"adminPassword,omitempty"`
	ClusterSoftwareAvailableVersions []string                      `json:"cluster_software_available_versions,omitempty" yaml:"clusterSoftwareAvailableVersions,omitempty"`
	ClusterSoftwareVersion           string                        `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
	Type                             string                        `json:"type,omitempty" yaml:"type,omitempty"`
}

type AppKubernetesWrapper struct {
	ClusterApp AppKubernetes `json:"cluster_app,omitempty" yaml:"clusterApp,omitempty"`
	Type       string        `json:"type,omitempty" yaml:"type,omitempty"`
}

type AppKubernetesEKSAConnectableClusters struct {
	Cluster          interface{} `json:"cluster,omitempty" yaml:"cluster,omitempty"`
	ContainerCluster interface{} `json:"container_cluster,omitempty" yaml:"containerCluster,omitempty"`
}

type AppKubernetesEKSA struct {
	KubernetesEKSAMgmt               []AppInstanceDetails                 `json:"eks_mgmt,omitempty" yaml:"eksMgmt,omitempty"`
	KubernetesMaster                 []AppInstanceDetails                 `json:"kubernetes_control_plane,omitempty" yaml:"kubernetesMaster,omitempty"`
	KubernetesWorker                 []AppInstanceDetails                 `json:"kubernetes_worker,omitempty" yaml:"kubernetesWorker,omitempty"`
	AdminUsername                    string                               `json:"admin_username,omitempty" yaml:"adminUsername,omitempty"`
	AdminPassword                    string                               `json:"admin_password,omitempty" yaml:"adminPassword,omitempty"`
	ClusterSoftwareAvailableVersions []string                             `json:"cluster_software_available_versions,omitempty" yaml:"clusterSoftwareAvailableVersions,omitempty"`
	ClusterSoftwareVersion           string                               `json:"cluster_software_version,omitempty" yaml:"clusterSoftwareVersion,omitempty"`
	ConnectableClusters              AppKubernetesEKSAConnectableClusters `json:"connectable_clusters,omitempty" yaml:"connectableClusters,omitempty"`
	Type                             string                               `json:"type,omitempty" yaml:"type,omitempty"`
}

type AppKubernetesEKSAWrapper struct {
	ClusterApp AppKubernetesEKSA `json:"cluster_app,omitempty" yaml:"clusterApp,omitempty"`
	Type       string            `json:"type,omitempty" yaml:"type,omitempty"`
}

const (
	CLUSTER_TYPE_VMWARE_VSPHERE  string = "vmware_vsphere"
	CLUSTER_TYPE_VMWARE_VCF      string = "vmware_vcf"
	CLUSTER_TYPE_KUBERNETES      string = "kubernetes"
	CLUSTER_TYPE_KUBERNETES_EKSA string = "kubernetes_eksa"
)

// clusterCreate creates an application cluster
func (c *Client) clusterCreate(infrastructureID id, cluster Cluster) (*Cluster, error) {
	var createdObject Cluster

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_create",
		infrastructureID,
		cluster)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// clusterCreate returns an cluster (app) with specified id
func (c *Client) clusterGet(clusterID id) (*Cluster, error) {
	var createdObject Cluster

	if err := checkID(clusterID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_get",
		clusterID)

	if err != nil {
		return nil, err
	}
	/*
		//doesn't seem to work, the passwords are returned in the wrong format
			if decryptCredentials {

				for k, v := range createdObject.ClusterCustom {

					val, err := c.decryptIfEncrypted(v)
					if err != nil {
						continue
					}

					createdObject.ClusterCustom[k] = val
				}

				var m map[string]string

				err := json.Unmarshal([]byte(createdObject.ClusterCustomJSON), &m)
				if err != nil {
					return nil, err
				}
				for k, v := range m {

					val, err := c.decryptIfEncrypted(v)
					if err != nil {
						continue // we ignore errors
					}

					m[k] = val
				}

				bytes, err := json.Marshal(m)
				if err != nil {
					return nil, err
				}

				createdObject.ClusterCustomJSON = string(bytes)

			}
	*/
	return &createdObject, nil
}

// clusterAppVMWare returns details for a vmware cluster
func (c *Client) clusterAppVMWareVSphere(clusterID id, decryptCredentials bool) (*AppVMWareVsphere, error) {

	var createdObject AppVMWareVsphereWrapper

	if err := checkID(clusterID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_app",
		clusterID)

	if err != nil {
		return nil, err
	}

	if decryptCredentials {
		createdObject.ClusterApp.AdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.AdminPassword)
		createdObject.ClusterApp.VCSAInitialPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.VCSAInitialPassword)
		for label, inst := range createdObject.ClusterApp.VSphereMaster {
			s, _ := c.decryptIfEncrypted(createdObject.ClusterApp.VSphereMaster[label].ESXIPassword)
			inst.ESXIPassword = s
			createdObject.ClusterApp.VSphereMaster[label] = inst
		}

		for label, inst := range createdObject.ClusterApp.VSphereWorker {
			s, _ := c.decryptIfEncrypted(createdObject.ClusterApp.VSphereWorker[label].ESXIPassword)
			inst.ESXIPassword = s
			createdObject.ClusterApp.VSphereWorker[label] = inst
		}
	}

	return &createdObject.ClusterApp, nil
}

// clusterAppVMWare returns details for a vmware cluster
func (c *Client) clusterAppVMWareVCF(clusterID id, decryptCredentials bool) (*AppVMWareVCF, error) {

	var createdObject AppVMWareVCFWrapper

	if err := checkID(clusterID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_app",
		clusterID)

	if err != nil {
		return nil, err
	}

	if decryptCredentials {
		createdObject.ClusterApp.AdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.AdminPassword)
		createdObject.ClusterApp.VCSAInitialPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.VCSAInitialPassword)
		createdObject.ClusterApp.CBAAdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.CBAAdminPassword)
		createdObject.ClusterApp.CBARootPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.CBARootPassword)
		createdObject.ClusterApp.NSXManagerAdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.NSXManagerAdminPassword)
		createdObject.ClusterApp.NSXManagerAuditPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.NSXManagerAuditPassword)
		createdObject.ClusterApp.NSXManagerRootPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.NSXManagerRootPassword)
		createdObject.ClusterApp.SDDCLocalAdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.SDDCLocalAdminPassword)
		createdObject.ClusterApp.SDDCRootPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.SDDCRootPassword)
		createdObject.ClusterApp.SDDCSecondPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.SDDCSecondPassword)
		createdObject.ClusterApp.VCenterRootPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.VCenterRootPassword)
		createdObject.ClusterApp.VCenterSSOAdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.VCenterSSOAdminPassword)

		for label, inst := range createdObject.ClusterApp.VSphereManagement {
			s, _ := c.decryptIfEncrypted(createdObject.ClusterApp.VSphereManagement[label].ESXIPassword)
			inst.ESXIPassword = s
			createdObject.ClusterApp.VSphereManagement[label] = inst
		}

		for label, inst := range createdObject.ClusterApp.VSphereWorkload {
			s, _ := c.decryptIfEncrypted(createdObject.ClusterApp.VSphereWorkload[label].ESXIPassword)
			inst.ESXIPassword = s
			createdObject.ClusterApp.VSphereWorkload[label] = inst
		}
	}

	return &createdObject.ClusterApp, nil
}

// clusterAppKubernetes returns details for a kubernetes cluster
func (c *Client) clusterAppKubernetes(clusterID id, decryptCredentials bool) (*AppKubernetes, error) {
	var createdObject AppKubernetesWrapper

	if err := checkID(clusterID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_app",
		clusterID)

	if err != nil {
		return nil, err
	}
	if decryptCredentials {
		createdObject.ClusterApp.AdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.AdminPassword)
	}

	return &createdObject.ClusterApp, nil
}

// clusterAppKubernetes returns details for a kubernetes EKSA cluster
func (c *Client) clusterAppKubernetesEKSA(clusterID id, decryptCredentials bool) (*AppKubernetesEKSA, error) {
	var createdObject AppKubernetesEKSAWrapper

	if err := checkID(clusterID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_app",
		clusterID)

	if err != nil {
		return nil, err
	}
	if decryptCredentials {
		createdObject.ClusterApp.AdminPassword, _ = c.decryptIfEncrypted(createdObject.ClusterApp.AdminPassword)
	}

	return &createdObject.ClusterApp, nil
}

// clusterDelete deletes an instance array. Requires deploy.
func (c *Client) clusterDelete(clusterID id) error {

	if err := checkID(clusterID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"cluster_delete",
		clusterID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// clusterEdit alterns a deployed instance array. Requires deploy.
func (c *Client) clusterEdit(clusterId id, clusterOperation ClusterOperation) (*Cluster, error) {
	var createdObject Cluster

	if err := checkID(clusterId); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"cluster_edit",
		clusterId,
		clusterOperation)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// clusterInstanceArrays returns the list of instance arrays belonging to this cluster
func (c *Client) clusters(infrastructureId id) (*map[string]Cluster, error) {

	if err := checkID(infrastructureId); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"clusters",
		infrastructureId,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Cluster{}
		return &m, nil
	}

	var createdObject map[string]Cluster

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// clusterInstanceArrays returns the list of instance arrays belonging to this cluster
func (c *Client) clusterInstanceArrays(clusterId id) (*map[string]InstanceArray, error) {

	if err := checkID(clusterId); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"cluster_instance_arrays",
		clusterId,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]InstanceArray{}
		return &m, nil
	}

	var createdObject map[string]InstanceArray

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}
