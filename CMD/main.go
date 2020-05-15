package main

import (
	"flag"
	"fmt"
)

var (
	flagHelp bool
	flagEcho string
)

func parseCmdLineFlags() {
	//flag.BoolVar(&flagHelp, "help", false, "if true, show help")
	flag.StringVar(&flagEcho, "echo", "", "echo somethin")
	flag.Parse()
}

func main() {
	parseCmdLineFlags()
	// if flagHelp {
	// 	flag.Usage()
	// 	os.Exit(0)
	// }
	// fmt.Printf("flag -echo: '%s'\n", flagEcho)

	remainingArgs := flag.Args()
	for _, arg := range remainingArgs {
		fmt.Printf("Remainig arg: '%s'\n", arg)
	}
}
