//go:generate go run helper/docgen.go - $GOFILE ./ SubnetPool,SubnetPoolUtilization SubnetPool
package metalcloud

import (
	"encoding/json"
	"fmt"
)

// SubnetPool represents a pool of subnets
type SubnetPool struct {
	//description: Id of the subnetpool
	SubnetPoolID int `json:"subnet_pool_id,omitempty" yaml:"id,omitempty"`
	//description: Label fo the Datacenter
	DatacenterName string `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	//description: ID fo the network equipment to which this subnet pool is associated
	NetworkEquipmentID int `json:"network_equipment_id,omitempty" yaml:"networkEquipmentID,omitempty"`
	//description: Owner of this subent pool
	UserID int `json:"user_id,omitempty" yaml:"user,omitempty"`
	//description: Prefix
	SubnetPoolPrefixHumanReadable string `json:"subnet_pool_prefix_human_readable,omitempty" yaml:"prefix,omitempty"`
	//description: Label of this subnet pool
	SubnetPoolLabel string `json:"subnet_pool_label,omitempty" yaml:"label,omitempty"`
	//description: Internal
	SubnetPoolPrefixHex string `json:"subnet_pool_prefix_hex,omitempty" yaml:"prefixHex,omitempty"`
	//description: Netmask
	SubnetPoolNetmaskHumanReadable string `json:"subnet_pool_netmask_human_readable,omitempty" yaml:"netmask,omitempty"`
	//description: Internal
	SubnetPoolNetmaskHex string `json:"subnet_pool_netmask_hex,omitempty" yaml:"netmaskHex,omitempty"`
	//description: Size of the prefix
	//example: 27
	SubnetPoolPrefixSize int `json:"subnet_pool_prefix_size,omitempty" yaml:"size,omitempty"`
	//description: Type
	//values:
	// - ipv4
	// - ipv6
	SubnetPoolType string `json:"subnet_pool_type,omitempty" yaml:"type,omitempty"`
	//description: If set to true this subnet pool will be used for subnets that are routed (usually to the internet)
	SubnetPoolRoutable bool `json:"subnet_pool_routable" yaml:"routable"`
	//description: What this subnet is intended for:
	//values:
	// - wan
	// - lan
	// - san
	// - oob
	// - quarantine
	// - loopback
	// - vtep
	// - disabled
	SubnetPoolDestination string `json:"subnet_pool_destination,omitempty" yaml:"destination,omitempty"`
	//description: Internal
	SubnetPoolUtilizationCachedJSON string `json:"subnet_pool_utilization_cached_json,omitempty" yaml:"currentUtilizationJSON,omitempty"`
	//description: Internal
	SubnetPoolUtilizationCachedUpdatedTimestamp string `json:"subnet_pool_cached_updated_timestamp,omitempty" yaml:"currentUtilizationLastUpdated,omitempty"`
	//description: If set to true this subnet will not be used for automatic allocation and will only be used in network profiles.
	SubnetPoolIsOnlyForManualAllocation bool `json:"subnet_pool_is_only_for_manual_allocation" yaml:"manualAllocationOnly"`
}

// SubnetPoolUtilization describes the current utilization of the subnet
type SubnetPoolUtilization struct {
	//description: Count of available addresses
	PrefixCountFree map[string]int `json:"prefix_count_free,omitempty" yaml:"availableSubnets,omitempty"`
	//description: Count of allocated addresses
	PrefixCountAllocated map[string]int `json:"prefix_count_allocated,omitempty" yaml:"allocatedSubnets,omitempty"`
	//description: Amount of usable addresses (excludes broadcast)
	IPAddressesUsableCountFree string `json:"ip_addresses_usable_count_free,omitempty" yaml:"availableUsableIps,omitempty"`
	//description: Amount of allocated from the usable set
	IPAddressesUsableCountAllocated string `json:"ip_addresses_usable_count_allocated,omitempty" yaml:"allocatedUsableIps,omitempty"`
	//description: Available usable (Percentage)
	IPAddressesUsableFreePercentOptimistic string `json:"ip_addresses_usable_free_percent_optimistic,omitempty" yaml:"availablePercentage,omitempty"`
}

// SearchResultForSubnetPools describes a search result for subnet pools search
type searchResultForSubnetPools struct {
	DurationMilliseconds int          `json:"duration_millisecnds,omitempty"`
	Rows                 []SubnetPool `json:"rows,omitempty"`
	RowsOrder            [][]string   `json:"rows_order,omitempty"`
	RowsTotal            int          `json:"rows_total,omitempty"`
}

// UnmarshalJSON to handle the shity [] to {} and 0 and "123123" cases
func (s *SubnetPoolUtilization) UnmarshalJSON(data []byte) error {

	var v struct {
		PrefixCountFree                        interface{} `json:"prefix_count_free,omitempty" yaml:"availableSubnets,omitempty"`
		PrefixCountAllocated                   interface{} `json:"prefix_count_allocated,omitempty" yaml:"allocatedSubnets,omitempty"`
		IPAddressesUsableCountFree             interface{} `json:"ip_addresses_usable_count_free,omitempty" yaml:"availableUsableIps,omitempty"`
		IPAddressesUsableCountAllocated        interface{} `json:"ip_addresses_usable_count_allocated,omitempty" yaml:"allocatedUsableIps,omitempty"`
		IPAddressesUsableFreePercentOptimistic interface{} `json:"ip_addresses_usable_free_percent_optimistic,omitempty" yaml:"availablePercentage,omitempty"`
	}

	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.IPAddressesUsableCountAllocated.(type) {
	case int:
		s.IPAddressesUsableCountAllocated = fmt.Sprintf("%d", v.IPAddressesUsableCountAllocated.(int))
	case string:

		s.IPAddressesUsableCountAllocated = v.IPAddressesUsableCountAllocated.(string)
	}

	switch v.IPAddressesUsableCountFree.(type) {
	case int:
		s.IPAddressesUsableCountFree = fmt.Sprintf("%d", v.IPAddressesUsableCountFree.(int))
	case string:

		s.IPAddressesUsableCountFree = v.IPAddressesUsableCountFree.(string)
	}

	switch v.IPAddressesUsableFreePercentOptimistic.(type) {
	case int:
		s.IPAddressesUsableFreePercentOptimistic = fmt.Sprintf("%d", v.IPAddressesUsableFreePercentOptimistic.(int))
	case string:

		s.IPAddressesUsableFreePercentOptimistic = v.IPAddressesUsableFreePercentOptimistic.(string)
	}

	s.PrefixCountFree = map[string]int{}
	if _, ok := v.PrefixCountFree.([]interface{}); !ok {
		for i, v := range v.PrefixCountFree.(map[string]interface{}) {
			s.PrefixCountFree[i] = int(v.(float64))
		}
	}

	s.PrefixCountAllocated = map[string]int{}
	if _, ok := v.PrefixCountAllocated.([]interface{}); !ok {
		for i, v := range v.PrefixCountAllocated.(map[string]interface{}) {
			s.PrefixCountAllocated[i] = int(v.(float64))
		}
	}

	return nil
}

// SubnetPoolCreate creates a new SubnetPool.
func (c *Client) SubnetPoolCreate(subnetPool SubnetPool) (*SubnetPool, error) {
	var createdObject SubnetPool

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_pool_create",
		[]interface{}{subnetPool})

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SubnetPoolGet retrieves information regarding a specified SubnetPool.
func (c *Client) SubnetPoolGet(subnetPoolID int) (*SubnetPool, error) {

	var createdObject SubnetPool

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_pool_get",
		subnetPoolID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SubnetPoolGetByLabel retrieves information regarding a specified SubnetPool by label.
func (c *Client) SubnetPoolGetByLabel(subnetPoolLabel string) (*SubnetPool, error) {

	subnetPools, err := c.SubnetPoolSearch(fmt.Sprintf("subnet_pool_label:%s", subnetPoolLabel))
	if err != nil {
		return nil, err
	}

	if len(*subnetPools) == 0 {
		return nil, fmt.Errorf("subnet pool with label %s not found", subnetPoolLabel)
	}

	return &(*subnetPools)[0], nil

}

// SubnetPoolPrefixSizesStats retrieves information regarding the utilization of a specified SubnetPool.
func (c *Client) SubnetPoolPrefixSizesStats(subnetPoolID int) (*SubnetPoolUtilization, error) {

	var createdObject SubnetPoolUtilization

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_pool_prefix_sizes_stats",
		subnetPoolID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// SubnetPoolDelete deletes the specified SubnetPool
func (c *Client) SubnetPoolDelete(subnetPoolID int) error {

	resp, err := c.rpcClient.Call("subnet_pool_delete", subnetPoolID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (c *Client) SubnetPoolDeleteByLabel(subnetPoolLabel string) error {
	subnetPool, err := c.SubnetPoolGetByLabel(subnetPoolLabel)
	if err != nil {
		return err
	}

	return c.SubnetPoolDelete(subnetPool.SubnetPoolID)
}

// SubnetPools retrieves all switch devices registered in the database.
func (c *Client) SubnetPools() (*[]SubnetPool, error) {
	return c.SubnetPoolSearch("*")
}

// SubnetPoolSearch retrieves all switch devices registered in the database with the specified filter
func (c *Client) SubnetPoolSearch(filter string) (*[]SubnetPool, error) {

	tables := []string{"_subnet_pools"}
	columns := map[string][]string{
		"_subnet_pools": {
			"subnet_pool_id",
			"subnet_pool_label",
			"subnet_pool_prefix_human_readable",
			"subnet_pool_prefix_hex",
			"subnet_pool_netmask_human_readable",
			"subnet_pool_netmask_hex",
			"subnet_pool_prefix_size",
			"subnet_pool_prefix_type",
			"subnet_pool_prefix_destination",
			"subnet_pool_routable",
			"user_id",
			"subnet_pool_destination",
			"datacenter_name",
			"network_equipment_id",
			"subnet_pool_utilization_cached_json",
			"subnet_pool_cached_updated_timestamp",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	var createdObject map[string]searchResultForSubnetPools

	resp, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]searchResultForSubnetPools{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []SubnetPool{}

	list = append(list, createdObject[tables[0]].Rows...)

	return &list, nil
}

// Validate implements interface Applier
func (s SubnetPool) Validate() error {
	if s.SubnetPoolLabel == "" && s.SubnetPoolID == 0 {
		return fmt.Errorf("id or label is required")
	}
	return nil
}

// CreateOrUpdate implements interface Applier
func (s SubnetPool) CreateOrUpdate(client MetalCloudClient) error {
	err := s.Validate()

	if err != nil {
		return err
	}

	if s.SubnetPoolLabel != "" {
		_, err = client.SubnetPoolGetByLabel(s.SubnetPoolLabel)
	} else if s.SubnetPoolID != 0 {
		_, err = client.SubnetPoolGet(s.SubnetPoolID)
	}

	if err != nil {
		_, err := client.SubnetPoolCreate(s)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (s SubnetPool) Delete(client MetalCloudClient) error {
	err := s.Validate()
	if err != nil {
		return err
	}

	if s.SubnetPoolLabel != "" {
		return client.SubnetPoolDeleteByLabel(s.SubnetPoolLabel)
	}

	return client.SubnetPoolDelete(s.SubnetPoolID)
}
