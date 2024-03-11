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
	Version = "1.1v hpm"
	// Path
	PackageLink = "https://hubschluftos.github.io/db/packages/%s.html"
	PackagesURL = "https://hubschluftos.github.io/pages/packages.html"
	PackageDir  = "/bin/"
	// Msg
	ErrorMsg   = "%s not successfully uninstalled! Could not find file or use"
	SuccessMsg = "%s successfully uninstalled!\n"
)

func Help() {
	fmt.Println(Bold + "Usage: pem [OPTION] [FILE]" + Reset +
		"\n\ninstall: installation of the entered package" +
		"\nremove: uninstalling of the entered package" +
		"\nlist: printing of all available packages" +
		"\nstats: print all information about the package manager" +
		"\n\nhelp: print the help menu" +
		"\nversion: print the version" +
		"\n\npem: <https://github.com/rendick/pem/>")
}

func Missing() {
	fmt.Println("hpm: missing arguments\nUsage: hpm [OPTION] [FILE]\nTry help for more information")
}
