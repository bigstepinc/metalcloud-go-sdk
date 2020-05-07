package metalcloud

import "fmt"

//go:generate go run helper/gen_exports.go

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

//Datacenters returns all datacenters
func (c *Client) Datacenters() (*map[string]Datacenter, error) {
	return c.DatacentersByUserID(nil, false)
}

//DatacentersOnlyActive returns all active datacenters
func (c *Client) DatacentersOnlyActive() (*map[string]Datacenter, error) {
	return c.DatacentersByUserID(nil, true)
}

//DatacentersByUserID returns datacenters belonging to a particular user
func (c *Client) DatacentersByUserID(userID id, onlyActive bool) (*map[string]Datacenter, error) {

	if err := checkID(userID); err != nil {
		return nil, err
	}

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
	return c.DatacenterGetForUser(datacenterName, nil)
}

//DatacenterGetForUser returns details of a specific datacenter
func (c *Client) DatacenterGetForUser(datacenterName string, userID id) (*Datacenter, error) {
	var datacenter Datacenter

	if err := checkID(userID); err != nil {
		return nil, err
	}

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
