package metalcloud

import "log"
import "fmt"

type InfrastructureOperation struct {
	InfrastructureDeployStatus  	string  `json:"infrastructure_deploy_status, omitempty"`
	InfrastructureDeployType    	string  `json:"infrastructure_deploy_type, omitempty"`
	InfrastructureLabel          	string  `json:"infrastructure_label"`
	InfrastructureSubdomain      	string  `json:"infrastructure_subdomain, omitempty"`
	DatacenterName                	string  `json:"datacenter_name"`
	InfrastructureID                int `json:"infrastructure_id,omitempty"`
	UserIDOwner                 	int `json:"user_id_owner,omitempty"`
	InfrastructureUpdatedTimestamp 	string  `json:"infrastructure_updated_timestamp,omitempty"`
	InfrastructureChangeID         	int `json:"infrastructure_change_id,omitempty"`
	InfrastructureDeployID         	int `json:"infrastructure_deploy_id,omitempty"`
}

type Infrastructure struct {
	InfrastructureLabel          string  `json:"infrastructure_label"`
	InfrastructureSubdomain      string  `json:"infrastructure_subdomain, omitempty"`
	DatacenterName               string  `json:"datacenter_name"`
	InfrastructureID             int `json:"infrastructure_id,omitempty"`
	UserIDowner                  int `json:"user_id_owner,omitempty"`
	UserEmailOwner               string  `json:"user_email_owner,omitempty"`
	InfrastructureTouchUnixtime  string  `json:"infrastructure_touch_unixtime,omitempty"`
	InfrastructureServiceStatus  string  `json:"infrastructure_touch_unixtime,omitempty"`
	InfrastructureCreatedTimestamp string  `json:"infrastructure_created_timestamp,omitempty"`
	InfrastructureUpdatedTimestamp string  `json:"infrastructure_updated_timestamp,omitempty"`
	InfrastructureChangeID         int `json:"infrastructure_change_id,omitempty"`
	InfrastructureDeployID         int `json:"infrastructure_deploy_id,omitempty"`
	InfrastructureDesignIsLocked   bool    `json:"infrastructure_design_is_locked,omitempty"`
	InfrastructureOperation 	   InfrastructureOperation `json:"infrastructure_operation,omitempty"`

}

type ShutdownOptions struct {
	Hard_shutdown_after_timeout   bool
	Attempt_soft_shutdown         bool
	Soft_shutdown_timeout_seconds int
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

func (c *MetalCloudClient) InfrastructureEdit(infrastructure_id int, infrastructure_operation InfrastructureOperation) (*Infrastructure, error) {
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


func (c *MetalCloudClient) InfrastructureDelete(infrastructureID int) error {
	_, err := c.rpcClient.Call("infrastructure_delete", infrastructureID)
	if err != nil {
		return err
	}

	return nil
}


func (c *MetalCloudClient) InfrastructureOperationCancel(infrastructureID int) error {
	_, err := c.rpcClient.Call(
		"infrastructure_operation_cancel",
		infrastructureID)

	if err != nil {
		return err
	}

	return nil
}

//TODO: add the rest of the options
func (c *MetalCloudClient) InfrastructureDeploy(infrastructureID int, shutdownOptions ShutdownOptions, allowDataLoss bool, skipAnsible bool) error {
	_, err := c.rpcClient.Call(
		"infrastructure_deploy",
		infrastructureID,
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
	res, err := c.rpcClient.Call(
		"infrastructures",
		c.user)
	
	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]Infrastructure{}
		return &m, nil
	}

	var created_object map[string]Infrastructure
	
	err2 := res.GetObject(&created_object)
	if err2 != nil {
			return nil, err2
	}

	return &created_object, nil
}

func (c *MetalCloudClient) InfrastructureGet(infrastructureID int) (*Infrastructure, error) {
	var infrastructure Infrastructure

	err := c.rpcClient.CallFor(&infrastructure, "infrastructure_get", infrastructureID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &infrastructure, nil
}


func (c *MetalCloudClient) InfrastructureUserLimits(infrastructureID int) (*map[string]interface{}, error) {
	var userLimits map[string]interface{}

	err := c.rpcClient.CallFor(&userLimits, "infrastructure_user_limits", infrastructureID)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	return &userLimits, nil
}
