package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/tarm/serial"
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
func CleanResponse(s string) string {
	s = strings.Replace(s, "OK", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "AT+CIMI", "", -1)
	return s
}

func ImsiCommand() string {
	config := &serial.Config{
		Name:        "/dev/ttyUSB2",
		Baud:        115200,
		ReadTimeout: time.Millisecond * 100,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
	}

	// Opens the serial port
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
		// fmt.Println(err)
	}
	defer port.Close()

	// Write AT commands to the modem
	imsiCommand := "AT+CIMI\r\n"
	_, err = port.Write([]byte(imsiCommand))
	if err != nil {
		fmt.Println("Cannot connect with device...")
		log.Fatal(err)
	}

	// Reads modem Responses
	buffer := make([]byte, 128)
	n, err := port.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// Print Response
	response := string(buffer[:n])
	return response
}
