package main

import (
	"io/ioutil"
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

func Banner() string {
	b, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// Removes spaces, tabs, newlines and AT echo from responses
func CleanImsi(s string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "\n", "", "AT+CIMI", "")
	s = replacer.Replace(s)
	return s
}

func CleanIccid(s string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "\n", "", "+ICCID:", "")
	s = replacer.Replace(s)
	return s
}

func CleanContacts(s string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "+CPBR: ", "", "\"", "", ",", "  ", "129", "", "0", "")
	s = replacer.Replace(s)
	return s
}

func CleanMsisdn(s string) string {
	replacer := strings.NewReplacer("\"", "", ",", "", "129", "", "+CNUM: ", "")
	s = replacer.Replace(s)
	return s
}
