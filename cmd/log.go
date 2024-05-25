package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	dirPath  = "/tmp/hpm"
	filePath = dirPath + "/hpm.log"
)

func Logs() {
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

func DisplayLogs() {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(strings.TrimSpace(string(data)))
	}
}

func DeleteLogs() {
	if err := os.RemoveAll("/tmp/hpm/"); err != nil {
		fmt.Printf("[0/1] error while removing the folder")
	}
}
