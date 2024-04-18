package metalcloud

import "fmt"

//go:generate go run helper/gen_exports.go
//go:generate go run helper/docgen.go - $GOFILE ./ Network,NetworkOperation Network

// Network object describes an high level connection construct
type Network struct {
	// description: The ID of the object.
	NetworkID int `json:"network_id,omitempty" yaml:"id,omitempty"`
	// description: The Label of the object.
	NetworkLabel string `json:"network_label,omitempty" yaml:"label,omitempty"`
	// description: A unique string representing the network. Deprecated
	NetworkSubdomain string `json:"network_subdomain,omitempty" yaml:"subdomain,omitempty"`
	// description: The network type
	// values:
	// - wan
	// - lan
	// - san
	NetworkType string `json:"network_type,omitempty" yaml:"type,omitempty"`
	// The infrastructure on which this network is defined
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the network was created.
	//example: 2013-11-29T13:00:01Z
	NetworkCreatedTimestamp string `json:"network_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the network was last updated.
	//example: 2013-11-29T13:00:01Z
	NetworkUpdatedTimestamp string `json:"network_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	//description: If set to true and if the network is of type LAN IPs will be automatically allocated from any attached subnet pool. Deprecated.
	NetworkLANAutoAllocateIPs bool `json:"network_lan_autoallocate_ips,omitempty" yaml:"LANAutoAllocateIPs,omitempty"`
	//description: The operation object.
	NetworkOperation *NetworkOperation `json:"network_operation,omitempty" yaml:"operation,omitempty"`
}

// NetworkOperation object describes the change(s) required to be applied to a Network
type NetworkOperation struct {
	// description: The ID of the object.
	NetworkID int `json:"network_id,omitempty" yaml:"id,omitempty"`
	// description: The Label of the object.
	NetworkLabel string `json:"network_label,omitempty" yaml:"label,omitempty"`
	// description: A unique string representing the network. Deprecated
	NetworkSubdomain string `json:"network_subdomain,omitempty" yaml:"subdomain,omitempty"`
	// description: The network type
	// values:
	// - wan
	// - lan
	// - san
	NetworkType string `json:"network_type,omitempty" yaml:"type,omitempty"`
	// The infrastructure on which this network is defined
	InfrastructureID int `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	//description: If set to true and if the network is of type LAN IPs will be automatically allocated from any attached subnet pool. Deprecated.
	NetworkLANAutoAllocateIPs bool `json:"network_lan_autoallocate_ips" yaml:"LANAutoAllocateIPs"`
	// description: The deploy type
	// values:
	//     - create
	// 	   - delete
	//     - edit
	// 	   - start
	//     - stop
	// 	   - suspend
	NetworkDeployType string `json:"network_deploy_type,omitempty" yaml:"deployType,omitempty"`
	// description: The id of the change operation. Readonly.
	NetworkChangeID int `json:"network_change_id,omitempty" yaml:"changeID,omitempty"`
}

// networkGet retrieves a network object
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

// networks returns a list of all network objects of an infrastructure
func (c *Client) networks(infrastructureID id) (*map[string]Network, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"networks",
		infrastructureID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Network{}
		return &m, nil
	}

	var createdObject map[string]Network

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// networkCreate creates a network
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

// networkEdit applies a change to an existing network
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

// networkDelete deletes a network.
func (c *Client) networkDelete(networkID id) error {

	if err := checkID(networkID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"network_delete",
		networkID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// networkJoin merges two specified Network objects.
func (c *Client) networkJoin(networkID id, networkToBeDeletedID id) error {

	if err := checkID(networkID); err != nil {
		return err
	}

	if err := checkID(networkToBeDeletedID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call(
		"network_join",
		networkID,
		networkToBeDeletedID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

func (n *Network) instanceToOperation(op *NetworkOperation) {
	operation := n.NetworkOperation
	operation.NetworkID = n.NetworkID
	operation.NetworkLabel = n.NetworkLabel
	operation.NetworkSubdomain = n.NetworkSubdomain
	operation.NetworkType = n.NetworkType
	operation.NetworkLANAutoAllocateIPs = n.NetworkLANAutoAllocateIPs
	operation.NetworkChangeID = op.NetworkChangeID
}

// CreateOrUpdate implements interface Applier
func (n Network) CreateOrUpdate(client MetalCloudClient) error {
	var result *Network
	var err error
	err = n.Validate()

	if err != nil {
		return err
	}

	if n.NetworkID != 0 {
		result, err = client.NetworkGet(n.NetworkID)
	} else {
		result, err = client.NetworkGetByLabel(n.NetworkLabel)
	}

	if err != nil {
		_, err = client.NetworkCreate(n.InfrastructureID, n)

		if err != nil {
			return err
		}
	} else {
		n.instanceToOperation(result.NetworkOperation)
		_, err = client.NetworkEdit(result.NetworkID, *n.NetworkOperation)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (n Network) Delete(client MetalCloudClient) error {
	var result *Network
	var id int
	err := n.Validate()

	if err != nil {
		return err
	}

	if n.NetworkLabel != "" {
		result, err = client.NetworkGetByLabel(n.NetworkLabel)
		if err != nil {
			return err
		}
		id = result.NetworkID
	} else {
		id = n.NetworkID
	}
	err = client.NetworkDelete(id)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (n Network) Validate() error {
	if n.NetworkID == 0 && n.NetworkLabel == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
