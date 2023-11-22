package main

import (
	"os"
	"time"
)

var hostname = mustGetHostname()

func mustGetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func stop() {
	go func() {
		// Shut down after some time to make sure the response is actually returned to the caller.
		time.Sleep(time.Second)
		os.Exit(0)
	}()
}
