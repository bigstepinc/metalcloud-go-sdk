package metalcloud

//go:generate go run helper/gen_exports.go

//Network object describes an high level connection construct
type Network struct {
	NetworkID                 int               `json:"network_id,omitempty"`
	NetworkLabel              string            `json:"network_label,omitempty"`
	NetworkSubdomain          string            `json:"network_subdomain,omitempty"`
	NetworkType               string            `json:"network_type,omitempty"`
	InfrastructureID          int               `json:"infrastructure_id,omitempty"`
	NetworkCreatedTimestamp   string            `json:"network_created_timestamp,omitempty"`
	NetworkUpdatedTimestamp   string            `json:"network_updated_timestamp,omitempty"`
	NetworkLANAutoAllocateIPs bool              `json:"network_lan_autoallocate_ips,omitempty"`
	NetworkOperation          *NetworkOperation `json:"network_operation,omitempty"`
}

//NetworkOperation object describes the change(s) required to be applied to a Network
type NetworkOperation struct {
	NetworkID                 int    `json:"network_id,omitempty"`
	NetworkLabel              string `json:"network_label,omitempty"`
	NetworkSubdomain          string `json:"network_subdomain,omitempty"`
	NetworkType               string `json:"network_type,omitempty"`
	InfrastructureID          int    `json:"infrastructure_id,omitempty"`
	NetworkLANAutoAllocateIPs bool   `json:"network_lan_autoallocate_ips,omitempty"`
	NetworkDeployType         string `json:"network_deploy_type,omitempty"`
	NetworkChangeID           int    `json:"network_change_id,omitempty"`
}

//networkGet retrieves a network object
func (c *Client) networkGet(networkID id) (*Network, error) {
	var createdObject Network

	if err := checkID(networkID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_get",
		networkID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//networks returns a list of all network objects of an infrastructure
func (c *Client) networks(infrastructureID id) (*map[string]Network, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	res, err := c.rpcClient.Call(
		"networks",
		infrastructureID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Network{}
		return &m, nil
	}

	var createdObject map[string]Network

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//networkCreate creates a network
func (c *Client) networkCreate(infrastructureID id, network Network) (*Network, error) {
	var createdObject Network

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_create",
		infrastructureID,
		network)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//networkEdit applies a change to an existing network
func (c *Client) networkEdit(networkID id, networkOperation NetworkOperation) (*Network, error) {
	var createdObject Network

	if err := checkID(networkID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_edit",
		networkID,
		networkOperation)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//networkDelete deletes a network.
func (c *Client) networkDelete(networkID id) error {

	if err := checkID(networkID); err != nil {
		return err
	}

	_, err := c.rpcClient.Call(
		"network_delete",
		networkID)

	if err != nil {
		return err
	}

	return nil
}

//networkJoin merges two specified Network objects.
func (c *Client) networkJoin(networkID id, networkToBeDeletedID id) error {

	if err := checkID(networkID); err != nil {
		return err
	}

	if err := checkID(networkToBeDeletedID); err != nil {
		return err
	}

	_, err := c.rpcClient.Call(
		"network_join",
		networkID,
		networkToBeDeletedID)

	if err != nil {
		return err
	}

	return nil
}
