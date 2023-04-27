package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var clear map[string]func()

// Initialized the clear/cls commands for appropritate runtime env
func init() {
	clear = make(map[string]func()) // Initialize
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Clears terminal window
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Unsupported platform. I can't clear your terminal screen :(")
	}
}

// Removes spaces, tabs, newlines and AT echo from responses
func CleanImsi(s string) string {
	s = strings.Replace(s, "OK", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "AT+CIMI", "", -1)
	return s
}

func CleanIccid(s string) string {
	s = strings.Replace(s, "OK", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "+ICCID:", "", -1)
	return s
}

func CleanContacts(s string) string {
	s = strings.Replace(s, "OK", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "+CPBR: ", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, ",", "  ", -1)
	s = strings.Replace(s, "129", "", -1)
	s = strings.Replace(s, " 0", "", -1)
	return s
}

func CleanMsisdn(s string) string {
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "129", "", -1)
	s = strings.Replace(s, "+CNUM: ", "", -1)
	return s
}
