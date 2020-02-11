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

//ServerComponent information about a server's components
type ServerComponent struct {
	ServerComponentID                              int      `json:"server_component_id,omitempty"`
	ServerID                                       int      `json:"server_id,omitempty"`
	ServerComponentName                            string   `json:"server_component_name,omitempty"`
	ServerComponentFirmwareVersion                 string   `json:"server_component_firmware_version,omitempty"`
	ServerComponentFirmwareUpdateable              bool     `json:"server_component_firmware_updateable,omitempty"`
	ServerComponentFirmwareJSON                    string   `json:"server_component_firmware_json,omitempty"`
	ServerComponentFirmwareUpdateAvailableVersions []string `json:"server_component_firmware_update_available_versions,omitempty"`
	ServerComponentFirmwareStatus                  string   `json:"server_component_firmware_status,omitempty"`
	ServerComponentType                            string   `json:"server_component_type,omitempty"`
	ServerComponentFirmwareUpdateTimestamp         string   `json:"server_component_firmware_update_timestamp,omitempty"`
	ServerComponentFirmwareTargetVersion           string   `json:"server_component_firmware_target_version,omitempty"`
	ServerComponentFirmwareScheduledTimestamp      string   `json:"server_component_firmware_scheduled_timestamp,omitempty"`
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

//ServerFirmwareComponentUpgrade Creates a firmware upgrading session for the specified component.
//If no strServerComponentFirmwareNewVersion or strFirmwareBinaryURL are provided the system will use the values from the database which should have been previously added
func (c *Client) ServerFirmwareComponentUpgrade(serverID int, serverComponentID int, serverComponentFirmwareNewVersion string, firmwareBinaryURL string) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_upgrade",
		serverID,
		serverComponentID,
		serverComponentFirmwareNewVersion,
		firmwareBinaryURL,
	)
	return err
}

//ServerFirmwareUpgrade creates a firmware upgrading session that affects all components from the specified server that have a target version set and are updatable.
func (c *Client) ServerFirmwareUpgrade(serverID int) error {

	_, err := c.rpcClient.Call(
		"server_firmware_upgrade",
		serverID,
	)
	return err
}

//ServerFirmwareComponentTargetVersionSet Sets a firmware target version for the upgrading process. The system will apply the upgrade at the next upgrading session.
func (c *Client) ServerFirmwareComponentTargetVersionSet(serverComponentID int, serverComponentFirmwareNewVersion string) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_target_version_set",
		serverComponentID,
		serverComponentFirmwareNewVersion,
	)
	return err
}

//ServerFirmwareComponentTargetVersionUpdate Updates for every component of the specified server the available firmware versions that can be used as target by the firmware upgrading process. The available versions are extracted from a vendor specific catalog.
func (c *Client) ServerFirmwareComponentTargetVersionUpdate(serverComponentID int) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_available_versions_update",
		serverComponentID,
	)
	return err
}

//ServerFirmwareComponentTargetVersionAdd Adds a new available firmware version for a server component along with the url of the binary. If the version already exists the old url will be overwritten.
func (c *Client) ServerFirmwareComponentTargetVersionAdd(serverComponentID int, version string, firmareBinaryURL string) error {

	_, err := c.rpcClient.Call(
		"server_firmware_component_available_versions_add",
		serverComponentID,
		version,
		firmareBinaryURL,
	)
	return err
}

//ServerComponentGet returns a server's component's details
func (c *Client) ServerComponentGet(serverComponentID int) (*ServerComponent, error) {

	var createdObject ServerComponent

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_get_internal",
		serverComponentID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}
