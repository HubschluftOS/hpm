package modules

import (
	"os/user"
)

func IsBool() bool {
	currentUser, err := user.Current()
	if err != nil {
		Error("Unable to get current user: %s", err)
	}
	return currentUser.Username == "root"
}
