package metalcloud

import (
	"encoding/json"
	"fmt"
)

// Subnet represents a subnet
type Subnet struct {
	SubnetID                                    int    `json:"subnet_id,omitempty" yaml:"id,omitempty"`
	SubnetType                                  string `json:"subnet_type,omitempty" yaml:"type,omitempty"`
	NetworkID                                   int    `json:"network_id,omitempty" yaml:"networkID,omitempty"`
	InfrastructureID                            int    `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	SubnetPoolID                                int    `json:"subnet_pool_id,omitempty" yaml:"subnetPoolID,omitempty"`
	ClusterID                                   int    `json:"cluster_id,omitempty" yaml:"clusterID,omitempty"`
	SubnetNetmaskHumanReadable                  string `json:"subnet_netmask_human_readable,omitempty" yaml:"netmask,omitempty"`
	SubnetGatewayHumanReadable                  string `json:"subnet_gateway_human_readable,omitempty" yaml:"gateway,omitempty"`
	SubnetPrefixSize                            int    `json:"subnet_prefix_size,omitempty" yaml:"prefixSize,omitempty"`
	SubnetChangeID                              int    `json:"subnet_change_id,omitempty" yaml:"changeID,omitempty"`
	SubnetServiceStatus                         string `json:"subnet_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	SubnetDestination                           string `json:"subnet_destination,omitempty" yaml:"destination,omitempty"`
	SubnetAutomaticAllocation                   bool   `json:"subnet_automatic_allocation" yaml:"automaticAllocation,omitempty"`
	SubnetLabel                                 string `json:"subnet_label,omitempty" yaml:"label,omitempty"`
	DNSSubdomainPermanentID                     int    `json:"dns_subdomain_permanent_id,omitempty" yaml:"dnsSubdomainPermanentID,omitempty"`
	SubnetSubdomain                             string `json:"subnet_subdomain,omitempty" yaml:"subdomain,omitempty"`
	SubnetSubdomainPermanent                    string `json:"subnet_subdomain_permanent,omitempty" yaml:"subdomainPermanent,omitempty"`
	SubnetTrafficDownloadBytes                  int    `json:"subnet_traffic_download_bytes,omitempty" yaml:"trafficDownloadBytes,omitempty"`
	SubnetTrafficUploadBytes                    int    `json:"subnet_traffic_upload_bytes,omitempty" yaml:"trafficUploadBytes,omitempty"`
	SubnetTrafficDatacenterBytes                int    `json:"subnet_traffic_datacenter_bytes,omitempty" yaml:"trafficDatacenterBytes,omitempty"`
	SubnetDirtyBit                              bool   `json:"subnet_dirty_bit,omitempty" yaml:"dirtyBit,omitempty"`
	SubnetUpdatedTimestamp                      string `json:"subnet_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	SubnetCreatedTimestamp                      string `json:"subnet_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	SubnetIsAPIPrivate                          bool   `json:"subnet_is_api_private,omitempty" yaml:"isAPIPrivate,omitempty"`
	SubnetTrafficDownloadBytesTemporary         int    `json:"subnet_traffic_download_bytes_temporary,omitempty" yaml:"trafficDownloadBytesTemporary,omitempty"`
	SubnetTrafficUploadBytesTemporary           int    `json:"subnet_traffic_upload_bytes_temporary,omitempty" yaml:"trafficUploadBytesTemporary,omitempty"`
	SubnetTrafficDatacenterBytesTemporary       int    `json:"subnet_traffic_datacenter_bytes_temporary,omitempty" yaml:"trafficDatacenterBytesTemporary,omitempty"`
	SubnetTrafficFetchedUntilTimestamp          string `json:"subnet_traffic_fetched_until_timestamp,omitempty" yaml:"trafficFetchedUntilTimestamp,omitempty"`
	SubnetTrafficFetchedUntilTimestampTemporary string `json:"subnet_traffic_fetched_until_timestamp_temporary,omitempty" yaml:"trafficFetchedUntilTimestampTemporary,omitempty"`
	SubnetFromSubnetPoolForcedOnly              bool   `json:"subnet_from_subnet_pool_forced_only,omitempty" yaml:"fromSubnetPoolForcedOnly,omitempty"`
	SubnetOverrideVLANID                        int    `json:"subnet_override_vlan_id,omitempty" yaml:"overrideVLANID,omitempty"`
	SubnetOverrideVLANAutoAllocationIndex       int    `json:"subnet_override_vlan_auto_allocation_index,omitempty" yaml:"overrideVLANAutoAllocationIndex,omitempty"`
	SubnetIsIPRange                             bool   `json:"subnet_is_ip_range,omitempty" yaml:"isIPRange,omitempty"`
}

// SubnetPool represents a pool of subnets
type SubnetPool struct {
	SubnetPoolID                                int    `json:"subnet_pool_id,omitempty" yaml:"id,omitempty"`
	DatacenterName                              string `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	NetworkEquipmentID                          int    `json:"network_equipment_id,omitempty" yaml:"networkEquipmentID,omitempty"`
	UserID                                      int    `json:"user_id,omitempty" yaml:"user,omitempty"`
	SubnetPoolPrefixHumanReadable               string `json:"subnet_pool_prefix_human_readable,omitempty" yaml:"prefix,omitempty"`
	SubnetPoolLabel                             string `json:"subnet_pool_label,omitempty" yaml:"label,omitempty"`
	SubnetPoolPrefixHex                         string `json:"subnet_pool_prefix_hex,omitempty" yaml:"prefixHex,omitempty"`
	SubnetPoolNetmaskHumanReadable              string `json:"subnet_pool_netmask_human_readable,omitempty" yaml:"netmask,omitempty"`
	SubnetPoolNetmaskHex                        string `json:"subnet_pool_netmask_hex,omitempty" yaml:"netmaskHex,omitempty"`
	SubnetPoolPrefixSize                        int    `json:"subnet_pool_prefix_size,omitempty" yaml:"size,omitempty"`
	SubnetPoolType                              string `json:"subnet_pool_type,omitempty" yaml:"type,omitempty"`
	SubnetPoolRoutable                          bool   `json:"subnet_pool_routable" yaml:"routable"`
	SubnetPoolDestination                       string `json:"subnet_pool_destination,omitempty" yaml:"destination,omitempty"`
	SubnetPoolUtilizationCachedJSON             string `json:"subnet_pool_utilization_cached_json,omitempty" yaml:"currentUtilizationJSON,omitempty"`
	SubnetPoolUtilizationCachedUpdatedTimestamp string `json:"subnet_pool_cached_updated_timestamp,omitempty" yaml:"currentUtilizationLastUpdated,omitempty"`
	SubnetPoolIsOnlyForManualAllocation         bool   `json:"subnet_pool_is_only_for_manual_allocation" yaml:"manualAllocationOnly"`
}

// SubnetPoolUtilization describes the current utilization of the subnet
type SubnetPoolUtilization struct {
	PrefixCountFree                        map[string]int `json:"prefix_count_free,omitempty" yaml:"availableSubnets,omitempty"`
	PrefixCountAllocated                   map[string]int `json:"prefix_count_allocated,omitempty" yaml:"allocatedSubnets,omitempty"`
	IPAddressesUsableCountFree             string         `json:"ip_addresses_usable_count_free,omitempty" yaml:"availableUsableIps,omitempty"`
	IPAddressesUsableCountAllocated        string         `json:"ip_addresses_usable_count_allocated,omitempty" yaml:"allocatedUsableIps,omitempty"`
	IPAddressesUsableFreePercentOptimistic string         `json:"ip_addresses_usable_free_percent_optimistic,omitempty" yaml:"availablePercentage,omitempty"`
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

// SearchResultForSubnets describes a search result for subnets search
type searchResultForSubnets struct {
	DurationMilliseconds int          `json:"duration_millisecnds,omitempty"`
	Rows                 []SubnetPool `json:"rows,omitempty"`
	RowsOrder            [][]string   `json:"rows_order,omitempty"`
	RowsTotal            int          `json:"rows_total,omitempty"`
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

func (c *Client) SubnetGet(subnetID int) (*Subnet, error) {
	var createdObject Subnet

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_get",
		subnetID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (c *Client) SubnetCreate(subnet Subnet) (*Subnet, error) {
	var createdObject Subnet

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_create",
		subnet.NetworkID,
		subnet)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (c *Client) SubnetDelete(subnetID int) error {
	resp, err := c.rpcClient.Call("subnet_delete", subnetID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
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
	var createdObject map[string]searchResultForSubnets

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
		createdObject = map[string]searchResultForSubnets{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []SubnetPool{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil

}

// CreateOrUpdate implements interface Applier
func (s SubnetPool) CreateOrUpdate(client MetalCloudClient) error {
	err := s.Validate()

	if err != nil {
		return err
	}

	_, err = client.SubnetPoolGet(s.SubnetPoolID)

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
	err = client.SubnetPoolDelete(s.SubnetPoolID)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (s SubnetPool) Validate() error {
	if s.SubnetPoolID == 0 {
		return fmt.Errorf("id is required")
	}
	return nil
}

func (s Subnet) CreateOrUpdate(client MetalCloudClient) error {
	err := s.Validate()

	if err != nil {
		return err
	}

	_, err = client.SubnetGet(s.SubnetID)

	if err != nil {
		_, err := client.SubnetCreate(s)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s Subnet) Delete(client MetalCloudClient) error {
	err := s.Validate()

	if err != nil {
		return err
	}
	err = client.SubnetDelete(s.SubnetID)

	if err != nil {
		return err
	}

	return nil
}

func (s Subnet) Validate() error {
	if s.SubnetID == 0 {
		return fmt.Errorf("id is required")
	}
	return nil
}

// SubnetPoolCreateOrUpdate creates or updates a subnet pool
func (c *Client) SubnetPoolCreateOrUpdate(subnetPool SubnetPool) (*SubnetPool, error) {
	err := subnetPool.Validate()

	if err != nil {
		return nil, err
	}

	_, err = c.SubnetPoolGet(subnetPool.SubnetPoolID)

	if err != nil {
		_, err := c.SubnetPoolCreate(subnetPool)

		if err != nil {
			return nil, err
		}
	}

	return &subnetPool, nil
}
