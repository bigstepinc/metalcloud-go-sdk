package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//Variable struct defines a Variable type
type Variable struct {
	VariableID               int    `json:"variable_id,omitempty"`
	UserIDOwner              int    `json:"user_id_owner,omitempty"`
	UserIDAuthenticated      int    `json:"user_id_authenticated,omitempty"`
	VariableName             string `json:"variable_name,omitempty"`
	VariableUsage            string `json:"variable_usage,omitempty"`
	VariableJSON             string `json:"variable_json,omitempty"`
	VariableCreatedTimestamp string `json:"variable_created_timestamp,omitempty"`
	VariableUpdatedTimestamp string `json:"variable_updated_timestamp,omitempty"`
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
