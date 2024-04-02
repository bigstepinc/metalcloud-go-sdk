package metalcloud

import (
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

func (s Subnet) Validate() error {
	if s.SubnetID == 0 && s.SubnetLabel != "" {
		return fmt.Errorf("id or label is required")
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
