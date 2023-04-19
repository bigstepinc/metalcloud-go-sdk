package metalcloud

//go:generate go run helper/gen_exports.go

import "fmt"

//NetworkProfile object describes a network profile
type NetworkProfile struct {
	NetworkProfileID               int                  `json:"network_profile_id,omitempty" yaml:"id,omitempty"`
	NetworkProfileLabel            string               `json:"network_profile_label,omitempty" yaml:"label,omitempty"`
	DatacenterName                 string               `json:"datacenter_name,omitempty" yaml:"dc,omitempty"`
	NetworkProfileIsPublic         bool                 `json:"network_profile_is_public,omitempty" yaml:"networkProfileIsPublic,omitempty"`
	NetworkType                    string               `json:"network_type,omitempty" yaml:"networkType,omitempty"`
	NetworkProfileVLANs            []NetworkProfileVLAN `json:"network_profile_vlans" yaml:"vlans"`
	NetworkProfileCreatedTimestamp string               `json:"nework_profile_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	NetworkProfileUpdatedTimestamp string               `json:"nework_profile_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

//NetworkProfileVLAN object describes a VLAN
type NetworkProfileVLAN struct {
	VlanID                  *int                       `json:"vlan_id" yaml:"vlanID"`
	PortMode                string                     `json:"port_mode,omitempty" yaml:"portMode,omitempty"`
	ProvisionSubnetGateways bool                       `json:"provision_subnet_gateways" yaml:"provisionSubnetGateways"`
	ProvisionVXLAN          bool                       `json:"provision_vxlan" yaml:"provisionVXLAN"`
	ExternalConnectionIDs   []int                      `json:"external_connection_ids" yaml:"extConnectionIDs"`
	SubnetPools             []NetworkProfileSubnetPool `json:"subnet_pools" yaml:"subnetPools"`
}

type NetworkProfileSubnetPool struct {
	SubnetPoolID   *int   `json:"subnet_pool_id" yaml:"subnetPoolID"`
	SubnetPoolType string `json:"subnet_pool_type" yaml:"subnetPoolType"`
}

//NetworkProfileGet returns a NetworkProfile with specified id
func (c *Client) networkProfileGet(networkProfileID id) (*NetworkProfile, error) {
	var createdObject NetworkProfile

	if err := checkID(networkProfileID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_profile_get",
		networkProfileID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//NetworkProfiles returns a list of network profiles for the specified datacenter
func (c *Client) NetworkProfiles(datacenterName string) (*map[int]NetworkProfile, error) {

	resp, err := c.rpcClient.Call(
		"network_profiles",
		datacenterName,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[int]NetworkProfile{}
		return &m, nil
	}

	var createdObject map[int]NetworkProfile

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//NetworkProfileCreate creates a network profile.
func (c *Client) NetworkProfileCreate(datacenterName string, networkProfile NetworkProfile) (*NetworkProfile, error) {
	var createdObject NetworkProfile

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_profile_create",
		datacenterName,
		networkProfile)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//NetworkProfileUpdate updates a network profile.
func (c *Client) networkProfileUpdate(networkProfileID id, networkProfile NetworkProfile) (*NetworkProfile, error) {
	var createdObject NetworkProfile

	if err := checkID(networkProfileID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_profile_update",
		networkProfileID,
		networkProfile)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//NetworkProfileDelete deletes a network profile.
func (c *Client) networkProfileDelete(networkProfileID id) error {

	if err := checkID(networkProfileID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"network_profile_delete",
		networkProfileID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (c *Client) InstanceArrayNetworkProfileSet(instanceArrayID int, networkID int, networkProfileID int) (*map[int]int, error) {
	if err := checkID(networkProfileID); err != nil {
		return nil, err
	}

	if err := checkID(instanceArrayID); err != nil {
		return nil, err
	}

	if err := checkID(networkID); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"instance_array_network_profile_set",
		instanceArrayID,
		networkID,
		networkProfileID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = make(map[int]int)
		return &m, nil
	}

	var createdObject map[int]int

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (c *Client) InstanceArrayNetworkProfileClear(instanceArrayID int, networkID int) error {
	if err := checkID(instanceArrayID); err != nil {
		return err
	}

	if err := checkID(networkID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"instance_array_network_profile_clear",
		instanceArrayID,
		networkID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (c *Client) NetworkProfileListByInstanceArray(instanceArrayID int) (*map[int]int, error) {
	resp, err := c.rpcClient.Call(
		"instance_array_network_profiles",
		instanceArrayID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = make(map[int]int)
		return &m, nil
	}

	var createdObject map[int]int

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}
