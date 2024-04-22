package metalcloud

//go:generate go run helper/docgen.go - $GOFILE ./ OSTemplate,OSTemplateCredentials OSTemplate

import (
	"fmt"
	"strings"
)

// OSTemplate A template can be created based on a drive and it has the same characteristics and holds the same information as the parent drive.
type OSTemplate struct {
	//description: The id of the object
	VolumeTemplateID int `json:"volume_template_id,omitempty" yaml:"id,omitempty"`
	//description: The label of the object
	VolumeTemplateLabel string `json:"volume_template_label,omitempty" yaml:"label,omitempty"`
	//description: The display name of the object
	VolumeTemplateDisplayName string `json:"volume_template_display_name,omitempty" yaml:"name,omitempty"`
	//description: If this is an image type os template, the size of the image
	VolumeTemplateSizeMBytes int `json:"volume_template_size_mbytes,omitempty" yaml:"sizeMBytes,omitempty"`
	//description: If set to true if this template can be installed to local disks
	VolumeTemplateLocalDiskSupported bool `json:"volume_template_local_disk_supported,omitempty" yaml:"localDisk,omitempty"`
	//description: If set to true this is an unattended install type template. If set to false this will be an image type template.
	VolumeTemplateIsOSTemplate bool `json:"volume_template_is_os_template,omitempty" yaml:"isOsTemplate,omitempty"`
	//description: If set to true this is a template for switches rather than servers.
	VolumeTemplateIsForSwitch bool `json:"volume_template_is_for_switch,omitempty"`
	//description: If set to true this is a template that requires a build using the image builder to package assets into the ISO
	VolumeTemplateImageBuildRequired bool `json:"volume_template_image_build_required,omitempty" yaml:"isImageBuildRequired,omitempty"`
	//description: If set to true this template can be installed via the OOB network (Virtual Media)
	VolumeTemplateProvisionViaOOB bool `json:"volume_template_provision_via_oob,omitempty" yaml:"provisionViaOOB,omitempty"`
	//description: If set to true this template can be installed via the OOB network (Virtual Media)
	//values:
	// - pxe_iscsi
	// - local_drives
	VolumeTemplateBootMethodsSupported string `json:"volume_template_boot_methods_supported,omitempty" yaml:"bootMethods,omitempty"`
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
	//description: Controls the type of boot method supported by this template
	//values:
	// - legacy_only
	// - uefi_only
	// - hybrid
	VolumeTemplateBootType string `json:"volume_template_boot_type,omitempty" yaml:"bootType,omitempty"`
	//description: A description of the template
	VolumeTemplateDescription string `json:"volume_template_description,omitempty" yaml:"description,omitempty"`
	// description: ISO 8601 timestamp which holds the date and time when the template was created. Readonly.
	// example: 2013-11-29T13:00:01Z
	VolumeTemplateCreatedTimestamp string `json:"volume_template_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	// description: ISO 8601 timestamp which holds the date and time when the template was updated. Readonly.
	// example: 2013-11-29T13:00:01Z
	VolumeTemplateUpdatedTimestamp string `json:"volume_template_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	// description: The user that owns this template
	UserID int `json:"user_id,omitempty" yaml:"userID,omitempty"`
	// description: The operating system details
	VolumeTemplateOperatingSystem *OperatingSystem `json:"volume_template_operating_system,omitempty" yaml:"os,omitempty"`
	//description: In some cases additional files are required and these are typically pulled from this URL
	VolumeTemplateRepoURL string `json:"volume_template_repo_url,omitempty" yaml:"repoURL,omitempty"`
	//description: Controls if this template can be used for new provisioning operations or just expanding of existing clusters
	//values:
	// - not_deprecated
	// - deprecated_allow_expand
	// - deprecated_deny_provision
	VolumeTemplateDeprecationStatus string `json:"volume_template_deprecation_status,omitempty" yaml:"deprecationStatus,omitempty"`
	//description: Credentials of the template
	OSTemplateCredentials *OSTemplateCredentials `json:"os_template_credentials,omitempty" yaml:"credentials,omitempty"`
	//description: Tags
	VolumeTemplateTags []string `json:"volume_template_tags,omitempty" yaml:"tags,omitempty"`
	//description: The ID of the bootloader asset to use during local install
	OSAssetBootloaderLocalInstall int `json:"os_asset_id_bootloader_local_install" yaml:"OSAssetIDBootloaderLocalInstall"`
	//description: The ID of the bootloader asset to use during initial boot (such as iPXE)
	OSAssetBootloaderOSBoot int `json:"os_asset_id_bootloader_os_boot" yaml:"OSAssetIDBootloaderOSBoot"`
	//description: Variables this template requires.
	//example: ['custom_var1','custom_var2']
	VolumeTemplateVariablesJSON string `json:"volume_template_variables_json,omitempty" yaml:"variablesJSON,omitempty"`
	//description: If this is an network operating system, here are the details
	VolumeTemplateNetworkOperatingSystem *NetworkOperatingSystem `json:"volume_template_network_operating_system,omitempty" yaml:"networkOS,omitempty"`
	//description: Version of the operating system
	//example: 20.04
	VolumeTemplateVersion string `json:"volume_template_version,omitempty"`
	//description: The method to determine if the OS is ready to use.
	//values:
	// - wait_for_ssh
	// - wait_for_signal_from_os
	// - wait_for_power_off
	// - wait_for_signal_from_os_and_wait_for_power_off
	VolumeTemplateOSReadyMethod string `json:"volume_template_os_ready_method,omitempty"`
}

// OSTemplateCredentials holds information needed to connect to an OS installed by an OSTemplate.
type OSTemplateCredentials struct {
	//description: The initial user
	//example: root
	OSTemplateInitialUser string `json:"os_template_initial_user,omitempty" yaml:"initialUser,omitempty"`
	//description: The initial password. Readonly
	//example: encrypted value
	OSTemplateInitialPasswordEncrypted string `json:"os_template_initial_password_encrypted,omitempty" yaml:"initialPasswordEncrypted,omitempty"`
	//description: The initial password. Write-only
	OSTemplateInitialPassword string `json:"os_template_initial_password,omitempty" yaml:"initialPassword,omitempty"`
	//description: The port to use for SSH
	OSTemplateInitialSSHPort int `json:"os_template_initial_ssh_port,omitempty" yaml:"initialSSHPort,omitempty"`
	//description: If set, and if the template supports it the password will be changed after deploy (this password will be used only during deploy).
	OSTemplateChangePasswordAfterDeploy bool `json:"os_template_change_password_after_deploy,omitempty" yaml:"changePasswordAfterDeploy,omitempty"`
	//description: If set, the password will be automatically generated
	OSTemplateUseAutogeneratedInitialPassword bool `json:"os_template_use_autogenerated_initial_password,omitempty" yaml:"useAutogeneratedInitialPassword,omitempty"`
}

// OSTemplateOSAssetData holds asset-template information
type OSTemplateOSAssetData struct {
	//description: The asset to use
	OSAsset *OSAsset `json:"os_asset,omitempty"`
	//description: The file path for this asset
	OSAssetFilePath string `json:"os_asset_file_path,omitempty"`
	// description: ISO 8601 timestamp which holds the date and time when the asset was updated. Readonly.
	// example: 2013-11-29T13:00:01Z
	OSTemplateOSAssetUpdatedTimestamp string `json:"volume_template_os_asset_updated_timestamp,omitempty"`
	// description: variables to use during deployment
	OSTemplateOSAssetVariablesJSON string `json:"volume_template_os_asset_variables_json,omitempty"`
}

// OSTemplateCreate creates a osTemplate object
func (c *Client) OSTemplateCreate(osTemplate OSTemplate) (*OSTemplate, error) {
	var createdObject OSTemplate

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_template_create",
		userID,
		osTemplate)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// OSTemplateDelete permanently destroys a OSTemplate.
func (c *Client) OSTemplateDelete(osTemplateID int) error {

	if err := checkID(osTemplateID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("os_template_delete", osTemplateID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// OSTemplateUpdate updates a osTemplate
func (c *Client) OSTemplateUpdate(osTemplateID int, osTemplate OSTemplate) (*OSTemplate, error) {
	var createdObject OSTemplate

	if err := checkID(osTemplateID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_template_update",
		osTemplateID,
		osTemplate)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

// OSTemplateGet returns a OSTemplate specified by nOSTemplateID. The OSTemplate's protected value is never returned.
func (c *Client) OSTemplateGet(osTemplateID int, decryptPasswd bool) (*OSTemplate, error) {

	var createdObject OSTemplate

	if err := checkID(osTemplateID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_template_get",
		osTemplateID)

	if err != nil {

		return nil, err
	}

	if decryptPasswd && createdObject.OSTemplateCredentials != nil {
		passwdComponents := strings.Split(createdObject.OSTemplateCredentials.OSTemplateInitialPassword, ":")
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

				createdObject.OSTemplateCredentials.OSTemplateInitialPassword = passwd
			}
		}
	}

	return &createdObject, nil
}

// OSTemplates retrieves a list of all the OSTemplate objects which a specified User is allowed to see through ownership or delegation. The OSTemplate objects never return the actual protected OSTemplate value.
func (c *Client) OSTemplates() (*map[string]OSTemplate, error) {

	userID := c.GetUserID()

	resp, err := c.rpcClient.Call(
		"os_templates",
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
		var m = map[string]OSTemplate{}
		return &m, nil
	}
	var createdObject map[string]OSTemplate

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// OSTemplatesNetwork retrieves a list of all the network OSTemplate objects which a specified User is allowed to see through ownership or delegation.
func (c *Client) OSTemplatesNetwork() (*map[string]OSTemplate, error) {

	userID := c.GetUserID()

	resp, err := c.rpcClient.Call(
		"os_templates_network",
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
		var m = map[string]OSTemplate{}
		return &m, nil
	}
	var createdObject map[string]OSTemplate

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// OSTemplateOSAssets returns the OSAssets assigned to an OSTemplate.
func (c *Client) OSTemplateOSAssets(osTemplateID int) (*map[string]OSTemplateOSAssetData, error) {
	if err := checkID(osTemplateID); err != nil {
		return nil, err
	}

	resp, err := c.rpcClient.Call(
		"os_template_os_assets",
		osTemplateID,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]OSTemplateOSAssetData{}
		return &m, nil
	}

	var createdObject map[string]OSTemplateOSAssetData

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

// OSTemplateAddOSAsset adds an asset to a template
func (c *Client) OSTemplateAddOSAsset(osTemplateID int, osAssetID int, path string, variablesJSON string) error {

	// var cond bool
	resp, err := c.rpcClient.Call(
		"os_template_add_os_asset",
		osTemplateID,
		osAssetID,
		path,
		variablesJSON)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// OSTemplateRemoveOSAsset removes an asset from a template
func (c *Client) OSTemplateRemoveOSAsset(osTemplateID int, osAssetID int) error {

	resp, err := c.rpcClient.Call(
		"os_template_remove_os_asset",
		osTemplateID,
		osAssetID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// OSTemplateUpdateOSAssetPath updates an asset mapping
func (c *Client) OSTemplateUpdateOSAssetPath(osTemplateID int, osAssetID int, path string) error {

	resp, err := c.rpcClient.Call(
		"os_template_update_os_asset_path",
		osTemplateID,
		osAssetID,
		path)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// OSTemplateUpdateOSAssetVariables updates an asset variable
func (c *Client) OSTemplateUpdateOSAssetVariables(osTemplateID int, osAssetID int, variablesJSON string) error {

	resp, err := c.rpcClient.Call(
		"os_template_update_os_asset_variables",
		osTemplateID,
		osAssetID,
		variablesJSON)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// OSTemplateMakePublic makes a template public
func (c *Client) OSTemplateMakePublic(osTemplateID int) error {
	resp, err := c.rpcClient.Call(
		"os_template_make_public",
		osTemplateID,
	)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

// OSTemplateMakePrivate makes a template private
func (c *Client) OSTemplateMakePrivate(osTemplateID int, userID int) error {
	resp, err := c.rpcClient.Call(
		"os_template_make_private",
		osTemplateID,
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

// CreateOrUpdate implements interface Applier
func (t OSTemplate) CreateOrUpdate(client MetalCloudClient) error {
	var err error
	var result *OSTemplate
	err = t.Validate()

	if err != nil {
		return err
	}

	if t.VolumeTemplateID != 0 {
		result, err = client.OSTemplateGet(t.VolumeTemplateID, false)
	} else {
		templates, err := client.OSTemplates()
		if err != nil {
			return err
		}

		for _, temp := range *templates {
			if temp.VolumeTemplateLabel == t.VolumeTemplateLabel {
				result = &temp
			}
		}
	}

	if err != nil {
		_, err = client.OSTemplateCreate(t)

		if err != nil {
			return err
		}
	} else {
		_, err = client.OSTemplateUpdate(result.VolumeTemplateID, t)

		if err != nil {
			return err
		}
	}

	return nil
}

// Delete implements interface Applier
func (t OSTemplate) Delete(client MetalCloudClient) error {
	var result *OSTemplate
	var id int
	err := t.Validate()

	if err != nil {
		return err
	}

	if t.VolumeTemplateID != 0 {
		id = t.VolumeTemplateID
	} else {
		templates, err := client.OSTemplates()
		if err != nil {
			return err
		}

		for _, temp := range *templates {
			if temp.VolumeTemplateLabel == t.VolumeTemplateLabel {
				result = &temp
			}
		}

		id = result.VolumeTemplateID
	}
	err = client.OSTemplateDelete(id)

	if err != nil {
		return err
	}

	return nil
}

// Validate implements interface Applier
func (t OSTemplate) Validate() error {
	if t.VolumeTemplateID == 0 && t.VolumeTemplateLabel == "" {
		return fmt.Errorf("id is required")
	}

	if t.VolumeTemplateDisplayName == "" {
		return fmt.Errorf("name is required")
	}

	if t.VolumeTemplateBootType == "" {
		return fmt.Errorf("bootType is required")
	}

	if t.VolumeTemplateOperatingSystem.OperatingSystemType == "" {
		return fmt.Errorf("type is required")
	}
	if t.VolumeTemplateOperatingSystem.OperatingSystemVersion == "" {
		return fmt.Errorf("version is required")
	}
	if t.VolumeTemplateOperatingSystem.OperatingSystemArchitecture == "" {
		return fmt.Errorf("architecture is required")
	}

	return nil
}
