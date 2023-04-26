package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) // Initialize
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Unsupported platform. I can't clear your terminal screen :(")
	}
}

// Removes spaces, tabs, newlines and AT echo from responses
func CleanResponse(s string) string {
	s = strings.Replace(s, "OK", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "AT+CIMI", "", -1)
	return s
}
