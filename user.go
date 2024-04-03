package metalcloud

import (
	"fmt"
)

//User describes user account specifications.
type User struct {
	UserID          int    `json:"user_id,omitempty"`
	UserDisplayName string `json:"user_display_name,omitempty"`
	UserEmail       string `json:"user_email,omitempty"`
}

type searchResultWrapperForUsers struct {
	DurationMilliseconds int                 `json:"duration_millisecnds,omitempty"`
	Rows                 []UsersSearchResult `json:"rows,omitempty"`
	RowsOrder            [][]string          `json:"rows_order,omitempty"`
	RowsTotal            int                 `json:"rows_total,omitempty"`
}

type UsersSearchResult struct {
	Franchise                            string `json:"franchise" yaml:"franchise"`
	UserAccessLevel                      string `json:"user_access_level" yaml:"user_access_level"`
	UserAuthFailedAttemptsSinceLastLogin int    `json:"user_auth_failed_attempts_since_last_login" yaml:"userFailedLoginAttempts"`
	UserAuthenticatorCreatedTimestamp    string `json:"user_authenticator_created_timestamp" yaml:"userCreatedTimestamp"`
	UserAuthenticatorIsMandatory         bool   `json:"user_authenticator_is_mandatory,omitempty" yaml:"userAuthenticatorIsMandatory,omitempty"`
	UserAuthenticatorMustChange          bool   `json:"user_authenticator_must_change,omitempty" yaml:"userAuthenticatorMustChange,omitempty"`
	UserBlocked                          bool   `json:"user_blocked" yaml:"userBlocked"`
	UserBrand                            string `json:"user_brand,omitempty" yaml:"userBrand,omitempty"`
	UserCreatedTimestamp                 string `json:"user_created_timestamp" yaml:"userCreatedTimestamp"`
	UserCustomPricesJson                 string `json:"user_custom_prices_json,omitempty" yaml:"userCustomPricesJson,omitempty"`
	UserDisplayName                      string `json:"user_display_name" yaml:"displayName"`
	UserEmail                            string `json:"user_email" yaml:"email"`
	UserEmailStatus                      string `json:"user_email_status,omitempty" yaml:"emailStatus,omitempty"`
	UserExcludeFromReports               bool   `json:"user_exclude_from_reports,omitempty" yaml:"userExcludeFromReports,omitempty"`
	UserExperimentalTagsJson             string `json:"user_experimental_tags_json,omitempty" yaml:"userExperimentalTagsJson,omitempty"`
	UserExternalIDsJson                  string `json:"user_external_ids_json,omitempty" yaml:"userExternalIDsJson,omitempty"`
	UserGuiSettingsJson                  string `json:"user_gui_settings_json,omitempty" yaml:"userGuiSettingsJson,omitempty"`
	UserID                               int    `json:"user_id" yaml:"id"`
	UserInfrastructureIDDefault          int    `json:"user_infrastructure_id_default,omitempty" yaml:"userInfrastructureIDDefault,omitempty"`
	UserIsBillable                       bool   `json:"user_is_billable,omitempty" yaml:"userIsBillable,omitempty"`
	UserIsBrandManager                   bool   `json:"user_is_brand_manager,omitempty" yaml:"userIsBrandManager,omitempty"`
	UserIsDatastorePublisher             bool   `json:"user_is_datastore_publisher,omitempty" yaml:"userIsDatastorePublisher,omitempty"`
	UserIsSuspended                      bool   `json:"user_is_suspended,omitempty" yaml:"userIsSuspended,omitempty"`
	UserIsTestAccount                    bool   `json:"user_is_test_account,omitempty" yaml:"userIsTestAccount,omitempty"`
	UserIsTestingMode                    bool   `json:"user_is_testing_mode,omitempty" yaml:"userIsTestingMode,omitempty"`
	UserKerberosPrincipalName            string `json:"user_kerberos_principal_name,omitempty" yaml:"userKerberosPrincipalName,omitempty"`
	UserLanguage                         string `json:"user_language,omitempty" yaml:"userLanguage,omitempty"`
	UserLastLoginTimestamp               string `json:"user_last_login_timestamp" yaml:"userLastLoginTimestamp"`
	UserLastLoginType                    string `json:"user_last_login_type,omitempty" yaml:"userLastLoginType,omitempty"`
	UserLimitsJson                       string `json:"user_limits_json,omitempty" yaml:"userLimitsJson,omitempty"`
	UserPasswordChangeRequired           bool   `json:"user_password_change_required,omitempty" yaml:"userPasswordChangeRequired,omitempty"`
	UserPermissionsJson                  string `json:"user_permissions_json,omitempty" yaml:"userPermissionsJson,omitempty"`
	UserPlanType                         string `json:"user_plan_type,omitempty" yaml:"userPlanType,omitempty"`
	UserPromotionTagsJson                string `json:"user_promotion_tags_json,omitempty" yaml:"userPromotionTagsJson,omitempty"`
}

//userGet describes returns user account specifications.
func (c *Client) userGet(userID id) (*User, error) {

	var createdObject User

	err := c.rpcClient.CallFor(
		&createdObject,
		"user_get",
		userID)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//UserGet describes returns user account specifications.
func (c *Client) UserGet(userID int) (*User, error) {
	return c.userGet(userID)
}

//UserGetByEmail describes returns user account specifications.
func (c *Client) UserGetByEmail(userLabel string) (*User, error) {
	userID, err := c.UserEmailToUserID(userLabel)
	if err != nil {
		return nil, err
	}
	
	return c.userGet(*userID)
}

//UserEmailToUserID returns the user id of an user given an email
func (c *Client) UserEmailToUserID(userEmail string) (*int, error) {

	var createdObject int

	err := c.rpcClient.CallFor(
		&createdObject,
		"user_email_to_user_id",
		userEmail)

	if err != nil {

		return nil, err
	}

	return &createdObject, nil
}

//UserSearch searches for users with filtering support
func (c *Client) UserSearch(filter string) (*[]UsersSearchResult, error) {

	tables := []string{"_users"}
	columns := map[string][]string{

		"_users": {
			"franchise",
			"user_access_level",
			"user_auth_failed_attempts_since_last_login",
			"user_authenticator_created_timestamp",
			"user_authenticator_is_mandatory",
			"user_authenticator_must_change",
			"user_blocked",
			"user_brand",
			"user_created_timestamp",
			"user_custom_prices_json",
			"user_display_name",
			"user_email",
			"user_email_status",
			"user_exclude_from_reports",
			"user_experimental_tags_json",
			"user_external_ids_json",
			"user_gui_settings_json",
			"user_id",
			"user_infrastructure_id_default",
			"user_is_billable",
			"user_is_brand_manager",
			"user_is_datastore_publisher",
			"user_is_suspended",
			"user_is_test_account",
			"user_is_testing_mode",
			"user_kerberos_principal_name",
			"user_language",
			"user_last_login_timestamp",
			"user_last_login_type",
			"user_limits_json",
			"user_password_change_required",
			"user_permissions_json",
			"user_plan_type",
			"user_promotion_tags_json",
		},
	}

	userID := c.GetUserID()

	collapseType := "array_row_span"
	sortBy := [][]string{
		{
			"user_id",
			"DESC",
		},
	}

	var createdObject map[string]searchResultWrapperForUsers

	resp, err := c.rpcClient.Call(
		"search",
		userID,
		filter,
		tables,
		columns,
		collapseType,
		sortBy,
	)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		createdObject = map[string]searchResultWrapperForUsers{}
	} else {
		err = resp.GetObject(&createdObject)

		if err != nil {
			return nil, err
		}
	}

	list := []UsersSearchResult{}
	list = append(list, createdObject[tables[0]].Rows...)

	return &list, nil
}