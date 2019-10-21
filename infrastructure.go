package metalcloud

import "log"
import "fmt"

type InfrastructureOperation struct {
	InfrastructureDeployStatus  string  `json:"infrastructure_deploy_status, omitempty"`
	InfrastructureDeployType    string  `json:"infrastructure_deploy_type, omitempty"`
	InfrastructureLabel          string  `json:"infrastructure_label"`
	InfrastructureSubdomain      string  `json:"infrastructure_subdomain, omitempty"`
	DatacenterName               string  `json:"datacenter_name"`
	InfrastructureID             float64 `json:"infrastructure_id,omitempty"`
	UserIDOwner                 float64 `json:"user_id_owner,omitempty"`
	InfrastructureUpdatedTimestamp string  `json:"infrastructure_updated_timestamp,omitempty"`
	InfrastructureChangeID         float64 `json:"infrastructure_change_id,omitempty"`
	InfrastructureDeployID         float64 `json:"infrastructure_deploy_id,omitempty"`
}

type Infrastructure struct {
	InfrastructureLabel          string  `json:"infrastructure_label"`
	InfrastructureSubdomain      string  `json:"infrastructure_subdomain, omitempty"`
	DatacenterName               string  `json:"datacenter_name"`
	InfrastructureID             float64 `json:"infrastructure_id,omitempty"`
	UserIDowner                 float64 `json:"user_id_owner,omitempty"`
	UserEmailOwner              string  `json:"user_email_owner,omitempty"`
	InfrastructureTouchUnixtime string  `json:"infrastructure_touch_unixtime,omitempty"`
	InfrastructureServiceStatus string  `json:"infrastructure_touch_unixtime,omitempty"`
	InfrastructureCreatedTimestamp string  `json:"infrastructure_created_timestamp,omitempty"`
	InfrastructureUpdatedTimestamp string  `json:"infrastructure_updated_timestamp,omitempty"`
	InfrastructureChangeID         float64 `json:"infrastructure_change_id,omitempty"`
	InfrastructureDeployID         float64 `json:"infrastructure_deploy_id,omitempty"`
	InfrastructureDesignIsLocked  bool    `json:"infrastructure_design_is_locked,omitempty"`
	InfrastructureOperation InfrastructureOperation `json:"infrastructure_operation,omitempty"`

}

type ShutdownOptions struct {
	Hard_shutdown_after_timeout   bool
	Attempt_soft_shutdown         bool
	Soft_shutdown_timeout_seconds float64
}



func (c *MetalCloudClient) InfrastructureCreate(infrastructure Infrastructure) (*Infrastructure, error) {
	var created_infrastructure Infrastructure

	err := c.rpcClient.CallFor(
		&created_infrastructure,
		"infrastructure_create",
		c.user,
		infrastructure)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &created_infrastructure, nil
}

func (c *MetalCloudClient) InfrastructureEdit(infrastructure_id float64, infrastructure_operation InfrastructureOperation) (*Infrastructure, error) {
	var created_infrastructure Infrastructure

	err := c.rpcClient.CallFor(
		&created_infrastructure,
		"infrastructure_edit",
		infrastructure_id,
		infrastructure_operation)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &created_infrastructure, nil
}


func (c *MetalCloudClient) InfrastructureDelete(infrastructure_id float64) error {
	_, err := c.rpcClient.Call("infrastructure_delete", infrastructure_id)
	if err != nil {
		return err
	}

	return nil
}


func (c *MetalCloudClient) InfrastructureOperationCancel(infrastructure_id float64) error {
	_, err := c.rpcClient.Call(
		"infrastructure_operation_cancel",
		infrastructure_id)

	if err != nil {
		return err
	}

	return nil
}

//TODO: add the rest of the options
func (c *MetalCloudClient) InfrastructureDeploy(infrastructure_id float64, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error {
	_, err := c.rpcClient.Call(
		"infrastructure_deploy",
		infrastructure_id,
		shutdownOptions,
		nil,
		allowDataLoss,
		skipAnsible,
	)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}

func (c *MetalCloudClient) InfrastructureGetByLabel(infrastructureLabel string) (*Infrastructure, error) {
	var infrastructures map[string]Infrastructure

	err := c.rpcClient.CallFor(&infrastructures, "infrastructures", c.user)
	if err != nil || infrastructures == nil {
		// rpc error handling goes here
		// check response.Error.Code, response.Error.Message and optional response.Error.Data
		log.Printf("%s", err)
		return nil, err
	}

	for _, infrastructure := range infrastructures {
		if infrastructure.InfrastructureLabel == infrastructureLabel {
			return &infrastructure, nil
		}
	}
	err = fmt.Errorf("could not find infrastructure with label %s", infrastructureLabel)
	log.Printf("%s", err)

	return nil, err
}

func (c *MetalCloudClient) Infrastructures() (*map[string]Infrastructure, error) {
	var infrastructures map[string]Infrastructure

	err := c.rpcClient.CallFor(&infrastructures, "infrastructures", c.user)
	if err != nil || infrastructures == nil {
		// rpc error handling goes here
		// check response.Error.Code, response.Error.Message and optional response.Error.Data
		log.Printf("%s", err)
		return nil, err
	}

	return &infrastructures, nil
}

func (c *MetalCloudClient) InfrastructureGet(infrastructureID float64) (*Infrastructure, error) {
	var infrastructure Infrastructure

	err := c.rpcClient.CallFor(&infrastructure, "infrastructure_get", infrastructureID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &infrastructure, nil
}


func (c *MetalCloudClient) InfrastructureUserLimits(infrastructureID float64) (*map[string]interface{}, error) {
	var userLimits map[string]interface{}

	err := c.rpcClient.CallFor(&userLimits, "infrastructure_user_limits", infrastructureID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &userLimits, nil
}
