package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func main() {
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
	testCommand := "AT\r\n"
	_, err = port.Write([]byte(testCommand))
	if err != nil {
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
	fmt.Println(cleanUp(string(colorBlack), response))

}

// Removes spaces, tabs, newlines and AT echo from output
func cleanUp(s string) string {
	s = strings.Replace(s, "*AT*", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return string(s)
}
