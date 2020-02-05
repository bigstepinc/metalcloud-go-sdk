package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//Secret struct defines a server type
type Secret struct {
	SecretID               int    `json:"secret_id,omitempty"`
	UserIDOwner            int    `json:"user_id_owner,omitempty"`
	UserIDAuthenticated    int    `json:"user_id_authenticated,omitempty"`
	SecretName             string `json:"secret_name,omitempty"`
	SecretUsage            string `json:"secret_usage,omitempty"`
	SecretBase64           string `json:"secret_base64,omitempty"`
	SecretCreatedTimestamp string `json:"secret_created_timestamp,omitempty"`
	SecretUpdatedTimestamp string `json:"secret_updated_timestamp,omitempty"`
}

//SecretCreate creates a secret
func (c *Client) SecretCreate(secret Secret) (*Secret, error) {
	var createdObject Secret

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"secret_create",
		userID,
		secret)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//SecretDelete Permanently destroys a Secret.
func (c *Client) SecretDelete(secretID int) error {

	if err := checkID(secretID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("secret_delete", secretID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//SecretUpdate This function allows updating the secret_usage, secret_label and secret_base64 of a Secret
func (c *Client) SecretUpdate(secretID int, secret Secret) (*Secret, error) {
	var createdObject Secret

	if err := checkID(secretID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"secret_update",
		secretID,
		secret)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//SecretGet returns a Secret specified by nSecretID. The secret's protected value is never returned.
func (c *Client) SecretGet(secretID int) (*Secret, error) {

	var createdObject Secret

	if err := checkID(secretID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"secret_get",
		secretID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//Secrets retrieves a list of all the Secret objects which a specified User is allowed to see through ownership or delegation. The secret objects never return the actual protected secret value.
func (c *Client) Secrets(usage string) (*map[string]Secret, error) {

	userID, err := c.UserEmailToUserID(c.user)
	if err != nil {
		return nil, err
	}
	var res *jsonrpc.RPCResponse
	if usage != "" {
		res, err = c.rpcClient.Call(
			"secrets",
			*userID,
			usage)
	} else {
		res, err = c.rpcClient.Call(
			"secrets",
			*userID)
	}

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Secret{}
		return &m, nil
	}

	var createdObject map[string]Secret

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}
