package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

var configJsonFileExample = `{
    "username": "mynickname",
    "installation_path": "7.1.0",
}
`

var (
	reset = "\033[0m"
	bold  = "\033[1m"
	red   = "\033[31m"
)

func ConfigurateManager() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("[0/1] %s\n", err)
	}

	configPath := currentUser.HomeDir + "/.config/hpm/"
	fmt.Printf("[1/1] checking the configuration directory\n")
	if _, err := os.Stat(configPath); err != nil {
		fmt.Printf("[1/1] %s directory does not exist\n", configPath)
		createDir()
		createFile()
		return
	} else {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf(red + bold + "[!/!] THIS COMMAND WILL SET THE ~/.config/hpm/ DIRECTORY TO THE DEFAULT!!!\n" + reset)
		fmt.Print("[?/?] you would like to reload you configuration file? [y/n] ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)

		if input == "" || input == "y" || input == "yes" {
			currentUser, err := user.Current()
			if err != nil {
				fmt.Printf("[0/1] %s\n", err)
			}

			configPath := currentUser.HomeDir + "/.config/hpm/"
			if err := os.RemoveAll(configPath); err != nil {
				fmt.Println("err")
				return
			}
			createDir()
			createFile()
		} else {
			fmt.Println("EXITING")
			return
		}
	}
}

func createDir() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("[0/1] %s\n", err)
	}

	configPath := currentUser.HomeDir + "/.config/hpm/"

	fmt.Printf("[1/1] directory does not exist, creating...\n")
	if err := os.Mkdir(configPath, 0755); err != nil {
		fmt.Printf("[0/1] %s\n", err)
		return
	}
	fmt.Printf("[1/1] directory successfully created\n")
}

func createFile() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("[0/1] %s\n", err)
	}

	configPath := currentUser.HomeDir + "/.config/hpm/"

	fmt.Printf("[1/1] configuration file does not exist, creating...\n")
	createConfigFile, err := os.Create(configPath + "config.json")
	if err != nil {
		fmt.Printf("[0/1] %s\n", err)
		return
	}
	defer createConfigFile.Close()
	fmt.Printf("[1/1] configuration file successfully created\n")

	writingJsonData, err := createConfigFile.WriteString(configJsonFileExample)
	if err != nil {
		fmt.Printf("[0/1] error writing data to the filo")
	}
	fmt.Println(writingJsonData)

	fmt.Printf("[1/1] config directory successfully configured\n")
}
