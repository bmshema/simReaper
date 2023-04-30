package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func GetApn() string {
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

	// Write APN and PDP Context Command
	imsi := "AT+CGDCONT?\r\n"
	_, err = port.Write([]byte(imsi))
	if err != nil {
		fmt.Println("Cannot connect with device...")
		log.Fatal(err)
	}

	// Reads modem Responses
	buffer := make([]byte, 256)
	n, err := port.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// Print Response
	apnResponse := string(buffer[:n])
	port.Flush()
	return apnResponse
}
