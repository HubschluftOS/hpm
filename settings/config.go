package config

import (
	"fmt"
	"time"
)

var (
	// style
	Red    = "\033[31m"
	Bold   = "\033[1m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Time   = time.Now().Format("2006-01-02 15:04:05")
)

const (
	// Version
	Version = "1.0v"
	// Path
	PackageLink = "https://pem-packages.vercel.app/packages/%s.html"
	LogPath     = "/var/log/pem/"
	PackagesURL = "https://pem-packages.vercel.app/"
	PackageDir  = "/bin/"
	// Msg
	ErrorMsg   = "%s not successfully uninstalled!\nCould not find file!\n"
	SuccessMsg = "%s successfully uninstalled!\n"
)

func Help() {
	fmt.Println(Bold + "Usage: pem [OPTION] [FILE]" + Reset +
		"\n\ninstall: installation of the entered package" +
		"\nremove: uninstalling of the entered package" +
		"\nlist: printing of all available packages" +
		"\n\npem: <https://github.com/rendick/pem/>")
}

func Missing() {
	fmt.Println("pem: missing arguments\nUsage: pem [OPTION]\nTry --help for more information")
}
