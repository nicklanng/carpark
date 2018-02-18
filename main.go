package main

import (
	"flag"
	"fmt"
)

var (
	version string // set by LDFlags
)

func main() {
	// check for version command
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "version" {
		fmt.Println(version)
		return
	}

	// be friendly
	fmt.Println("Hello, carpark.")
}
