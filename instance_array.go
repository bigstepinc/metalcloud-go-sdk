package metalcloud

//go:generate go run helper/gen_exports.go
//go:generate go run helper/docgen.go - $GOFILE ./ InstanceArray,InstanceArrayOperation,InstanceArrayInterface,InstanceArrayInterfaceOperation InstanceArray

import "fmt"

// InstanceArray object describes a collection of identical instances
type InstanceArray struct {
	// description: The ID of the object.
	InstanceArrayID int `json:"instance_array_id,omitempty" yaml:"instanceID,omitempty"`
	// description: The label of the object. Must be unique per infrastructure. Must follow DNS naming rules: Pattern: ^[a-zA-Z]{1,1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1,1}|[a-zA-Z]{1,1}$
	InstanceArrayLabel string `json:"instance_array_label,omitempty" yaml:"label,omitempty"`
	// description: |
	//		User editable DNS record that gets created for this instance array in the built-in DNS
	//		service and associated with all the primary IP address on the WAN network. Must adhere to DNS naming rules such
	//      as:  only "-", lowercase alphanumeric characters and not start with a number.
	//      Pattern:
	InstanceArraySubdomain string `json:"instance_array_subdomain,omitempty" yaml:"subdomain,omitempty"`
	// description: The booth method to use. Note that the pxe_iscsi booth method is deprecated.
	//values:
	//	- local_disks
	//  - pxe_iscsi
	InstanceArrayBootMethod string `json:"instance_array_boot_method,omitempty" yaml:"bootMethod,omitempty"`
	// description: The number of instances in this array.
	InstanceArrayInstanceCount int `json:"instance_array_instance_count" yaml:"instanceCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of RAM expressed in GB.
	InstanceArrayRAMGbytes int `json:"instance_array_ram_gbytes,omitempty" yaml:"ramGBytes,omitempty"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of CPU sockets.
	InstanceArrayProcessorCount int `json:"instance_array_processor_count" yaml:"processorCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum GPU frequency.
	InstanceArrayProcessorCoreMHZ int `json:"instance_array_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum core count (hyperthreaded).
	InstanceArrayProcessorCoreCount int `json:"instance_array_processor_core_count" yaml:"processorCoreCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk count.
	InstanceArrayDiskCount int `json:"instance_array_disk_count" yaml:"diskCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk size.
	InstanceArrayDiskSizeMBytes int `json:"instance_array_disk_size_mbytes,omitempty" yaml:"diskSizeMBytes,omitempty"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this disk type. This assumes all disks are identical.
	// values:
	// - ssh
	// - hdd
	InstanceArrayDiskTypes []string `json:"instance_array_disk_types,omitempty" yaml:"diskTypes,omitempty"`
	// description: The id of the infrastructure on which this infrastructure is created on.
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	// description: The service status.  Read only.
	// values:
	//  - ordered #designed but not yet deployed
	//  - active #deployed
	//  - suspended #deployed but suspended in a way that cannot be resumed by the user. (typically this means shutdown and disconnected from the internet).
	//  - stopped #deployed but suspended in a way that can be resumed by the user.
	//  - deleted #deleted permanently
	InstanceArrayServiceStatus string `json:"instance_array_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	// description: The instance array interfaces configuration
	InstanceArrayInterfaces []InstanceArrayInterface `json:"instance_array_interfaces,omitempty" yaml:"interfaces,omitempty"`
	// description: The cluster (such as Kubernetes, VMWare vSphere etc) of which this instance array is part of. A vanilla cluster is created for all instance arrays not added to any application cluster.
	ClusterID int `json:"cluster_id,omitempty" yaml:"clusterID,omitempty"`
	// description: If part of an app cluster this field will receive the role that this instance array has such as `master` or `worker` which is application specific.
	ClusterRoleGroup string `json:"cluster_role_group,omitempty" yaml:"clusterRoleGroup,omitempty"`
	// description: If set to true, the firewall will be configured based on rules provided in the InstanceArrayFirewallRules field. Note that for this to work the following conditions must be fufilled:
	// a. The OS template should have the template set capability (for the first configuration of the firewall, at install time)
	// b. The in-band site controller agent should be enabled with in-band access to the operating system over SSH.
	InstanceArrayFirewallManaged bool `json:"instance_array_firewall_managed" yaml:"firewallManaged"`
	// description: The list of firewall rules to configure
	InstanceArrayFirewallRules []FirewallRule `json:"instance_array_firewall_rules,omitempty" yaml:"firewallRules,omitempty"`
	// description: The operating system template to use.
	VolumeTemplateID int `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	// description: Used when changing an instance array. It captures the operation that needs to happen on the instance array.
	InstanceArrayOperation *InstanceArrayOperation `json:"instance_array_operation,omitempty" yaml:"operation,omitempty"`
	// description: Information about additional ips to be assigned to the WAN interfaces. Used internally.
	InstanceArrayAdditionalWanIPv4JSON string `json:"instance_array_additional_wan_ipv4_json,omitempty" yaml:"additionalWanIPv4,omitempty"`
	// description: Custom variables and variable overrides to be pushed to the operating system deployment process.
	InstanceArrayCustomVariables interface{} `json:"instance_array_custom_variables,omitempty" yaml:"customVariables,omitempty"`
	// description: Firmware policies to apply. Deprecated. Use baselines.
	InstanceArrayFirmwarePolicies []int `json:"instance_array_firmware_policies,omitempty" yaml:"firmwarePolicies,omitempty"`
	// description: When iSCSI boot is used this is the id of the drive array that will be the boot device.
	DriveArrayIDBoot int `json:"drive_array_id_boot,omitempty" yaml:"drive_array_id_boot,omitempty"`
}

// InstanceArrayOperation object describes the changes that will be applied to an instance array
type InstanceArrayOperation struct {
	// description: The ID of the object.
	InstanceArrayID int `json:"instance_array_id,omitempty" yaml:"id,omitempty"`
	// description: The label of the object. Must be unique per infrastructure. Must follow DNS naming rules: Pattern: ^[a-zA-Z]{1,1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1,1}|[a-zA-Z]{1,1}$
	InstanceArrayLabel string `json:"instance_array_label,omitempty" yaml:"label,omitempty"`
	// description: |
	//		User editable DNS record that gets created for this instance array in the built-in DNS
	//		service and associated with all the primary IP address on the WAN network. Must adhere to DNS naming rules such
	//      as:  only "-", lowercase alphanumeric characters and not start with a number.
	//      Pattern:
	InstanceArraySubdomain string `json:"instance_array_subdomain,omitempty" yaml:"subdomain,omitempty"`
	// description: The booth method to use. Note that the pxe_iscsi booth method is deprecated.
	//values:
	//	- local_disks
	//  - pxe_iscsi
	InstanceArrayBootMethod string `json:"instance_array_boot_method,omitempty" yaml:"bootMethod,omitempty"`
	// description: The number of instances in this array.
	InstanceArrayInstanceCount int `json:"instance_array_instance_count" yaml:"instanceCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of RAM expressed in GB.
	InstanceArrayRAMGbytes int `json:"instance_array_ram_gbytes,omitempty" yaml:"ramGBytes,omitempty"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of CPU sockets.
	InstanceArrayProcessorCount int `json:"instance_array_processor_count" yaml:"processorCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum GPU frequency.
	InstanceArrayProcessorCoreMHZ int `json:"instance_array_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum core count (hyperthreaded).
	InstanceArrayProcessorCoreCount int `json:"instance_array_processor_core_count" yaml:"processorCoreCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk count.
	InstanceArrayDiskCount int `json:"instance_array_disk_count" yaml:"diskCount"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk size.
	InstanceArrayDiskSizeMBytes int `json:"instance_array_disk_size_mbytes,omitempty" yaml:"diskSizeMBytes,omitempty"`
	// description: When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this disk type. This assumes all disks are identical.
	// values:
	// - ssh
	// - hdd
	InstanceArrayDiskTypes []string `json:"instance_array_disk_types,omitempty" yaml:"diskTypes,omitempty"`
	// description: The service status.  Read only.
	// values:
	//  - ordered #designed but not yet deployed
	//  - active #deployed
	//  - suspended #deployed but suspended in a way that cannot be resumed by the user. (typically this means shutdown and disconnected from the internet).
	//  - stopped #deployed but suspended in a way that can be resumed by the user.
	//  - deleted #deleted permanently
	InstanceArrayServiceStatus string `json:"instance_array_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	// description: The instance array interfaces configuration
	InstanceArrayInterfaces []InstanceArrayInterfaceOperation `json:"instance_array_interfaces,omitempty" yaml:"interfaces,omitempty"`
	// description: The cluster (such as Kubernetes, VMWare vSphere etc) of which this instance array is part of. A vanilla cluster is created for all instance arrays not added to any application cluster.
	ClusterID int `json:"cluster_id,omitempty" yaml:"clusterID,omitempty"`
	// description: If part of an app cluster this field will receive the role that this instance array has such as master or worker which is application specific.
	ClusterRoleGroup string `json:"cluster_role_group,omitempty" yaml:"clusterRoleGroup,omitempty"`
	// description: If set to true, the firewall will be configured based on rules provided in the InstanceArrayFirewallRules field. Note that for this to work the following conditions must be fufilled:
	// a. The OS template should have the template set capability (for the first configuration of the firewall, at install time)
	// b. The in-band site controller agent should be enabled with in-band access to the operating system over SSH.
	InstanceArrayFirewallManaged bool `json:"instance_array_firewall_managed" yaml:"firewallManaged"`
	// description: The list of firewall rules to configure
	InstanceArrayFirewallRules []FirewallRule `json:"instance_array_firewall_rules,omitempty" yaml:"firewallRules,omitempty"`
	// description: The operating system template to use.
	VolumeTemplateID int `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	// description: The deploy type, one of:
	// values:
	//     - create
	// 	   - delete
	//     - edit
	// 	   - start
	//     - stop
	// 	   - suspend
	InstanceArrayDeployType string `json:"instance_array_deploy_type,omitempty" yaml:"deployType,omitempty"`
	// description: The status of the deployment
	// values:
	//     - not_started
	//     - ongoing
	//     - finished
	InstanceArrayDeployStatus string `json:"instance_array_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	// description: The id of the change operation. Readonly.
	InstanceArrayChangeID int `json:"instance_array_change_id,omitempty" yaml:"changeID,omitempty"`
	// description: Information about additional ips to be assigned to the WAN interfaces. Used internally.
	InstanceArrayAdditionalWanIPv4JSON string `json:"instance_array_additional_wan_ipv4_json,omitempty" yaml:"additionalWanIPv4,omitempty"`
	// description: Custom variables and variable overrides to be pushed to the operating system deployment process.
	InstanceArrayCustomVariables interface{} `json:"instance_array_custom_variables,omitempty" yaml:"customVariables,omitempty"`
	// description: Firmware policies to apply. Deprecated. Use baselines instead.
	InstanceArrayFirmwarePolicies []int `json:"instance_array_firmware_policies" yaml:"firmwarePolicies"`
	// description: When iSCSI boot is used this is the id of the drive array that will be the boot device.
	DriveArrayIDBoot int `json:"drive_array_id_boot,omitempty" yaml:"drive_array_id_boot,omitempty"`
}

// FirewallRule describes a firewall rule that is to be applied on all instances of an array
type FirewallRule struct {
	//description: Description of the rule
	FirewallRuleDescription string `json:"firewall_rule_description,omitempty" yaml:"description,omitempty"`
	//description: Port range start
	//example: 22
	FirewallRulePortRangeStart int `json:"firewall_rule_port_range_start,omitempty" yaml:"portRangeStart,omitempty"`
	//description: Port range end
	//example: 22
	FirewallRulePortRangeEnd int `json:"firewall_rule_port_range_end,omitempty" yaml:"portRangeEnd,omitempty"`
	//description: Source Ip Address Range Start
	//example: 192.168.0.1
	FirewallRuleSourceIPAddressRangeStart string `json:"firewall_rule_source_ip_address_range_start,omitempty" yaml:"sourceIPAddressRangeStart,omitempty"`
	//description: Source Ip Address Range End
	//example: 192.168.0.100
	FirewallRuleSourceIPAddressRangeEnd string `json:"firewall_rule_source_ip_address_range_end,omitempty" yaml:"sourceIPAddressRangeEnd,omitempty"`
	//description: Source Ip Address Range Start
	//example: 192.168.0.1
	FirewallRuleDestinationIPAddressRangeStart string `json:"firewall_rule_destination_ip_address_range_start,omitempty" yaml:"destinationIPAddressRangeStart,omitempty"`
	//description: Source Ip Address Range End
	//example: 192.168.0.100
	FirewallRuleDestinationIPAddressRangeEnd string `json:"firewall_rule_destination_ip_address_range_end,omitempty" yaml:"destinationIPAddressRangeEnd,omitempty"`
	//description: Source Ip Address Range Start
	//example: tcp
	FirewallRuleProtocol string `json:"firewall_rule_protocol,omitempty" yaml:"protocol,omitempty"`
	//description: IP address type
	//values:
	//	- ipv4
	//	- ipv6
	FirewallRuleIPAddressType string `json:"firewall_rule_ip_address_type,omitempty" yaml:"IPAddressType,omitempty"`
	//description: Set to true if the rule is enabled
	//values:
	//	- true
	//	- false
	FirewallRuleEnabled bool `json:"firewall_rule_enabled" yaml:"enabled"`
}

// InstanceArrayInterface describes a network interface of the array.
// It's properties will be applied to all InstanceInterfaces of the array's instances.
type InstanceArrayInterface struct {
	//description: Label of the interface
	InstanceArrayInterfaceLabel string `json:"instance_array_interface_label,omitempty" yaml:"label,omitempty"`
	//description: An unique string describing the interface.
	InstanceArrayInterfaceSubdomain string `json:"instance_array_interface_subdomain,omitempty" yaml:"subdomain,omitempty"`
	//description: Interface ID. A unique id of the interface.
	//example: 234
	InstanceArrayInterfaceID int `json:"instance_array_interface_id,omitempty" yaml:"id,omitempty"`
	//description: The instance array to which this interface is associated to
	InstanceArrayID int `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	//description: The ID of the Network to which this interface is connected to. Can be 0 if the interface is not connected.
	NetworkID int `json:"network_id,omitempty" yaml:"networkID,omitempty"`
	//description: Used internally. Readonly.
	InstanceArrayInterfaceLAGGIndexes []interface{} `json:"instance_array_interface_lagg_indexes,omitempty" yaml:"LAGGIndexes,omitempty"`
	//description: |
	//	The index of the interface in the server. This is 0-based and configured based on the lexicographic sorting of the switch and switch ports
	//	thus NOT based on the PCI slots or anything like that. This ensures consistent ordering regardless of cabling and/or NIC seating.
	InstanceArrayInterfaceIndex int `json:"instance_array_interface_index,omitempty" yaml:"index,omitempty"`
	// description: The service status.  Read only.
	// values:
	//  - ordered #designed but not yet deployed
	//  - active #deployed
	//  - suspended #deployed but suspended in a way that cannot be resumed by the user. (typically this means shutdown and disconnected from the internet).
	//  - stopped #deployed but suspended in a way that can be resumed by the user.
	//  - deleted #deleted permanently
	InstanceArrayInterfaceServiceStatus string `json:"instance_array_interface_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	// description: The creation date and time in ISO 8601 format.
	// example: 2013-11-29T13:00:01Z
	InstanceArrayInterfaceCreatedTimestamp string `json:"instance_array_interface_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	// description: The last update date and time in ISO 8601 format.
	// example: 2013-11-29T13:00:01Z
	InstanceArrayInterfaceUpdatedTimestamp string `json:"instance_array_interface_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	// description: The operation object. Must be set to alter the configuration of an interface.
	InstanceArrayInterfaceOperation *InstanceArrayInterfaceOperation `json:"instance_array_interface_operation,omitempty" yaml:"operation,omitempty"`
	// description: Ongoing change ID. Readonly.
	InstanceArrayInterfaceChangeID int `json:"instance_array_interface_change_id,omitempty" yaml:"instance_array_interface_change_id,omitempty"`
}

// InstanceArrayInterfaceOperation describes changes to a network array interface
type InstanceArrayInterfaceOperation struct {
	//description: Label of the interface
	InstanceArrayInterfaceLabel string `json:"instance_array_interface_label,omitempty" yaml:"label,omitempty"`
	//description: An unique string describing the interface.
	InstanceArrayInterfaceSubdomain string `json:"instance_array_interface_subdomain,omitempty" yaml:"subdomain,omitempty"`
	//description: Interface ID. A unique id of the interface.
	//example: 234
	InstanceArrayInterfaceID int `json:"instance_array_interface_id,omitempty" yaml:"id,omitempty"`
	//description: The instance array to which this interface is associated to
	InstanceArrayID int `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	//description: The ID of the Network to which this interface is connected to. Can be 0 if the interface is not connected.
	NetworkID int `json:"network_id,omitempty" yaml:"networkID,omitempty"`
	//description: Used internally. Readonly.
	InstanceArrayInterfaceLAGGIndexes []interface{} `json:"instance_array_interface_lagg_indexes,omitempty" yaml:"LAGGIndexes,omitempty"`
	//description: |
	//	The index of the interface in the server. This is 0-based and configured based on the lexicographic sorting of the switch and switch ports
	//	thus NOT based on the PCI slots or anything like that. This ensures consistent ordering regardless of cabling and/or NIC seating.
	InstanceArrayInterfaceIndex int `json:"instance_array_interface_index,omitempty" yaml:"index,omitempty"`
	// description: The service status.  Read only.
	// values:
	//  - ordered #designed but not yet deployed
	//  - active #deployed
	//  - suspended #deployed but suspended in a way that cannot be resumed by the user. (typically this means shutdown and disconnected from the internet).
	//  - stopped #deployed but suspended in a way that can be resumed by the user.
	//  - deleted #deleted permanently
	InstanceArrayInterfaceServiceStatus string `json:"instance_array_interface_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	// description: The creation date and time in ISO 8601 format.
	// example: 2013-11-29T13:00:01Z
	InstanceArrayInterfaceCreatedTimestamp string `json:"instance_array_interface_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	// description: The last update date and time in ISO 8601 format.
	// example: 2013-11-29T13:00:01Z
	InstanceArrayInterfaceUpdatedTimestamp string `json:"instance_array_interface_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	// description: Ongoing change ID. Readonly.
	InstanceArrayInterfaceChangeID int `json:"instance_array_interface_change_id,omitempty" yaml:"changeID,omitempty"`
}

// ServerTypeMatches used in InstanceArrayEdit operations to specify which server types to use
type ServerTypeMatches struct {
	ServerTypes map[int]ServerTypeMatch `json:"server_types,omitempty"`
}

// ServerTypeMatch what exact server types to use
type ServerTypeMatch struct {
	ServerCount int `json:"server_count,omitempty"`
}

// instanceArrayGet returns an InstanceArray with specified id
func (c *Client) instanceArrayGet(instanceArrayID id) (*InstanceArray, error) {
	var createdObject InstanceArray

	if err := checkID(instanceArrayID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_get",
		instanceArrayID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// instanceArrays returns list of instance arrays of specified infrastructure
func (c *Client) instanceArrays(infrastructureID id) (*map[string]InstanceArray, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"instance_arrays",
		infrastructureID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]InstanceArray{}
		return &m, nil
	}

	var createdObject map[string]InstanceArray

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// instanceArrayCreate creates an instance array (collection of identical instances). Requires Deploy.
func (c *Client) instanceArrayCreate(infrastructureID id, instanceArray InstanceArray) (*InstanceArray, error) {
	var createdObject InstanceArray

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_create",
		infrastructureID,
		instanceArray)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// instanceArrayEdit edits a deployed instance array. Requires deploy.
func (c *Client) instanceArrayEdit(instanceArrayID id, instanceArrayOperation InstanceArrayOperation, bSwapExistingInstancesHardware *bool, bKeepDetachingDrives *bool, objServerTypeMatches *ServerTypeMatches, arrInstancesToBeDeleted *[]int) (*InstanceArray, error) {
	var createdObject InstanceArray

	if err := checkID(instanceArrayID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_edit",
		instanceArrayID,
		instanceArrayOperation,
		bSwapExistingInstancesHardware,
		bKeepDetachingDrives,
		objServerTypeMatches,
		arrInstancesToBeDeleted)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// instanceArrayDelete deletes an instance array. Requires deploy.
func (c *Client) instanceArrayDelete(instanceArrayID id) error {

	if err := checkID(instanceArrayID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"instance_array_delete",
		instanceArrayID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// InstanceArrayInterfaceAttachNetwork attaches an InstanceArrayInterface to a Network
func (c *Client) InstanceArrayInterfaceAttachNetwork(instanceArrayID int, instanceArrayInterfaceIndex int, networkID int) (*InstanceArray, error) {
	var createdObject InstanceArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_interface_attach_network",
		instanceArrayID,
		instanceArrayInterfaceIndex,
		networkID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// InstanceArrayInterfaceDetach detaches an InstanceArrayInterface from any Network element that is attached to.
func (c *Client) InstanceArrayInterfaceDetach(instanceArrayID int, instanceArrayInterfaceIndex int) (*InstanceArray, error) {
	var createdObject InstanceArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_interface_detach",
		instanceArrayID,
		instanceArrayInterfaceIndex)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// instanceArrayStop stops a specified InstanceArray.
func (c *Client) instanceArrayStop(instanceArrayID id) (*InstanceArray, error) {

	var createdObject InstanceArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_stop",
		instanceArrayID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// instanceArrayStart starts a specified InstanceArray.
func (c *Client) instanceArrayStart(instanceArrayID id) (*InstanceArray, error) {

	var createdObject InstanceArray

	err := c.rpcClient.CallFor(
		&createdObject,
		"instance_array_start",
		instanceArrayID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

func (ia *InstanceArray) instanceToOperation(op *InstanceArrayOperation) {
	operation := ia.InstanceArrayOperation
	operation.InstanceArrayID = ia.InstanceArrayID
	operation.InstanceArrayLabel = ia.InstanceArrayLabel
	operation.InstanceArraySubdomain = ia.InstanceArraySubdomain
	operation.InstanceArrayBootMethod = ia.InstanceArrayBootMethod
	operation.InstanceArrayInstanceCount = ia.InstanceArrayInstanceCount
	operation.InstanceArrayRAMGbytes = ia.InstanceArrayRAMGbytes
	operation.InstanceArrayProcessorCount = ia.InstanceArrayProcessorCount
	operation.InstanceArrayProcessorCoreMHZ = ia.InstanceArrayProcessorCoreMHZ
	operation.InstanceArrayProcessorCoreCount = ia.InstanceArrayProcessorCoreCount
	operation.InstanceArrayDiskCount = ia.InstanceArrayDiskCount
	operation.InstanceArrayDiskSizeMBytes = ia.InstanceArrayDiskSizeMBytes
	operation.InstanceArrayFirewallManaged = ia.InstanceArrayFirewallManaged
	operation.VolumeTemplateID = ia.VolumeTemplateID
	operation.InstanceArrayChangeID = op.InstanceArrayChangeID
}

// CreateOrUpdate implements interface Applier
func (ia InstanceArray) CreateOrUpdate(client MetalCloudClient) error {
	var result *InstanceArray
	var err error

	if ia.InstanceArrayID != 0 {
		result, err = client.InstanceArrayGet(ia.InstanceArrayID)
	} else {
		result, err = client.InstanceArrayGetByLabel(ia.InstanceArrayLabel)
	}

	if err != nil {
		_, err = client.InstanceArrayCreate(ia.InfrastructureID, ia)

		if err != nil {
			return err
		}
	} else {
		var bKeepDetachingDrives, bSwapExistingInstancesHardware bool = false, false

		ia.instanceToOperation(result.InstanceArrayOperation)
		_, err = client.InstanceArrayEditByLabel(result.InstanceArrayLabel, *ia.InstanceArrayOperation, &bSwapExistingInstancesHardware, &bKeepDetachingDrives, nil, nil)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (ia InstanceArray) Delete(client MetalCloudClient) error {
	err := ia.Validate()
	var result *InstanceArray
	var id int

	if err != nil {
		return err
	}

	if ia.InstanceArrayLabel != "" {
		result, err = client.InstanceArrayGetByLabel(ia.InstanceArrayLabel)

		if err != nil {
			return err
		}

		id = result.InstanceArrayID
	} else {
		id = ia.InstanceArrayID
	}
	err = client.InstanceArrayDelete(id)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (ia InstanceArray) Validate() error {
	if ia.InstanceArrayID == 0 && ia.InstanceArrayLabel == "" {
		return fmt.Errorf("id is required")
	}

	return nil
}
