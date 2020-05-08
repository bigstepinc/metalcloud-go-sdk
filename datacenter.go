package metalcloud

import (
	"fmt"
	"strings"
)

//Datacenter - datacenter description
type Datacenter struct {
	DatacenterName             string   `json:"datacenter_name,omitempty"`
	DatacenterNameParent       string   `json:"datacenter_name_parent,omitempty"`
	UserID                     int      `json:"user_id,omitempty"`
	DatacenterDisplayName      string   `json:"datacenter_display_name,omitempty"`
	DatacenterIsMaster         bool     `json:"datacenter_is_master,omitempty"`
	DatacenterIsMaintenance    bool     `json:"datacenter_is_maintenance,omitempty"`
	DatacenterType             string   `json:"datacenter_type,omitempty"`
	DatacenterCreatedTimestamp string   `json:"datacenter_created_timestamp,omitempty"`
	DatacenterUpdatedTimestamp string   `json:"datacenter_updated_timestamp,omitempty"`
	DatacenterHidden           bool     `json:"datacenter_hidden,omitempty"`
	DatacenterTags             []string `json:"datacenter_tags,omitempty"`
}

//DatacenterConfig - datacenter configuration
type DatacenterConfig struct {
	SANRoutedSubnet                       string          `json:"SANRoutedSubnet,omitempty"`
	BSIVRRPListenIPv4                     string          `json:"BSIVRRPListenIPv4,omitempty"`
	BSIMachineListenIPv4List              []string        `json:"BSIMachineListenIPv4List,omitempty"`
	BSIExternallyVisibleIPv4              string          `json:"BSIExternallyVisibleIPv4,omitempty"`
	RepoURLRoot                           string          `json:"repoURLRoot,omitempty"`
	RepoURLRootQuarantineNetwork          string          `json:"repoURLRootQuarantineNetwork,omitempty"`
	NTPServers                            []string        `json:"NTPServers,omitempty"`
	DNSServers                            []string        `json:"DNSServers,omitempty"`
	KMS                                   string          `json:"KMS,omitempty"`
	TFTPServerWANVRRPListenIPv4           string          `json:"TFTPServerWANVRRPListenIPv4,omitempty"`
	DataLakeEnabled                       bool            `json:"dataLakeEnabled,omitempty"`
	MonitoringGraphitePlainTextSocketHost string          `json:"monitoringGraphitePlainTextSocketHost,omitempty"`
	MonitoringGraphiteRenderURLHost       string          `json:"monitoringGraphiteRenderURLHost,omitempty"`
	Latitude                              float64         `json:"latitude,omitempty"`
	Longitude                             float64         `json:"longitude,omitempty"`
	Address                               string          `json:"address,omitempty"`
	VLANProvisioner                       VLANProvisioner `json:"VLANProvisioner,omitempty"`
}

//VLANProvisioner - defines settings for the networking provisioning architecture that uses vlans
type VLANProvisioner struct {
	LANVLANRange     string `json:"LANVLANRange,omitempty"`
	WANVLANRange     string `json:"WANVLANRange,omitempty"`
	QuarantineVLANID int    `json:"quarantineVLANID,omitempty"`
}

//VPLSProvisioner - defines settings for the networking provisioning architecture that uses vpls
type VPLSProvisioner struct {
	ACLSAN            string `json:"ACLSAN,omitempty"`
	ACLWAN            string `json:"ACLWAN,omitempty"`
	SANACLRange       string `json:"SANACLRange,omitempty"`
	ToRLANVLANRange   string `json:"ToRLANVLANRange,omitempty"`
	ToRSANVLANRange   string `json:"ToRSANVLANRange,omitempty"`
	ToRWANVLANRange   string `json:"ToRWANVLANRange,omitempty"`
	QuarantineVLANID  int    `json:"quarantineVLANID,omitempty"`
	NorthWANVLANRange string `json:"NorthWANVLANRange,omitempty"`
}

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

	res, err := c.rpcClient.Call(
		"datacenters",
		userID,
		onlyActive,
		false)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Datacenter{}
		return &m, nil
	}

	var createdObject map[string]Datacenter

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
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
		if len(passwdComponents) != 2 {
			return "", fmt.Errorf("Password not returned with proper components")
		}
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

	return agentConfigURL, nil
}
