package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func GetSms() string {
	config := &serial.Config{
		Name:        "/dev/ttyUSB2",
		Baud:        115200,
		ReadTimeout: time.Millisecond * 200,
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

	// Changes storage mode to SM
	sm := "AT+CPBS=\"SM\"\r\n"
	_, err = port.Write([]byte(sm))
	if err != nil {
		fmt.Println("Cannot connect with device...")
		log.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)

	// Reads Stored SMS entries
	rd := "AT+CMGL=\"ALL\"\r\n"
	_, err = port.Write([]byte(rd))
	if err != nil {
		fmt.Println("Cannot connect with device...")
		log.Fatal(err)
	}

	port.Flush()

	// Reads modem Responses
	buffer := make([]byte, 256)
	n, err := port.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// Print Response
	replacer := strings.NewReplacer("\t", "", "\r", "", "\n", "")
	sms := string(buffer[:n])
	sms = replacer.Replace(sms)

	if strings.HasPrefix(sms, "") {
		port.Flush()
		return string("\nNo messages stored...\n")
	} else {
		port.Flush()
		return sms
	}

	port.Flush()
	return sms
}
