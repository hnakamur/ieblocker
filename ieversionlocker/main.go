package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hnakamur/ieversionlocker"
)

var lockvar, unlockvar bool

func init() {
	flag.BoolVar(&lockvar, "l", false, "lock IE version")
	flag.BoolVar(&unlockvar, "u", false, "unlock IE version")
}

func main() {
	flag.Parse()

	if (lockvar && unlockvar) || (!lockvar && !unlockvar) {
		fmt.Println("Please specify one of -l or -u")
		os.Exit(1)
	}

	version, err := ieversionlocker.CurrentVersion()
	if err != nil {
		fmt.Println("Failed detect IE version: %s", err)
		os.Exit(1)
	}

	if lockvar {
		err = ieversionlocker.Lock(version)
		if err != nil {
			fmt.Println("Failed to lock IE version: %s", err)
			os.Exit(1)
		}
	} else if unlockvar {
		err = ieversionlocker.Unlock(version)
		if err != nil {
			fmt.Println("Failed to unlock IE version: %s", err)
			os.Exit(1)
		}
	}
}
