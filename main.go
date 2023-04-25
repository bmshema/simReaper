package main

import (
	"log"
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

}
