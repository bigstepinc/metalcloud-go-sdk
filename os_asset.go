package metalcloud

import "fmt"

//OSAsset struct defines a server type
type OSAsset struct {
	OSAssetID                    int      `json:"os_asset_id,omitempty" yaml:"id,omitempty"`
	UserIDOwner                  int      `json:"user_id_owner,omitempty" yaml:"ownerID,omitempty"`
	UserIDAuthenticated          int      `json:"user_id_authenticated,omitempty" yaml:"userIDAuthenticated,omitempty"`
	OSAssetFileName              string   `json:"os_asset_filename,omitempty" yaml:"fileName,omitempty"`
	OSAssetFileSizeBytes         int      `json:"os_asset_file_size_bytes,omitempty" yaml:"fileSizeBytes,omitempty"`
	OSAssetFileMime              string   `json:"os_asset_file_mime,omitempty" yaml:"fileMime,omitempty"`
	OSAssetContentsBase64        string   `json:"os_asset_contents_base64,omitempty" yaml:"contentBase64,omitempty"`
	OSAssetContentsSHA256Hex     string   `json:"os_asset_contents_sha256_hex,omitempty" yaml:"contentSHA256Hex,omitempty"`
	OSAssetUsage                 string   `json:"os_asset_usage,omitempty" yaml:"usage,omitempty"`
	OSAssetSourceURL             string   `json:"os_asset_source_url,omitempty" yaml:"sourceURL,omitempty"`
	OSAssetVariableNamesRequired []string `json:"os_asset_variable_names_required,omitempty" yaml:"variables,omitempty"`
	OSAssetTags                  []string `json:"os_asset_tags,omitempty" yaml:"tags,omitempty"`
	OSAssetCreatedTimestamp      string   `json:"os_asset_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	OSAssetUpdatedTimestamp      string   `json:"os_asset_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
}

//OSAssetCreate creates a osAsset object
func (c *Client) OSAssetCreate(osAsset OSAsset) (*OSAsset, error) {
	var createdObject OSAsset

	userID := c.GetUserID()

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_asset_create",
		userID,
		osAsset)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//OSAssetDelete permanently destroys a OSAsset.
func (c *Client) OSAssetDelete(osAssetID int) error {

	if err := checkID(osAssetID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("os_asset_delete", osAssetID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//OSAssetUpdate updates a osAsset
func (c *Client) OSAssetUpdate(osAssetID int, osAsset OSAsset) (*OSAsset, error) {
	var createdObject OSAsset

	if err := checkID(osAssetID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_asset_update",
		osAssetID,
		osAsset)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//OSAssetGet returns a OSAsset specified by nOSAssetID. The OSAsset's protected value is never returned.
func (c *Client) OSAssetGet(osAssetID int) (*OSAsset, error) {

	var createdObject OSAsset

	if err := checkID(osAssetID); err != nil {
		return nil, err
	}

	err := c.rpcClient.CallFor(
		&createdObject,
		"os_asset_get",
		osAssetID)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//OSAssets retrieves a list of all the OSAsset objects which a specified User is allowed to see through ownership or delegation. The OSAsset objects never return the actual protected OSAsset value.
func (c *Client) OSAssets() (*map[string]OSAsset, error) {

	userID := c.GetUserID()

	res, err := c.rpcClient.Call(
		"os_assets",
		userID)

	if err != nil {
		return nil, err
	}

	_, ok := res.Result.([]interface{})
	if ok {
		var m = map[string]OSAsset{}
		return &m, nil
	}

	var createdObject map[string]OSAsset

	err2 := res.GetObject(&createdObject)
	if err2 != nil {
		return nil, err2
	}

	return &createdObject, nil
}

//CreateOrUpdate implements interface Applier
func (asset OSAsset) CreateOrUpdate(c interface{}) error {
	client := c.(*Client)

	var err error
	var result *OSAsset

	if asset.OSAssetID != 0 {
		result, err = client.OSAssetGet(asset.OSAssetID)
	} else if asset.OSAssetFileName != "" {
		assets, err := client.OSAssets()
		if err != nil {
			return err
		}

		for _, a := range *assets {
			if a.OSAssetFileName == asset.OSAssetFileName {
				result = &a
			}
		}
	} else {
		return fmt.Errorf("id is required")
	}

	if result == nil {
		_, err = client.OSAssetCreate(asset)

		if err != nil {
			return err
		}
	} else {
		_, err = client.OSAssetUpdate(asset.OSAssetID, asset)

		if err != nil {
			return err
		}
	}

	return nil
}

//Delete implements interface Applier
func (asset OSAsset) Delete(c interface{}) error {
	client := c.(*Client)

	err := client.OSAssetDelete(asset.OSAssetID)

	if err != nil {
		return err
	}

	return nil
}

//Validate implements interface Applier
func (asset OSAsset) Validate() error {
	return nil
}
