package cmd

import (
	"fmt"
	"os"
)

func Logs() {
	dirPath := "/tmp/hpm"
	filePath := dirPath + "/hpm.log"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, 0755); err != nil {
			fmt.Println(err)
			return
		}
	}

	openLogFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer openLogFile.Close()

	_, err = openLogFile.WriteString(PackageInput)
	if err != nil {
		fmt.Println(err)
		return
	}

}
