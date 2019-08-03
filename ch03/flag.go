package main

import (
	"flag"
	"fmt"
)

// to replace this version
// export GIT_VER=`git describe --tags`; go build -ldflags "-X flag.version=${GIT_VER}" flag.go
var version = "1.0.0"

func main() {
	var showVersion bool
	// if command line option -v -version detected then showVersion is true.
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse() // parse argument options
	if showVersion {
		// print version number and program ends
		fmt.Println("version: ", version)
		return
	}
}
