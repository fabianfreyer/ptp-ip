package main

import (
	"fmt"
	"github.com/malc0mn/ptp-ip/ip"
	"os"
	"path/filepath"
)

var (
	Version   = "0.0.0"
	BuildTime = "unknown"
	exe       string
)

func main() {
	exe = filepath.Base(os.Args[0])

	if noArgs := len(os.Args) < 2; noArgs || help == true {
		usage()
		exit := 0
		if noArgs {
			exit = 1
		}
		os.Exit(exit)
	}

	initFlags()

	if version == true {
		fmt.Printf("%s version %s built on %s\n", exe, Version, BuildTime)
		os.Exit(0)
	}

	if file != "" {
		loadConfig()
	}

	/*sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)*/

	cl, err := ip.NewClient(conf.host, uint16(conf.port), conf.fname, conf.guid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating PTP/IP client - %s\n", err)
		os.Exit(4)
	}
	defer cl.Close()

	fmt.Printf("Created new client with name '%s' and GUID '%s'.\n", cl.InitiatorFriendlyName(), cl.InitiatorGUIDAsString())
	fmt.Printf("Attempting to connect to %s\n", cl.String())
	err = cl.Dial()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to responder - %s\n", err)
		os.Exit(5)
	}

	if server == true {
		launchServer()
	}

	os.Exit(0)
}
