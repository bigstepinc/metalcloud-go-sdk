package metalcloud

//go:generate go run helper/docgen.go - $GOFILE ./ Server,ServerDisk,ServerInterface,ServerNICDetails Server

import (
	"fmt"
	"strings"
)

// ServerSearchResult represents a server in a datacenter.
type ServerSearchResult struct {
	ServerID                       int               `json:"server_id,omitempty" yaml:"id,omitempty"`
	ServerUUID                     string            `json:"server_uuid,omitempty" yaml:"uuid,omitempty"`
	ServerStatus                   string            `json:"server_status,omitempty" yaml:"status,omitempty"`
	ServerSerialNumber             string            `json:"server_serial_number,omitempty" yaml:"serial_number,omitempty"`
	ServerVendor                   string            `json:"server_vendor,omitempty" yaml:"vendor,omitempty"`
	ServerNetworkTotalCapacityMbps int               `json:"server_network_total_capacity_mbps,omitempty" yaml:"networkTotalCapacityMbps,omitempty"`
	ServerBootType                 string            `json:"server_boot_type,omitempty" yaml:"bootType,omitempty"`
	ServerPowerStatus              string            `json:"server_power_status,omitempty" yaml:"powerStatus,omitempty"`
	ServerProcessorName            string            `json:"server_processor_name,omitempty" yaml:"processorName,omitempty"`
	ServerProcessorCoreCount       int               `json:"server_processor_core_count,omitempty" yaml:"processorCoreCount,omitempty"`
	ServerProcessorCoreMhz         int               `json:"server_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	ServerProcessorCount           int               `json:"server_processor_count,omitempty" yaml:"processorCount,omitempty"`
	ServerProcessorThreads         int               `json:"server_processor_threads,omitempty" yaml:"processorThreads,omitempty"`
	ServerProcessorCPUMark         int               `json:"server_processor_cpu_mark" yaml:"processorCPUMark"`
	ServerRAMGbytes                int               `json:"server_ram_gbytes,omitempty" yaml:"ramGbytes,omitempty"`
	ServerDiskCount                int               `json:"server_disk_count,omitempty" yaml:"diskCount,omitempty"`
	ServerDiskSizeMbytes           int               `json:"server_disk_size_mbytes" yaml:"diskSizeMbytes,omitempty"`
	ServerDiskType                 string            `json:"server_disk_type,omitempty" yaml:"diskType,omitempty"`
	ServerProductName              string            `json:"server_product_name,omitempty" yaml:"productName,omitempty"`
	ServerClass                    string            `json:"server_class,omitempty" yaml:"class,omitempty"`
	ServerTypeID                   int               `json:"server_type_id,omitempty" yaml:"typeID,omitempty"`
	ServerTypeName                 string            `json:"server_type_name,omitempty" yaml:"type,omitempty"`
	ServerTypeBootType             string            `json:"server_type_boot_type,omitempty" yaml:"serverBootType,omitempty"`
	ServerInterfaces               []ServerInterface `json:"server_interfaces,omitempty" yaml:"interfaces,omitempty"`
	ServerRackName                 string            `json:"server_rack_name" yaml:"rackName"`
	ServerRackPositionLowerUnit    string            `json:"server_rack_position_lower_unit,omitempty" yaml:"rackPositionLowerUnit,omitempty"`
	ServerRackPositionUpperUnit    string            `json:"server_rack_position_upper_unit,omitempty" yaml:"rackPositionUpperUnit,omitempty"`
	ServerInventoryId              string            `json:"server_inventory_id,omitempty" yaml:"inventoryId,omitempty"`
	ServerDisks                    []ServerDisk      `json:"server_disks,omitempty" yaml:"disks,omitempty"`
	ServerSupportsOOBProvisioning  bool              `json:"server_supports_oob_provisioning" yaml:"supportsOOBProvisioning"`
	ServerTags                     []string          `json:"server_tags,omitempty" yaml:"tags,omitempty"`
	ServerIPMIChannel              int               `json:"server_ipmi_channel,omitempty" yaml:"IPMIChannel,omitempty"`
	ServerIPMIHost                 string            `json:"server_ipmi_host,omitempty" yaml:"IPMIHostname,omitempty"`
	ServerIPMInternalUsername      string            `json:"server_ipmi_internal_username,omitempty" yaml:"IPMIUsername,omitempty"`
	ServerIPMInternalPassword      string            `json:"server_ipmi_internal_password,omitempty" yaml:"IPMIPassword,omitempty"`
	ServerIPMCredentialsNeedUpdate bool              `json:"server_ipmi_credentials_need_update,omitempty" yaml:"IPMICredentialsNeedUpdate,omitempty"`
	ServerVendorSKUID              string            `json:"server_vendor_sku_id,omitempty" yaml:"vendorSKUID,omitempty"`
	ServerComments                 string            `json:"server_comments,omitempty" yaml:"comments,omitempty"`
	InstanceLabel                  []string          `json:"instance_label,omitempty" yaml:"instanceLabel,omitempty"`
	InstanceID                     []int             `json:"instance_id,omitempty" yaml:"instanceID,omitempty"`
	InstanceArrayID                []int             `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	InfrastructureID               []int             `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	UserEmail                      [][]string        `json:"user_email,omitempty" yaml:"userEmail,omitempty"`
	UserID                         [][]int           `json:"user_id,omitempty" yaml:"users,omitempty"`
	DatacenterName                 string            `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	ServerSubmodel                 string            `json:"server_submodel,omitempty" yaml:"submodel,omitempty"`
}

// Server represents a server in a datacenter. Many of these fields are readonly.
// Servers should always be registered via the registration process but in certain circumstances such as brownfield setups
// it may be required to create these entries by hand.
type Server struct {
	//description: id of the server
	ServerID int `json:"server_id,omitempty" yaml:"id,omitempty"`
	//description: UUID of the server. Readonly.
	ServerUUID string `json:"server_uuid,omitempty" yaml:"UUID,omitempty"`
	//description: Status of a server. Note that not all transitions between states are possible.
	//values:
	// - registering
	// - available
	// - cleaning_required
	// - cleaning
	// - used
	// - used_registering
	// - unavailable
	// - defective
	// - removed_from_rack
	// - decommissioned
	// - updating_firmware
	// - used_diagnostics
	ServerStatus string `json:"server_status,omitempty" yaml:"status,omitempty"`
	//description: serial number
	ServerSerialNumber string `json:"server_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	//description: Vendor. Must be one of the supported values.
	//values:
	// - HP
	// - HPE
	// - Dell Inc.
	// - Lenovo
	// - Supermicro
	// - BULL
	// - QEMU
	// - BSI
	ServerVendor string `json:"server_vendor,omitempty" yaml:"vendor,omitempty"`
	//description: Datacenter where this datacenter is registered
	DatacenterName string `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	//description: Summed up value of total network capacities of all in-band network connections of the server.
	ServerNetworkTotalCapacityMbps int `json:"server_network_total_capacity_mbps,omitempty" yaml:"networkTotalCapacityMbps,omitempty"`
	//description: Type of boot this server uses
	//values:
	// - classic
	// - uefi
	ServerBootType string `json:"server_boot_type,omitempty" yaml:"bootType,omitempty"`
	//description: Server power status
	//values:
	// - on
	// - off
	// - unknown
	ServerPowerStatus string `json:"server_power_status,omitempty" yaml:"powerStatus,omitempty"`
	//description: Server processor name
	//example: Intel(R) Xeon(R) CPU E5345  @ 2.33GHz"
	ServerProcessorName string `json:"server_processor_name,omitempty" yaml:"processorName,omitempty"`
	//description: Server total core count
	ServerProcessorCoreCount int `json:"server_processor_core_count,omitempty" yaml:"processorCoreCount,omitempty"`
	//description: Server maximum frequency
	ServerProcessorCoreMhz int `json:"server_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	//description: Server CPU count
	ServerProcessorCount int `json:"server_processor_count,omitempty" yaml:"processorCount,omitempty"`
	//description: Server threads per core count
	ServerProcessorThreads int `json:"server_processor_threads,omitempty" yaml:"processorThreads,omitempty"`
	//description: Deprecated
	ServerProcessorCPUMark int `json:"server_processor_cpu_mark" yaml:"processorCPUMark"`
	//description: Total amount of RAM in Gbytes
	ServerRAMGbytes int `json:"server_ram_gbytes,omitempty" yaml:"ramGbytes,omitempty"`
	//description: The list of disks
	ServerDisks []ServerDisk `json:"server_disks" yaml:"disks,omitempty"`
	//description: The count of disks
	ServerDiskCount int `json:"server_disk_count" yaml:"diskCount,omitempty"`
	//description: If all disks are identical this is the size of the disk. Deprecated.
	ServerDiskSizeMbytes int `json:"server_disk_size_mbytes" yaml:"diskSizeMbytes"`
	//description: If all disks are identical this type of disk. Deprecated.
	//values:
	// - none
	// - NVME
	// - HDD
	// - SSD
	ServerDiskType string `json:"server_disk_type,omitempty" yaml:"diskType,omitempty"`
	//description: Name of rack (metadata)
	ServerRackName *string `json:"server_rack_name" yaml:"rackName"`
	//description: Rack position lower U (metadata)
	ServerRackPositionLowerUnit *string `json:"server_rack_position_lower_unit" yaml:"rackPositionLowerUnit"`
	//description: Rack position upper U (metadata)
	ServerRackPositionUpperUnit *string `json:"server_rack_position_upper_unit" yaml:"rackPositionUpperUnit"`
	//description: Rack ID (metadata)
	ServerRackId *string `json:"server_rack_id" yaml:"rackID"`
	//description: ID of the chassis if any
	ChassisRackId *int `json:"chassis_rack_id" yaml:"chassisRackID"`
	//description: Inventory ID (metadata)
	ServerInventoryId *string `json:"server_inventory_id" yaml:"inventoryId"`
	//description: Name of the server as returned by the server itself
	//example: PowerEdge 1950
	ServerProductName string `json:"server_product_name,omitempty" yaml:"productName,omitempty"`
	//description: Class of server. Deprecated
	ServerClass string `json:"server_class,omitempty" yaml:"serverClass,omitempty"`
	//description: The ID of the server type. Note that the server types are normally automatically determined during registration. However this value can be changed by the client if needed although it is not recommended.
	ServerTypeID int `json:"server_type_id,omitempty" yaml:"serverTypeID,omitempty"`
	//description: Descriptions of the server's interfaces
	ServerInterfaces []ServerInterface `json:"server_interfaces,omitempty" yaml:"interfaces,omitempty"`
	//description: Set to true if the server can be deployed via the out-of-band network (currently via virtual media)
	ServerSupportsOOBProvisioning bool `json:"server_supports_oob_provisioning" yaml:"supportsOOBProvisioning"`
	//description: Tags (metadata)
	ServerTags []string `json:"server_tags" yaml:"tags"`
	//description: The IPMI channel to use when communicating on the IPMI protocol. Used in Legacy scenarios.
	ServerIPMIChannel int `json:"server_ipmi_channel" yaml:"IPMIChannel"`
	//description: The IPMI host  to use when communicating on the IPMI protocol. Used in Legacy scenarios.
	ServerIPMIHost string `json:"server_ipmi_host,omitempty" yaml:"IPMIHostname,omitempty"`
	//description: The IPMI username to use when communicating on the IPMI protocol. Used in Legacy scenarios.
	ServerIPMInternalUsername string `json:"server_ipmi_internal_username,omitempty" yaml:"IPMIUsername,omitempty"`
	//description: The IPMI username to use when communicating on the IPMI protocol. Used in Legacy scenarios. Use this field to set the password. Write-only.
	ServerIPMInternalPassword string `json:"server_ipmi_internal_password,omitempty" yaml:"IPMIPassword,omitempty"`
	//description: The IPMI username to use when communicating on the IPMI protocol. Used in Legacy scenarios. Readonly. Encrypted.
	ServerIPMInternalPasswordEncrypted string `json:"server_ipmi_internal_password_encrypted,omitempty" yaml:"IPMIPasswordEncrypted,omitempty"`
	//description: Internal field denoting that the credentials will be updated.
	ServerIPMCredentialsNeedUpdate bool `json:"server_ipmi_credentials_need_update" yaml:"IPMICredentialsNeedUpdate"`
	//description: SKU ID returned by the server. Note that not all vendors return this.
	ServerVendorSKUID string `json:"server_vendor_sku_id,omitempty" yaml:"vendorSKU,omitempty"`
	//description: Comments on the server (metadata)
	ServerComments string `json:"server_comments" yaml:"comments,omitempty"`
	//description: Internal
	ServerBIOSInfoJSON string `json:"server_bios_info_json" yaml:"BIOSInfoJson"`
	//description: Internal
	ServerCustomJSON *string `json:"server_custom_json" yaml:"CustomJSON"`
	//description: Internal
	ServerInfoJSON *string `json:"server_info_json" yaml:"infoJSON"`
	//description: Internal
	ServerDetailXML string `json:"server_details_xml" yaml:"serverDetailsXML"`
	//description: Internal
	ServerInstanceCustomJSON *string `json:"server_instance_custom_json" yaml:"instanceCustomJSON"`
	//description: Internal
	ServerSupportsSOL bool `json:"server_supports_sol" yaml:"supportsSOL"`
	//description: Internal
	ServerILOResetTimestamp string `json:"server_ilo_reset_timestamp" yaml:"ILOResetTimestamp"`
	//description: Timestamp value denoting the last known boot.
	ServerBootLastUpdateTimestamp string `json:"server_boot_last_update_timestamp" yaml:"BootLastUpdateTimestamp"`
	//description: Timestamp value denoting the last power status update.
	ServerPowerStatusUpdateTimestamp string `json:"server_power_status_last_update_timestamp" yaml:"PowerStatusUpdateTimestamp"`
	//description: What OOB subnet is used to allocate ips to the BMC of this server
	SubnetOOBID int `json:"subnet_oob_id" yaml:"subnetOOBID"`
	//description: DHCP server's behavior in relation to this server.
	//values:
	// - quarantine
	// - deny_requests
	// - ansible_provision
	// - allow_requests
	ServerDHCPStatus string `json:"server_dhcp_status" yaml:"subnetDHCPStatus"`
	//description: MAC address of the server's BMC
	ServerBMCMACAddress string `json:"server_bmc_mac_address" yaml:"BMCMACAddress"`
	//description: SNMP Community password encrypted. Internal.
	ServerCommunityPasswordDCEncrypted string `json:"snmp_community_password_dcencrypted" yaml:"SNMPCommunityPaswordDCencrypted"`
	//description: SNMP Community password encrypted
	ServerMgmtSNMPCommunityPasswordDCEncrypted string `json:"server_mgmt_snmp_community_password_dcencrypted" yaml:"MGMTNMPCommunityPasswordDCEncrypted"`
	//description: SNMP Community version
	ServerMgmtSNMPVersion int `json:"server_mgmt_snmp_version" yaml:"MGMTSNMPVersion"`
	//description: SNMP Community port
	ServerMgmtSNMPPort int `json:"server_mgmt_snmp_port" yaml:"MGMTSNMPPort"`
	//description: If the Secure Boot is enabled
	ServerSecureBootIsEnabled bool `json:"server_secure_boot_is_enabled" yaml:"secureBootIsEnabled"`
	//description: Internal
	ServerOOBIndex int `json:"subnet_oob_index" yaml:"subnetOOBIndex"`
	//description: The version of the IPMI protocol when IPMI is used.
	ServerIPMIVersion string `json:"server_ipmi_version" yaml:"subnetIPMIVersion"`
	//description: ISO 8601 timestamp which holds the date and time when the server record was created. Readonly.
	//example: 2013-11-29T13:00:01Z
	ServerCreatedTimestamp string `json:"server_created_timestamp" yaml:"createdTimestamp"`
	//description: Last cleanup timestamp
	//example: 2013-11-29T13:00:01Z
	ServerLastCleanupStart string `json:"server_last_cleanup_start" yaml:"lastCleanupStart"`
	//description: Version information as returned by the server.
	ServerVersionInfoJSON string `json:"server_vendor_info_json" yaml:"vendorInfoJSON"`
	//description: Internal
	ServerChipsetName string `json:"server_chipset_name" yaml:"chipsetName"`
	//description: Set if the server is considered 'dirty' and needs re-registration.
	ServerRequiresReregister bool `json:"server_requires_reregister" yaml:"requiresReregiter"`
	//description: Count of GPUs
	ServerGPUCount int `json:"server_gpu_count" yaml:"GPUCount"`
	//description: Model of GPUs if all are identical
	//example: H100
	ServerGPUModel string `json:"server_gpu_model" yaml:"GPUModel"`
	//description: Vendor of GPUs if all are identical
	//example: NVidia
	ServerGPUVendor string `json:"server_gpu_vendor" yaml:"GPUVendor"`
	//description: When the server was allocated to an infrastructure.
	//example: 2013-11-29T13:00:01Z
	ServerAllocationTimestamp string `json:"server_allocation_timestamp" yaml:"allocationTimestamp"`
	//description:  If enabled, the DHCP server will use verify the DHCP option82 for added security in interpreting the DHCP packet. Used for diagnosing issues with the legacy protocol.
	ServerDHCPRelaySecurityIsEnabled bool `json:"server_dhcp_relay_security_is_enabled" yaml:"DHCPRelaySecurityIsEnabled"`
	//description: If true, some steps of the cleanup process could not be performed automatically.
	ServerRequiresManualCleaning bool `json:"server_requires_manual_cleaning" yaml:"requiresManualCleaning"`
	//description: If true, a cleanup operation is ongoing.
	ServerCleanupInProgress bool `json:"server_cleanup_in_progress" yaml:"cleanupInProgress"`
	//description: If true, server supports virtual media-based deployment
	ServerSupportsVirtualMedia bool `json:"server_supports_virtual_media" yaml:"serverSupportsVirtualMedia"`
	//description: If true, server is ongoing a diagnosis. Deprecated.
	ServerIsInDiagnostics bool `json:"server_is_in_diagnostics" yaml:"serverIsInDiagnostics"`
	//description: If true, server will be wiped.
	ServerDiskWipe bool `json:"server_disk_wipe" yaml:"diskWipe"`
	//description: Internal.
	ServerMetricsMetadataJSON *string `json:"server_metrics_metadata_json" yaml:"metricsMetadataJSON"`
	//description: Id of the site controller agent managing this server.
	AgentID int `json:"agent_id" yaml:"agentID"`
	//description: If set to true the DHCP server will record a log of the dialog with the server. Useful for diagnosing issues with the legacy protocol.
	ServerDHCPPacketSniffingIsEnabled bool `json:"server_dhcp_packet_sniffing_is_enabled" yaml:"DHCPPacketSniffingIsEnabled"`
	//description: Internal
	ServerBDKDEbug bool `json:"server_bdk_debug" yaml:"BDKDebug"`
	//description: Internal
	ServerKeysJSON string `json:"server_keys_json" yaml:"keysJSON"`
	//description: Details about the network interfaces
	NICDetails map[string]ServerNICDetails `json:"nic_details,omitempty" yaml:"NICDetails,omitempty"`
	//description: Model details as returned by the server
	ServerSubmodel string `json:"server_submodel,omitempty" yaml:"submodel,omitempty"`
}

// ServerDisk describes a disk
type ServerDisk struct {
	//description: The id of the object
	ServerDiskID int `json:"server_disk_id,omitempty" yaml:"id,omitempty"`
	//description: The model of the object
	ServerDiskModel string `json:"server_disk_model,omitempty" yaml:"model,omitempty"`
	//description: The type of the disk
	//values:
	// - NVME
	// - SSD
	// - HDD
	ServerDiskType string `json:"server_disk_type,omitempty" yaml:"type,omitempty"`
	//description: The vendor of the disk
	ServerDiskVendor string `json:"server_disk_vendor,omitempty" yaml:"vendor,omitempty"`
	//description: The status of the disk
	ServerDiskStatus string `json:"server_disk_status,omitempty" yaml:"status,omitempty"`
	//description: The serial number of the disk
	ServerDiskSerial string `json:"server_disk_serial,omitempty" yaml:"serial_number,omitempty"`
	//description: The size of the disk
	ServerDiskSizeGB int `json:"server_disk_size_gb,omitempty" yaml:"sizeGB,omitempty"`
}

// ServerInterface contains server connectivity information.
type ServerInterface struct {
	//MAC address of the interface
	ServerInterfaceMACAddress string `json:"server_interface_mac_address,omitempty" yaml:"macAddress,omitempty"`
}

// SearchResultForServers describes a serach result
type SearchResultForServers struct {
	DurationMilliseconds int                  `json:"duration_millisecnds,omitempty"`
	Rows                 []ServerSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string           `json:"rows_order,omitempty"`
	RowsTotal            int                  `json:"rows_total,omitempty"`
}

// ServerComponent information about a server's components
type ServerComponent struct {
	ServerComponentID                              int      `json:"server_component_id,omitempty" yaml:"id,omitempty"`
	ServerID                                       int      `json:"server_id,omitempty" yaml:"serverID,omitempty"`
	ServerComponentName                            string   `json:"server_component_name,omitempty" yaml:"componentName,omitempty"`
	ServerComponentFirmwareVersion                 string   `json:"server_component_firmware_version,omitempty" yaml:"firmwareVersion,omitempty"`
	ServerComponentFirmwareUpdateable              bool     `json:"server_component_firmware_updateable,omitempty" yaml:"firmwareUpdateable,omitempty"`
	ServerComponentFirmwareJSON                    string   `json:"server_component_firmware_json,omitempty" yaml:"firmwareJSON,omitempty"`
	ServerComponentFirmwareUpdateAvailableVersions []string `json:"server_component_firmware_update_available_versions,omitempty" yaml:"firmwareUpdateAvailableVersions,omitempty"`
	ServerComponentFirmwareStatus                  string   `json:"server_component_firmware_status,omitempty" yaml:"firmwareStatus,omitempty"`
	ServerComponentType                            string   `json:"server_component_type,omitempty" yaml:"type,omitempty"`
	ServerComponentFirmwareUpdateTimestamp         string   `json:"server_component_firmware_update_timestamp,omitempty" yaml:"firmwareUpdateTimestamp,omitempty"`
	ServerComponentFirmwareTargetVersion           string   `json:"server_component_firmware_target_version,omitempty" yaml:"firmwareTargetVersion,omitempty"`
	ServerComponentFirmwareScheduledTimestamp      string   `json:"server_component_firmware_scheduled_timestamp,omitempty" yaml:"firmwareScheduledTimestamp,omitempty"`
}

// SearchResultForServerComponents describes a search result
type SearchResultForServerComponents struct {
	DurationMilliseconds int               `json:"duration_millisecnds,omitempty"`
	Rows                 []ServerComponent `json:"rows,omitempty"`
	RowsOrder            [][]string        `json:"rows_order,omitempty"`
	RowsTotal            int               `json:"rows_total,omitempty"`
}

type ServerNICDetails struct {
	//description: LLDP information for this interface
	NetworkEquipmentInterfaceLLDPInformation string `json:"network_equipment_interface_lldp_information,omitempty" yaml:"networkEquipmentInterfaceLLDPInformation,omitempty"`
	//description: MAC Address of the switch interface
	NetworkEquipmentInterfaceMACAddress string `json:"network_equipment_interface_mac_address,omitempty" yaml:"networkEquipmentInterfaceMACAddress,omitempty"`
	//description: Switch ID of the switch to which this interface is connected to
	SwitchPortID string `json:"switch_port_id,omitempty" yaml:"switchPortID,omitempty"`
	//description: Switch port name
	SwitchPortDescription string `json:"switch_port_description,omitempty" yaml:"switchPortDescription,omitempty"`
	//description: Hostname of the switch
	SwitchHostname string `json:"switch_hostname,omitempty" yaml:"switchHostname,omitempty"`
	//description: Description of the switch
	NetworkEquipmentDescription string `json:"network_equipment_description,omitempty" yaml:"networkEquipmentDescription,omitempty"`
	//description: VLAN ID allocated as native to this interface
	SwitchVLANID string `json:"switch_vlan_id,omitempty" yaml:"switchVLANID,omitempty"`
	//description: Index of this interface on the server. This is determined based on the switch hostname and port description.
	SwitchInterfaceIndex int `json:"server_interface_index,omitempty" yaml:"switchInterfaceIndex,omitempty"`
	//description: MAC Address of the server interface
	ServerInterfaceMACAddress string `json:"server_interface_mac_address,omitempty" yaml:"serverInterfaceMACAddress,omitempty"`
	//description: MAX Speed of the interface
	ServerInterfaceCapacityMBPs int `json:"server_interface_capacity_mbps,omitempty" yaml:"serverInterfaceCapacityMBPs,omitempty"`
}

type ServerEditRack struct {
	ServerRackName              *string `json:"server_rack_name" yaml:"rackName"`
	ServerRackPositionLowerUnit *string `json:"server_rack_position_lower_unit" yaml:"rackPositionLowerUnit"`
	ServerRackPositionUpperUnit *string `json:"server_rack_position_upper_unit" yaml:"rackPositionUpperUnit"`
}

type ServerEditInventory struct {
	ServerInventoryId *string `json:"server_inventory_id" yaml:"inventoryId"`
}

// description: Record used to register a server
type ServerCreateAndRegister struct {
	//description: Datacenter to use
	DatacenterName string `json:"datacenter_name" yaml:"datacenter"`
	//description: Vendor. Must be one of the supported values.
	//values:
	// - HP
	// - HPE
	// - Dell Inc.
	// - Lenovo
	// - Supermicro
	// - BULL
	// - QEMU
	// - BSI
	ServerVendor string `json:"server_vendor" yaml:"vendor"`
	//description: IP of the BMC of the server
	ServerManagementAddress string `json:"server_ipmi_host" yaml:"address"`
	//description: Username of the BMC
	ServerManagementUser string `json:"server_ipmi_user" yaml:"user"`
	//description: Password of the BMC
	ServerManagementPassword string `json:"server_ipmi_password" yaml:"pass"`
}

type ServerCreateUnmanaged struct {
	DatacenterName           string                  `json:"datacenter_name" yaml:"datacenter"`
	ServerSerialNumber       string                  `json:"server_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	ServerManagementAddress  string                  `json:"server_ipmi_host,omitempty" yaml:"address,omitempty"`
	ServerManagementUser     string                  `json:"server_ipmi_user,omitempty" yaml:"user,omitempty"`
	ServerManagementPassword string                  `json:"server_ipmi_password,omitempty" yaml:"pass,omitempty"`
	ServerInterfaces         []ServerInterfaceCreate `json:"server_interfaces,omitempty" yaml:"interfaces,omitempty"`
	ServerTypeID             int                     `json:"server_type_id,omitempty" yaml:"serverTypeID,omitempty"`
}

type ServerInterfaceCreate struct {
	NetworkEquipmentIdentifierString          string `json:"network_equipment_identifier_string,omitempty" yaml:"switch,omitempty"`
	ServerInterfaceMACAddress                 string `json:"server_interface_mac_address,omitempty" yaml:"mac,omitempty"`
	NetworkEquipmentInterfaceIdentifierString string `json:"network_equipment_interface_identifier_string,omitempty" yaml:"switchInterface,omitempty"`
}

type ServerDefaultCredentials struct {
	ServerDefaultCredentialsID       int    `json:"server_default_credentials_id,omitempty" yaml:"id,omitempty"`
	DatacenterName                   string `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	ServerSerialNumber               string `json:"server_serial_number,omitempty" yaml:"serialNumber,omitempty"`
	ServerBMCMACAddress              string `json:"server_bmc_mac_address" yaml:"BMCMACAddress"`
	ServerDefaultCredentialsUsername string `json:"server_default_credentials_username" yaml:"username"`
	ServerDefaultCredentialsPassword string `json:"server_default_credentials_password" yaml:"password"`
}

// ServersSearch searches for servers matching certain filter
func (c *Client) ServersSearch(filter string) (*[]ServerSearchResult, error) {

	tables := []string{"_servers_instances"}
	columns := map[string][]string{
		"_servers_instances": {
			"server_id",
			"server_type_name",
			"server_type_boot_type",
			"server_product_name",
			"datacenter_name",
			"server_status",
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
			"server_inventory_id",
			"server_rack_name",
			"server_rack_position_lower_unit",
			"server_rack_position_upper_unit",
			"server_ipmi_host",
			"server_custom_json",
			"server_ipmi_internal_username",
			"server_ipmi_internal_password",
			"server_processor_name",
			"server_processor_count",
			"server_processor_core_count",
			"server_processor_core_mhz",
			"server_processor_threads",
			"server_processor_cpu_mark",
			"server_supports_oob_provisioning",
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
			"server_submodel",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	var createdObject map[string]SearchResultForServers

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
		createdObject = map[string]SearchResultForServers{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	servers := []ServerSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		servers = append(servers, s)
	}

	return &servers, nil
}

// ServerGetByUUID retrieves information about a specified Server by using the server's UUID
func (c *Client) ServerGetByUUID(serverUUID string, decryptPasswd bool) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_with_uuid_get",
		serverUUID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.ServerIPMInternalPassword, ":")

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
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
			}
		}
	} else {
		createdObject.ServerIPMInternalPassword = ""
	}

	return &createdObject, nil

}

// ServerGet returns a server's details
func (c *Client) ServerGet(serverID int, decryptPasswd bool) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_get",
		serverID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		var internalSrvObject Server
		err := c.rpcClient.CallFor(
			&internalSrvObject,
			"server_get_internal",
			serverID)

		if err != nil {
			return nil, err
		}

		passwdComponents := strings.Split(internalSrvObject.ServerIPMInternalPassword, ":")

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
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
			}
		}
	} else {
		createdObject.ServerIPMInternalPassword = ""
	}

	return &createdObject, nil
}

// ServerCreate manually creates a server record. DEPRECATED
func (c *Client) ServerCreate(server Server, autoGenerate bool) (int, error) {

	var createdObject int

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_create",
		server,
		autoGenerate,
	)

	if err != nil {
		return 0, err
	}

	return createdObject, nil
}

// ServerUnmanagedImport creates an unmanaged server
func (c *Client) ServerUnmanagedImport(server ServerCreateUnmanaged) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_unmanaged_import",
		[]ServerCreateUnmanaged{server},
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// ServerUnmanagedImportBatch Imports multiple unmanaged servers
func (c *Client) ServerUnmanagedImportBatch(servers []ServerCreateUnmanaged) (*map[string]Server, error) {

	var createdObject map[string]Server

	resp, err := c.rpcClient.Call(
		"server_unmanaged_import_batch",
		[][]ServerCreateUnmanaged{servers},
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]Server{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	return &createdObject, nil
}

// ServerCreateAndRegister manually creates and registers a server
func (c *Client) ServerCreateAndRegister(serverCreateAndRegister ServerCreateAndRegister) (int, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_create_and_register",
		serverCreateAndRegister,
		"server",
	)

	if err != nil {
		return 0, err
	}

	return createdObject.ServerID, nil
}

// ServerEditComplete - perform a complete edit
func (c *Client) ServerEditComplete(serverID int, server Server) (*Server, error) {
	return c.ServerEdit(serverID, "complete", server)
}

// ServerEditIPMI - edit only IPMI settings
func (c *Client) ServerEditIPMI(serverID int, server Server, serverUpdateInBMC bool) (*Server, error) {
	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_update_ipmi_credentials",
		serverID,
		server.ServerIPMIHost,
		server.ServerIPMInternalUsername,
		server.ServerIPMInternalPassword,
		serverUpdateInBMC,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// ServerEditAvailability - edit only server availability settings
func (c *Client) ServerEditAvailability(serverID int, server Server) (*Server, error) {
	return c.ServerEdit(serverID, "availability", server)
}

// ServerEdit edits a server record
func (c *Client) ServerEdit(serverID int, serverEditType string, server Server) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_edit",
		serverID,
		server,
		serverEditType,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// ServerEditProperty edits a specific property from the server record
// this is used instead of the server edit function as it does not require a marshal-unmarshal to and form the server object
// which sometimes is broken due to frequent changes on the server side
func (c *Client) ServerEditProperty(serverID int, serverPropertyToEdit string, serverPropertyValue interface{}) error {

	var createdObject map[string]interface{}

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_get",
		serverID)

	if err != nil {
		return err
	}

	createdObject[serverPropertyToEdit] = serverPropertyValue

	err = c.rpcClient.CallFor(
		&createdObject,
		"server_edit",
		serverID,
		createdObject,
		"complete",
	)

	if err != nil {
		return err
	}

	return nil
}

// ServerDelete deletes all the information about a specified Server.
func (c *Client) ServerDelete(serverID int, skipIPMI bool) error {

	resp, err := c.rpcClient.Call("server_delete", serverID, skipIPMI)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerDecomission decomissions the server row and deletes all child rows.
func (c *Client) ServerDecomission(serverID int, skipIPMI bool) error {

	resp, err := c.rpcClient.Call("server_decomission", serverID, skipIPMI)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerFirmwareComponentUpgrade Creates a firmware upgrading session for the specified component.
// If no strServerComponentFirmwareNewVersion or strFirmwareBinaryURL are provided the system will use the values from the database which should have been previously added
func (c *Client) ServerFirmwareComponentUpgrade(serverID int, serverComponentID int, serverComponentFirmwareNewVersion string, firmwareBinaryURL string) error {

	resp, err := c.rpcClient.Call(
		"server_firmware_component_upgrade",
		serverID,
		serverComponentID,
		serverComponentFirmwareNewVersion,
		firmwareBinaryURL,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerFirmwareUpgrade creates a firmware upgrading session that affects all components from the specified server that have a target version set and are updatable.
func (c *Client) ServerFirmwareUpgrade(serverID int) error {

	resp, err := c.rpcClient.Call(
		"server_firmware_upgrade",
		serverID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerFirmwareComponentTargetVersionSet Sets a firmware target version for the upgrading process. The system will apply the upgrade at the next upgrading session.
func (c *Client) ServerFirmwareComponentTargetVersionSet(serverComponentID int, serverComponentFirmwareNewVersion string) error {

	resp, err := c.rpcClient.Call(
		"server_firmware_component_target_version_set",
		serverComponentID,
		serverComponentFirmwareNewVersion,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerFirmwareComponentTargetVersionUpdate Updates for every component of the specified server the available firmware versions that can be used as target by the firmware upgrading process. The available versions are extracted from a vendor specific catalog.
func (c *Client) ServerFirmwareComponentTargetVersionUpdate(serverComponentID int) error {

	resp, err := c.rpcClient.Call(
		"server_firmware_component_available_versions_update",
		serverComponentID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}
	return nil
}

// ServerFirmwareComponentTargetVersionAdd Adds a new available firmware version for a server component along with the url of the binary. If the version already exists the old url will be overwritten.
func (c *Client) ServerFirmwareComponentTargetVersionAdd(serverComponentID int, version string, firmareBinaryURL string) error {

	resp, err := c.rpcClient.Call(
		"server_firmware_component_available_versions_add",
		serverComponentID,
		version,
		firmareBinaryURL,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}
	return nil
}

// ServerComponentGet returns a server's component's details
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

// ServerComponents searches for servers matching certain filter
func (c *Client) ServerComponents(serverID int, filter string) (*[]ServerComponent, error) {

	tables := []string{"_server_components"}
	columns := map[string][]string{
		"_server_components": {
			"server_component_name",
			"server_component_id",
			"server_component_firmware_json",
			"server_component_type",
			"server_component_firmware_update_timestamp",
			"server_component_firmware_status",
			"server_component_firmware_update_available_versions",
			"server_component_firmware_updateable",
			"server_component_firmware_version",
			"server_component_firmware_status",
		},
	}

	userID := c.GetUserID()

	collapseType := "none"
	filterWithServerID := fmt.Sprintf("+server_id:%d %s", serverID, filter)
	res, err := c.rpcClient.Call(
		"search",
		userID,
		filterWithServerID,
		tables,
		columns,
		collapseType)

	if err != nil {
		return nil, err
	}

	var createdObject map[string]SearchResultForServerComponents

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	list := []ServerComponent{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil
}

// ServerPowerSet reboots or powers on a server
func (c *Client) ServerPowerSet(serverID int, operation string) error {
	if err := checkID(serverID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"server_power_set",
		serverID,
		operation)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerReregister triggers a re-register of a server
func (c *Client) ServerReregister(serverID int, bSkipIPMI bool, bUseBDKAgent bool) error {
	if err := checkID(serverID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"server_reregister",
		serverID,
		bSkipIPMI,
		bUseBDKAgent)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// ServerStatusUpdate alters the status of a server
func (c *Client) ServerStatusUpdate(serverID int, status string) error {
	if err := checkID(serverID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"server_status_update",
		serverID,
		status)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// CreateOrUpdate implements interface Applier
func (s Server) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var result *Server
	err = s.Validate()

	if err != nil {
		return err
	}

	if s.ServerID != 0 {
		result, err = client.ServerGet(s.ServerID, false)
	} else {
		result, err = client.ServerGetByUUID(s.ServerUUID, false)
	}

	if err != nil {
		_, err = client.ServerCreate(s, false)

		if err != nil {
			return err
		}
	} else {
		_, err = client.ServerEditComplete(result.ServerID, s)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (s Server) Delete(client MetalCloudClient) error {
	var result *Server
	var id int
	err := s.Validate()

	if err != nil {
		return err
	}

	if s.ServerID != 0 {
		id = s.ServerID
	} else {
		result, err = client.ServerGetByUUID(s.ServerUUID, false)
		if err != nil {
			return err
		}
		id = result.ServerID
	}

	err = client.ServerDelete(id, true)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (s Server) Validate() error {
	if s.ServerUUID == "" && s.ServerID == 0 {
		return fmt.Errorf("id is required")
	}
	return nil
}

// ServerEditRack returns a server's rack info details
func (c *Client) ServerEditRack(serverID int, serverEditRack ServerEditRack) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_edit_rack",
		serverID,
		serverEditRack,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// ServerEditInventory returns a server's inventory details
func (c *Client) ServerEditInventory(serverID int, serverEditInventory ServerEditInventory) (*Server, error) {

	var createdObject Server

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_edit_inventory",
		serverID,
		serverEditInventory,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// InstanceServerReplace replaces a server associated to an instance. Returns an AFC Group ID to be used in the AFC Deploy Viewer.
func (c *Client) InstanceServerReplace(instanceID int, serverID int) (int, error) {

	var createdObject int

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_server_replace",
		instanceID,
		serverID,
	)

	if err != nil {
		return 0, err
	}

	return createdObject, nil
}

// ServerDefaultCredentialsAdd Adds BMC credentials to the default credentials list for the ZTP process
func (c *Client) ServerDefaultCredentialsAdd(credentials []ServerDefaultCredentials) error {

	var createdObject int
	err := c.rpcClient.CallFor(
		&createdObject,
		"server_default_credentials_add",
		[][]ServerDefaultCredentials{credentials},
	)

	if err != nil {
		return err
	}

	return nil
}

// ServerDefaultCredentials retrieves the default credentials for server BMCs for the ZTP process for a given datacenter
func (c *Client) ServerDefaultCredentials(datacenter_name string, decryptPasswd bool) (*[]ServerDefaultCredentials, error) {

	var createdObjects []ServerDefaultCredentials

	err := c.rpcClient.CallFor(
		&createdObjects,
		"server_default_credentials",
		datacenter_name)

	if err != nil {
		return nil, err
	}

	for i, createdObject := range createdObjects {

		if decryptPasswd {

			passwdComponents := strings.Split(createdObject.ServerDefaultCredentialsPassword, ":")

			if len(passwdComponents) == 2 {
				if strings.Contains(passwdComponents[0], "Not authorized") {
					return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
				} else {
					var passwd string

					err = c.rpcClient.CallFor(
						&passwd,
						"password_decrypt",
						passwdComponents[1],
					)
					if err != nil {
						return nil, err
					}

					createdObjects[i].ServerDefaultCredentialsPassword = passwd
				}
			}
		} else {
			createdObject.ServerDefaultCredentialsPassword = ""
		}
	}
	return &createdObjects, nil
}

// ServerDefaultCredentialsRemove Removes BMC credentials to the default credentials list for the ZTP process
func (c *Client) ServerDefaultCredentialsRemove(default_credentials_id []int) error {

	var createdObject int
	err := c.rpcClient.CallFor(
		&createdObject,
		"server_default_credentials_remove",
		[][]int{default_credentials_id},
	)

	if err != nil {
		return err
	}

	return nil
}
