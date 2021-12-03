package metalcloud

import (
	"fmt"
)

//ServerFirmwareUpgradePolicy represents a server firmware policy.
type ServerFirmwareUpgradePolicy struct {
	ServerFirmwareUpgradePolicyID     int                               `json:"server_firmware_upgrade_policy_id,omitempty" yaml:"id,omitempty"`
	ServerFirmwareUpgradePolicyLabel  string                            `json:"server_firmware_upgrade_policy_label,omitempty" yaml:"label,omitempty"`
	ServerFirmwareUpgradePolicyRules  []ServerFirmwareUpgradePolicyRule `json:"server_firmware_upgrade_policy_rules,omitempty" yaml:"rules,omitempty"`
	ServerFirmwareUpgradePolicyAction string                            `json:"server_firmware_upgrade_policy_action,omitempty" yaml:"action,omitempty"`
	InstanceArrayIDList               []int                             `json:"instance_array_ids,omitempty" yaml:"instanceArrayList,omitempty"`
}

//ServerFirmwareUpgradePolicyRule describes a policy rule.
type ServerFirmwareUpgradePolicyRule struct {
	Operation string `json:"operation,omitempty" yaml:"operation,omitempty"`
	Property  string `json:"property,omitempty" yaml:"property,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}

//ServerFirmwarePolicyGet returns a server policy's details
func (c *Client) ServerFirmwarePolicyGet(serverFirmwarePolicyID int) (*ServerFirmwareUpgradePolicy, error) {

	var createdObject ServerFirmwareUpgradePolicy

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_firmware_policy_get",
		serverFirmwarePolicyID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ServerFirmwareUpgradePolicyCreate creates a server firmware policy.
func (c *Client) ServerFirmwareUpgradePolicyCreate(serverFirmwarePolicy *ServerFirmwareUpgradePolicy) (*ServerFirmwareUpgradePolicy, error) {

	var createdObject *ServerFirmwareUpgradePolicy
	var policyID int
	var action interface{}

	if serverFirmwarePolicy.ServerFirmwareUpgradePolicyAction == "" {
		action = nil
	} else {
		action = serverFirmwarePolicy.ServerFirmwareUpgradePolicyAction
	}

	err := c.rpcClient.CallFor(
		&policyID,
		"server_firmware_policy_create",
		serverFirmwarePolicy.ServerFirmwareUpgradePolicyLabel,
		action,
		serverFirmwarePolicy.ServerFirmwareUpgradePolicyRules,
	)

	if err != nil {
		return nil, err
	}

	createdObject, err = c.ServerFirmwarePolicyGet(policyID)

	if err != nil {
		return nil, err
	}
	return createdObject, nil
}

//ServerFirmwarePolicyAddRule add a new rule for a policy.
func (c *Client) ServerFirmwarePolicyAddRule(serverFirmwarePolicyID int, serverRule *ServerFirmwareUpgradePolicyRule) (*ServerFirmwareUpgradePolicy, error) {

	var createdObject ServerFirmwareUpgradePolicy

	err := c.rpcClient.CallFor(
		&createdObject,
		"server_firmware_policy_add_rule",
		serverFirmwarePolicyID,
		serverRule,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ServerFirmwarePolicyDeleteRule deletes a rule from a policy.
func (c *Client) ServerFirmwarePolicyDeleteRule(serverFirmwarePolicyID int, serverRule *ServerFirmwareUpgradePolicyRule) error {
	resp, err := c.rpcClient.Call(
		"server_firmware_policy_rule_delete",
		serverFirmwarePolicyID,
		serverRule,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	if err != nil {
		return err
	}

	return nil
}

//ServerFirmwareUpgradePolicyDelete deletes all the information about a specified ServerFirmwareUpgradePolicy.
func (c *Client) ServerFirmwareUpgradePolicyDelete(serverFirmwarePolicyID int) error {

	resp, err := c.rpcClient.Call("server_firmware_policy_delete", serverFirmwarePolicyID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (c *Client) ServerFirmwareUgradePolicyInstanceArraySet(serverFirmwarePolicyID int, instanceArrayList []int) error {
	resp, err := c.rpcClient.Call("server_firmware_policy_instance_arrays_set", serverFirmwarePolicyID, instanceArrayList)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
