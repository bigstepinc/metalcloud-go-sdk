package metalcloud

import "log"

type InstanceArray struct  {
	InstanceArrayID 			float64 	`json:"instance_array_id, omitempty"`
	InstanceArrayLabel 			string		`json:"instance_array_label, omitempty"`
	InstanceArraySubdomain 		string 		`json:"instance_array_subdomain, omitempty"`
	InstanceArrayInstanceCount  float64 	`json:"instance_array_instance_count, omitempty"`
	InstanceArrayRamGbytes 		float64 	`json:"instance_array_ram_gbytes, omitempty"`
	InstanceArrayProcessorCount  float64 	`json:"instance_array_processor_count, omitempty"`
	InstanceArrayProcessorCoreMHZ float64 	`json:"instance_array_processor_core_mhz, omitempty"`
	InstanceArrayProcessorCoreCount float64 `json:"instance_array_processor_core_count, omitempty"`
	InstanceArrayDiskCount 		float64 	`json:"instance_array_disk_count, omitempty"`
	InstanceArrayDiskSizeMBytes float64     `json:"instance_array_disk_size_mbytes, omitempty"`
	InstanceArrayDiskTypes 		[]string 	`json:"instance_array_disk_types, omitempty"`
	InfrastructureID			float64 	`json:"infrastructure_id"`
	InstanceArrayServiceStatus  string 		`json:"instance_array_service_status, omitempty"`

//	instance_array_operation = None;
//	instance_array_interfaces = [];

	ClusterID 						float64 	`json:"cluster_id, omitempty"`			
	ClusterRoleGroup 				string 		`json:"cluster_role_group, omitempty"`			
	InstanceArrayChangeId		float64 	`json:"instance_array_change_id, omitempty"`			
	InstanceArrayFirewallManaged bool 		`json:"instance_array_firewall_managed, omitempty"`
	InstanceArrayFirewallRules   []FirewallRule `json:"instance_array_firewall_rules, omitempty"`;
	VolumeTemplateID 				*string 		`json:"volume_template_id, omitempty"`;
}


type FirewallRule struct {
	FirewallRuleDescription 				string `json:"firewall_rule_description, omitempty "`
	FirewallRulePortRangeStart  			float64 `json:"firewall_rule_port_range_start, omitempty "`
	FirewallRulePortRangeEnd  				float64 `json:"firewall_rule_port_range_end, omitempty "`
	FirewallRuleSourceIPAddressRangeStart 	string `json:"firewall_rule_source_ip_address_range_start, omitempty "`
	FirewallRuleSourceIPAddressRangeEnd   	string `json:"firewall_rule_source_ip_address_range_end, omitempty "`
	FirewallRuleProtocol 					string `json:"firewall_rule_protocol, omitempty "`
	FirewallRuleIPAddressType 			    string `json:"firewall_rule_ip_address_type, omitempty "`
	FirewallRuleEnabled 					bool   `json:"firewall_rule_ip_address_type, omitempty "`
}


func (c *MetalCloudClient) InstanceArrayGet(instanceArrayID float64) (*InstanceArray, error) {
	var created_object InstanceArray

	err := c.rpcClient.CallFor(
		&created_object,
		"instance_array_get",
		instanceArrayID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &created_object, nil
}


func (c *MetalCloudClient) InstanceArrays(InfrastructureID float64) (*map[string]InstanceArray, error) {
	var created_object map[string]InstanceArray

	err := c.rpcClient.CallFor(
		&created_object,
		"instance_arrays",
		InfrastructureID)
	if err != nil {
		return nil, err
	}

	return &created_object, nil
}



func (c *MetalCloudClient) InstanceArrayCreate(infrastructureID float64, instanceArray InstanceArray) (*InstanceArray, error) {
	var created_object InstanceArray

	err := c.rpcClient.CallFor(
		&created_object,
		"instance_array_create",
		infrastructureID,
		instanceArray)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &created_object, nil
}