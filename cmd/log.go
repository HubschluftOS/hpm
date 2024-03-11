package cmd

import (
	"log"
	"os"
	"os/user"
)

func Logs() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	logDir := currentUser.HomeDir + "/.config/hpm/"
	_, err = os.Stat(logDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0750)
		if err != nil && !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}

	logFilePath := logDir + "hpm.log"
	OpenLog, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = OpenLog.WriteString(Log)
	if err != nil {
		log.Fatal(err)
	}

	defer OpenLog.Close()
}
