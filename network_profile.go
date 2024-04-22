package metalcloud

//go:generate go run helper/gen_exports.go
//go:generate go run helper/docgen.go - $GOFILE ./ NetworkProfile,NetworkProfileVLAN,NetworkProfileSubnetPool NetworkProfile

import "fmt"

//NetworkProfile  A network profile modifies the default network configuration of an instance array when attached to a specific Network.
type NetworkProfile struct {
	//description: The id of the network profile
	NetworkProfileID int `json:"network_profile_id,omitempty" yaml:"id,omitempty"`
	//description: The label of the network profile
	NetworkProfileLabel string `json:"network_profile_label,omitempty" yaml:"label,omitempty"`
	//description: The label of the datacenter on which this network profile applies
	DatacenterName string `json:"datacenter_name,omitempty" yaml:"dc,omitempty"`
	//description: If set to true any of the users can use this network profile. If set to false the network profile needs to be
	//explicitly allowed in the user limits for a user to be able to use it
	NetworkProfileIsPublic bool `json:"network_profile_is_public,omitempty" yaml:"networkProfileIsPublic,omitempty"`
	//description: Type of network to which this network profile can be applied
	//values:
	// - wan
	// - lan
	// - san
	NetworkType string `json:"network_type,omitempty" yaml:"networkType,omitempty"`
	//description: VLAN (L2 network) entries
	NetworkProfileVLANs []NetworkProfileVLAN `json:"network_profile_vlans" yaml:"vlans"`
	//description: ISO 8601 timestamp which holds the date and time when the network profile was created.
	//example: 2013-11-29T13:00:01Z
	NetworkProfileCreatedTimestamp string `json:"nework_profile_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the network profile was updated.
	//example: 2013-11-29T13:00:01Z
	NetworkProfileUpdatedTimestamp string `json:"nework_profile_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

//NetworkProfileVLAN object describes an overlay L2 network which usually translates into a VLAN and
// associated VXLAN VNIs or EPGs depending on the vendor
type NetworkProfileVLAN struct {
	//description: the VLAN ID to use (on EVPNVXLAN provisioner this translates to a VLAN ID to be used as VTEP for VXLAN).
	//On VLAN provisioner the VLAN will be used end-to-end. On SDN provisioner with CISCO ACI it will be translated into a VLAN attached to an EPG.
	//If set to null it will be automatically allocated based on the settings in the Datacenter config object.
	VLANID *int `json:"vlan_id" yaml:"vlanID"`
	//description: the label of this VLAN entry
	VLANLabel *int `json:"vlan_label" yaml:"label"`
	//description: Type of VLAN
	//values:
	//  - native
	//  - trunk
	//  - access
	PortMode string `json:"port_mode,omitempty" yaml:"portMode,omitempty"`
	//description: if true will provision subnet gateways on the switches. This depends on the provisioner and vendor: Symmetric IRB, VRRP-based redundant gateway configuration etc.
	ProvisionSubnetGateways bool `json:"provision_subnet_gateways" yaml:"provisionSubnetGateways"`
	//description: if false it will not configure the vxlan tunnel. Equivalent to the VLAN provisioner.
	ProvisionVXLAN bool `json:"provision_vxlan" yaml:"provisionVXLAN"`
	//description: If any external connection id is configured in the array the L2 network will be extended to that external connection point
	ExternalConnectionIDs []int `json:"external_connection_ids" yaml:"extConnectionIDs"`
	//description: Subnet pools to allocate IPs from
	SubnetPools []NetworkProfileSubnetPool `json:"subnet_pools" yaml:"subnetPools"`
	//description: If set to a non-null value this will help ensure that two different network profiles can get the same auto-allocated VLAN ID
	VLANAutoAllocationIndex *int `json:"vlan_auto_allocation_index,omitempty" yaml:"vlanAutoAllocationIndex,omitempty"`
	//description: The VRF ID to use. If set to null it will fall back to automatically allocated VRFs
	VRFID *int `json:"vrf_id,omitempty" yaml:"vrfID,omitempty"`
}

type NetworkProfileSubnetPool struct {
	//description: The ID of the subnet pool to use. Set to null for auto in which case the type of subnet pool will be used to pick a subnet pool
	SubnetPoolID *int `json:"subnet_pool_id" yaml:"subnetPoolID"`
	//description: The type of subnet pool to use
	//values:s
	// - ipv4
	// - ipv6
	SubnetPoolType string `json:"subnet_pool_type" yaml:"subnetPoolType"`
	//description: If set to true this subnet pool will be used to set the default route
	SubnetPoolProvidesDefaultRoute bool `json:"subnet_pool_provides_default_route" yaml:"subnetPoolProvidesDefaultRoute"`
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

//CreateOrUpdate implements interface Applier
func (np NetworkProfile) CreateOrUpdate(client MetalCloudClient) error {
	var result *NetworkProfile
	var err error
	err = np.Validate()

	if err != nil {
		return err
	}

	if np.NetworkProfileID != 0 {
		result, err = client.NetworkProfileGet(np.NetworkProfileID)
	} else {
		result, err = client.NetworkProfileGetByLabel(np.NetworkProfileLabel)
	}

	if err != nil {
		_, err = client.NetworkProfileCreate(np.DatacenterName, np)

		if err != nil {
			return err
		}
	} else {
		_, err = client.NetworkProfileUpdate(result.NetworkProfileID, np)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (np NetworkProfile) Delete(client MetalCloudClient) error {
	var result *NetworkProfile
	var id int
	err := np.Validate()

	if err != nil {
		return err
	}

	if np.NetworkProfileLabel != "" {
		result, err = client.NetworkProfileGetByLabel(np.NetworkProfileLabel)
		if err != nil {
			return err
		}
		id = result.NetworkProfileID
	} else {
		id = np.NetworkProfileID
	}
	err = client.NetworkProfileDelete(id)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (np NetworkProfile) Validate() error {
	if np.NetworkProfileID == 0 && np.NetworkProfileLabel == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
