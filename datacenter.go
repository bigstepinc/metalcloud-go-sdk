package metalcloud

import (
	"fmt"
	"strings"
)

//Datacenter - datacenter description
type Datacenter struct {
	DatacenterID               int               `json:"datacenter_id,omitempty" yaml:"id,omitempty"`
	DatacenterName             string            `json:"datacenter_name,omitempty" yaml:"name,omitempty"`
	DatacenterNameParent       string            `json:"datacenter_name_parent,omitempty" yaml:"parentName,omitempty"`
	UserID                     int               `json:"user_id,omitempty" yaml:"userid,omitempty"`
	DatacenterDisplayName      string            `json:"datacenter_display_name,omitempty" yaml:"displayname,omitempty"`
	DatacenterIsMaster         bool              `json:"datacenter_is_master" yaml:"ismaster"`
	DatacenterIsMaintenance    bool              `json:"datacenter_is_maintenance" yaml:"ismaintenance"`
	DatacenterType             string            `json:"datacenter_type,omitempty" yaml:"type,omitempty"`
	DatacenterCreatedTimestamp string            `json:"datacenter_created_timestamp,omitempty" yaml:"createdtimestamp,omitempty"`
	DatacenterUpdatedTimestamp string            `json:"datacenter_updated_timestamp,omitempty" yaml:"updatedtimestamp,omitempty"`
	DatacenterHidden           bool              `json:"datacenter_hidden" yaml:"ishidden"`
	DatacenterTags             []string          `json:"datacenter_tags,omitempty" yaml:"tags,omitempty"`
	DatacenterConfig           *DatacenterConfig `json:"datacenter_config_json,omitempty" yaml:"config,omitempty"`
}

//DatacenterConfig - datacenter configuration
type DatacenterConfig struct {
	BSIMachinesSubnetIPv4CIDR                          string                 `json:"BSIMachinesSubnetIPv4CIDR,omitempty" yaml:"BSIMachinesSubnetIPv4CIDR,omitempty"`
	BSIVRRPListenIPv4                                  string                 `json:"BSIVRRPListenIPv4,omitempty" yaml:"BSIVRRPListenIPv4,omitempty"`
	BSIMachineListenIPv4List                           []string               `json:"BSIMachineListenIPv4List,omitempty" yaml:"BSIMachineListenIPv4List,omitempty"`
	BSIExternallyVisibleIPv4                           string                 `json:"BSIExternallyVisibleIPv4,omitempty" yaml:"BSIExternallyVisibleIPv4,omitempty"`
	RepoURLRoot                                        string                 `json:"repoURLRoot,omitempty" yaml:"repoURLRoot,omitempty"`
	RepoURLRootQuarantineNetwork                       string                 `json:"repoURLRootQuarantineNetwork,omitempty" yaml:"repoURLRootQuarantineNetwork,omitempty"`
	SANRoutedSubnet                                    string                 `json:"SANRoutedSubnet,omitempty" yaml:"SANRoutedSubnet,omitempty"`
	NTPServers                                         []string               `json:"NTPServers,omitempty" yaml:"NTPServers,omitempty"`
	DNSServers                                         []string               `json:"DNSServers,omitempty" yaml:"DNSServers,omitempty"`
	KMS                                                string                 `json:"KMS,omitempty" yaml:"KMS,omitempty"`
	TFTPServerWANVRRPListenIPv4                        string                 `json:"TFTPServerWANVRRPListenIPv4,omitempty" yaml:"TFTPServerWANVRRPListenIPv4,omitempty"`
	DataLakeEnabled                                    bool                   `json:"dataLakeEnabled" yaml:"dataLakeEnabled"`
	MonitoringGraphitePlainTextSocketHost              string                 `json:"monitoringGraphitePlainTextSocketHost,omitempty" yaml:"monitoringGraphitePlainTextSocketHost,omitempty"`
	MonitoringGraphiteRenderURLHost                    string                 `json:"monitoringGraphiteRenderURLHost,omitempty" yaml:"monitoringGraphiteRenderURLHost,omitempty"`
	Latitude                                           float64                `json:"latitude,omitempty" yaml:"latitude,omitempty"`
	Longitude                                          float64                `json:"longitude,omitempty" yaml:"longitude,omitempty"`
	Address                                            string                 `json:"address,omitempty" yaml:"address,omitempty"`
	ServerRegisterUsingGeneratedIPMICredentialsEnabled bool                   `json:"serverRegisterUsingGeneratedIPMICredentialsEnabled,omitempty" yaml:"serverRegisterUsingGeneratedIPMICredentialsEnabled,omitempty"`
	ServerRegisterUsingProvidedIPMICredentialsEnabled  bool                   `json:"serverRegisterUsingProvidedIPMICredentialsEnabled,omitempty" yaml:"serverRegisterUsingProvidedIPMICredentialsEnabled,omitempty"`
	SwitchProvisioner                                  map[string]interface{} `json:"switchProvisioner,omitempty" yaml:"switchProvisioner,omitempty"`
	EnableTenantAccessToIPMI                           bool                   `json:"enableTenantAccessToIPMI" yaml:"enableTenantAccessToIPMI"`
	AllowVLANOverrides                                 bool                   `json:"allowVLANOverrides" yaml:"allowVLANOverrides"`
	AllowNetworkProfiles                               bool                   `json:"allowNetworkProfiles" yaml:"allowNetworkProfiles"`
	UseSecondarySANVLAN                                bool                   `json:"useSecondarySANVLAN" yaml:"useSecondarySANVLAN"`
	ExtraInternalIPsPerSubnet                          int                    `json:"extraInternalIPsPerSubnet" yaml:"extraInternalIPsPerSubnet"`
	ExtraInternalIPsPerSANSubnet                       int                    `json:"extraInternalIPsPerSANSubnet" yaml:"extraInternalIPsPerSANSubnet"`
	ServerRAIDConfigurationEnabled                     bool                   `json:"serverRAIDConfigurationEnabled" yaml:"serverRAIDConfigurationEnabled"`
	WebProxy                                           *WebProxy              `json:"webProxy" yaml:"webProxy"`
	IsKubernetesDeployment                             bool                   `json:"isKubernetesDeployment" yaml:"isKubernetesDeployment"`
	AllowInstanceArrayFirmwarePolicies                 bool                   `json:"allowInstanceArrayFirmwarePolicies" yaml:"allowInstanceArrayFirmwarePolicies"`
	ProvisionUsingTheQuarantineNetwork                 bool                   `json:"provisionUsingTheQuarantineNetwork" yaml:"provisionUsingTheQuarantineNetwork"`
	EnableDHCPRelaySecurityForQuarantineNetwork        bool                   `json:"enableDHCPRelaySecurityForQuarantineNetwork" yaml:"enableDHCPRelaySecurityForQuarantineNetwork"`
	EnableDHCPRelaySecurityForClientNetworks           bool                   `json:"enableDHCPRelaySecurityForClientNetworks" yaml:"enableDHCPRelaySecurityForClientNetworks"`
	EnableDHCPBMCMACAddressWhitelist                   bool                   `json:"enableDHCPBMCMACAddressWhitelist" yaml:"enableDHCPBMCMACAddressWhitelist"`
	DHCPBMCMACAddressWhitelist                         []string               `json:"dhcpBMCMACAddressWhitelist" yaml:"dhcpBMCMACAddressWhitelist"`
}

type WebProxy struct {
	WebProxyServerIP   string `json:"web_proxy_server_ip,omitempty" yaml:"ip,omitempty"`
	WebProxyServerPort int    `json:"web_proxy_server_port,omitempty" yaml:"port,omitempty"`
	WebProxyUsername   string `json:"web_proxy_username,omitempty" yaml:"username,omitempty"`
	WebProxyPassword   string `json:"web_proxy_password,omitempty" yaml:"password,omitempty"`
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

//Datacenters returns datacenters for all users
func (c *Client) Datacenters(onlyActive bool) (*map[string]Datacenter, error) {
	return c.datacenters(nil, onlyActive)
}

//DatacentersByUserID returns datacenters for specific user
func (c *Client) DatacentersByUserID(userID int, onlyActive bool) (*map[string]Datacenter, error) {
	return c.datacenters(userID, onlyActive)
}

//DatacentersByUserEmail returns datacenters by email
func (c *Client) DatacentersByUserEmail(userEmail string, onlyActive bool) (*map[string]Datacenter, error) {
	return c.datacenters(userEmail, onlyActive)
}

//datacenters returns datacenters
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

//DatacenterGet returns details of a specific datacenter
func (c *Client) DatacenterGet(datacenterName string) (*Datacenter, error) {
	return c.datacenterGetForUser(datacenterName, nil)
}

//DatacenterGetForUserByEmail returns details of a specific datacenter
func (c *Client) DatacenterGetForUserByEmail(datacenterName string, userID string) (*Datacenter, error) {
	return c.datacenterGetForUser(datacenterName, userID)
}

//DatacenterGetForUserByID returns details of a specific datacenter
func (c *Client) DatacenterGetForUserByID(datacenterName string, userID int) (*Datacenter, error) {
	return c.datacenterGetForUser(datacenterName, userID)
}

//DatacenterGetForUser returns details of a specific datacenter
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

//DatacenterConfigGet returns details of a specific datacenter
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

//DatacenterConfigUpdate Updates configuration information for a specified Datacenter.
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

//DatacenterCreate creates a new Datacenter
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

//bsideveloper.datacenter_agents_config_json_download_url('uk-reading')

//structure to hold the return for datacenter_agents_config_json_download_url
type datacenterConfigJSONURL struct {
	URL string `json:"datacenter_agents_config_json_download_url,omitempty"`
}

//DatacenterAgentsConfigJSONDownloadURL returns the agent url (and automatically decrypts it)
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
				return "", fmt.Errorf("Permission missing. %s", passwdComponents[1])
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

//CreateOrUpdate implements interface Applier
func (dc Datacenter) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	err = dc.Validate()

	if err != nil {
		return err
	}

	_, err = client.DatacenterGet(dc.DatacenterName)

	if err != nil {
		_, err = client.DatacenterCreate(dc, *dc.DatacenterConfig)

		if err != nil {
			return err
		}
	} else {
		err = client.DatacenterConfigUpdate(dc.DatacenterName, *dc.DatacenterConfig)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (dc Datacenter) Delete(client MetalCloudClient) error {
	return nil
}

//Validate implements interface Applier
func (dc Datacenter) Validate() error {
	if dc.DatacenterName == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}
