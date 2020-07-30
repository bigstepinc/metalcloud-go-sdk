package metalcloud

import (
	"fmt"
	"strings"
)

//OSTemplate A template can be created based on a drive and it has the same characteristics and holds the same information as the parent drive.
type OSTemplate struct {
	VolumeTemplateID                   int                    `json:"volume_template_id,omitempty"`
	VolumeTemplateLabel                string                 `json:"volume_template_label,omitempty"`
	VolumeTemplateDisplayName          string                 `json:"volume_template_display_name,omitempty"`
	VolumeTemplateSizeMBytes           int                    `json:"volume_template_size_mbytes,omitempty"`
	VolumeTemplateLocalDiskSupported   bool                   `json:"volume_template_local_disk_supported,omitempty"`
	VolumeTemplateIsOSTemplate         bool                   `json:"volume_template_is_os_template,omitempty"`
	VolumeTemplateBootMethodsSupported string                 `json:"volume_template_boot_methods_supported,omitempty"`
	VolumeTemplateBootType             string                 `json:"volume_template_boot_type,omitempty"`
	VolumeTemplateDescription          string                 `json:"volume_template_description,omitempty"`
	VolumeTemplateCreatedTimestamp     string                 `json:"volume_template_created_timestamp,omitempty"`
	VolumeTemplateUpdatedTimestamp     string                 `json:"volume_template_updated_timestamp,omitempty"`
	UserID                             int                    `json:"user_id,omitempty"`
	VolumeTemplateOperatingSystem      *OperatingSystem       `json:"volume_template_operating_system,omitempty"`
	VolumeTemplateRepoURL              string                 `json:"volume_template_repo_url,omitempty"`
	VolumeTemplateDeprecationStatus    string                 `json:"volume_template_deprecation_status,omitempty"`
	OSTemplateCredentials              *OSTemplateCredentials `json:"os_template_credentials,omitempty"`
	VolumeTemplateTags                 []string               `json:"volume_template_tags,omitempty"`
	OSTemplatePreBootArchitecture      string                 `json:"os_template_pre_boot_architecture,omitempty"`
	OSAssetBootloaderLocalInstall      int                    `json:"os_asset_id_bootloader_local_install,omitempty"`
	OSAssetBootloaderOSBoot            int                    `json:"os_asset_id_bootloader_os_boot,omitempty"`
	VolumeTemplateVariablesJSON        string                 `json:"volume_template_variables_json,omitempty"`
}

//OSTemplateCredentials holds information needed to connect to an OS installed by an OSTemplate.
type OSTemplateCredentials struct {
	OSTemplateInitialUser               string `json:"os_template_initial_user,omitempty"`
	OSTemplateInitialPasswordEncrypted  string `json:"os_template_initial_password_encrypted,omitempty"`
	OSTemplateInitialPassword           string `json:"os_template_initial_password,omitempty"`
	OSTemplateInitialSSHPort            int    `json:"os_template_initial_ssh_port,omitempty"`
	OSTemplateChangePasswordAfterDeploy bool   `json:"os_template_change_password_after_deploy,omitempty"`
}

//OSTemplateOSAssetData holds asset-template information
type OSTemplateOSAssetData struct {
	OSAsset                           *OSAsset `json:"os_asset,omitempty"`
	OSAssetFilePath                   string   `json:"os_asset_file_path,omitempty"`
	OSTemplateOSAssetUpdatedTimestamp string   `json:"volume_template_os_asset_updated_timestamp,omitempty"`
	OSTemplateOSAssetVariablesJSON    string   `json:"volume_template_os_asset_variables_json,omitempty"`
}

//OSTemplateCreate creates a osTemplate object
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

//OSTemplateDelete permanently destroys a OSTemplate.
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

//OSTemplateUpdate updates a osTemplate
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

//OSTemplateGet returns a OSTemplate specified by nOSTemplateID. The OSTemplate's protected value is never returned.
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

	if decryptPasswd {

		passwdComponents := strings.Split(createdObject.OSTemplateCredentials.OSTemplateInitialPassword, ":")
		if len(passwdComponents) != 2 {
			return nil, fmt.Errorf("Password not returned with proper components")
		}

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

	return &createdObject, nil
}

//OSTemplates retrieves a list of all the OSTemplate objects which a specified User is allowed to see through ownership or delegation. The OSTemplate objects never return the actual protected OSTemplate value.
func (c *Client) OSTemplates() (*map[string]OSTemplate, error) {

	userID := c.GetUserID()
	res, err := c.rpcClient.Call(
		"os_templates",
		userID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]OSTemplate{}
		return &m, nil
	}

	var createdObject map[string]OSTemplate

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//OSTemplateOSAssets returns the OSAssets assigned to an OSTemplate.
func (c *Client) OSTemplateOSAssets(osTemplateID int) (*map[string]OSTemplateOSAssetData, error) {

	res, err := c.rpcClient.Call(
		"os_template_os_assets",
		osTemplateID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]OSTemplateOSAssetData{}
		return &m, nil
	}

	var createdObject map[string]OSTemplateOSAssetData

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//OSTemplateAddOSAsset adds an asset to a template
func (c *Client) OSTemplateAddOSAsset(osTemplateID int, osAssetID int, path string, variablesJSON string) error {

	_, err := c.rpcClient.Call(
		"os_template_add_os_asset",
		osTemplateID,
		osAssetID,
		path,
		variablesJSON)

	if err != nil {
		return err
	}

	return nil
}

//OSTemplateRemoveOSAsset removes an asset from a template
func (c *Client) OSTemplateRemoveOSAsset(osTemplateID int, osAssetID int) error {

	_, err := c.rpcClient.Call(
		"os_template_remove_os_asset",
		osTemplateID,
		osAssetID)

	if err != nil {
		return err
	}

	return nil
}

//OSTemplateUpdateOSAssetPath updates an asset mapping
func (c *Client) OSTemplateUpdateOSAssetPath(osTemplateID int, osAssetID int, path string) error {

	_, err := c.rpcClient.Call(
		"os_template_update_os_asset_path",
		osTemplateID,
		osAssetID,
		path)

	if err != nil {
		return err
	}

	return nil
}

//OSTemplateUpdateOSAssetVariables updates an asset variable
func (c *Client) OSTemplateUpdateOSAssetVariables(osTemplateID int, osAssetID int, variablesJSON string) error {

	_, err := c.rpcClient.Call(
		"os_template_update_os_asset_variables",
		osTemplateID,
		osAssetID,
		variablesJSON)

	if err != nil {
		return err
	}

	return nil
}
