package homedir

import (
	"os/user"
)

// GetHomeDir returns the user's home directory.
func GetHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.HomeDir, nil
}
