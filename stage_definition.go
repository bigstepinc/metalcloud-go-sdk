package metalcloud

import (
	"encoding/json"
	"fmt"

	"github.com/ybbus/jsonrpc"
)

//StageDefinition contains a JavaScript file, HTTP request url and options, an AnsibleBundle or an API call template.
type StageDefinition struct {
	StageDefinitionID                     int         `json:"stage_definition_id,omitempty" yaml:"id,omitempty"`
	UserIDOwner                           int         `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserIDAuthenticated                   int         `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	StageDefinitionLabel                  string      `json:"stage_definition_label,omitempty" yaml:"label,omitempty"`
	IconAssetDataURI                      string      `json:"icon_asset_data_uri,omitempty" yaml:"iconAssetDataURI,omitempty"`
	StageDefinitionTitle                  string      `json:"stage_definition_title,omitempty" yaml:"title,omitempty"`
	StageDefinitionDescription            string      `json:"stage_definition_description,omitempty" yaml:"description,omitempty"`
	StageDefinitionType                   string      `json:"stage_definition_type,omitempty" yaml:"type,omitempty"`
	StageDefinitionVariablesNamesRequired []string    `json:"stage_definition_variable_names_required,omitempty" yaml:"variableNames,omitempty"`
	StageDefinition                       interface{} `json:"stage_definition,omitempty" yaml:"stageDefinition,omitempty"`
	StageDefinitionCreatedTimestamp       string      `json:"stage_definition_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	StageDefinitionUpdatedTimestamp       string      `json:"stage_definition_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

//HTTPRequest represents an HTTP request definition compatible with the standard Web Fetch API.
type HTTPRequest struct {
	URL     string              `json:"url,omitempty"`
	Options WebFetchAAPIOptions `json:"options,omitempty"`
	Type    string              `json:"type,omitempty"`
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
	Type                               string `json:"type,omitempty"`
}

//WorkflowReference points to a Workflow object via its workflow_id. To be used as a stage definition.
type WorkflowReference struct {
	WorkflowID int    `json:"workflow_id,omitempty"`
	Type       string `json:"type,omitempty"`
}

//SSHExec executes a command on a remote server using the SSH exec functionality (not through a shell).
type SSHExec struct {
	Command   string           `json:"command,omitempty"`
	SSHTarget SSHClientOptions `json:"ssh_target,omitempty"`
	Timeout   int              `json:"timeout,omitempty"`
	Type      string           `json:"type,omitempty"`
}

//SSHClientOptions defines an ssh cnnection such as the host, port, user, password, private keys, etc. All properties support template-like variables; for example, ${{instance_credentials_password}} may be used as value for the password property.
type SSHClientOptions struct {
	Host         string        `json:"host,omitempty"`
	Port         int           `json:"port,omitempty"`
	ForceIPv4    bool          `json:"forceIPv4,omitempty"`
	ForceIPv6    bool          `json:"forceIPv6,omitempty"`
	HostHash     string        `json:"hostHash,omitempty"`
	HashedKey    string        `json:"hashedKey,omitempty"`
	Username     string        `json:"username,omitempty"`
	Password     string        `json:"password,omitempty"`
	PrivateKey   string        `json:"privateKey,omitempty"`
	Passphrase   string        `json:"passphrase,omitempty"`
	ReadyTimeout int           `json:"readyTimeout,omitempty"`
	StrictVendor bool          `json:"strictVendor,omitempty"`
	Algorithms   SSHAlgorithms `json:"algorithms,omitempty"`
	Compress     string        `json:"compress,omitempty"`
}

//SSHAlgorithms defines algorithms that can be used during an ssh session
type SSHAlgorithms struct {
	Kex           []string `json:"kex,omitempty"`
	Cipher        []string `json:"cipher,omitempty"`
	ServerHostKey []string `json:"serverHostKey,omitempty"`
	HMAC          []string `json:"hmac,omitempty"`
	Compress      []string `json:"compress,omitempty"`
}

//Copy defines the source and destination of a SCP operation. The source may be of various types. SCP and HTTP requests are streamed so they are recommended as sources. The destination has to be a SCP resource.
type Copy struct {
	Source                     interface{}         `json:"source,omitempty"`
	Destination                SCPResourceLocation `json:"destination,omitempty"`
	TimeoutMinutes             int                 `json:"timeoutMinutes,omitempty"`
	IfDestinationAlreadyExists string              `json:"ifDestinationAlreadyExists,omitempty"`
	Type                       string              `json:"type,omitempty"`
}

//SCPResourceLocation defines a file path and SSH client connection options for use with Secure Copy Protocol (SCP).
type SCPResourceLocation struct {
	Path      string           `json:"path,omitempty"`
	SSHTarget SSHClientOptions `json:"ssh_target,omitempty"`
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
		obj.Type = "AnsibleBundle"
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
		obj.Type = "HTTPRequest"
		s.StageDefinition = obj

	case "WorkflowReference":
		var obj WorkflowReference
		b, err := json.Marshal(w.StageDefinition)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		obj.Type = "WorkflowReference"
		s.StageDefinition = obj

	case "SSHExec":
		var obj SSHExec
		b, err := json.Marshal(w.StageDefinition)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		obj.Type = "SSHExec"
		s.StageDefinition = obj

	case "Copy":
		var obj Copy
		b, err := json.Marshal(w.StageDefinition)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
		obj.Type = "Copy"
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

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"stage_definition_create",
		userID,
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

	userID := c.GetUserID()

	var res *jsonrpc.RPCResponse
	var err error

	res, err = c.rpcClient.Call(
		"stage_definitions",
		userID)

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

//CreateOrUpdate implements interface Applier
func (s StageDefinition) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var result *StageDefinition
	err = s.Validate()

	if err != nil {
		return err
	}

	if s.StageDefinitionID != 0 {
		result, err = client.StageDefinitionGet(s.StageDefinitionID)
	} else {
		definitions, err := client.StageDefinitions()
		if err != nil {
			return err
		}

		for _, def := range *definitions {
			if def.StageDefinitionLabel == s.StageDefinitionLabel {
				result = &def
			}
		}
	}

	if err != nil {
		_, err = client.StageDefinitionCreate(s)

		if err != nil {
			return err
		}
	} else {
		_, err = client.StageDefinitionUpdate(result.StageDefinitionID, s)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (s StageDefinition) Delete(client MetalCloudClient) error {
	var result *StageDefinition
	var id int
	err := s.Validate()

	if err != nil {
		return err
	}

	if s.StageDefinitionID != 0 {
		id = s.StageDefinitionID
	} else {
		definitions, err := client.StageDefinitions()
		if err != nil {
			return err
		}

		for _, def := range *definitions {
			if def.StageDefinitionLabel == s.StageDefinitionLabel {
				result = &def
			}
		}

		id = result.StageDefinitionID
	}
	err = client.StageDefinitionDelete(id)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (s StageDefinition) Validate() error {
	if s.StageDefinitionID == 0 && s.StageDefinitionLabel == "" {
		return fmt.Errorf("id is required")
	}

	if s.StageDefinitionType == "" {
		return fmt.Errorf("type is required")
	}

	if s.StageDefinitionTitle == "" {
		return fmt.Errorf("title is required")
	}

	return nil
}
