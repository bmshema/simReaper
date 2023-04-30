package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func StartCommands() {
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

	// Disables AT command echo
	echoOff := "ATE0\r\n"
	_, err = port.Write([]byte(echoOff))
	if err != nil {
		fmt.Println("Cannot connect with device...")
		log.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	// Enables SIM Query Mode Status
	queryMode := "AT#QSS=1\r\n"
	_, err = port.Write([]byte(queryMode))
	if err != nil {
		fmt.Println("Cannot connect with device...")
		log.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	// // Changes storage mode to SM
	// sm := "AT+CPBS=\"SM\"\r\n"
	// _, err = port.Write([]byte(sm))
	// if err != nil {
	// 	fmt.Println("Cannot connect with device...")
	// 	log.Fatal(err)
	// }

	// time.Sleep(100 * time.Millisecond)

	port.Flush()
}
