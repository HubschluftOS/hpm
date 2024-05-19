package cmd

import (
	"fmt"
	"os"
)

func UpdateSystem() {
	fmt.Println("soon")
}

func UpdatePackage(pkg string) {
	jsonPath := "/home/rendick/.config/hpm/" + pkg + ".json"
	fmt.Println(jsonPath)

	if _, err := os.Stat(jsonPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[1/1] %s directory does not exist\n", jsonPath)
			return
		} else {
			fmt.Printf("Error accessing %s: %v\n", jsonPath, err)
			return
		}
	}

	fmt.Printf("[1/1] %s directory exists\n", jsonPath)
	return
}
