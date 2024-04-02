package metalcloud

import (
	"fmt"
)

// Subnet represents a subnet
type SubnetOOB struct {
	SubnetOOBID                                int      `json:"subnet_oob_id,omitempty" yaml:"id,omitempty"`
	SubnetOOBLabel                             string   `json:"subnet_oob_label,omitempty" yaml:"label,omitempty"`
	SubnetOOBType                              string   `json:"subnet_oob_type,omitempty" yaml:"type,omitempty"`
	SubnetOOBAllocateForResourceType           string   `json:"subent_oob_allocate_for_resource_type,omitempty" yaml:"forResourceType,omitempty"`
	SubnetOOBBlackList                         []string `json:"subnet_oob_blacklist,omitempty" yaml:"blacklist,omitempty"`
	SubnetOOBGatewayHex                        string   `json:"subnet_oob_gateway_hex,omitempty" yaml:"gatewayHex,omitempty"`
	SubnetOOBGatewayHumanReadable              string   `json:"subnet_oob_gateway_human_readable,omitempty" yaml:"gateway,omitempty"`
	SubnetOOBNetmaskHex                        string   `json:"subnet_oob_netmask_hex,omitempty" yaml:"netmaskHex,omitempty"`
	SubnetOOBNetmaskHumanReadable              string   `json:"subnet_oob_netmask_human_readable,omitempty" yaml:"netmask,omitempty"`
	SubnetOOBPrefixSize                        int      `json:"subnet_oob_prefix_size,omitempty" yaml:"size,omitempty"`
	SubnetOOBRangeStartHex                     string   `json:"subnet_oob_range_start_hex,omitempty" yaml:"rangeStartHex,omitempty"`
	SubnetOOBRangeStartHumanReadable           string   `json:"subnet_oob_range_start_human_readable,omitempty" yaml:"rangeStart,omitempty"`
	SubnetOOBRangeStartHumanReadableCompressed string   `json:"subnet_oob_range_start_human_readable_compressed,omitempty" yaml:"rangeStartCompressed,omitempty"`
	SubnetOOBRangeEndHex                       string   `json:"subnet_oob_range_end_hex,omitempty" yaml:"rangeEndHex,omitempty"`
	SubnetOOBRangeEndHumanReadable             string   `json:"subnet_oob_range_end_human_readable,omitempty" yaml:"rangeEnd,omitempty"`
	SubnetOOBRangeEndHumanReadableCompressed   string   `json:"subnet_oob_range_end_human_readable_compressed,omitempty" yaml:"rangeEndCompressed,omitempty"`
	SubnetOOBUseForAutoAllocation              bool     `json:"subnet_oob_use_for_auto_allocation,omitempty" yaml:"useForAutoAllocation,omitempty"`
	DatacenterName                             string   `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
}

// searchResultForSubnetsOOB describes a search result for OOB subnet search
type searchResultForSubnetsOOB struct {
	DurationMilliseconds float32     `json:"duration_milliseconds,omitempty"`
	Rows                 []SubnetOOB `json:"rows,omitempty"`
	RowsOrder            [][]string  `json:"rows_order,omitempty"`
	RowsTotal            int         `json:"rows_total,omitempty"`
}

func (c *Client) SubnetOOBGet(subnetOOBID int) (*SubnetOOB, error) {
	var createdObject SubnetOOB

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_oob_get",
		subnetOOBID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SubnetOOBGetByLabel retrieves information regarding a specified Subnet by label.
func (c *Client) SubnetOOBGetByLabel(subnetOOBLabel string) (*SubnetOOB, error) {

	subnets, err := c.SubnetOOBSearch(fmt.Sprintf("subnet_oob_label:%s", subnetOOBLabel))
	if err != nil {
		return nil, err
	}

	if len(*subnets) == 0 {
		return nil, fmt.Errorf("oob subnet with label %s not found", subnetOOBLabel)
	}

	return &(*subnets)[0], nil

}

func (c *Client) SubnetOOBCreate(subnetOOB SubnetOOB) (*SubnetOOB, error) {
	var createdObject SubnetOOB

	err := c.rpcClient.CallFor(
		&createdObject,
		"subnet_oob_create",
		[]interface{}{subnetOOB})

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (c *Client) SubnetOOBDelete(subnetOOBID int) error {
	resp, err := c.rpcClient.Call("subnet_oob_delete", subnetOOBID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (c *Client) SubnetOOBDeleteByLabel(subnetOOBLabel string) error {
	subnetOOB, err := c.SubnetOOBGetByLabel(subnetOOBLabel)
	if err != nil {
		return err
	}

	return c.SubnetOOBDelete(subnetOOB.SubnetOOBID)
}

// SubnetOOBSearch retrieves all OOB subnets registered in the database with the specified filter
func (c *Client) SubnetOOBSearch(filter string) (*[]SubnetOOB, error) {

	tables := []string{"_subnets_oob"}
	columns := map[string][]string{
		"_subnets_oob": {
			"subnet_oob_id",
			"subnet_oob_label",
			"subnet_oob_netmask_human_readable",
			"subnet_oob_prefix_size",
			"user_id",
			"datacenter_name",
			"subnet_oob_range_start_human_readable",
			"subnet_oob_range_end_human_readable",
			"subnet_oob_gateway_human_readable",
			"subnet_oob_type",
			"subnet_oob_use_for_auto_allocation",
			"subent_oob_allocate_for_resource_type",
			"subnet_oob_blacklist",
			"subnet_oob_blacklist_json",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	var createdObject map[string]searchResultForSubnetsOOB

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
		createdObject = map[string]searchResultForSubnetsOOB{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []SubnetOOB{}

	list = append(list, createdObject[tables[0]].Rows...)

	return &list, nil
}

func (s SubnetOOB) Validate() error {
	if s.SubnetOOBID == 0 && s.SubnetOOBLabel == "" {
		return fmt.Errorf("id or label is required")
	}
	return nil
}

// CreateOrUpdate implements interface Applier
func (s SubnetOOB) CreateOrUpdate(client MetalCloudClient) error {
	err := s.Validate()

	if err != nil {
		return err
	}

	if s.SubnetOOBLabel != "" {
		_, err = client.SubnetOOBGetByLabel(s.SubnetOOBLabel)
	} else if s.SubnetOOBID != 0 {
		_, err = client.SubnetOOBGet(s.SubnetOOBID)
	}

	if err != nil {
		_, err := client.SubnetOOBCreate(s)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (s SubnetOOB) Delete(client MetalCloudClient) error {
	err := s.Validate()
	if err != nil {
		return err
	}

	if s.SubnetOOBLabel != "" {
		return client.SubnetOOBDeleteByLabel(s.SubnetOOBLabel)
	}

	return client.SubnetOOBDelete(s.SubnetOOBID)
}
