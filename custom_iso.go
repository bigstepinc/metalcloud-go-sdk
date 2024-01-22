package metalcloud

import "fmt"

type CustomISO struct {
	CustomISOID               int    `json:"custom_iso_id,omitempty" yaml:"customISOID,omitempty"`
	UserIDOwner               int    `json:"user_id_owner,omitempty" yaml:"userIDOwner,omitempty"`
	UserIDAuthenticated       int    `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	CustomISODisplayName      string `json:"custom_iso_display_name,omitempty" yaml:"customISODisplayName,omitempty"`
	CustomISOName             string `json:"custom_iso_name,omitempty" yaml:"customISOName,omitempty"`
	CustomISOType             string `json:"custom_iso_type,omitempty" yaml:"customISOType,omitempty"`
	CustomISOIsPublic         bool   `json:"custom_iso_is_public,omitempty" yaml:"customISOIsPublic,omitempty"`
	CustomISOAccessURL        string `json:"custom_iso_access_url,omitempty" yaml:"customISOAccessURL,omitempty"`
	CustomISOAccessUsername   string `json:"custom_iso_access_username,omitempty" yaml:"customISOAccessUsername,omitempty"`
	CustomISOAccessPassword   string `json:"custom_iso_access_password,omitempty" yaml:"customISOAccessPassword,omitempty"`
	CustomISOCreatedTimestamp string `json:"custom_iso_created_timestamp,omitempty" yaml:"customISOCreatedTimestamp,omitempty"`
	CustomISOUpdatedTimestamp string `json:"custom_iso_updated_timestamp,omitempty" yaml:"customISOUpdatedTimestamp,omitempty"`
}

// CustomISOs returns custom ISOs for user
func (c *Client) CustomISOs(userID id) (*map[string]CustomISO, error) {
	resp, err := c.rpcClient.Call(
		"custom_isos",
		userID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]CustomISO{}
		return &m, nil
	}

	var createdObject map[string]CustomISO

	err = resp.GetObject(&createdObject)
	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// CustomISOCreate creates a custom ISO record
func (c *Client) CustomISOCreate(customISO CustomISO) (*CustomISO, error) {
	var createdObject CustomISO

	err := c.rpcClient.CallFor(
		&createdObject,
		"custom_iso_create",
		customISO)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// CustomISOCreate creates a custom ISO record
func (c *Client) CustomISOGet(customISOID int) (*CustomISO, error) {
	var createdObject CustomISO

	if err := checkID(customISOID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"custom_iso_get",
		customISOID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// CustomISODelete deletes a CustomISO with specified id
func (c *Client) CustomISODelete(customISOID int) error {

	resp, err := c.rpcClient.Call(
		"custom_iso_delete",
		customISOID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// CustomISODelete deletes a CustomISO with specified id
func (c *Client) CustomISOUpdate(customISOID int, customISO CustomISO) (*CustomISO, error) {

	var createdObject CustomISO

	err := c.rpcClient.CallFor(
		&createdObject,
		"custom_iso_update",
		customISOID,
		customISO,
	)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// CustomISOBootIntoServer boots a server with an iso. Returns an AFC group id
func (c *Client) CustomISOBootIntoServer(customISOID int, serverID int) (int, error) {

	var createdObject int

	err := c.rpcClient.CallFor(
		&createdObject,
		"custom_iso_boot_into_server",
		customISOID,
		serverID,
	)

	if err != nil {

		return 0, err
	}

	return createdObject, nil
}
