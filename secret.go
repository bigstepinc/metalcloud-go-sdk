package metalcloud

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//Secret struct defines a server type
type Secret struct {
	SecretID               int    `json:"secret_id,omitempty" yaml:"id,omitempty"`
	UserIDOwner            int    `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserIDAuthenticated    int    `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	SecretName             string `json:"secret_name,omitempty" yaml:"name,omitempty"`
	SecretUsage            string `json:"secret_usage,omitempty" yaml:"usage,omitempty"`
	SecretBase64           string `json:"secret_base64,omitempty" yaml:"base64,omitempty"`
	SecretCreatedTimestamp string `json:"secret_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	SecretUpdatedTimestamp string `json:"secret_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
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

	userID := c.GetUserID()

	var err error
	var res *jsonrpc.RPCResponse
	if usage != "" {
		res, err = c.rpcClient.Call(
			"secrets",
			userID,
			usage)
	} else {
		res, err = c.rpcClient.Call(
			"secrets",
			userID)
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

//CreateOrUpdate implements interface Applier
func (s Secret) CreateOrUpdate(c interface{}) error {
	client := c.(*Client)

	var err error

	if s.SecretID != 0 {
		_, err = client.SecretGet(s.SecretID)
	} else if s.SecretName != "" {
		secrets, err := client.Secrets("")
		if err != nil {
			return err
		}
		err = fmt.Errorf("secret not found")

		for _, secret := range *secrets {
			if secret.SecretName == s.SecretName {
				err = nil
			}
		}
	} else {
		return fmt.Errorf("id is required")
	}

	if err != nil {
		_, err = client.SecretCreate(s)

		if err != nil {
			return err
		}
	} else {
		_, err = client.SecretUpdate(s.SecretID, s)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (s Secret) Delete(c interface{}) error {
	client := c.(*Client)

	err := client.SecretDelete(s.SecretID)

	if err != nil {
		return err
	}

	return nil
}
