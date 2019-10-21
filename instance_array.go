package metalcloud

import "log"
import "fmt"


type InstanceArray struct  {
	InstanceArrayID 			int 	`json:"instance_array_id, omitempty"`
	InstanceArrayLabel 			string		`json:"instance_array_label, omitempty"`
	InstanceArraySubdomain 		string 		`json:"instance_array_subdomain, omitempty"`
	InstanceArrayInstanceCount  int 	`json:"instance_array_instance_count, omitempty"`
	InstanceArrayRamGbytes 		int 	`json:"instance_array_ram_gbytes, omitempty"`
	InstanceArrayProcessorCount  int 	`json:"instance_array_processor_count, omitempty"`
	InstanceArrayProcessorCoreMHZ int 	`json:"instance_array_processor_core_mhz, omitempty"`
	InstanceArrayProcessorCoreCount int `json:"instance_array_processor_core_count, omitempty"`
	InstanceArrayDiskCount 		int 	`json:"instance_array_disk_count, omitempty"`
	InstanceArrayDiskSizeMBytes int     `json:"instance_array_disk_size_mbytes, omitempty"`
	InstanceArrayDiskTypes 		[]string 	`json:"instance_array_disk_types, omitempty"`
	InfrastructureID			int 		`json:"infrastructure_id"`
	InstanceArrayServiceStatus  string 		`json:"instance_array_service_status, omitempty"`

//	instance_array_operation = None;
//	instance_array_interfaces = [];

	ClusterID 						int 	`json:"cluster_id, omitempty"`			
	ClusterRoleGroup 				string 		`json:"cluster_role_group, omitempty"`			
	InstanceArrayChangeId			int 	`json:"instance_array_change_id, omitempty"`			
	InstanceArrayFirewallManaged 	bool 		`json:"instance_array_firewall_managed, omitempty"`
	InstanceArrayFirewallRules   	[]FirewallRule `json:"instance_array_firewall_rules, omitempty"`;
	VolumeTemplateID 				int 		`json:"volume_template_id, omitempty"`;
}


type FirewallRule struct {
	FirewallRuleDescription 				string `json:"firewall_rule_description, omitempty "`
	FirewallRulePortRangeStart  			int `json:"firewall_rule_port_range_start, omitempty "`
	FirewallRulePortRangeEnd  				int `json:"firewall_rule_port_range_end, omitempty "`
	FirewallRuleSourceIPAddressRangeStart 	string `json:"firewall_rule_source_ip_address_range_start, omitempty "`
	FirewallRuleSourceIPAddressRangeEnd   	string `json:"firewall_rule_source_ip_address_range_end, omitempty "`
	FirewallRuleProtocol 					string `json:"firewall_rule_protocol, omitempty "`
	FirewallRuleIPAddressType 			    string `json:"firewall_rule_ip_address_type, omitempty "`
	FirewallRuleEnabled 					bool   `json:"firewall_rule_ip_address_type, omitempty "`
}


func (c *MetalCloudClient) InstanceArrayGet(instanceArrayID int) (*InstanceArray, error) {
	var created_object InstanceArray

	err := c.rpcClient.CallFor(
		&created_object,
		"instance_array_get",
		instanceArrayID)

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	return &created_object, nil
}


func (c *MetalCloudClient) InstanceArrays(infrastructureID int) (*map[string]InstanceArray, error) {
	
	res, err := c.rpcClient.Call(
		"instance_arrays",
		infrastructureID)
	
	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]InstanceArray{}
		return &m, nil
	}

	var created_object map[string]InstanceArray

	err2 := res.GetObject(&created_object)
	if err2 != nil {
			return nil, err2
	}

	return &created_object, nil
}



func (c *MetalCloudClient) InstanceArrayCreate(infrastructureID int, instanceArray InstanceArray) (*InstanceArray, error) {
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