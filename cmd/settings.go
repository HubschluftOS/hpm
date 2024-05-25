package cmd

import "time"

var currentTime = time.Now().Format("2006-01-02 15:04:05")

// main db
const (
	db         = "https://hubshluft.github.io/db/"
	configPath = "/etc/hpm/"
)

// package information. sync.go
const ContinueMSG = `
Packages:		%s
Version:		%s
Maintainer:		%s
Dependencies:		%s
Size:			%d MiB
Source:			%s

`

// config file example. config.go
const configJsonFileExample = `{
    "root_user": "mynickname",
    "installation_path": "/usr/bin/"
}
`

const packageInfo = `
%s is available.
----------------------
%s%s%s%s
`

// colors *.go
const (
	reset = "\033[0m"
	bold  = "\033[1m"
	red   = "\033[31m"
)
