package main

import (
	"os/exec"
	"strconv"
)

// Check asserts connectivity.
func Check(ping Ping) bool {
	cmd := exec.Command("ping", "-w", strconv.Itoa(ping.Timeout), ping.Address)
	err := cmd.Run()
	return err == nil
}
