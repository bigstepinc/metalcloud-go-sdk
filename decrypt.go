package metalcloud

import (
	"fmt"
	"strings"
)

//decryptIfEncrypted will decrypt a password if encrypted
//using the serverside decrypt functionality
func (c *Client) decryptIfEncrypted(potentiallyEncryptedPassword string) (string, error) {
	passwdComponents := strings.Split(potentiallyEncryptedPassword, ":")

	if len(passwdComponents) == 2 {
		if strings.Contains(passwdComponents[0], "Not authorized") {
			return "", fmt.Errorf("Permission missing. %s", passwdComponents[1])
		} else {
			var passwd string

			err := c.rpcClient.CallFor(
				&passwd,
				"password_decrypt",
				passwdComponents[1],
			)
			if err != nil {
				return "", err
			}

			return passwd, nil
		}
	}
	fmt.Printf("Password did not have a ':' it is %s", potentiallyEncryptedPassword)
	return potentiallyEncryptedPassword, nil
}
