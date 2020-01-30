package metalcloud

//User describes user account specifications.
type User struct {
	UserID          int    `json:"user_id,omitempty"`
	UserDisplayName string `json:"user_display_name,omitempty"`
	UserEmail       string `json:"user_email,omitempty"`
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
	return c.userGet(userLabel)
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
