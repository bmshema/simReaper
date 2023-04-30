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
	replacer := strings.NewReplacer("\t", "", "\r", "", "\n", "", "OK", "")
	msisdnResponse := string(buffer[:n])
	msisdnResponse = replacer.Replace(msisdnResponse)

	if strings.HasPrefix(msisdnResponse, "+CNUM:") {
		port.Flush()
		return msisdnResponse
	} else {
		port.Flush()
		return string("Not Stored")
	}
}
