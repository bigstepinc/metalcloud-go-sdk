package metalcloud

//go:generate go run helper/gen_exports.go

import "fmt"

//searchResultResponseWrapperForAFC describes a search result for AFC
type searchResultResponseWrapperForAFC struct {
	DurationMilliseconds int               `json:"duration_millisecnds,omitempty"`
	Rows                 []AFCSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string        `json:"rows_order,omitempty"`
	RowsTotal            int               `json:"rows_total,omitempty"`
}

//AFCSearchResult Represents an AFC search result
type AFCSearchResult struct {
	AFCID                     int    `json:"afc_id,omitempty" yaml:"AFCID,omitempty"`
	AFCType                   string `json:"afc_type,omitempty" yaml:"afcType,omitempty"`
	AFCResponseJSON           string `json:"afc_response_json,omitempty" yaml:"AFCResponseJSON,omitempty"`
	AFCExceptionJSON          string `json:"afc_exception_json,omitempty" yaml:"AFCExceptionJSON,omitempty"`
	InfrastructureID          int    `json:"infrastructure_id,omitempty" yaml:"infrastructureID,omitempty"`
	AFCIPAddressHumanReadable string `json:"afc_ip_address_human_readable,omitempty" yaml:"AFCIPAddressHumanReadable,omitempty"`
	AFCIsBlocked              int    `json:"afc_id_blocked,omitempty" yaml:"AFCIsBlocked,omitempty"`
	AFCIsBlockedBy            int    `json:"afc_id_blocked_by,omitempty" yaml:"AFCIsBlockedBy,omitempty"`
	AFCGroupID                int    `json:"afc_group_id,omitempty" yaml:"AFCGroupID,omitempty"`
	DatacenterName            string `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
	InstanceID                int    `json:"instance_id,omitempty" yaml:"instanceID,omitempty"`
	ServerID                  int    `json:"server_id,omitempty" yaml:"serverID,omitempty"`
	AFCExecuteEngine          string `json:"afc_execute_engine,omitempty" yaml:"AFCExecuteEngine,omitempty"`
	AFCCallCount              int    `json:"afc_call_count,omitempty" yaml:"AFCCallCount,omitempty"`
	AFCRetryMax               int    `json:"afc_retry_max,omitempty" yaml:"AFCRetryMax,omitempty"`
	AFCRetryCount             int    `json:"afc_retry_count,omitempty" yaml:"AFCRetryCount,omitempty"`
	AFCStartTimestamp         string `json:"afc_start_timestamp,omitempty" yaml:"AFCStartTimestamp,omitempty"`
	AFCRetryMinSec            int    `json:"afc_retry_min_sec,omitempty" yaml:"AFCRetryMinSec,omitempty"`
	AFCDurationMs             int    `json:"afc_duration_milliseconds,omitempty" yaml:"AFCDurationMs,omitempty"`
	AFCUpdatedTimestamp       string `json:"afc_updated_timestamp,omitempty" yaml:"AFCUpdatedTimestamp,omitempty"`
	AFCCreatedTimestamp       string `json:"afc_created_timestamp,omitempty" yaml:"AFCCreatedTimestamp,omitempty"`
	AFCFunctionName           string `json:"afc_function_name,omitempty" yaml:"AFCFunctionName,omitempty"`
	AFCParamsJSON             string `json:"afc_params_json,omitempty" yaml:"AFCParamsJSON,omitempty"`
	AFCStatus                 string `json:"afc_status,omitempty" yaml:"AFCStatus,omitempty"`
}

//AFC Represents an AFC result
type AFC struct {
	AFCID                     int    `json:"afc_id,omitempty" yaml:"AFCID,omitempty"`
	AFCType                   string `json:"afc_type,omitempty" yaml:"afcType,omitempty"`
	AFCResponseJSON           string `json:"afc_response_json,omitempty" yaml:"AFCResponseJSON,omitempty"`
	AFCExceptionJSON          string `json:"afc_exception_json,omitempty" yaml:"AFCExceptionJSON,omitempty"`
	AFCInfrastructureID       int    `json:"infrastructure_id,omitempty" yaml:"AFCInfrastructureID,omitempty"`
	AFCIPAddressHumanReadable string `json:"afc_ip_address_human_readable,omitempty" yaml:"AFCIPAddressHumanReadable,omitempty"`
	AFCIsBlocked              int    `json:"afc_id_blocked,omitempty" yaml:"AFCIsBlocked,omitempty"`
	AFCIsBlockedBy            int    `json:"afc_id_blocked_by,omitempty" yaml:"AFCIsBlockedBy,omitempty"`
	AFCGroupID                int    `json:"afc_group_id,omitempty" yaml:"AFCGroupID,omitempty"`
	DatacenterName            string `json:"datacenter_name,omitempty" yaml:"datacenterName,omitempty"`
	InstanceID                int    `json:"instance_id,omitempty" yaml:"instanceID,omitempty"`
	ServerID                  int    `json:"server_id,omitempty" yaml:"serverID,omitempty"`
	AFCExecuteEngine          string `json:"afc_execute_engine,omitempty" yaml:"AFCExecuteEngine,omitempty"`
	AFCCallCount              int    `json:"afc_call_count,omitempty" yaml:"AFCCallCount,omitempty"`
	AFCRetryMax               int    `json:"afc_retry_max,omitempty" yaml:"AFCRetryMax,omitempty"`
	AFCRetryCount             int    `json:"afc_retry_count,omitempty" yaml:"AFCRetryCount,omitempty"`
	AFCStartTimestamp         string `json:"afc_start_timestamp,omitempty" yaml:"AFCStartTimestamp,omitempty"`
	AFCRetryMinSec            int    `json:"afc_retry_min_sec,omitempty" yaml:"AFCRetryMinSec,omitempty"`
	AFCDurationMs             int    `json:"afc_duration_milliseconds,omitempty" yaml:"AFCDurationMs,omitempty"`
	AFCUpdatedTimestamp       string `json:"afc_updated_timestamp,omitempty" yaml:"AFCUpdatedTimestamp,omitempty"`
	AFCCreatedTimestamp       string `json:"afc_created_timestamp,omitempty" yaml:"AFCCreatedTimestamp,omitempty"`
	AFCFunctionName           string `json:"afc_function_name,omitempty" yaml:"AFCFunctionName,omitempty"`
	AFCParamsJSON             string `json:"afc_params_json,omitempty" yaml:"AFCParamsJSON,omitempty"`
	AFCStatus                 string `json:"afc_status,omitempty" yaml:"AFCStatus,omitempty"`
}

//AFCSearch searches for AFCs
func (c *Client) AFCSearch(filter string, page_start int, page_end int) (*[]AFCSearchResult, error) {

	tables := []string{"_afc_queue"}
	columns := map[string][]string{
		"_afc_queue": {
			"afc_type",
			"afc_response_json",
			"afc_exception_json",
			"afc_ip_address_human_readable",
			"afc_id_blocked",
			"afc_id_blocked_by",
			"infrastructure_id",
			"afc_group_id",
			"datacenter_name",
			"instance_id",
			"server_id",
			"afc_execute_engine",
			"afc_id",
			"afc_call_count",
			"afc_retry_max",
			"afc_retry_count",
			"afc_start_timestamp",
			"afc_retry_min_sec",
			"afc_duration_milliseconds",
			"afc_updated_timestamp",
			"afc_created_timestamp",
			"afc_function_name",
			"afc_params_json",
			"afc_status",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	var createdObject map[string]searchResultResponseWrapperForAFC
	sort := [][]string{{"afc_status", "ASC"}, {"afc_created_timestamp", "DESC"}}
	pagination := []int{page_start, page_end}

	resp, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType,
		sort,
		pagination,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]searchResultResponseWrapperForAFC{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []AFCSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil
}

//AFCGet Returns an AFC
func (c *Client) AFCGet(afcID int) (*AFC, error) {

	var createdObject AFC

	err := c.rpcClient.CallFor(
		&createdObject,
		"afc_get",
		afcID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//AFCRetryCall Retries an AFC
func (c *Client) AFCRetryCall(afcID int) error {

	resp, err := c.rpcClient.Call("afc_retry_call", afcID, true)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//AFCRetryCall Skips an AFC
func (c *Client) AFCSkip(afcID int) error {

	resp, err := c.rpcClient.Call("afc_skip", afcID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//AFCDelete Skips an AFC
func (c *Client) AFCDelete(afcID int) error {

	resp, err := c.rpcClient.Call("afc_delete", afcID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//AFCMarkForDeath Tries to kill an AFC
func (c *Client) AFCMarkForDeath(afcID int, typeOfMark string) error {

	resp, err := c.rpcClient.Call("afc_mark_for_death", afcID, typeOfMark)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
