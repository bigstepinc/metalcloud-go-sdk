package metalcloud

import (
	"fmt"
	"log"
)

//Network object describes an high level connection construct
type Network struct {
	NetworkID                 int    `json:"network_id,omitempty"`
	NetworkLabel              string `json:"network_label,omitempty"`
	NetworkSubdomain          string `json:"network_subdomain,omitempty"`
	NetworkType               string `json:"network_type,omitempty"`
	InfrastructureID          int    `json:"infrastructure_id,omitempty"`
	NetworkServiceStatus      string `json:"network_service_status,omitempty"`
	NetworkCreatedTimestamp   string `json:"network_created_timestamp,omitempty"`
	NetworkUpdatedTimestamp   string `json:"network_updated_timestamp,omitempty"`
	NetworkSuspendStatus      string `json:"network_suspend_status,omitempty"`
	NetworkLANAutoAllocateIPs bool   `json:"network_lan_autoallocate_ips,omitempty"`

	NetworkOperation *NetworkOperation `json:"network_operation,omitempty"`
}

//NetworkOperation object describes the change(s) required to be applied to a Network
type NetworkOperation struct {
	NetworkID               int    `json:"network_id,omitempty"`
	NetworkLabel            string `json:"network_label,omitempty"`
	NetworkSubdomain        string `json:"network_subdomain,omitempty"`
	NetworkType             string `json:"network_type,omitempty"`
	InfrastructureID        int    `json:"infrastructure_id,omitempty"`
	NetworkServiceStatus    string `json:"network_service_status,omitempty"`
	NetworkCreatedTimestamp string `json:"network_created_timestamp,omitempty"`
	NetworkUpdatedTimestamp string `json:"network_updated_timestamp,omitempty"`

	NetworkLANAutoAllocateIPs bool `json:"network_lan_autoallocate_ips,omitempty"`

	NetworkDeployType string `json:"network_deploy_type,omitempty"`
	NetworkChangeID   string `json:"network_change_id,omitempty"`
}

//NetworkGet retrieves a network object
func (c *Client) NetworkGet(NetworkID int) (*Network, error) {
	var createdObject Network

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_get",
		NetworkID)

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	return &createdObject, nil
}

//Networks returns a list of all network objects of an infrastructure
func (c *Client) Networks(infrastructureID int) (*map[string]Network, error) {

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

//NetworkCreate creates a network
func (c *Client) NetworkCreate(infrastructureID int, network Network) (*Network, error) {
	var createdObject Network

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_create",
		infrastructureID,
		network)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &createdObject, nil
}

//NetworkEdit applies a change to an existing network
func (c *Client) NetworkEdit(networkID int, networkOperation NetworkOperation) (*Network, error) {
	var createdObject Network

	err := c.rpcClient.CallFor(
		&createdObject,
		"network_edit",
		networkID,
		networkOperation)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &createdObject, nil
}

//NetworkDelete deletes a network.
func (c *Client) NetworkDelete(networkID int) error {

	_, err := c.rpcClient.Call(
		"network_delete",
		networkID)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}

//NetworkJoin merges two specified Network objects.
func (c *Client) NetworkJoin(networkID int, networkToBeDeletedID int) error {

	_, err := c.rpcClient.Call(
		"network_join",
		networkID,
		networkToBeDeletedID)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}
