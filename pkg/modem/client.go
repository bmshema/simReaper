package modem

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tarm/serial"
)

// Command constants
const (
	CmdEchoOff      = "ATE0\r\n"
	CmdQueryMode    = "AT#QSS=1\r\n"
	CmdGetIMSI      = "AT+CIMI\r\n"
	CmdGetICCID     = "AT+ICCID\r\n"
	CmdGetMSISDN    = "AT+CNUM\r\n"
	CmdGetAPN       = "AT+CGDCONT?\r\n"
	CmdSetStorageSM = "AT+CPBS=\"SM\"\r\n"
	CmdReadContacts = "AT+CPBR=1,250\r\n"
	CmdReadSMS      = "AT+CMGL=\"ALL\"\r\n"
)

// ModemClient represents a connection to a modem device
type ModemClient struct {
	port       *serial.Port
	config     *serial.Config
	isOpen     bool
	devicePath string
}

// NewModemClient creates a new modem client with default configuration
func NewModemClient(devicePath string) *ModemClient {
	return &ModemClient{
		devicePath: devicePath,
		config: &serial.Config{
			Name:        devicePath,
			Baud:        115200,
			ReadTimeout: time.Millisecond * 200,
			Size:        8,
			Parity:      serial.ParityNone,
			StopBits:    serial.Stop1,
		},
		isOpen: false,
	}
}

// Open establishes a connection to the modem
func (c *ModemClient) Open() error {
	if c.isOpen {
		return nil
	}

	port, err := serial.OpenPort(c.config)
	if err != nil {
		return fmt.Errorf("failed to open serial port: %w", err)
	}

	c.port = port
	c.isOpen = true
	return nil
}

// Close closes the connection to the modem
func (c *ModemClient) Close() error {
	if !c.isOpen || c.port == nil {
		return nil
	}

	err := c.port.Close()
	if err != nil {
		return fmt.Errorf("failed to close serial port: %w", err)
	}

	c.isOpen = false
	c.port = nil
	return nil
}

// SendCommand sends an AT command to the modem and returns the response
func (c *ModemClient) SendCommand(cmd string) (string, error) {
	if !c.isOpen || c.port == nil {
		return "", errors.New("modem connection not open")
	}

	// Write command to modem
	_, err := c.port.Write([]byte(cmd))
	if err != nil {
		return "", fmt.Errorf("failed to send command: %w", err)
	}

	// Read response with a more robust approach
	var response strings.Builder
	buffer := make([]byte, 256)

	// Wait for complete response
	time.Sleep(300 * time.Millisecond)

	// Read in chunks until timeout or buffer is empty
	for {
		n, err := c.port.Read(buffer)
		if err != nil {
			// Timeout or EOF is normal at the end of the response
			break
		}
		if n == 0 {
			break
		}
		response.Write(buffer[:n])

		// Check if response is complete (ends with OK or ERROR)
		if strings.Contains(response.String(), "OK\r\n") ||
			strings.Contains(response.String(), "ERROR\r\n") {
			break
		}

		// Short pause between reads
		time.Sleep(50 * time.Millisecond)
	}

	c.port.Flush()
	return response.String(), nil
}

// InitializeModem sets up the modem with basic configuration
func (c *ModemClient) InitializeModem() error {
	if err := c.Open(); err != nil {
		return err
	}

	// Disable command echo
	_, err := c.SendCommand(CmdEchoOff)
	if err != nil {
		return fmt.Errorf("failed to disable echo: %w", err)
	}

	time.Sleep(100 * time.Millisecond)

	// Enable SIM Query Mode
	_, err = c.SendCommand(CmdQueryMode)
	if err != nil {
		return fmt.Errorf("failed to set query mode: %w", err)
	}

	time.Sleep(100 * time.Millisecond)

	return nil
}
