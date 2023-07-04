package metalcloud

import (
	"fmt"
	"strings"
)

// SwitchDeviceController represents a switch controller installed in a datacenter.
type SwitchDeviceController struct {
	DatacenterName                               string      `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
	NetworkEquipmentControllerDescription        string      `json:"network_equipment_controller_description,omitempty" yaml:"description,omitempty"`
	NetworkEquipmentControllerDriver             string      `json:"network_equipment_controller_driver,omitempty" yaml:"driver,omitempty"`
	NetworkEquipmentControllerID                 int         `json:"network_equipment_controller_id,omitempty" yaml:"id,omitempty"`
	NetworkEquipmentControllerIdentifierString   string      `json:"network_equipment_controller_identifier_string,omitempty" yaml:"identifierString,omitempty"`
	NetworkEquipmentControllerManagementAddress  string      `json:"network_equipment_controller_management_address,omitempty" yaml:"managementAddress,omitempty"`
	NetworkEquipmentControllerManagementPassword string      `json:"network_equipment_controller_management_password,omitempty" yaml:"managementPassword,omitempty"`
	NetworkEquipmentControllerManagementPort     int         `json:"network_equipment_controller_management_port,omitempty" yaml:"managementPort,omitempty"`
	NetworkEquipmentControllerManagementUsername string      `json:"network_equipment_controller_management_username,omitempty" yaml:"managementUsername,omitempty"`
	NetworkEquipmentControllerOptions            interface{} `json:"network_equipment_controller_options,omitempty" yaml:"options,omitempty"`
	NetworkEquipmentControllerOptionsJSON        string      `json:"network_equipment_controller_options_json,omitempty" yaml:"optionsJSON,omitempty"`
	NetworkEquipmentControllerProvisionerType    string      `json:"network_equipment_controller_provisioner_type,omitempty" yaml:"provisionerType,omitempty"`
}

// SwitchDeviceControllerGet retrieves information regarding a specified SwitchDeviceController.
func (c *Client) SwitchDeviceControllerGet(networkEquipmentControllerID int, decryptPasswd bool) (*SwitchDeviceController, error) {
	var createdObject SwitchDeviceController

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_controller_get",
		networkEquipmentControllerID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.NetworkEquipmentControllerManagementPassword, ":")

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
				var passwd string

				err = c.rpcClient.CallFor(
					&passwd,
					"password_decrypt",
					passwdComponents[1],
				)
				if err != nil {
					return nil, err
				}

				createdObject.NetworkEquipmentControllerManagementPassword = passwd
			}
		}
	} else {
		createdObject.NetworkEquipmentControllerManagementPassword = ""
	}

	return &createdObject, nil
}

// SwitchDeviceControllerGetByIdentifierString retrieves information regarding a specified SwitchDeviceController by identifier string.
func (c *Client) SwitchDeviceControllerGetByIdentifierString(networkEquipmentIdentifierString string, decryptPasswd bool) (*SwitchDeviceController, error) {
	var createdObject SwitchDeviceController

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_controller_get",
		networkEquipmentIdentifierString)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {
		passwdComponents := strings.Split(createdObject.NetworkEquipmentControllerManagementPassword, ":")

		if len(passwdComponents) == 2 {
			if strings.Contains(passwdComponents[0], "Not authorized") {
				return nil, fmt.Errorf("Permission missing. %s", passwdComponents[1])
			} else {
				var passwd string

				err = c.rpcClient.CallFor(
					&passwd,
					"password_decrypt",
					passwdComponents[1],
				)
				if err != nil {
					return nil, err
				}

				createdObject.NetworkEquipmentControllerManagementPassword = passwd
			}
		}
	} else {
		createdObject.NetworkEquipmentControllerManagementPassword = ""
	}

	return &createdObject, nil
}

// SwitchDeviceControllerCreate creates a record for a new SwitchDeviceController and for the switches that were detected and created.
// The return value is a map of switch devices belonging to the newly created SwitchDeviceController.
func (c *Client) SwitchDeviceControllerCreate(switchDevice SwitchDevice) (*map[int]SwitchDevice, error) {
	var createdObject map[int]SwitchDevice

	// When making a call with a single object parameter, we have to put it into an array.
	resp, err := c.rpcClient.Call(
		"switch_device_create_from_cisco_aci",
		[]SwitchDevice{switchDevice},
	)

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[int]SwitchDevice{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}
	return &createdObject, nil
}

// SwitchDeviceControllers retrieves all switch controller devices registered in the database.
// If a datacenter is specified, only the switch device controllers for that datacenter are returned.
func (c *Client) SwitchDeviceControllers(datacenter string) (*map[int]SwitchDeviceController, error) {
	var dc *string
	if datacenter != "" {
		dc = &datacenter
	}

	var createdObject map[int]SwitchDeviceController

	resp, err := c.rpcClient.Call(
		"switch_device_controllers",
		dc,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[int]SwitchDeviceController{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// SwitchDeviceControllers retrieves all switch devices registered in the database.
func (c *Client) SwitchDeviceControllerSwitches(networkEquipmentIdentifierString string) (map[int]SwitchDevice, error) {
	switchDeviceController, err := c.SwitchDeviceControllerGetByIdentifierString(networkEquipmentIdentifierString, false)

	if err != nil {
		return nil, err
	}

	switchDevices, err := c.SwitchDevices(switchDeviceController.DatacenterName, "")

	if err != nil {
		return nil, err
	}

	controllerSwitchDevices := map[int]SwitchDevice{}

	for _, switchDevice := range *switchDevices {
		if switchDevice.NetworkEquipmentControllerID == switchDeviceController.NetworkEquipmentControllerID {
			controllerSwitchDevices[switchDevice.NetworkEquipmentID] = switchDevice
		}
	}

	return controllerSwitchDevices, nil
}

// SwitchDeviceUpdate updates an existing switch configuration
func (c *Client) SwitchDeviceControllerUpdate(networkEquipmentControllerID int, networkEquipmentControllerData interface{}) (*SwitchDeviceController, error) {
	var createdObject SwitchDeviceController

	err := c.rpcClient.CallFor(
		&createdObject,
		"switch_device_controller_update",
		networkEquipmentControllerID,
		networkEquipmentControllerData,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}
