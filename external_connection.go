package metalcloud

//go:generate go run helper/gen_exports.go

import "fmt"

//ExternalConnection object describes an external connection
type ExternalConnection struct {
	ExternalConnectionID          int    `json:"external_connection_id,omitempty" yaml:"id,omitempty"`
	ExternalConnectionLabel       string `json:"external_connection_label,omitempty" yaml:"label,omitempty"`
	DatacenterName                string `json:"datacenter_name,omitempty" yaml:"dc,omitempty"`
	ExternalConnectionHidden      bool   `json:"external_connection_hidden" yaml:"hidden"`
	ExternalConnectionDescription string `json:"external_connection_description,omitempty" yaml:"description,omitempty"`
}

//ExternalConnectionGet returns an external connection with specified id
func (c *Client) externalConnectionGet(externalConnectionID id) (*ExternalConnection, error) {
	var createdObject ExternalConnection

	if err := checkID(externalConnectionID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"external_connection_get",
		externalConnectionID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ExternalConnections returns a list of external connections for the specified datacenter
func (c *Client) ExternalConnections(datacenterName string) (*[]ExternalConnection, error) {

	resp, err := c.rpcClient.Call(
		"external_connections",
		datacenterName,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	var createdObject []ExternalConnection

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ExternalConnectionCreate creates an external connection.
func (c *Client) ExternalConnectionCreate(externalConnection ExternalConnection) (*ExternalConnection, error) {
	var createdObject ExternalConnection

	err := c.rpcClient.CallFor(
		&createdObject,
		"external_connection_create",
		externalConnection,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ExternalConnectionEdit updates an external connection.
func (c *Client) externalConnectionEdit(externalConnectionID id, externalConnection ExternalConnection) (*ExternalConnection, error) {
	var createdObject ExternalConnection

	if err := checkID(externalConnectionID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"external_connection_edit",
		externalConnectionID,
		externalConnection,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//ExternalConnectionDelete deletes an external connection.
func (c *Client) externalConnectionDelete(externalConnectionID id) error {

	if err := checkID(externalConnectionID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"external_connection_delete",
		externalConnectionID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
