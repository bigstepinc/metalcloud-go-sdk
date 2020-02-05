package metalcloud

import (
	"fmt"
	"strings"
)

//ServerSearchResult represents a server in a datacenter.
type ServerSearchResult struct {
	ServerID                           int               `json:"server_id,omitempty"`
	ServerUUID                         string            `json:"server_uuid,omitempty"`
	ServerNetworkTotalCapacityMbps     int               `json:"server_network_total_capacity_mbps,omitempty"`
	ServerPowerStatus                  string            `json:"server_power_status,omitempty"`
	ServerProcessorCoreCount           int               `json:"server_processor_core_count,omitempty"`
	ServerProcessorCoreMhz             int               `json:"server_processor_core_mhz,omitempty"`
	ServerProcessorCount               int               `json:"server_processor_count,omitempty"`
	ServerRAMGbytes                    int               `json:"server_ram_gbytes,omitempty"`
	ServerDiskCount                    int               `json:"server_disk_count,omitempty"`
	ServerDiskSizeMbytes               int               `json:"server_disk_size_mbytes,omitempty"`
	ServerDiskType                     string            `json:"server_disk_type,omitempty"`
	ServerProcessorName                string            `json:"server_processor_name,omitempty"`
	ServerProductName                  string            `json:"server_product_name,omitempty"`
	ServerTypeID                       int               `json:"server_type_id,omitempty"`
	ServerTypeName                     string            `json:"server_type_name,omitempty"`
	ServerTypeBootType                 string            `json:"server_type_boot_type,omitempty"`
	ServerInterfaces                   []ServerInterface `json:"server_interfaces,omitempty"`
	ServerDisks                        []ServerDisk      `json:"server_disks,omitempty"`
	ServerTags                         []string          `json:"server_tags,omitempty"`
	ServerIPMIHost                     string            `json:"server_ipmi_host,omitempty"`
	ServerIPMInternalUsername          string            `json:"server_ipmi_internal_username,omitempty"`
	ServerIPMInternalPasswordEncrypted string            `json:"server_ipmi_internal_password_encrypted,omitempty"`
	ServerStatus                       string            `json:"server_status,omitempty"`
	ServerSerialNumber                 string            `json:"server_serial_number,omitempty"`
	ServerVendor                       string            `json:"server_vendor,omitempty"`
	ServerVendorSKUID                  string            `json:"server_vendor_sku_id,omitempty"`
	ServerComments                     string            `json:"server_comments,omitempty"`
	InstanceLabel                      string            `json:"instance_label,omitempty"`
	InstanceID                         int               `json:"instance_id,omitempty"`
	InstanceArrayID                    int               `json:"instance_array_id,omitempty"`
	InfrastructureID                   int               `json:"infrastructure_id,omitempty"`
	DatacenterName                     string            `json:"datacenter_name,omitempty"`
}

//Server represents a server in a datacenter.
type Server struct {
	ServerID                       int               `json:"server_id,omitempty"`
	ServerUUID                     string            `json:"server_uuid,omitempty"`
	ServerNetworkTotalCapacityMbps int               `json:"server_network_total_capacity_mbps,omitempty"`
	ServerPowerStatus              string            `json:"server_power_status,omitempty"`
	ServerProcessorCoreCount       int               `json:"server_processor_core_count,omitempty"`
	ServerProcessorCoreMhz         int               `json:"server_processor_core_mhz,omitempty"`
	ServerProcessorCount           int               `json:"server_processor_count,omitempty"`
	ServerRAMGbytes                int               `json:"server_ram_gbytes,omitempty"`
	ServerDiskCount                int               `json:"server_disk_count,omitempty"`
	ServerDiskSizeMbytes           int               `json:"server_disk_size_mbytes,omitempty"`
	ServerDiskType                 string            `json:"server_disk_type,omitempty"`
	ServerProcessorName            string            `json:"server_processor_name,omitempty"`
	ServerProductName              string            `json:"server_product_name,omitempty"`
	ServerTypeID                   int               `json:"server_type_id,omitempty"`
	ServerInterfaces               []ServerInterface `json:"server_interfaces,omitempty"`
	ServerDisks                    []ServerDisk      `json:"server_disks,omitempty"`
	ServerTags                     []string          `json:"server_tags,omitempty"`
	ServerIPMIHost                 string            `json:"server_ipmi_host,omitempty"`
	ServerIPMInternalUsername      string            `json:"server_ipmi_internal_username,omitempty"`
	ServerIPMInternalPassword      string            `json:"server_ipmi_internal_password,omitempty"`
	ServerStatus                   string            `json:"server_status,omitempty"`
	ServerSerialNumber             string            `json:"server_serial_number,omitempty"`
	ServerVendor                   string            `json:"server_vendor,omitempty"`
	ServerVendorSKUID              string            `json:"server_vendor_sku_id,omitempty"`
	ServerComments                 string            `json:"server_comments,omitempty"`
	DatacenterName                 string            `json:"datacenter_name,omitempty"`
}

//await bsideveloper.search(2, '*',["_servers_instances"])

//ServerDisk describes a disk
type ServerDisk struct {
	ServerDiskType   string `json:"server_disk_type,omitempty"`
	ServerDiskSizeGB string `json:"server_disk_size_gb,omitempty"`
}

//SearchResultForServers describes a serach result
type SearchResultForServers struct {
	DurationMilliseconds int                  `json:"duration_millisecnds,omitempty"`
	Rows                 []ServerSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string           `json:"rows_order,omitempty"`
	RowsTotal            int                  `json:"rows_total,omitempty"`
}

//ServersSearch searches for servers matching certain filter
func (c *Client) ServersSearch(filter string) (*[]ServerSearchResult, error) {

	tables := []string{"_servers_instances"}
	columns := map[string][]string{
		"_servers_instances": {
			"server_id",
			"server_type_name",
			"server_type_boot_type",
			"server_product_name",
			"server_status",
			"datacenter_name",
			"server_class",
			"server_created_timestamp",
			"server_vendor",
			"server_serial_number",
			"server_uuid",
			"server_bios_version",
			"server_vendor_sku_id",
			"server_boot_type",
			"server_allocation_timestamp",
			"instance_label",
			"instance_id",
			"instance_array_id",
			"infrastructure_id",
			"server_ipmi_host",
			"server_custom_json",
			"server_ipmi_internal_username",
			"server_ipmi_internal_password",
			"server_processor_name",
			"server_processor_count",
			"server_processor_core_count",
			"server_processor_core_mhz",
			"server_disk_type",
			"server_disk_count",
			"server_disk_size_mbytes",
			"server_ram_gbytes",
			"server_network_total_capacity_mbps",
			"server_dhcp_status",
			"server_dhcp_packet_sniffing_is_enabled",
			"server_dhcp_relay_security_is_enabled",
			"server_disk_wipe",
			"server_power_status",
			"server_power_status_last_update_timestamp",
			"user_id",
			"user_id_owner",
			"user_email",
		},
	}

	userID := c.GetUserID()

	collapseType := "none"

	res, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType)

	if err != nil {
		return nil, err
	}

	var createdObject map[string]SearchResultForServers

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	servers := []ServerSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		servers = append(servers, s)
	}

	return &servers, nil
}

//ServerGet returns a server's details
func (c *Client) ServerGet(serverID int, decryptPasswd bool) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_get_internal",
		serverID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.ServerIPMInternalPassword, ":")
		if len(passwdComponents) != 2 {
			return nil, fmt.Errorf("Password not returned with proper components")
		}
		var passwd string
		err = c.rpcClient.CallFor(
			&passwd,
			"password_decrypt",
			passwdComponents[1],
		)
		if err != nil {
			return nil, err
		}
		createdObject.ServerIPMInternalPassword = passwd
	} else {
		createdObject.ServerIPMInternalPassword = ""
	}

	return &createdObject, nil
}
