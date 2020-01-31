package metalcloud

import (
	"encoding/json"
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//StageDefinition contains a JavaScript file, HTTP request url and options, an AnsibleBundle or an API call template.
type StageDefinition struct {
	StageDefinitionID                     int         `json:"stage_definition_id,omitempty"`
	UserIDOwner                           int         `json:"user_id_owner,omitempty"`
	UserIDAuthenticated                   int         `json:"user_id_authenticated,omitempty"`
	StageDefinitionLabel                  string      `json:"stage_definition_label,omitempty"`
	IconAssetDataURI                      string      `json:"icon_asset_data_uri,omitempty"`
	StageDefinitionTitle                  string      `json:"stage_definition_title,omitempty"`
	StageDefinitionDescription            string      `json:"stage_definition_description,omitempty"`
	StageDefinitionType                   string      `json:"stage_definition_type,omitempty"`
	StageDefinitionVariablesNamesRequired []string    `json:"stage_definition_variable_names_required,omitempty"`
	StageDefinition                       interface{} `json:"stage_definition,omitempty"`
	StageDefinitionCreatedTimestamp       string      `json:"stage_definition_created_timestamp,omitempty"`
	StageDefinitionUpdatedTimestamp       string      `json:"stage_definition_updated_timestamp,omitempty"`
}

//HTTPRequest represents an HTTP request definition compatible with the standard Web Fetch API.
type HTTPRequest struct {
	URL     string              `json:"url,omitempty"`
	Options WebFetchAAPIOptions `json:"options,omitempty"`
}

//WebFetchAAPIOptions represents node-fetch options which is follows the Web API Fetch specification. See https://github.com/node-fetch/node-fetch
type WebFetchAAPIOptions struct {
	Method           string                    `json:"method,omitempty"`
	Redirect         string                    `json:"redirect,omitempty"`
	Follow           int                       `json:"follow,omitempty"`
	Compress         bool                      `json:"compress,omitempty"`
	Timeout          int                       `json:"timeout,omitempty"`
	Size             int                       `json:"size,omitempty"`
	Headers          WebFetchAPIRequestHeaders `json:"headers,omitempty"`
	Body             string                    `json:"body,omitempty"`
	BodyBufferBase64 string                    `json:"bodyBufferBase64,omitempty"`
}

//WebFetchAPIRequestHeaders HTTP request headers. null means undefined (the default for most) so the header will not be included with the request.
type WebFetchAPIRequestHeaders struct {
	Accept             string `json:"Accept,omitempty"`
	UserAgent          string `json:"User-Agent,omitempty"`
	ContentType        string `json:"Content-Type,omitempty"`
	Cookie             string `json:"Cookie,omitempty"`
	Authorization      string `json:"Authorization,omitempty"`
	ProxyAuthorization string `json:"Proxy-Authorization,omitempty"`
	ContentMD5         string `json:"Content-MD5,omitempty"`
}

//AnsibleBundle contains an Ansible project as a single archive file, usually .zip
type AnsibleBundle struct {
	AnsibleBundleArchiveFilename       string `json:"ansible_bundle_archive_filename,omitempty"`
	AnsibleBundleArchiveContentsBase64 string `json:"ansible_bundle_archive_contents_base64,omitempty"`
}

//UnmarshalJSON custom json marshaling
func (s *StageDefinition) UnmarshalJSON(b []byte) error {
	type Alias StageDefinition
	var w Alias
	err := json.Unmarshal(b, &w)
	if err != nil {
		return err
	}

	switch w.StageDefinitionType {
	case "AnsibleBundle":
		var obj AnsibleBundle
		b, err := json.Marshal(w.StageDefinition)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		s.StageDefinition = obj
	case "HTTPRequest":
		var obj HTTPRequest
		b, err := json.Marshal(w.StageDefinition)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		s.StageDefinition = obj
	}
	s.StageDefinitionID = w.StageDefinitionID
	s.UserIDOwner = w.UserIDOwner
	s.UserIDAuthenticated = w.UserIDAuthenticated
	s.StageDefinitionLabel = w.StageDefinitionLabel
	s.IconAssetDataURI = w.IconAssetDataURI
	s.StageDefinitionTitle = w.StageDefinitionTitle
	s.StageDefinitionDescription = w.StageDefinitionDescription
	s.StageDefinitionType = w.StageDefinitionType
	s.StageDefinitionVariablesNamesRequired = w.StageDefinitionVariablesNamesRequired
	s.StageDefinitionCreatedTimestamp = w.StageDefinitionCreatedTimestamp
	s.StageDefinitionUpdatedTimestamp = w.StageDefinitionUpdatedTimestamp

	return err
}

//StageDefinitionCreate creates a stageDefinition
func (c *Client) StageDefinitionCreate(stageDefinition StageDefinition) (*StageDefinition, error) {
	var createdObject StageDefinition

	userID, err := c.UserEmailToUserID(c.user)
	if err != nil {
		return nil, err
	}

	err = c.rpcClient.CallFor(
		&createdObject,
		"stage_definition_create",
		*userID,
		stageDefinition)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//StageDefinitionDelete Permanently destroys a StageDefinition.
func (c *Client) StageDefinitionDelete(stageDefinitionID int) error {

	if err := checkID(stageDefinitionID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("stage_definition_delete", stageDefinitionID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//StageDefinitionUpdate This function allows updating the stageDefinition_usage, stageDefinition_label and stageDefinition_base64 of a StageDefinition
func (c *Client) StageDefinitionUpdate(stageDefinitionID int, stageDefinition StageDefinition) (*StageDefinition, error) {
	var createdObject StageDefinition

	if err := checkID(stageDefinitionID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"stage_definition_update",
		stageDefinitionID,
		stageDefinition)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//StageDefinitionGet returns a StageDefinition specified by nStageDefinitionID. The stageDefinition's protected value is never returned.
func (c *Client) StageDefinitionGet(stageDefinitionID int) (*StageDefinition, error) {

	var createdObject StageDefinition

	if err := checkID(stageDefinitionID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"stage_definition_get",
		stageDefinitionID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//StageDefinitions retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
func (c *Client) StageDefinitions() (*map[string]StageDefinition, error) {

	userID, err := c.UserEmailToUserID(c.user)
	if err != nil {
		return nil, err
	}
	var res *jsonrpc.RPCResponse

	res, err = c.rpcClient.Call(
		"stage_definitions",
		*userID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]StageDefinition{}
		return &m, nil
	}

	var createdObject map[string]StageDefinition

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}
