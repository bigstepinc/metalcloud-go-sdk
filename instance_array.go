package metalcloud

//go:generate go run helper/gen_exports.go

import "fmt"

//InstanceArray object describes a collection of identical instances
type InstanceArray struct {
	InstanceArrayID                    int                      `json:"instance_array_id,omitempty" yaml:"instanceID,omitempty"`
	InstanceArrayLabel                 string                   `json:"instance_array_label,omitempty" yaml:"label,omitempty"`
	InstanceArraySubdomain             string                   `json:"instance_array_subdomain,omitempty" yaml:"subdomain,omitempty"`
	InstanceArrayBootMethod            string                   `json:"instance_array_boot_method,omitempty" yaml:"bootMethod,omitempty"`
	InstanceArrayInstanceCount         int                      `json:"instance_array_instance_count" yaml:"instanceCount"`
	InstanceArrayRAMGbytes             int                      `json:"instance_array_ram_gbytes,omitempty" yaml:"ramGBytes,omitempty"`
	InstanceArrayProcessorCount        int                      `json:"instance_array_processor_count" yaml:"processorCount"`
	InstanceArrayProcessorCoreMHZ      int                      `json:"instance_array_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	InstanceArrayProcessorCoreCount    int                      `json:"instance_array_processor_core_count" yaml:"processorCoreCount"`
	InstanceArrayDiskCount             int                      `json:"instance_array_disk_count" yaml:"diskCount"`
	InstanceArrayDiskSizeMBytes        int                      `json:"instance_array_disk_size_mbytes,omitempty" yaml:"diskSizeMBytes,omitempty"`
	InstanceArrayDiskTypes             []string                 `json:"instance_array_disk_types,omitempty" yaml:"diskTypes,omitempty"`
	InfrastructureID                   int                      `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	InstanceArrayServiceStatus         string                   `json:"instance_array_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	InstanceArrayInterfaces            []InstanceArrayInterface `json:"instance_array_interfaces,omitempty" yaml:"interfaces,omitempty"`
	ClusterID                          int                      `json:"cluster_id,omitempty" yaml:"clusterID,omitempty"`
	ClusterRoleGroup                   string                   `json:"cluster_role_group,omitempty" yaml:"clusterRoleGroup,omitempty"`
	InstanceArrayFirewallManaged       bool                     `json:"instance_array_firewall_managed" yaml:"firewallManaged,omitempty"`
	InstanceArrayFirewallRules         []FirewallRule           `json:"instance_array_firewall_rules,omitempty" yaml:"firewallRules,omitempty"`
	VolumeTemplateID                   int                      `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	InstanceArrayOperation             *InstanceArrayOperation  `json:"instance_array_operation,omitempty" yaml:"operation,omitempty"`
	InstanceArrayAdditionalWanIPv4JSON string                   `json:"instance_array_additional_wan_ipv4_json,omitempty" yaml:"additionalWanIPv4,omitempty"`
	InstanceArrayCustomVariables       interface{}              `json:"instance_array_custom_variables,omitempty" yaml:"customVariables,omitempty"`
	InstanceArrayFirmwarePolicies      []int                    `json:"instance_array_firmware_policies,omitempty" yaml:"firmwarePolicies,omitempty"`
}

//InstanceArrayOperation object describes the changes that will be applied to an instance array
type InstanceArrayOperation struct {
	InstanceArrayID                    int                               `json:"instance_array_id,omitempty" yaml:"id,omitempty"`
	InstanceArrayLabel                 string                            `json:"instance_array_label,omitempty" yaml:"label,omitempty"`
	InstanceArraySubdomain             string                            `json:"instance_array_subdomain,omitempty" yaml:"subdomain,omitempty"`
	InstanceArrayBootMethod            string                            `json:"instance_array_boot_method,omitempty" yaml:"bootMethod,omitempty"`
	InstanceArrayInstanceCount         int                               `json:"instance_array_instance_count" yaml:"instanceCount"`
	InstanceArrayRAMGbytes             int                               `json:"instance_array_ram_gbytes,omitempty" yaml:"ramGBytes,omitempty"`
	InstanceArrayProcessorCount        int                               `json:"instance_array_processor_count" yaml:"processorCount"`
	InstanceArrayProcessorCoreMHZ      int                               `json:"instance_array_processor_core_mhz,omitempty" yaml:"processorCoreMhz,omitempty"`
	InstanceArrayProcessorCoreCount    int                               `json:"instance_array_processor_core_count" yaml:"processorCoreCount"`
	InstanceArrayDiskCount             int                               `json:"instance_array_disk_count" yaml:"diskCount"`
	InstanceArrayDiskSizeMBytes        int                               `json:"instance_array_disk_size_mbytes,omitempty" yaml:"diskSizeMBytes,omitempty"`
	InstanceArrayDiskTypes             []string                          `json:"instance_array_disk_types,omitempty" yaml:"diskTypes,omitempty"`
	InstanceArrayServiceStatus         string                            `json:"instance_array_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	InstanceArrayInterfaces            []InstanceArrayInterfaceOperation `json:"instance_array_interfaces,omitempty" yaml:"interfaces,omitempty"`
	ClusterID                          int                               `json:"cluster_id,omitempty" yaml:"clusterID,omitempty"`
	ClusterRoleGroup                   string                            `json:"cluster_role_group,omitempty" yaml:"clusterRoleGroup,omitempty"`
	InstanceArrayFirewallManaged       bool                              `json:"instance_array_firewall_managed" yaml:"firewallManaged,omitempty"`
	InstanceArrayFirewallRules         []FirewallRule                    `json:"instance_array_firewall_rules,omitempty" yaml:"firewallRules,omitempty"`
	VolumeTemplateID                   int                               `json:"volume_template_id,omitempty" yaml:"volumeTemplateID,omitempty"`
	InstanceArrayDeployType            string                            `json:"instance_array_deploy_type,omitempty" yaml:"deployType,omitempty"`
	InstanceArrayDeployStatus          string                            `json:"instance_array_deploy_status,omitempty" yaml:"deployStatus,omitempty"`
	InstanceArrayChangeID              int                               `json:"instance_array_change_id,omitempty" yaml:"changeID,omitempty"`
	InstanceArrayAdditionalWanIPv4JSON string                            `json:"instance_array_additional_wan_ipv4_json,omitempty" yaml:"additionalWanIPv4,omitempty"`
	InstanceArrayCustomVariables       interface{}                       `json:"instance_array_custom_variables,omitempty" yaml:"customVariables,omitempty"`
	InstanceArrayFirmwarePolicies      []int                             `json:"instance_array_firmware_policies" yaml:"firmwarePolicies"`
}

//FirewallRule describes a firewall rule that is to be applied on all instances of an array
type FirewallRule struct {
	FirewallRuleDescription                    string `json:"firewall_rule_description,omitempty" yaml:"description,omitempty"`
	FirewallRulePortRangeStart                 int    `json:"firewall_rule_port_range_start,omitempty" yaml:"portRangeStart,omitempty"`
	FirewallRulePortRangeEnd                   int    `json:"firewall_rule_port_range_end,omitempty" yaml:"portRangeEnd,omitempty"`
	FirewallRuleSourceIPAddressRangeStart      string `json:"firewall_rule_source_ip_address_range_start,omitempty" yaml:"sourceIPAddressRangeStart,omitempty"`
	FirewallRuleSourceIPAddressRangeEnd        string `json:"firewall_rule_source_ip_address_range_end,omitempty" yaml:"sourceIPAddressRangeEnd,omitempty"`
	FirewallRuleDestinationIPAddressRangeStart string `json:"firewall_rule_destination_ip_address_range_start,omitempty" yaml:"destinationIPAddressRangeStart,omitempty"`
	FirewallRuleDestinationIPAddressRangeEnd   string `json:"firewall_rule_destination_ip_address_range_end,omitempty" yaml:"destinationIPAddressRangeEnd,omitempty"`
	FirewallRuleProtocol                       string `json:"firewall_rule_protocol,omitempty" yaml:"protocol,omitempty"`
	FirewallRuleIPAddressType                  string `json:"firewall_rule_ip_address_type,omitempty" yaml:"IPAddressType,omitempty"`
	FirewallRuleEnabled                        bool   `json:"firewall_rule_enabled,omitempty" yaml:"enabled,omitempty"`
}

//InstanceArrayInterface describes a network interface of the array.
//It's properties will be applied to all InstanceInterfaces of the array's instances.
type InstanceArrayInterface struct {
	InstanceArrayInterfaceLabel            string                           `json:"instance_array_interface_label,omitempty" yaml:"label,omitempty"`
	InstanceArrayInterfaceSubdomain        string                           `json:"instance_array_interface_subdomain,omitempty" yaml:"subdomain,omitempty"`
	InstanceArrayInterfaceID               int                              `json:"instance_array_interface_id,omitempty" yaml:"id,omitempty"`
	InstanceArrayID                        int                              `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	NetworkID                              int                              `json:"network_id,omitempty" yaml:"networkID,omitempty"`
	InstanceArrayInterfaceLAGGIndexes      []interface{}                    `json:"instance_array_interface_lagg_indexes,omitempty" yaml:"LAGGIndexes,omitempty"`
	InstanceArrayInterfaceIndex            int                              `json:"instance_array_interface_index,omitempty" yaml:"index,omitempty"`
	InstanceArrayInterfaceServiceStatus    string                           `json:"instance_array_interface_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	InstanceArrayInterfaceCreatedTimestamp string                           `json:"instance_array_interface_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	InstanceArrayInterfaceUpdatedTimestamp string                           `json:"instance_array_interface_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	InstanceArrayInterfaceOperation        *InstanceArrayInterfaceOperation `json:"instance_array_interface_operation,omitempty" yaml:"operation,omitempty"`
}

//InstanceArrayInterfaceOperation describes changes to a network array interface
type InstanceArrayInterfaceOperation struct {
	InstanceArrayInterfaceLabel            string        `json:"instance_array_interface_label,omitempty" yaml:"label,omitempty"`
	InstanceArrayInterfaceSubdomain        string        `json:"instance_array_interface_subdomain,omitempty" yaml:"subdomain,omitempty"`
	InstanceArrayInterfaceID               int           `json:"instance_array_interface_id,omitempty" yaml:"id,omitempty"`
	InstanceArrayID                        int           `json:"instance_array_id,omitempty" yaml:"instanceArrayID,omitempty"`
	NetworkID                              int           `json:"network_id,omitempty" yaml:"networkID,omitempty"`
	InstanceArrayInterfaceLAGGIndexes      []interface{} `json:"instance_array_interface_lagg_indexes,omitempty" yaml:"LAGGIndexes,omitempty"`
	InstanceArrayInterfaceIndex            int           `json:"instance_array_interface_index,omitempty" yaml:"index,omitempty"`
	InstanceArrayInterfaceServiceStatus    string        `json:"instance_array_interface_service_status,omitempty" yaml:"serviceStatus,omitempty"`
	InstanceArrayInterfaceCreatedTimestamp string        `json:"instance_array_interface_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	InstanceArrayInterfaceUpdatedTimestamp string        `json:"instance_array_interface_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	InstanceArrayInterfaceChangeID         int           `json:"instance_array_interface_change_id,omitempty" yaml:"changeID,omitempty"`
}

//ServerTypeMatches used in InstanceArrayEdit operations to specify which server types to use
type ServerTypeMatches struct {
	ServerTypes map[int]ServerTypeMatch `json:"server_types,omitempty"`
}

//ServerTypeMatch what exact server types to use
type ServerTypeMatch struct {
	ServerCount int `json:"server_count,omitempty"`
}

//instanceArrayGet returns an InstanceArray with specified id
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

//instanceArrays returns list of instance arrays of specified infrastructure
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

//instanceArrayCreate creates an instance array (colletion of identical instances). Requires Deploy.
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

//instanceArrayEdit alterns a deployed instance array. Requires deploy.
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

//instanceArrayDelete deletes an instance array. Requires deploy.
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

//InstanceArrayInterfaceAttachNetwork attaches an InstanceArrayInterface to a Network
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

//InstanceArrayInterfaceDetach detaches an InstanceArrayInterface from any Network element that is attached to.
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

//instanceArrayStop stops a specified InstanceArray.
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

//instanceArrayStart starts a specified InstanceArray.
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

//CreateOrUpdate implements interface Applier
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

//Delete implements interface Applier
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

//Validate implements interface Applier
func (ia InstanceArray) Validate() error {
	if ia.InstanceArrayID == 0 && ia.InstanceArrayLabel == "" {
		return fmt.Errorf("id is required")
	}

	return nil
}
