package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func GetMsisdn() string {
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
	imsi := "AT+CNUM\r\n"
	_, err = port.Write([]byte(imsi))
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
	msisdnResponse := string(buffer[:n])
	msisdnResponse = strings.Replace(msisdnResponse, "\t", "", -1)
	msisdnResponse = strings.Replace(msisdnResponse, "\r", "", -1)
	msisdnResponse = strings.Replace(msisdnResponse, "\n", "", -1)
	msisdnResponse = strings.Replace(msisdnResponse, "OK", "", -1)

	if strings.HasPrefix(msisdnResponse, "+CNUM:") {
		port.Flush()
		return msisdnResponse
	} else {
		port.Flush()
		return string("Not Stored")
	}
}
