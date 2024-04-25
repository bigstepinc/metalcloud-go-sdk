//go:generate go run helper/docgen.go - $GOFILE ./ StageDefinition,HTTPRequest,WebFetchAAPIOptions,WebFetchAPIRequestHeaders,AnsibleBundle,WorkflowReference,SSHExec,SSHClientOptions,SSHAlgorithms,Copy,SCPResourceLocation StageDefinition

package metalcloud

import (
	"encoding/json"
	"fmt"
)

// StageDefinition Also called a workflow task contains a JavaScript file, HTTP request url and options, an AnsibleBundle or an API call template.
type StageDefinition struct {
	//description: Id of the object
	StageDefinitionID int `json:"stage_definition_id,omitempty" yaml:"id,omitempty"`
	//description: Id of the owner
	UserIDOwner int `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	//description: Id of the user who created the task. Internal.
	UserIDAuthenticated int `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	//description: Label of the task
	StageDefinitionLabel string `json:"stage_definition_label,omitempty" yaml:"label,omitempty"`
	//description: Internal
	IconAssetDataURI string `json:"icon_asset_data_uri,omitempty" yaml:"iconAssetDataURI,omitempty"`
	//description: Title (Name) of the task
	StageDefinitionTitle string `json:"stage_definition_title,omitempty" yaml:"title,omitempty"`
	//description: Description of the task
	StageDefinitionDescription string `json:"stage_definition_description,omitempty" yaml:"description,omitempty"`
	//description: Type of task
	//values:
	// - AnsibleBundle
	// - JavaScript
	// - APICall
	// - HTTPRequest
	// - Copy
	// - WorkflowReference
	StageDefinitionType string `json:"stage_definition_type,omitempty" yaml:"type,omitempty"`
	//description: list of variable names required when executing this task.
	StageDefinitionVariablesNamesRequired []string `json:"stage_definition_variable_names_required,omitempty" yaml:"variableNames,omitempty"`
	//description: More details depending on the type of task
	// values:
	// - HTTPRequest
	// - AnsibleBundle
	// - WorkflowReference
	// - SSHExec
	// - Copy
	StageDefinition interface{} `json:"stage_definition,omitempty" yaml:"stageDefinition,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the stage definition record was created. Readonly.
	//example: 2013-11-29T13:00:01Z
	StageDefinitionCreatedTimestamp string `json:"stage_definition_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	//description: ISO 8601 timestamp which holds the date and time when the stage definition record was updated. Readonly.
	//example: 2013-11-29T13:00:01Z
	StageDefinitionUpdatedTimestamp string `json:"stage_definition_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	//description: Internal. Readonly.
	//values:
	// - local
	// - global
	StageDefinitionContext string `json:"stage_definition_context,omitempty" yaml:"context,omitempty"`
	//description: If set to 1 it will be added to all infrastructures at the pre-deploy stage
	StageDefinitionAutomaticallyAddedToPreDeploys int `json:"stage_definition_automatically_added_to_pre_deploys,omitempty" yaml:"automaticallyAddedToPreDeploys,omitempty"`
	//description: If set to 1 it will be added to all infrastructures at the post-deploy stage
	StageDefinitionAutomaticallyAddedToPostDeploys int `json:"stage_definition_automatically_added_to_post_deploys,omitempty" yaml:"automaticallyAddedToPostDeploys,omitempty"`
}

// HTTPRequest represents an HTTP request definition compatible with the standard Web Fetch API.
type HTTPRequest struct {
	//description: URL to call
	URL string `json:"url,omitempty" yaml:"url,omitempty"`
	//description: Options
	Options WebFetchAAPIOptions `json:"options,omitempty" yaml:"options,omitempty"`
	//description: Internal
	Type string `json:"type,omitempty" yaml:"url,omitempty"`
}

// WebFetchAAPIOptions represents node-fetch options which is follows the Web API Fetch specification. See https://github.com/node-fetch/node-fetch
type WebFetchAAPIOptions struct {
	//description: HTTP Method to use
	//example: "get"
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
	//description: Defaults to 'follow'. Set to `manual` to extract redirect headers, `error` to reject redirect
	//values:
	// - 'follow'
	// - 'manual'
	// - 'error'
	Redirect string `json:"redirect,omitempty" yaml:"redirect,omitempty"`
	//description: Maximum redirect count. 0 to not follow redirect. Defaults to 20
	Follow int `json:"follow,omitempty" yaml:"follow,omitempty"`
	//description: If se to true it will support compression
	Compress bool `json:"compress" yaml:"compress,omitempty"`
	//description: Timeout setting in seconds
	Timeout int `json:"timeout,omitempty"  yaml:"timeout,omitempty"`
	//description: Maximum response body size in bytes. 0 to disable (default)
	Size int `json:"size,omitempty" yaml:"size,omitempty"`
	//description: Request headers
	Headers WebFetchAPIRequestHeaders `json:"headers,omitempty" yaml:"headers,omitempty"`
	//description: Body of the request
	Body string `json:"body,omitempty" yaml:"body,omitempty"`
	//description: Body of the requested encoded base64
	BodyBufferBase64 string `json:"bodyBufferBase64,omitempty" yaml:"bodyBase64,omitempty"`
}

// WebFetchAPIRequestHeaders HTTP request headers. null means undefined (the default for most) so the header will not be included with the request.
type WebFetchAPIRequestHeaders struct {
	//description: accept header
	Accept string `json:"Accept,omitempty" yaml:"accept,omitempty"`
	//description: user-agent header
	UserAgent string `json:"User-Agent,omitempty" yaml:"userAgent,omitempty"`
	//description: content-type header
	ContentType string `json:"Content-Type,omitempty" yaml:"contentType,omitempty"`
	//description: cookie header
	Cookie string `json:"Cookie,omitempty" yaml:"cookie,omitempty"`
	//description: authorization header
	Authorization string `json:"Authorization,omitempty" yaml:"authorization,omitempty"`
	//description: proxy-authorization header
	ProxyAuthorization string `json:"Proxy-Authorization,omitempty" yaml:"proxyAuthorization,omitempty"`
	//description: content-md5 header
	ContentMD5 string `json:"Content-MD5,omitempty" yaml:"contentMD5,omitempty"`
}

// AnsibleBundle contains an Ansible project as a single archive file, usually .zip
type AnsibleBundle struct {
	//description: file name
	AnsibleBundleArchiveFilename string `json:"ansible_bundle_archive_filename,omitempty"  yaml:"filename,omitempty"`
	//description: Content in base64
	AnsibleBundleArchiveContentsBase64 string `json:"ansible_bundle_archive_contents_base64,omitempty" yaml:"contentsBase64,omitempty"`
	//description: internal
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// WorkflowReference points to a Workflow object via its workflow_id. To be used as a stage definition.
type WorkflowReference struct {
	//description: id of the workflow to reference
	WorkflowID int `json:"workflow_id,omitempty" yaml:"workflowId,omitempty"`
	//description: internal
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// SSHExec executes a command on a remote server using the SSH exec functionality (not through a shell).
type SSHExec struct {
	//description: Command to execute
	Command string `json:"command,omitempty" yaml:"command,omitempty"`
	//description: Details on how to connect to the target system
	SSHTarget SSHClientOptions `json:"ssh_target,omitempty" yaml:"target,omitempty"`
	//description: Timeout in seconds
	Timeout int `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	//description: Internal
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// SSHClientOptions defines an ssh cnnection such as the host, port, user, password, private keys, etc. All properties support template-like variables; for example, ${{instance_credentials_password}} may be used as value for the password property.
type SSHClientOptions struct {
	//description: Host
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
	//description: Port
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
	//description: Force the use of ipv4
	ForceIPv4 bool `json:"forceIPv4,omitempty" yaml:"forceIPv4,omitempty"`
	//description: Force the use of ipv6
	ForceIPv6 bool `json:"forceIPv6,omitempty" yaml:"forceIPv6,omitempty"`
	//description: Hash of the host
	HostHash string `json:"hostHash,omitempty" yaml:"hostHash,omitempty"`
	//description: Hash key of the host
	HashedKey string `json:"hashedKey,omitempty" yaml:"hashedKey,omitempty"`
	//description: Username
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	//description: Password
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	//description: Private key to use
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
	//description: Private key passphrase to use
	Passphrase string `json:"passphrase,omitempty" yaml:"passphrase,omitempty"`
	//description: Timeout in seconds
	ReadyTimeout int `json:"readyTimeout,omitempty" yaml:"readyTimeout,omitempty"`
	//description: Internal
	StrictVendor bool `json:"strictVendor,omitempty" yaml:"strictVendor,omitempty"`
	//description: SSH Algorithms to use
	Algorithms SSHAlgorithms `json:"algorithms,omitempty" yaml:"algorithms,omitempty"`
	//description: Internal
	Compress interface{} `json:"compress,omitempty" yaml:"compress,omitempty"`
}

// SSHAlgorithms defines algorithms that can be used during an ssh session
type SSHAlgorithms struct {
	//description: Kex
	Kex []string `json:"kex,omitempty" yaml:"kex,omitempty"`
	//description: Ciphers accepted
	Cipher []string `json:"cipher,omitempty" yaml:"cipher,omitempty"`
	//description: Host keys accepted
	ServerHostKey []string `json:"serverHostKey,omitempty" yaml:"serverHostKey,omitempty"`
	//description: HMAC accepted
	HMAC []string `json:"hmac,omitempty" yaml:"hmac,omitempty"`
	//description: Compress
	Compress []string `json:"compress,omitempty" yaml:"compress,omitempty"`
}

// Copy defines the source and destination of a SCP operation. The source may be of various types. SCP and HTTP requests are streamed so they are recommended as sources. The destination has to be a SCP resource.
type Copy struct {
	//description: the source of the file
	//values:
	// - SCPResourceLocation
	Source interface{} `json:"source,omitempty" yaml:"source,omitempty"`
	//description: the destination
	Destination SCPResourceLocation `json:"destination,omitempty" yaml:"destination,omitempty"`
	//description: timeout in minutes
	TimeoutMinutes int `json:"timeoutMinutes,omitempty" yaml:"timeoutMinutes,omitempty"`
	//description: What to do if file exists at the destination. Defaults to 'overwrite'
	//values:
	// - error
	// - overwrite
	// - errorIfNotSameSize
	// - overwriteIfNotSameSize
	// - ignore
	IfDestinationAlreadyExists string `json:"ifDestinationAlreadyExists,omitempty" yaml:"ifDestinationAlreadyExists,omitempty"`
	//description: internal
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// SCPResourceLocation defines a file path and SSH client connection options for use with Secure Copy Protocol (SCP).
type SCPResourceLocation struct {
	//description: path
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	//description: SSH Options for the target
	SSHTarget SSHClientOptions `json:"ssh_target,omitempty" yaml:"target,omitempty"`
}

// UnmarshalJSON custom json marshaling
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
	s.StageDefinitionContext = w.StageDefinitionContext
	s.StageDefinitionAutomaticallyAddedToPreDeploys = w.StageDefinitionAutomaticallyAddedToPreDeploys
	s.StageDefinitionAutomaticallyAddedToPostDeploys = w.StageDefinitionAutomaticallyAddedToPostDeploys

	return err
}

// StageDefinitionCreate creates a stageDefinition
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

// StageDefinitionDelete Permanently destroys a StageDefinition.
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

// StageDefinitionUpdate This function allows updating the stageDefinition_usage, stageDefinition_label and stageDefinition_base64 of a StageDefinition
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

// StageDefinitionGet returns a StageDefinition specified by nStageDefinitionID. The stageDefinition's protected value is never returned.
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

// StageDefinitions retrieves a list of all the StageDefinition objects which a specified User is allowed to see through ownership or delegation. The stageDefinition objects never return the actual protected stageDefinition value.
func (c *Client) StageDefinitions() (*map[string]StageDefinition, error) {

	userID := c.GetUserID()
	var createdObject map[string]StageDefinition

	resp, err := c.rpcClient.Call(
		"stage_definitions",
		userID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]StageDefinition{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// CreateOrUpdate implements interface Applier
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

// Delete implements interface Applier
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

// Validate implements interface Applier
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
