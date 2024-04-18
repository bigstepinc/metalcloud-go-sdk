//go:generate go run helper/docgen.go - $GOFILE ./ DatacenterWithConfig,Datacenter,DatacenterConfig,Option82ToIPMapping,WebProxy DatacenterWithConfig

package metalcloud

import (
	"encoding/json"
	"fmt"
	"strings"
)

// description: A data center object that contains both metadata and configuration
// examples:
//   - value: exampleDCYaml
type DatacenterWithConfig struct {
	// description: The datacenter part of the object
	Metadata Datacenter `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	//description: The datacenter configuration part of the object
	Config DatacenterConfig `json:"config,omitempty" yaml:"config,omitempty"`
}

// description: Datacenter metadata
type Datacenter struct {
	//description: The ID of this datacenter.
	DatacenterID int `json:"datacenter_id,omitempty" yaml:"id,omitempty"`
	//description: The name (label) of this datacenter. Once set it cannot be changed.
	DatacenterName string `json:"datacenter_name,omitempty" yaml:"name,omitempty"`
	//description: The name (label) of the parent datacenter. This is useful in hierarchical setups where one datacenter needs to access it's parent's resources.
	DatacenterNameParent string `json:"datacenter_name_parent,omitempty" yaml:"parentName,omitempty"`
	//description: The owner of a datacenter.
	UserID int `json:"user_id,omitempty" yaml:"userid,omitempty"`
	//description: The display name of a data center. Can be changed.
	DatacenterDisplayName string `json:"datacenter_display_name,omitempty" yaml:"displayname,omitempty"`
	//description: Deprecated.
	DatacenterIsMaster bool `json:"datacenter_is_master" yaml:"ismaster"`
	//description: If set to true no new operations can happen on this datacenter.
	DatacenterIsMaintenance bool `json:"datacenter_is_maintenance" yaml:"ismaintenance"`
	//description: The datacenter type. Deprecated. Currently the only supported value is metal_cloud.
	//values:
	// - metal_cloud
	DatacenterType string `json:"datacenter_type,omitempty" yaml:"type,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the datacenter was created.
	//example: 2013-11-29T13:00:01Z
	DatacenterCreatedTimestamp string `json:"datacenter_created_timestamp,omitempty" yaml:"createdtimestamp,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the datacenter was updated.
	//example: 2013-11-29T13:00:01Z
	DatacenterUpdatedTimestamp string `json:"datacenter_updated_timestamp,omitempty" yaml:"updatedtimestamp,omitempty"`
	//description: If set the datacenter will not be visible in the UI
	DatacenterHidden bool `json:"datacenter_hidden" yaml:"ishidden"`
	//description: An array of tags (strings)
	DatacenterTags []string `json:"datacenter_tags,omitempty" yaml:"tags,omitempty"`
}

// DatacenterConfig - datacenter configuration
type DatacenterConfig struct {
	//description: The ip address of the Global Controller. Deprecated.
	//example: 192.168.137.1
	BSIMachinesSubnetIPv4CIDR string `json:"BSIMachinesSubnetIPv4CIDR,omitempty" yaml:"BSIMachinesSubnetIPv4CIDR,omitempty"`
	//description: The ip address on which all datacenter agents listen for connections. Deprecated.
	//example: 192.168.137.1
	BSIVRRPListenIPv4 string `json:"BSIVRRPListenIPv4,omitempty" yaml:"BSIVRRPListenIPv4,omitempty"`
	//description: Site Controller's secondary ip addresses. Deprecated.
	//example:  192.168.137.1,192.168.137.2,192.168.137.4
	BSIMachineListenIPv4List []string `json:"BSIMachineListenIPv4List,omitempty" yaml:"BSIMachineListenIPv4List,omitempty"`
	//description: The agent's IP that is visible from the controller. Deprecated.
	//example:  89.36.24.2
	BSIExternallyVisibleIPv4 string `json:"BSIExternallyVisibleIPv4,omitempty" yaml:"BSIExternallyVisibleIPv4,omitempty"`
	//description: The repository to use
	//example: https://uk-reading-repo.example.com
	RepoURLRoot string `json:"repoURLRoot,omitempty" yaml:"repoURLRoot,omitempty"`
	//description: The repository to use during legacy (PXE) provisioning process. Same as repoURLRoot, with an IP address for the hostname, required in networks where DNS is not available.
	//example: https://192.178.1.1
	RepoURLRootQuarantineNetwork string `json:"repoURLRootQuarantineNetwork,omitempty" yaml:"repoURLRootQuarantineNetwork,omitempty"`
	//description: The SAN subnet in CIDR format.
	//example: 100.96.0.0/16
	SANRoutedSubnet string `json:"SANRoutedSubnet,omitempty" yaml:"SANRoutedSubnet,omitempty"`
	//description: IP addresses of NTP servers.
	//example: 8.8.8.8, 8.8.4.4
	NTPServers []string `json:"NTPServers" yaml:"NTPServers"`
	//description: IP addresses of DNS servers to be used in the DHCP response.
	//example: 8.8.8.8, 8.8.4.4
	DNSServers []string `json:"DNSServers,omitempty" yaml:"DNSServers,omitempty"`
	//description: Host (IP:port) of the Windows machine hosting the Key Management Service. Set to empty string to disable.
	//example: 84.40.58.70:1688
	KMS string `json:"KMS,omitempty" yaml:"KMS,omitempty"`
	//description: The IP of the Site Controller TFTP service used during the legacy (PXE) deployment process.
	//example: 89.36.24.2
	TFTPServerWANVRRPListenIPv4 string `json:"TFTPServerWANVRRPListenIPv4,omitempty" yaml:"TFTPServerWANVRRPListenIPv4,omitempty"`
	//description: If set to true, the datalake service is enabled in this environment. Deprecated
	//values:
	// - true
	// - false
	DataLakeEnabled bool `json:"dataLakeEnabled" yaml:"dataLakeEnabled"`
	//description: Graphite host (IPv4:port) for the plain text protocol socket. Set to empty string to disable. Deprecated
	MonitoringGraphitePlainTextSocketHost string `json:"monitoringGraphitePlainTextSocketHost,omitempty" yaml:"monitoringGraphitePlainTextSocketHost,omitempty"`
	//description: Graphite host (IPv4:port) for the HTTP Render URL API. Set to empty string to disable. Deprecated
	//example: 192.168.137.2:80
	MonitoringGraphiteRenderURLHost string `json:"monitoringGraphiteRenderURLHost,omitempty" yaml:"monitoringGraphiteRenderURLHost,omitempty"`
	//description: The Datacenter's latitude. Use negative numbers for the south hemisphere
	//example: 41.8426146
	Latitude float64 `json:"latitude,omitempty" yaml:"latitude,omitempty"`
	//description: The data center's longitude: Use negative numbers for areas west of Greenwich (UK)
	//example: -87.6695014
	Longitude float64 `json:"longitude,omitempty" yaml:"longitude,omitempty"`
	//description: The data center's address
	//example: 2800 S Ashland Ave, Chicago, IL 60608, United States
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
	//description: If set to true the system will configure a randomly generated username and password on the server's BMC(ILO/IDRAC etc.)
	//values:
	// - true
	// - false
	ServerRegisterUsingGeneratedIPMICredentialsEnabled bool `json:"serverRegisterUsingGeneratedIPMICredentialsEnabled" yaml:"serverRegisterUsingGeneratedIPMICredentialsEnabled"`
	//description: If set to true the system will ask for credentials during server registration.
	//values:
	// - true
	// - false
	ServerRegisterUsingProvidedIPMICredentialsEnabled bool `json:"serverRegisterUsingProvidedIPMICredentialsEnabled" yaml:"serverRegisterUsingProvidedIPMICredentialsEnabled"`
	//description: The provisioner (fabric) to use when provisioning the network on switch devices
	//values:
	// - VLAN
	// - EVPNVXLANL2
	// - VPLS
	// - LAN
	// - SDN
	SwitchProvisioner map[string]interface{} `json:"switchProvisioner,omitempty" yaml:"switchProvisioner,omitempty"`
	//description: If set to true the tenants will receive credentials for accessing the server's BMC with a special user.
	EnableTenantAccessToIPMI bool `json:"enableTenantAccessToIPMI" yaml:"enableTenantAccessToIPMI"`
	//description: Allows the end-user to force a VLAN ID (or EPG in CISCO ACI environments). This enables the user to connect to pre-existing VLANs in the established infrastructure. WARNING: This enables a tenant to access unauthorized VLANs.
	AllowVLANOverrides bool `json:"allowVLANOverrides" yaml:"allowVLANOverrides"`
	//description: Allows the usage of network profiles for customizing InstanceArray network connections.
	AllowNetworkProfiles bool `json:"allowNetworkProfiles" yaml:"allowNetworkProfiles"`
	//description: If set enables in-band triggered registration via the legacy (PXE) mechanism.
	EnableServerRegistrationStartedByInBandDHCP bool `json:"enableServerRegistrationStartedByInBandDHCP" yaml:"enableServerRegistrationStartedByInBandDHCP"`
	//description: Extra ips to reserve on each subnet for WAN networks. Certain fabrics (such as VRRP-based L3 SVIs need more than one IP to be allocated on each subnet). This option will force the system to reserve this number of IPs from each subnet.
	ExtraInternalIPsPerSubnet int `json:"extraInternalIPsPerSubnet" yaml:"extraInternalIPsPerSubnet"`
	//description: Extra ips to reserve on each subnet for SAN networks. Certain fabrics (such as VRRP-based L3 SVIs need more than one IP to be allocated on each subnet). This option will force the system to reserve this number of IPs from each subnet.
	ExtraInternalIPsPerSANSubnet int `json:"extraInternalIPsPerSANSubnet" yaml:"extraInternalIPsPerSANSubnet"`
	//description: If enabled RAID configurations are set on servers
	ServerRAIDConfigurationEnabled bool `json:"serverRAIDConfigurationEnabled" yaml:"serverRAIDConfigurationEnabled"`
	//description: If configured the proxy will be used by all operations.
	WebProxy *WebProxy `json:"webProxy" yaml:"webProxy"`
	//description: Deprecated.
	IsKubernetesDeployment bool `json:"isKubernetesDeployment" yaml:"isKubernetesDeployment"`
	//description: If set it allows  the use of firmware policies. Note that for baselines to function this needs to be enabled.
	AllowInstanceArrayFirmwarePolicies bool `json:"allowInstanceArrayFirmwarePolicies" yaml:"allowInstanceArrayFirmwarePolicies"`
	//description: If set to true, during the legacy registration process (PXE) the system will configure special provisioning VLAN on server ports prior to performing the deployment
	ProvisionUsingTheQuarantineNetwork bool `json:"provisionUsingTheQuarantineNetwork" yaml:"provisionUsingTheQuarantineNetwork"`
	//description: If set to true, during the legacy registration process (PXE) the system will enforce DHCP option 82 security.
	EnableDHCPRelaySecurityForQuarantineNetwork bool `json:"enableDHCPRelaySecurityForQuarantineNetwork" yaml:"enableDHCPRelaySecurityForQuarantineNetwork"`
	//description: If set to true, the DHCP server will ignore requests that do not respect DHCP option 82 for regular networks.
	EnableDHCPRelaySecurityForClientNetworks bool `json:"enableDHCPRelaySecurityForClientNetworks" yaml:"enableDHCPRelaySecurityForClientNetworks"`
	//description: If enabled, the DHCPBMCMACAddressWhitelist will be used to whitelist certain MAC addresses in order to ensure that only certain servers get registered during the ZTP process.
	EnableDHCPBMCMACAddressWhitelist bool `json:"enableDHCPBMCMACAddressWhitelist" yaml:"enableDHCPBMCMACAddressWhitelist"`
	//description: The mac addresses of the servers that are to be allowed to be registered via ZTP. This is useful during initial testing.
	DHCPBMCMACAddressWhitelist []string `json:"dhcpBMCMACAddressWhitelist" yaml:"dhcpBMCMACAddressWhitelist"`
	//description: If set the server cleanup policy will be the policy with the specified id instead of the default one (which is 0)
	DefaultServerCleanupPolicyID int `json:"defaultServerCleanupPolicyID" yaml:"defaultServerCleanupPolicyID"`
	//description: If set, this will be the default network profile instead of no network profile.
	DefaultWANNetworkProfileID int `json:"defaultWANNetworkProfileID" yaml:"defaultWANNetworkProfileID"`
	//description: Deployment mechanism used in case a server supports both Virtual Media and legacy (PXE).
	// values:
	// - virtual_media
	// - pxe
	DefaultDeploymentMechanism string `json:"defaultDeploymentMechanism" yaml:"defaultDeploymentMechanism"`
	//description: The cleanup and register mechanism used in case a server supports both BMC-only and BDK mechanisms. Defaults to BMC.
	// values:
	// - bmc
	// - bdk
	DefaultCleanupAndRegistrationMechanism string `json:"defaultCleanupAndRegistrationMechanism" yaml:"defaultCleanupAndRegistrationMechanism"`
	//description: The NFS server to use for server OS deployment (the IP of the site controller as seen from the server's BMC). Should be an IP to avoid DNS resolutions.
	//example: IP address of the NFS server
	NFSServer string `json:"NFSServer" yaml:"NFSServer"`
	//description: Can be used to set a mapping between Option82 and IPs that the DHCP server allocates to servers during registration.
	Option82ToIPMapping Option82ToIPMapping `json:"Option82ToIPMapping" yaml:"Option82ToIPMapping"`
}

// description: Defines web proxy configuration
type WebProxy struct {
	//description: Ip fo the web proxy
	WebProxyServerIP string `json:"web_proxy_server_ip,omitempty" yaml:"ip,omitempty"`
	//description: Port fo the web proxy
	WebProxyServerPort int `json:"web_proxy_server_port,omitempty" yaml:"port,omitempty"`
	//description: Username of the web proxy
	WebProxyUsername string `json:"web_proxy_username,omitempty" yaml:"username,omitempty"`
	//description: Password to use for the web proxy
	WebProxyPassword string `json:"web_proxy_password,omitempty" yaml:"password,omitempty"`
}

type Option82ToIPMapping map[string]string

func (m *Option82ToIPMapping) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && string(data) == "[]" {
		*m = make(Option82ToIPMapping)
		return nil
	}
	return json.Unmarshal(data, (*map[string]string)(m))
}

/*
//SwitchProvisioner provisioner base struct
type SwitchProvisioner struct {
	Provisioner interface{}
	Type        string
}

//VLANProvisioner - defines settings for the networking provisioning architecture that uses vlans
type VLANProvisioner struct {
	LANVLANRange     string `json:"LANVLANRange,omitempty"`
	WANVLANRange     string `json:"WANVLANRange,omitempty"`
	QuarantineVLANID int    `json:"quarantineVLANID,omitempty"`
	Type             string `json:"type,omitempty"`
}

//VPLSProvisioner - defines settings for the networking provisioning architecture that uses vpls
type VPLSProvisioner struct {
	ACLSAN            int    `json:"ACLSAN,omitempty"`
	ACLWAN            int    `json:"ACLWAN,omitempty"`
	SANACLRange       string `json:"SANACLRange,omitempty"`
	ToRLANVLANRange   string `json:"ToRLANVLANRange,omitempty"`
	ToRSANVLANRange   string `json:"ToRSANVLANRange,omitempty"`
	ToRWANVLANRange   string `json:"ToRWANVLANRange,omitempty"`
	QuarantineVLANID  int    `json:"quarantineVLANID,omitempty"`
	NorthWANVLANRange string `json:"NorthWANVLANRange,omitempty"`
	Type              string `json:"type,omitempty"`
}

//EVPNVXLANL2Provisioner - defines settings for the networking provisioning architecture that uses evpnvxlanl2
type EVPNVXLANL2Provisioner struct {
	allocateDefaultWANVLAN bool `json:"allocateDefaultWANVLAN,omitempty"`
	allocateDefaultSANVLAN bool `json:"allocateDefaultSANVLAN,omitempty"`
	allocateDefaultLANVLAN bool `json:"allocateDefaultLANVLAN,omitempty"`
	preventCleanupForVLANs []int `json:"preventCleanupForVLANs,omitempty"`
	LANVLANRange     	string `json:"LANVLANRange,omitempty"`
	SANVLANRange     	string `json:"SANVLANRange,omitempty"`
	WANVLANRange     	string `json:"WANVLANRange,omitempty"`
	QuarantineVLANID 	int    `json:"quarantineVLANID,omitempty"`
	StorageHasSeparateFabric bool `json:"storageHasSeparateFabric,omitempty"`
	Type             	string `json:"type,omitempty"`
}

//UnmarshalJSON custom unmarshaling
func (o *SwitchProvisioner) UnmarshalJSON(data []byte) error {
	var p struct {
		Type string `json:"type,omitempty"`
	}

	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	o.Type = p.Type

	switch p.Type {
	case "VLANProvisioner":

		var provisioner VLANProvisioner
		err := json.Unmarshal(data, &provisioner)
		if err != nil {
			return err
		}
		o.Provisioner = provisioner

	case "VPLSProvisioner":
		var provisioner VPLSProvisioner

		err := json.Unmarshal(data, &provisioner)
		if err != nil {
			return err
		}

		o.Provisioner = provisioner
	default:
		return fmt.Errorf("Cannot unmarshal unsupported provisioner type %s", p.Type)
	}

	return nil
}

//MarshalJSON custom marshaling
func (o *SwitchProvisioner) MarshalJSON() ([]byte, error) {

	switch o.Type {
	case "VLANProvisioner":

		provisioner := o.Provisioner.(VLANProvisioner)
		provisioner.Type = o.Type

		return json.Marshal(&provisioner)

	case "VPLSProvisioner":
		provisioner := o.Provisioner.(VPLSProvisioner)
		provisioner.Type = o.Type
		return json.Marshal(&provisioner)

	default:
		return nil, fmt.Errorf("Cannot marshal unsupported provisioner type %s", o.Type)
	}

}
*/

// Datacenters returns datacenters for all users
func (c *Client) Datacenters(onlyActive bool) (*map[string]Datacenter, error) {
	return c.datacenters(nil, onlyActive)
}

// DatacentersByUserID returns datacenters for specific user
func (c *Client) DatacentersByUserID(userID int, onlyActive bool) (*map[string]Datacenter, error) {
	return c.datacenters(userID, onlyActive)
}

// DatacentersByUserEmail returns datacenters by email
func (c *Client) DatacentersByUserEmail(userEmail string, onlyActive bool) (*map[string]Datacenter, error) {
	return c.datacenters(userEmail, onlyActive)
}

// datacenters returns datacenters
func (c *Client) datacenters(userID id, onlyActive bool) (*map[string]Datacenter, error) {
	resp, err := c.rpcClient.Call(
		"datacenters",
		userID,
		onlyActive,
		false,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Datacenter{}
		return &m, nil
	}

	var createdObject map[string]Datacenter

	err = resp.GetObject(&createdObject)
	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// DatacenterGet returns details of a specific datacenter
func (c *Client) DatacenterGet(datacenterName string) (*Datacenter, error) {
	return c.datacenterGetForUser(datacenterName, nil)
}

// DatacenterGetForUserByEmail returns details of a specific datacenter
func (c *Client) DatacenterGetForUserByEmail(datacenterName string, userID string) (*Datacenter, error) {
	return c.datacenterGetForUser(datacenterName, userID)
}

// DatacenterGetForUserByID returns details of a specific datacenter
func (c *Client) DatacenterGetForUserByID(datacenterName string, userID int) (*Datacenter, error) {
	return c.datacenterGetForUser(datacenterName, userID)
}

// DatacenterGetForUser returns details of a specific datacenter
func (c *Client) datacenterGetForUser(datacenterName string, userID id) (*Datacenter, error) {
	var datacenter Datacenter

	err := c.rpcClient.CallFor(&datacenter,
		"datacenter_get",
		userID,
		datacenterName)

	if err != nil {
		return nil, err
	}

	return &datacenter, nil
}

// DatacenterConfigGet returns details of a specific datacenter
func (c *Client) DatacenterConfigGet(datacenterName string) (*DatacenterConfig, error) {
	var datacenterConfig DatacenterConfig

	err := c.rpcClient.CallFor(
		&datacenterConfig,
		"datacenter_config",
		datacenterName)

	if err != nil {
		return nil, err
	}

	return &datacenterConfig, nil
}

// DatacenterWithConfigGet returns details of a specific datacenter as a single object that contains the config as well
func (c *Client) DatacenterWithConfigGet(datacenterName string) (*DatacenterWithConfig, error) {
	metadata, err := c.DatacenterGet(datacenterName)
	if err != nil {
		return nil, err
	}

	config, err := c.DatacenterConfigGet(datacenterName)
	if err != nil {
		return nil, err
	}

	dc := DatacenterWithConfig{
		Metadata: *metadata,
		Config:   *config,
	}

	return &dc, nil
}

// DatacenterConfigUpdate Updates configuration information for a specified Datacenter.
func (c *Client) DatacenterConfigUpdate(datacenterName string, datacenterConfig DatacenterConfig) error {

	resp, err := c.rpcClient.Call(
		"datacenter_config_update",
		datacenterName,
		datacenterConfig,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// DatacenterCreate creates a new Datacenter
func (c *Client) DatacenterCreate(datacenter Datacenter, datacenterConfig DatacenterConfig) (*Datacenter, error) {
	var createdObj Datacenter

	err := c.rpcClient.CallFor(
		&createdObj,
		"datacenter_create",
		datacenter,
		datacenterConfig)

	if err != nil {
		return nil, err
	}

	return &createdObj, nil
}

func (c *Client) DatacenterCreateFromDatacenterWithConfig(datacenter DatacenterWithConfig) (*DatacenterWithConfig, error) {
	_, err := c.DatacenterCreate(datacenter.Metadata, datacenter.Config)
	if err != nil {
		return nil, err
	}
	return c.DatacenterWithConfigGet(datacenter.Metadata.DatacenterName)
}

func (c *Client) DatacenterUpdateFromDatacenterWithConfig(datacenter DatacenterWithConfig) (*DatacenterWithConfig, error) {

	err := c.DatacenterConfigUpdate(datacenter.Metadata.DatacenterName, datacenter.Config)
	if err != nil {
		return nil, err
	}
	return c.DatacenterWithConfigGet(datacenter.Metadata.DatacenterName)
}

//bsideveloper.datacenter_agents_config_json_download_url('uk-reading')

// structure to hold the return for datacenter_agents_config_json_download_url
type datacenterConfigJSONURL struct {
	URL string `json:"datacenter_agents_config_json_download_url,omitempty"`
}

// DatacenterAgentsConfigJSONDownloadURL returns the agent url (and automatically decrypts it)
func (c *Client) DatacenterAgentsConfigJSONDownloadURL(datacenterName string, decrypt bool) (string, error) {
	var createdObj datacenterConfigJSONURL

	err := c.rpcClient.CallFor(
		&createdObj,
		"datacenter_agents_config_json_download_url",
		datacenterName)

	if err != nil {
		return "", err
	}

	agentConfigURL := createdObj.URL

	if decrypt {
		passwdComponents := strings.Split(createdObj.URL, ":")

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return "", fmt.Errorf("permission missing. %s", passwdComponents[1])
			} else {
				var decryptedURL string

				err = c.rpcClient.CallFor(
					&decryptedURL,
					"password_decrypt",
					passwdComponents[1],
				)
				if err != nil {
					return "", err
				}

				agentConfigURL = decryptedURL
			}
		}
	}

	return agentConfigURL, nil
}

// CreateOrUpdate implements interface Applier
func (dc DatacenterWithConfig) CreateOrUpdate(client MetalCloudClient) error {

	err := dc.Validate()
	if err != nil {
		return err
	}

	datacenters, err := client.Datacenters(false)
	if err != nil {
		return err
	}

	found := false
	for _, d := range *datacenters {
		if d.DatacenterName == dc.Metadata.DatacenterName {
			found = true
			break
		}
	}

	if found {

		_, err = client.DatacenterUpdateFromDatacenterWithConfig(dc)

		if err != nil {
			return err
		}
	} else {

		_, err = client.DatacenterCreateFromDatacenterWithConfig(dc)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (dc DatacenterWithConfig) Delete(client MetalCloudClient) error {
	return fmt.Errorf("delete not currently implemented for the Datacenter object")
}

// Validate implements interface Applier
func (dc DatacenterWithConfig) Validate() error {
	if dc.Metadata.DatacenterName == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}

const exampleDCYaml = `
kind: DatacenterWithConfig
apiVersion: 1.0
metadata:
  id: 332
  name: test-us03-chi-qts01-dc-3
  displayname: US03 DC
  ismaster: false
  ismaintenance: false
  type: metal_cloud
  createdtimestamp: "2023-03-06T11:12:20Z"
  updatedtimestamp: "2024-04-12T12:09:25Z"
  ishidden: false
  tags:
  - ""
config:
  BSIMachinesSubnetIPv4CIDR: 172.18.38.16/29
  BSIVRRPListenIPv4: 172.18.38.22
  BSIMachineListenIPv4List:
  - 172.18.38.22
  BSIExternallyVisibleIPv4: 176.223.248.11
  repoURLRoot: http://repointegrationpublic.bigstepcloud.com
  repoURLRootQuarantineNetwork: http://repointegrationpublic.bigstepcloud.com
  SANRoutedSubnet: 100.96.0.0/16
  NTPServers:
  - 45.55.58.103
  DNSServers:
  - 1.1.1.1
  - 8.8.8.8
  KMS: 84.40.58.70:1688
  TFTPServerWANVRRPListenIPv4: 172.18.38.22
  dataLakeEnabled: false
  latitude: 45
  longitude: 51.509865
  serverRegisterUsingGeneratedIPMICredentialsEnabled: false
  serverRegisterUsingProvidedIPMICredentialsEnabled: true
  switchProvisioner:
    ASNRanges: []
    LAGRanges:
    - 11-19
    LANVLANRange: 200-299
    MLAGRanges:
    - 20-25
    SANVLANRange: 300-399
    SANVNIPrefix: 2
    VRFL3VNIPrefix: 9
    VRFVLANRanges: []
    WANVLANRange: 100-199
    WANVNIPrefix: 1
    allocateDefaultLANVLAN: false
    allocateDefaultSANVLAN: true
    allocateDefaultWANVLAN: true
    disableQuarantineNetwork: false
    leafSwitchesHaveMLAGPairs: false
    preventCleanupForVLANs: []
    preventCleanupForVLANsFromExternalConnectionUplinks: []
    preventCleanupForVRFs: []
    preventUsageOfVLANs:
    - 3204
    preventUsageOfVRFs: []
    quarantineVLANID: 5
    storageHasSeparateFabric: false
    type: EVPNVXLANL2Provisioner
    zeroTouchRegistrationEnabled: false
  enableTenantAccessToIPMI: true
  allowVLANOverrides: true
  allowNetworkProfiles: true
  enableServerRegistrationStartedByInBandDHCP: false
  extraInternalIPsPerSubnet: 0
  extraInternalIPsPerSANSubnet: 0
  serverRAIDConfigurationEnabled: true
  webProxy: null
  isKubernetesDeployment: false
  allowInstanceArrayFirmwarePolicies: true
  provisionUsingTheQuarantineNetwork: true
  enableDHCPRelaySecurityForQuarantineNetwork: false
  enableDHCPRelaySecurityForClientNetworks: false
  enableDHCPBMCMACAddressWhitelist: false
  dhcpBMCMACAddressWhitelist: []
  defaultServerCleanupPolicyID: 3
  defaultWANNetworkProfileID: 864
  defaultDeploymentMechanism: virtual_media
  defaultCleanupAndRegistrationMechanism: bmc
  NFSServer: 172.18.38.22
  Option82ToIPMapping:
    Eth1/1: 172.18.32.1
`
