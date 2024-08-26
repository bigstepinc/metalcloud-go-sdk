package metalcloud

import "fmt"

//go:generate go run helper/gen_exports.go
//go:generate go run helper/docgen.go - $GOFILE ./ OperatingSystem,NetworkOperatingSystem OperatingSystem

// OperatingSystem describes an OS
type OperatingSystem struct {
	// description: operating system type
	// examples:
	//	- Ubuntu
	//  - CentOS
	OperatingSystemType string `json:"operating_system_type,omitempty" yaml:"type,omitempty"`
	// description: operating system type
	// example: 20.04
	OperatingSystemVersion string `json:"operating_system_version,omitempty" yaml:"version,omitempty"`
	//description: architecture of the operating system
	//values:
	// - none
	// - unknown
	// - x86
	// - x86_64
	OperatingSystemArchitecture string `json:"operating_system_architecture,omitempty" yaml:"architecture,omitempty"`
}

type NetworkOperatingSystem struct {
	//description: architecture of the operating system
	//values:
	// - arm
	// - powerpc
	// - x86_64
	OperatingSystemArchitecture string `json:"operating_system_architecture,omitempty" yaml:"architecture,omitempty"`
	//description: The datacenter for this operating system.
	OperatingSystemDatacenterName string `json:"operating_system_datacenter_name,omitempty" yaml:"datacenter_name,omitempty"`
	//description: A switch model number to be matched against an ONIE or POAP request during ztp.
	//example: s5232f_c3538
	OperatingSystemMachine string `json:"operating_system_machine,omitempty" yaml:"machine,omitempty"`
	//description: The switch driver to use.
	OperatingSystemSwitchDriver string `json:"operating_system_switch_driver,omitempty" yaml:"switchDriver,omitempty"`
	//description: The role of a switch
	//values:
	// - leaf
	// - spine
	OperatingSystemSwitchRole string `json:"operating_system_switch_role,omitempty" yaml:"switchRole,omitempty"`
	//description: The vendor of the switch
	//examples:
	// - dellemc
	// - juniper
	OperatingSystemVendor string `json:"operating_system_vendor,omitempty" yaml:"vendor,omitempty"`
	//description: The version of the operating system
	// example: 10.22
	OperatingSystemVersion string `json:"operating_system_version,omitempty" yaml:"version,omitempty"`
}

// VolumeTemplate describes an OS template, stored as an image on disk. Note that this object is deprecated, it will be merged with OS template in the future
type VolumeTemplate struct {
	//description: ID of the object
	VolumeTemplateID int `json:"volume_template_id,omitempty" yaml:"id,omitempty"`
	//description: Label of the object
	VolumeTemplateLabel string `json:"volume_template_label,omitempty" yaml:"label,omitempty"`
	//description: Size of the template on disk
	VolumeTemplateSizeMBytes int `json:"volume_template_size_mbytes,omitempty" yaml:"sizeMbytes,omitempty"`
	//description: Display name of the volume
	VolumeTemplateDisplayName string `json:"volume_template_display_name,omitempty" yaml:"displayName,omitempty"`
	//description: Description of the volume
	VolumeTemplateDescription string `json:"volume_template_description,omitempty" yaml:"description,omitempty"`
	//description: If se tto true this template can be used to install on local disks
	VolumeTemplateLocalDiskSupported bool `json:"volume_template_local_supported,omitempty"`
	//description: If set to true this template can be installed via the OOB network (Virtual Media)
	//values:
	// - pxe_iscsi
	// - local_drives
	VolumeTemplateBootMethodsSupported string `json:"volume_template_boot_methods_supported,omitempty"`
	//description: Controls the type of boot method supported by this template
	//values:
	// - legacy_only
	// - uefi_only
	// - hybrid
	VolumeTemplateBootType string `json:"volume_template_boot_type,omitempty"`
	//description: Controls if this template can be used for new provisioning operations or just expanding of existing clusters
	//values:
	// - not_deprecated
	// - deprecated_allow_expand
	// - deprecated_deny_provision
	VolumeTemplateDeprecationStatus string `json:"volume_template_deprecation_status,omitempty"`
	//description: In some cases additional files are required and these are typically pulled from this URL
	VolumeTemplateRepoURL string `json:"volume_template_repo_url,omitempty"`
	//description: The operating system details
	VolumeTemplateOperatingSystem OperatingSystem `json:"volume_template_operating_system,omitempty"`
	//description: Tags
	VolumeTemplateTags []string `json:"volume_template_tags,omitempty"`
	//description: Defines how the template is initialized.
	//values:
	//  - "provisioner_os_bootstrap_dummy",
	//  - "provisioner_os_bootstrap_generic",
	//  - "provisioner_os_bootstrap_centos",
	//  - "provisioner_os_bootstrap_ubuntu",
	//  - "provisioner_os_cloudinit_prepare_centos",
	//  - "provisioner_os_cloudinit_prepare_rhel",
	//  - "provisioner_os_cloudinit_prepare_ubuntu",
	//  - "provisioner_os_cloudbaseinit_prepare_windows",
	//  - "provisioner_os_cloudinit_prepare_generic_unix",
	//  - "provisioner_os_bootstrap_generic_non_cloudinit
	VolumeTemplateOsBootstrapFunctionName string `json:"volume_template_os_bootstrap_function_name,omitempty"`
	//description: If this is a network operating system, these are the details
	VolumeTemplateNetworkOperatingSystem NetworkOperatingSystem `json:"volume_template_network_operating_system,omitempty"`
	//description: Version of this operating system
	VolumeTemplateVersion string `json:"volume_template_version,omitempty"`
	//description: If set to true, only users with experimental flag set will be able to use it
	VolumeTemplateIsExperimental bool `json:"volume_template_is_experimental,omitempty"`
	//description: If set to true, this template is for switches rather than servers
	VolumeTemplateIsForSwitch bool `json:"volume_template_is_for_switch,omitempty"`
	//description: The method used to determine if the OS is ready to use.
	//values:
	// - wait_for_ssh
	// - wait_for_signal_from_os
	// - wait_for_power_off
	// - wait_for_signal_from_os_and_wait_for_power_off
	VolumeTemplateOSReadyMethod string `json:"volume_template_os_ready_method,omitempty"`
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
