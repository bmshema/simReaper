package modem

import (
	"fmt"
	"time"
)

// GetIMSI retrieves the IMSI from the modem
func (c *ModemClient) GetIMSI() (string, error) {
	if err := c.Open(); err != nil {
		return "", err
	}
	defer c.Close()

	return c.SendCommand(CmdGetIMSI)
}

// GetICCID retrieves the ICCID from the modem
func (c *ModemClient) GetICCID() (string, error) {
	if err := c.Open(); err != nil {
		return "", err
	}
	defer c.Close()

	return c.SendCommand(CmdGetICCID)
}

// GetMSISDN retrieves the MSISDN from the modem
func (c *ModemClient) GetMSISDN() (string, error) {
	if err := c.Open(); err != nil {
		return "", err
	}
	defer c.Close()

	return c.SendCommand(CmdGetMSISDN)
}

// GetAPN retrieves the APN settings from the modem
func (c *ModemClient) GetAPN() (string, error) {
	if err := c.Open(); err != nil {
		return "", err
	}
	defer c.Close()

	return c.SendCommand(CmdGetAPN)
}

// GetContacts retrieves contacts from the SIM
func (c *ModemClient) GetContacts() (string, error) {
	if err := c.Open(); err != nil {
		return "", err
	}
	defer c.Close()

	// Set storage to SIM memory
	_, err := c.SendCommand(CmdSetStorageSM)
	if err != nil {
		return "", fmt.Errorf("failed to set storage mode: %w", err)
	}

	time.Sleep(100 * time.Millisecond)

	// Read contacts
	return c.SendCommand(CmdReadContacts)
}

// GetSMS retrieves SMS messages from the SIM
func (c *ModemClient) GetSMS() (string, error) {
	if err := c.Open(); err != nil {
		return "", err
	}
	defer c.Close()

	// Set storage to SIM memory
	_, err := c.SendCommand(CmdSetStorageSM)
	if err != nil {
		return "", fmt.Errorf("failed to set storage mode: %w", err)
	}

	time.Sleep(100 * time.Millisecond)

	// Read SMS messages
	return c.SendCommand(CmdReadSMS)
}
