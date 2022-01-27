package metalcloud

import (
	"fmt"
	"strings"
)

//StoragePoolSearchResult represents a storage appliance in a datacenter.
type StoragePoolSearchResult struct {
	StoragePoolID                              int    `json:"storage_pool_id,omitempty" yaml:"id,omitempty"`
	StoragePoolName                            string `json:"storage_pool_name,omitempty" yaml:"name,omitempty"`
	StoragePoolStatus                          string `json:"storage_pool_status,omitempty" yaml:"status,omitempty"`
	StoragePoolInMaintenance                   bool   `json:"storage_pool_in_maintenance,omitempty" yaml:"maintenance,omitempty"`
	DatacenterName                             string `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	StorageType                                string `json:"storage_type,omitempty" yaml:"type,omitempty"`
	UserID                                     int    `json:"user_id,omitempty" yaml:"user,omitempty"`
	StoragePoolISCSIHost                       string `json:"storage_pool_iscsi_host,omitempty" yaml:"host,omitempty"`
	StoragePoolISCSIPort                       int    `json:"storage_pool_iscsi_port,omitempty" yaml:"port,omitempty"`
	StoragePoolEndpoint                        string `json:"storage_pool_endpoint,omitempty" yaml:"endpoint,omitempty"`
	StoragePoolCapacityTotalCachedRealMbytes   int    `json:"storage_pool_capacity_total_cached_real_mbytes,omitempty" yaml:"capacityTotalRealMbytes,omitempty"`
	StoragePoolCapacityUsableCachedRealMbytes  int    `json:"storage_pool_capacity_usable_cached_real_mbytes,omitempty" yaml:"capacityUsableRealMbytes,omitempty"`
	StoragePoolCapacityFreeCachedRealMbytes    int    `json:"storage_pool_capacity_free_cached_real_mbytes,omitempty" yaml:"capacityFreeRealMbytes,omitempty"`
	StoragePoolCapacityUsedCachedVirtualMbytes int    `json:"storage_pool_capacity_used_cached_virtual_mbytes,omitempty" yaml:"capacityUsedRealMbytes,omitempty"`
}

//searchResultResponseWrapperForStoragePoolSearchResult describes a search result for storage pools
type searchResultResponseWrapperForStoragePoolSearchResult struct {
	DurationMilliseconds int                       `json:"duration_millisecnds,omitempty"`
	Rows                 []StoragePoolSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string                `json:"rows_order,omitempty"`
	RowsTotal            int                       `json:"rows_total,omitempty"`
}

//StoragePool represents a storage appliance in a datacenter.
type StoragePool struct {
	StoragePoolID                              int      `json:"storage_pool_id,omitempty" yaml:"id,omitempty"`
	StoragePoolName                            string   `json:"storage_pool_name,omitempty" yaml:"name,omitempty"`
	StoragePoolStatus                          string   `json:"storage_pool_status,omitempty" yaml:"status,omitempty"`
	StoragePoolInMaintenance                   bool     `json:"storage_pool_in_maintenance,omitempty" yaml:"maintenance,omitempty"`
	StoragePoolIsExperimental                  bool     `json:"storage_pool_is_experimental,omitempty" yaml:"experimental,omitempty"`
	DatacenterName                             string   `json:"datacenter_name,omitempty" yaml:"datacenter,omitempty"`
	StorageType                                string   `json:"storage_type,omitempty" yaml:"type,omitempty"`
	UserID                                     int      `json:"user_id,omitempty" yaml:"user,omitempty"`
	StoragePoolISCSIHost                       string   `json:"storage_pool_iscsi_host,omitempty" yaml:"host,omitempty"`
	StoragePoolISCSIPort                       int      `json:"storage_pool_iscsi_port,omitempty" yaml:"port,omitempty"`
	StoragePoolCapacityTotalCachedRealMbytes   int      `json:"storage_pool_capacity_total_cached_real_mbytes,omitempty" yaml:"capacityTotalRealMbytes,omitempty"`
	StoragePoolCapacityUsableCachedRealMbytes  int      `json:"storage_pool_capacity_usable_cached_real_mbytes,omitempty" yaml:"capacityUsableRealMbytes,omitempty"`
	StoragePoolCapacityFreeCachedRealMbytes    int      `json:"storage_pool_capacity_free_cached_real_mbytes,omitempty" yaml:"capacityFreeRealMbytes,omitempty"`
	StoragePoolCapacityUsedCachedVirtualMbytes int      `json:"storage_pool_capacity_used_cached_virtual_mbytes,omitempty" yaml:"capacityUsedRealMbytes,omitempty"`
	StoragePoolCreatedTimestamp                string   `json:"storage_pool_created_timestamp,omitempty" yaml:"createdTimestamp,omitempty"`
	StoragePoolUpdatedTimestamp                string   `json:"storage_pool_updated_timestamp,omitempty" yaml:"updatedTimestamp,omitempty"`
	StoragePoolDrivePriority                   int      `json:"storage_pool_drive_priority,omitempty" yaml:"drivePriority,omitempty"`
	StoragePoolEndpoint                        string   `json:"storage_pool_endpoint,omitempty" yaml:"endpoint,omitempty"`
	StoragePoolOptionsJSON                     string   `json:"storage_pool_options_json,omitempty" yaml:"optionsJSON,omitempty"`
	StoragePoolUsername                        string   `json:"storage_pool_username,omitempty" yaml:"username,omitempty"`
	StoragePoolPassword                        string   `json:"storage_pool_password,omitempty" yaml:"password,omitempty"`
	StoragePoolPortGroupAllocationOrderJSON    string   `json:"storage_pool_port_group_allocation_order_json,omitempty" yaml:"portGroupAllocationOrderJSON,omitempty"`
	StoragePoolPortGroupPhysicalPortsJSON      string   `json:"storage_pool_port_group_physical_ports_json,omitempty" yaml:"portGroupPhysicalPortsJSON,omitempty"`
	StoragePoolSharedDrivePriority             int      `json:"storage_pool_shared_drive_priority,omitempty" yaml:"sharedDrivePriority,omitempty"`
	StoragePoolTags                            []string `json:"storage_pool_tags,omitempty" yaml:"tags,omitempty"`
	StoragePoolTargetIQN                       string   `json:"storage_pool_target_iqn,omitempty" yaml:"targetIQN,omitempty"`
}

//StoragePoolSearch searches for storage pools matching certain filter
func (c *Client) StoragePoolSearch(filter string) (*[]StoragePoolSearchResult, error) {

	tables := []string{"_storages"}
	columns := map[string][]string{
		"_storages": {
			"storage_pool_id",
			"storage_pool_name",
			"storage_pool_status",
			"storage_pool_in_maintenance",
			"server_serial_number",
			"datacenter_name",
			"storage_type",
			"storage_pool_endpoint",
			"user_id",
			"storage_pool_iscsi_host",
			"storage_pool_iscsi_port",
			"storage_pool_capacity_total_cached_real_mbytes",
			"storage_pool_capacity_usable_cached_real_mbytes",
			"storage_pool_capacity_free_cached_real_mbytes",
			"storage_pool_capacity_used_cached_virtual_mbytes",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	var createdObject map[string]searchResultResponseWrapperForStoragePoolSearchResult

	resp, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]searchResultResponseWrapperForStoragePoolSearchResult{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []StoragePoolSearchResult{}
	for _, s := range createdObject[tables[0]].Rows {
		list = append(list, s)
	}

	return &list, nil
}

//StoragePoolGet returns a storage pool's details
func (c *Client) StoragePoolGet(serverID int, decryptPasswd bool) (*StoragePool, error) {

	var createdObject StoragePool

	err := c.rpcClient.CallFor(
		&createdObject,
		"storage_pool_get",
		serverID)

	if err != nil {
		return nil, err
	}

	if decryptPasswd {

		passwdComponents := strings.Split(createdObject.StoragePoolPassword, ":")

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

				createdObject.StoragePoolPassword = passwd
			}
		}
	} else {
		createdObject.StoragePoolPassword = ""
	}

	return &createdObject, nil
}
