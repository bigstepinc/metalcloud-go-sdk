package metalcloud

import "fmt"

//go:generate go run helper/gen_exports.go

//Instance object describes an instance
type Instance struct {
	InstanceID                 int                 `json:"instance_id,omitempty"`
	InstanceLabel              string              `json:"instance_label,omitempty"`
	InstanceSubdomain          string              `json:"instance_subdomain,omitempty"`
	InstanceSubdomainPermanent string              `json:"instance_subdomain_permanent,omitempty"`
	InstanceArrayID            int                 `json:"instance_array_id,omitempty"`
	ServerID                   int                 `json:"server_id,omitempty"`
	ServerTypeID               int                 `json:"server_type_id,omitempty"`
	InstanceServiceStatus      string              `json:"instance_service_status,omitempty"`
	InstanceCredentials        InstanceCredentials `json:"instance_credentials,omitempty"`
	InstanceOperation          InstanceOperation   `json:"instance_operation,omitempty"`
	InstanceInterfaces         []InstanceInterface `json:"instance_interfaces,omitempty"`
	InstanceCreatedTimestamp   string              `json:"instance_created_timestamp,omitempty"`
	InstanceUpdatedTimestamp   string              `json:"instance_updated_timestamp,omitempty"`
	DriveIDBootable            int                 `json:"drive_id_bootable,omitempty"`
	InstanceChangeID           int                 `json:"instance_change_id,omitempty"`
	TemplateIDOrigin           int                 `json:"template_id_origin,omitempty"`
}

//InstanceOperation contains information regarding the changes that are to be made to a product. Edit and deploy functions have to be called in order to apply the changes. The operation type and status are unique to each operation object.
type InstanceOperation struct {
	InstanceID                 int    `json:"instance_id,omitempty"`
	InstanceDeployType         string `json:"instance_deploy_type,omitempty"`
	InstanceDeployStatus       string `json:"instance_deploy_status,omitempty"`
	InstanceLabel              string `json:"instance_label,omitempty"`
	InstanceSubdomain          string `json:"instance_subdomain,omitempty"`
	InstanceSubdomainPermanent string `json:"instance_subdomain_permanent,omitempty"`
	InstanceArrayID            int    `json:"instance_array_id,omitempty"`
	ServerID                   int    `json:"server_id,omitempty"`
	ServerTypeID               int    `json:"server_type_id,omitempty"`
	InstanceChangeID           int    `json:"instance_change_id,omitempty"`
	TemplateIDOrigin           int    `json:"template_id_origin,omitempty"`
}

//InstanceInterface objects are created automatically when instances are created. Subnets are added on networks and then IP addresses are associated automatically or manually through the API to instance interfaces.
type InstanceInterface struct {
	InstanceInterfaceLabel         string                     `json:"instance_interface_label,omitempty"`
	InstanceInterfaceSubdomain     string                     `json:"instance_interface_subdomain,omitempty"`
	InstanceInterfaceID            int                        `json:"instance_interface_id,omitempty"`
	InstanceID                     int                        `json:"instance_id,omitempty"`
	NetworkID                      int                        `json:"network_id,omitempty"`
	InstanceInterfaceLaggIndexes   []int                      `json:"instance_interface_lagg_indexes,omitempty"`
	InstanceInterfaceIndex         int                        `json:"instance_interface_index,omitempty"`
	InstanceInterfaceCapacityMbps  int                        `json:"instance_interface_capacity_mbps,omitempty"`
	InstanceInterfaceServiceStatus string                     `json:"instance_interface_service_status,omitempty"`
	ServerInterface                ServerInterface            `json:"server_interface,omitempty"`
	InstanceInterfaceOperation     InstanceInterfaceOperation `json:"instance_interface_operation,omitempty"`
	InstanceInterfaceIPs           []IP                       `json:"instance_interface_ips,omitempty"`
	InstanceInterfaceChangeID      int                        `json:"instance_interface_change_id,omitempty"`
}

//InstanceInterfaceOperation objects are created automatically when instances are created. Subnets are added on networks and then IP addresses are associated automatically or manually through the API to instance interfaces.
type InstanceInterfaceOperation struct {
	InstanceInterfaceLabel        string `json:"instance_interface_label,omitempty"`
	InstanceInterfaceSubdomain    string `json:"instance_interface_subdomain,omitempty"`
	InstanceInterfaceDeployStatus string `json:"instance_interface_deploy_status,omitempty"`
	InstanceInterfaceDeployType   string `json:"instance_interface_deploy_type,omitempty"`
	InstanceInterfaceID           int    `json:"instance_interface_id,omitempty"`
	InstanceID                    int    `json:"instance_id,omitempty"`
	NetworkID                     int    `json:"network_id,omitempty"`
	InstanceInterfaceLaggIndexes  []int  `json:"instance_interface_lagg_indexes,omitempty"`
	InstanceInterfaceIndex        int    `json:"instance_interface_index,omitempty"`
	InstanceInterfaceCapacityMbps int    `json:"instance_interface_capacity_mbps,omitempty"`
	InstanceInterfaceChangeID     int    `json:"instance_interface_change_id,omitempty"`
}

//ServerInterface contains server connectivity information.
type ServerInterface struct {
	ServerInterfaceMACAddress string `json:"server_interface_mac_address,omitempty"`
}

//InstanceCredentials contains information needed to connect to the server via IPMI, iLO etc.
type InstanceCredentials struct {
	SSH                *SSH            `json:"ssh,omitempty"`
	RDP                *RDP            `json:"rdp,omitempty"`
	IPMI               *IPMI           `json:"ipmi,omitempty"`
	ILO                *ILO            `json:"ilo,omitempty"`
	IDRAC              *IDRAC          `json:"idrac,omitempty"`
	ISCSI              *ISCSIInitiator `json:"iscsi,omitempty"`
	RemoteConsole      *RemoteConsole  `json:"remote_console,omitempty"`
	SharedDrives       []ISCSI         `json:"shared_drives,omitempty"`
	IPAddressesPublic  []IP            `json:"ip_addresses_public,omitempty"`
	IPAddressesPrivate []IP            `json:"ip_addresses_private,omitempty"`
}

//SSH credentials for the installed OS.
type SSH struct {
	Port            int     `json:"port,omitempty"`
	Username        string  `json:"username,omitempty"`
	InitialPassword string  `json:"initial_password,omitempty"`
	InitialSSHKeys  *SSHKey `json:"initial_ssh_keys,omitempty"`
}

//SSHKey represents an SSH key added by a user
type SSHKey struct {
	UserSSHKeyID               int    `json:"user_ssh_key_id,omitempty"`
	UserID                     int    `json:"user_id,omitempty"`
	UserSSHKey                 string `json:"user_ssh_key,omitempty"`
	UserSSHKeyCreatedTimeStamp string `json:"user_ssh_key_created_timestamp,omitempty"`
	UserSSHKeyStatus           string `json:"user_ssh_key_status,omitempty"`
}

//RDP credentials for the installed OS.
type RDP struct {
	Port            int    `json:"port,omitempty"`
	Username        string `json:"username,omitempty"`
	InitialPassword string `json:"initial_password,omitempty"`
}

//IPMI credentials.
type IPMI struct {
	IPAddress       string `json:"ip_address,omitempty"`
	Version         string `json:"version,omitempty"`
	Username        string `json:"username,omitempty"`
	InitialPassword string `json:"initial_password,omitempty"`
}

//ILO control panel credentials
type ILO struct {
	ControlPanelURL string `json:"control_panel_url,omitempty"`
	Username        string `json:"username,omitempty"`
	InitialPassword string `json:"initial_password,omitempty"`
}

//IDRAC control panel credentials.
type IDRAC struct {
	ControlPanelURL string `json:"control_panel_url,omitempty"`
	Username        string `json:"username,omitempty"`
	InitialPassword string `json:"initial_password,omitempty"`
}

//ISCSIInitiator provides initiator IQN, username and password and other iSCSI connection details.
type ISCSIInitiator struct {
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	InitiatorIQN       string `json:"initiator_iqn,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	InitiatorIPAddress string `json:"initiator_ip_address,omitempty"`
}

//RemoteConsole provides credentials needed to connect to the server via the HTML interface
type RemoteConsole struct {
	RemoteProtocol        string `json:"remote_protocol,omitempty"`
	TunnelPathURL         string `json:"tunnel_path_url,omitempty"`
	RemoteControlPanelURL string `json:"remote_control_panel_url,omitempty"`
}

//ISCSI provides target IQN, IP address, port number and the LUN ID.
type ISCSI struct {
	TargetIQN        string `json:"target_iqn,omitempty"`
	StorageIPAddress string `json:"storage_ip_address,omitempty"`
	StoragePort      int    `json:"storage_port,omitempty"`
	LunID            int    `json:"lun_id,omitempty"`
}

//IP object contains information regarding an IP address.
type IP struct {
	IPID                       int         `json:"ip_id,omitempty"`
	IPType                     string      `json:"ip_type,omitempty"`
	IPHumanReadable            string      `json:"ip_human_readable,omitempty"`
	IPHex                      string      `json:"ip_hex,omitempty"`
	IPLeaseExpires             string      `json:"ip_lease_expires,omitempty"`
	IPOperation                IPOperation `json:"ip_operation,omitempty"`
	SubnetID                   int         `json:"subnet_id,omitempty"`
	SubnetDestination          string      `json:"subnet_destination,omitempty"`
	SubnetGatewayHumanReadable string      `json:"subnet_gateway_human_readable,omitempty"`
	SubnetNetmaskHumanReadable string      `json:"subnet_netmask_human_readable,omitempty"`
	InstanceInterfaceID        int         `json:"instance_interface_id,omitempty"`
	IPChangeID                 int         `json:"ip_change_id,omitempty"`
}

//IPOperation contains information regarding the changes that are to be made to a product. Edit and deploy functions have to be called in order to apply the changes. The operation type and status are unique to each operation object.
type IPOperation struct {
	InstanceInterfaceID int    `json:"instance_interface_id,omitempty"`
	IPDeployStatus      string `json:"ip_deploy_status,omitempty"`
	IPDeployType        string `json:"ip_deploy_type,omitempty"`
	IPID                int    `json:"ip_id,omitempty"`
	IPType              string `json:"ip_type,omitempty"`
	IPHumanReadable     string `json:"ip_human_readable,omitempty"`
	IPHex               string `json:"ip_hex,omitempty"`
	IPLabel             string `json:"ip_label,omitempty"`
	IPSubdomain         string `json:"ip_subdomain,omitempty"`
	IPLeaseExpires      string `json:"ip_lease_expires,omitempty"`
	IPUpdatedTimestamp  string `json:"ip_updated_timestamp,omitempty"`
	SubnetID            int    `json:"subnet_id,omitempty"`
	IPChangeID          int    `json:"ip_change_id,omitempty"`
}

//instanceArrayInstances retrieves a list of all the Instance objects associated with a specified InstanceArray.
func (c *Client) instanceArrayInstances(instanceArrayID id) (*map[string]Instance, error) {

	if err := checkID(instanceArrayID); err != nil {
		return nil, err
	}

	res, err := c.rpcClient.Call(
		"instance_array_instances",
		instanceArrayID,
		nil)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Instance{}
		return &m, nil
	}

	var createdObject map[string]Instance

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//instanceGet returns a specific instance by id
func (c *Client) instanceGet(instanceID id) (*Instance, error) {
	var instance Instance

	if err := checkID(instanceID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&instance, "instance_get", instanceID)

	if err != nil {
		return nil, err
	}

	return &instance, nil
}

//instanceServerPowerSet reboots or powers on an instance
func (c *Client) instanceServerPowerSet(instanceID id, operation string) error {
	if err := checkID(instanceID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("instance_server_power_set", instanceID, operation)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//instanceServerPowerGet returns the power status of an instance
func (c *Client) instanceServerPowerGet(instanceID id) (*string, error) {
	var power string

	if err := checkID(instanceID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(&power, "instance_server_power_get", instanceID)

	if err != nil {
		return nil, err
	}

	return &power, nil
}

//instanceServerPowerGetBatch returns the power status of multiple instances
func (c *Client) instanceServerPowerGetBatch(infrastructureID id, instanceIDs []int) (*map[string]string, error) {

	if err := checkID(infrastructureID); err != nil {
		return nil, err
	}

	res, err := c.rpcClient.Call("instance_server_power_get_batch", infrastructureID, instanceIDs)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]string{}
		return &m, nil
	}

	var createdObject map[string]string

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}
