package metalcloud

import (
	"fmt"
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

type AppKubernetes struct {
	KubernetesNodes                  map[string]AppInstanceDetails `json:"vsphere_worker,omitempty" yaml:"vsphereWorker,omitempty"`
	KubernetesMaster                 map[string]AppInstanceDetails `json:"vsphere_master,omitempty" yaml:"vsphereWorker,omitempty"`
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

const (
	CLUSTER_TYPE_VMWARE_VSPHERE  string = "vmware_vsphere"
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
