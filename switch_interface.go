package metalcloud

import (
	"fmt"
)

// searchResultResponseWrapperForSwitchInterfaces describes a search result for switch interfaces
type searchResultResponseWrapperForSwitchInterfaces struct {
	DurationMilliseconds int                           `json:"duration_millisecnds,omitempty"`
	Rows                 []SwitchInterfaceSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string                    `json:"rows_order,omitempty"`
	RowsTotal            int                           `json:"rows_total,omitempty"`
}

// SwitchInterfaceSearchResult Represents a switch interface-to-server interface mapping.
type SwitchInterfaceSearchResult struct {
	ServerID                                  int        `json:"server_id,omitempty" yaml:"serverID,omitempty"`
	ServerIPMIHost                            string     `json:"server_ipmi_host,omitempty" yaml:"serverIPMIHost,omitempty"`
	ServerSerialNumber                        string     `json:"server_serial_number,omitempty" yaml:"serverSerialNumber,omitempty"`
	ServerTypeID                              int        `json:"server_type_id,omitempty" yaml:"ServerTypeID,omitempty"`
	NetworkEquipmentID                        int        `json:"network_equipment_id,omitempty" yaml:"networkEquipmentID,omitempty"`
	NetworkEquipmentIdentifierString          string     `json:"network_equipment_identifier_string,omitempty" yaml:"networkEquipmentIdentifierString,omitempty"`
	NetworkEquipmentManagementAddress         string     `json:"network_equipment_management_address,omitempty" yaml:"networkEquipmentManagementAddress,omitempty"`
	NetworkEquipmentInterfaceID               int        `json:"network_equipment_interface_id,omitempty" yaml:"networkEquipmentInterfaceID,omitempty"`
	NetworkEquipmentInterfaceMACAddress       string     `json:"network_equipment_interface_mac_address,omitempty" yaml:"networkEquipmentInterfaceMACAddress,omitempty"`
	NetworkEquipmentInterfaceIdentifierString string     `json:"network_equipment_interface_identifier_string,omitempty" yaml:"networkEquipmentInterfaceIdentifierString,omitempty"`
	ServerInterfaceMACAddress                 string     `json:"server_interface_mac_address,omitempty" yaml:"serverInterfaceMACAddress,omitempty"`
	ServerInterfaceIndex                      int        `json:"server_interface_index,omitempty" yaml:"serverInterfaceIndex,omitempty"`
	ServerInterfaceCapacityMBPs               int        `json:"server_interface_capacity_mbps,omitempty" yaml:"serverInterfaceCapacityMBPs,omitempty"`
	IP                                        [][]string `json:"ip_human_readable,omitempty" yaml:"ip,omitempty"`
	SubnetRangeStart                          [][]string `json:"subnet_range_start_human_readable,omitempty" yaml:"rangeStart,omitempty"`
	NetworkType                               []string   `json:"network_type,omitempty" yaml:"networkType,omitempty"`
	NetworkID                                 []int      `json:"network_id,omitempty" yaml:"networkID,omitempty"`
	SubnetPrefixSize                          [][]int    `json:"subnet_prefix_size,omitempty" yaml:"subnetPrefixSize,omitempty"`
	SubnetPoolID                              [][]int    `json:"subnet_pool_id,omitempty" yaml:"subnetPoolID,omitempty"`
}

// SwitchInterfaceSearch searches for server interfaces filtering on various elements such as switch id or server id
func (c *Client) SwitchInterfaceSearch(filter string) (*[]SwitchInterfaceSearchResult, error) {

	tables := []string{"_switch_interfaces"}
	columns := map[string][]string{
		"_switch_interfaces": {
			"server_id",
			"server_ipmi_host",
			"server_serial_number",
			"server_type_id",
			"network_equipment_interface_id",
			"network_equipment_identifier_string",
			"network_equipment_management_address",
			"network_equipment_id",
			"network_equipment_interface_identifier_string",
			"network_equipment_interface_mac_address",
			"server_interface_mac_address",
			"server_interface_index",
			"server_interface_capacity_mbps",
			"ip_human_readable",
			"subnet_range_start_human_readable",
			"network_type",
			"network_id",
			"subnet_prefix_size",
			"subnet_pool_id",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	var createdObject map[string]searchResultResponseWrapperForSwitchInterfaces

	resp, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]searchResultResponseWrapperForSwitchInterfaces{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []SwitchInterfaceSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil
}
