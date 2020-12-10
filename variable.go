package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//Variable struct defines a Variable type
type Variable struct {
	VariableID               int    `json:"variable_id,omitempty" yaml:"id,omitempty"`
	UserIDOwner              int    `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserIDAuthenticated      int    `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	VariableName             string `json:"variable_name,omitempty" yaml:"name,omitempty"`
	VariableUsage            string `json:"variable_usage,omitempty" yaml:"usage,omitempty"`
	VariableJSON             string `json:"variable_json,omitempty" yaml:"json,omitempty"`
	VariableCreatedTimestamp string `json:"variable_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	VariableUpdatedTimestamp string `json:"variable_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

//VariableCreate creates a variable object
func (c *Client) VariableCreate(variable Variable) (*Variable, error) {
	var createdObject Variable

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"variable_create",
		userID,
		variable)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//VariableDelete permanently destroys a Variable.
func (c *Client) VariableDelete(variableID int) error {

	if err := checkID(variableID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("variable_delete", variableID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//VariableUpdate updates a variable
func (c *Client) VariableUpdate(variableID int, variable Variable) (*Variable, error) {
	var createdObject Variable

	if err := checkID(variableID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"variable_update",
		variableID,
		variable)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//VariableGet returns a Variable specified by nVariableID. The Variable's protected value is never returned.
func (c *Client) VariableGet(variableID int) (*Variable, error) {

	var createdObject Variable

	if err := checkID(variableID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"variable_get",
		variableID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//Variables retrieves a list of all the Variable objects which a specified User is allowed to see through ownership or delegation. The Variable objects never return the actual protected Variable value.
func (c *Client) Variables(usage string) (*map[string]Variable, error) {

	userID := c.GetUserID()

	var res *jsonrpc.RPCResponse
	if usage != "" {
		v, err := c.rpcClient.Call(
			"variables",
			userID,
			usage)
		if err != nil {
			return nil, err
		}
		res = v
	} else {
		v, err := c.rpcClient.Call(
			"variables",
			userID)
		if err != nil {
			return nil, err
		}
		res = v
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Variable{}
		return &m, nil
	}

	var createdObject map[string]Variable

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//CreateOrUpdate implements interface Applier
func (v Variable) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var result *Variable
	err = v.Validate()

	if err != nil {
		return err
	}

	if v.VariableID != 0 {
		result, err = client.VariableGet(v.VariableID)
	} else {
		vars, err := client.Variables("")
		if err != nil {
			return err
		}

		for _, variable := range *vars {
			if variable.VariableName == v.VariableName {
				result = &variable
			}
		}
	}

	if err != nil {
		_, err = client.VariableCreate(v)

		if err != nil {
			return err
		}
	} else {
		_, err = client.VariableUpdate(result.VariableID, v)
		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (v Variable) Delete(client MetalCloudClient) error {
	err := v.Validate()

	if err != nil {
		return err
	}

	err = client.VariableDelete(v.VariableID)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (v Variable) Validate() error {
	if v.VariableID == 0 && v.VariableName == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
