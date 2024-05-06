package cmd

import (
	"fmt"
	"os"
	"os/user"
)

func ConfigurateManager() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("[0/1] %s\n", err)
	}

	configPath := currentUser.HomeDir + "/.config/hpm/"
	fmt.Printf("[1/1] checking the configuration directory\n")
	if _, err := os.Stat(configPath); err != nil {
		if os.IsExist(err) {
			fmt.Printf("[1/1] %s directory exists\n", configPath)
		} else {
			if err := os.Mkdir(configPath, 0755); err != nil {
				fmt.Printf("[0/1] %s\n", err)
			}

			fmt.Printf("[1/1] config directory successfully configured\n")
		}
	} else {
		fmt.Printf("[1/1] config directory successfully checked\n")
	}
}
