package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// config path
const configPath = "/etc/hpm/"

func ConfigurateManager() {
	configPath := "/etc/hpm/"
	fmt.Printf("[1/1] Checking the configuration directory\n")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("[1/1] %s directory does not exist\n", configPath)
		createDir()
		createFile()
	}

	fmt.Printf("[1/1] Everything is okay\n")
	return
}

func createDir() {
	fmt.Printf("[1/1] Directory does not exist, creating...\n")

	if err := os.Mkdir(configPath, 0755); err != nil {
		fmt.Printf("[0/1] %s\n", err)
		return
	}

	fmt.Printf("[1/1] Directory successfully created\n")
	return
}

func createFile() {
	fmt.Printf("[1/1] Configuration file does not exist, creating...\n")
	createConfigFile, err := os.Create(configPath + "config.json")
	if err != nil {
		fmt.Printf("[0/1] %s\n", err)
		return
	}
	defer createConfigFile.Close()

	_, err = createConfigFile.WriteString(configJsonFileExample)
	if err != nil {
		fmt.Printf("[0/1] Error writing data to the file\n")
		return
	}

	fmt.Printf("[1/1] Configuration file successfully created\n")
	fmt.Printf("[1/1] Config directory successfully configured\n")
	return
}

func reloadConfig() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(red + bold + "[!/!] THIS COMMAND WILL SET THE ~/.config/hpm/ DIRECTORY TO THE DEFAULT!!!\n" + reset)
	fmt.Print("[?/?] Would you like to reload your configuration file? [y/n] ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	if input == "" || input == "y" || input == "yes" {
		if err := os.RemoveAll(configPath); err != nil {
			fmt.Println("Error removing config directory:", err)
			return
		}
		createDir()
		createFile()
		return
	} else {
		fmt.Printf("[0/1] exiting\n")
		return
	}
}

func removeConfig() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(red + bold + "[!/!] THIS COMMAND WILL DELETE THE ~/.config/hpm/ DIRECTORY!!!\n" + reset)
	fmt.Print("[?/?] Would you like to delete your configuration file? [y/n] ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	if input == "" || input == "y" || input == "yes" {
		if err := os.RemoveAll(configPath); err != nil {
			fmt.Println("Error removing config directory:", err)
			return
		}
		return
	} else {
		fmt.Printf("[0/1] exiting\n")
		return
	}
}
