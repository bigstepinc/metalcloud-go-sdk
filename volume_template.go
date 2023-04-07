package metalcloud

import "fmt"

//go:generate go run helper/gen_exports.go

// OperatingSystem describes an OS
type OperatingSystem struct {
	OperatingSystemType         string `json:"operating_system_type,omitempty" yaml:"type,omitempty"`
	OperatingSystemVersion      string `json:"operating_system_version,omitempty" yaml:"version,omitempty"`
	OperatingSystemArchitecture string `json:"operating_system_architecture,omitempty" yaml:"architecture,omitempty"`
}

type NetworkOperatingSystem struct {
	OperatingSystemArchitecture   string `json:"operating_system_architecture,omitempty" yaml:"architecture,omitempty"`
	OperatingSystemDatacenterName string `json:"operating_system_datacenter_name,omitempty" yaml:"datacenter_name,omitempty"`
	OperatingSystemMachine        string `json:"operating_system_machine,omitempty" yaml:"machine,omitempty"`
	OperatingSystemSwitchDriver   string `json:"operating_system_switch_driver,omitempty" yaml:"switchDriver,omitempty"`
	OperatingSystemSwitchRole     string `json:"operating_system_switch_role,omitempty" yaml:"switchRole,omitempty"`
	OperatingSystemVendor         string `json:"operating_system_vendor,omitempty" yaml:"vendor,omitempty"`
	OperatingSystemVersion        string `json:"operating_system_version,omitempty" yaml:"version,omitempty"`
}

// VolumeTemplate describes an OS template
type VolumeTemplate struct {
	VolumeTemplateID                      int                    `json:"volume_template_id,omitempty"`
	VolumeTemplateLabel                   string                 `json:"volume_template_label,omitempty"`
	VolumeTemplateSizeMBytes              int                    `json:"volume_template_size_mbytes,omitempty"`
	VolumeTemplateDisplayName             string                 `json:"volume_template_display_name,omitempty"`
	VolumeTemplateDescription             string                 `json:"volume_template_description,omitempty"`
	VolumeTemplateLocalDiskSupported      bool                   `json:"volume_template_local_supported,omitempty"`
	VolumeTemplateBootMethodsSupported    string                 `json:"volume_template_boot_methods_supported,omitempty"`
	VolumeTemplateBootType                string                 `json:"volume_template_boot_type,omitempty"`
	VolumeTemplateDeprecationStatus       string                 `json:"volume_template_deprecation_status,omitempty"`
	VolumeTemplateRepoURL                 string                 `json:"volume_template_repo_url,omitempty"`
	VolumeTemplateOperatingSystem         OperatingSystem        `json:"volume_template_operating_system,omitempty"`
	VolumeTemplateTags                    []string               `json:"volume_template_tags,omitempty"`
	VolumeTemplateOsBootstrapFunctionName string                 `json:"volume_template_os_bootstrap_function_name,omitempty"`
	VolumeTemplateNetworkOperatingSystem  NetworkOperatingSystem `json:"volume_template_network_operating_system,omitempty"`
	VolumeTemplateVersion                 string                 `json:"volume_template_version,omitempty"`
	VolumeTemplateIsExperimental          bool                   `json:"volume_template_is_experimental,omitempty"`
	VolumeTemplateIsForSwitch             bool                   `json:"volume_template_is_for_switch,omitempty"`
	VolumeTemplateOSReadyMethod           string                 `json:"volume_template_os_ready_method,omitempty"`
}

// VolumeTemplates retrives the list of available templates
func (c *Client) VolumeTemplates() (*map[string]VolumeTemplate, error) {
	var createdObject map[string]VolumeTemplate
	userID := c.GetUserID()

	resp, err := c.rpcClient.Call(
		"volume_templates",
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
		var m = map[string]VolumeTemplate{}
		return &m, nil
	}

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// volumeTemplateGet returns the specified volume template
func (c *Client) volumeTemplateGet(volumeTemplateID id) (*VolumeTemplate, error) {
	var createdObject VolumeTemplate

	if err := checkID(volumeTemplateID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"volume_template_get",
		volumeTemplateID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// volumeTemplateCreate creates a private volume template from a drive
func (c *Client) volumeTemplateCreateFromDrive(driveID id, objVolumeTemplate VolumeTemplate) (*VolumeTemplate, error) {
	var createdObject VolumeTemplate

	if err := checkID(driveID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"volume_template_create_from_drive",
		driveID,
		objVolumeTemplate)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// VolumeTemplateMakePublic makes a template public
func (c *Client) VolumeTemplateMakePublic(volumeTemplateID int, bootstrapFunctionName string) error {
	resp, err := c.rpcClient.Call(
		"volume_template_make_public",
		volumeTemplateID,
		bootstrapFunctionName,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// VolumeTemplateMakePrivate makes a template private
func (c *Client) VolumeTemplateMakePrivate(volumeTemplateID int, userID int) error {
	resp, err := c.rpcClient.Call(
		"volume_template_make_private",
		volumeTemplateID,
		userID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
