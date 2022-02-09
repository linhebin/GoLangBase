package main

import (
	"flag"
	"fmt"

	"go.etcd.io/etcd/version"
)

var (
	help    *bool = flag.Bool("h", false, "to show help")
	showVer *bool = flag.Bool("v", false, "to show version")
)

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *showVer {
		fmt.Printf("version %s \n", version.Version)
		return
	}
}
