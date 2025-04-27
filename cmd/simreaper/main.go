package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bmshema/simReaper/pkg/display"
	"github.com/bmshema/simReaper/pkg/modem"
)

const (
	DefaultDevicePath = "/dev/ttyUSB2"
	BannerPath        = "banner.txt"
)

func main() {
	// Set up logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stderr)

	// Parse command line arguments
	devicePath := DefaultDevicePath
	if len(os.Args) > 1 {
		devicePath = os.Args[1]
	}

	// Create modem client
	client := modem.NewModemClient(devicePath)

	// Initialize display
	err := display.ClearScreen()
	if err != nil {
		log.Printf("Warning: could not clear screen: %v", err)
	}

	// Load and display banner
	banner, err := display.LoadBanner(BannerPath)
	if err != nil {
		log.Printf("Warning: could not load banner: %v", err)
	} else {
		fmt.Println(banner)
	}

	// Initialize modem
	err = client.InitializeModem()
	if err != nil {
		log.Fatalf("Failed to initialize modem: %v", err)
	}

	// Get and display identifiers
	imsi, err := client.GetIMSI()
	if err != nil {
		log.Fatalf("Failed to get IMSI: %v", err)
	}
	formattedIMSI := display.FormatIMSI(imsi)

	iccid, err := client.GetICCID()
	if err != nil {
		log.Fatalf("Failed to get ICCID: %v", err)
	}
	formattedICCID := display.FormatICCID(iccid)

	msisdn, err := client.GetMSISDN()
	if err != nil {
		log.Fatalf("Failed to get MSISDN: %v", err)
	}
	formattedMSISDN := display.FormatMSISDN(msisdn)

	display.DisplayIdentifiers(formattedIMSI, formattedICCID, formattedMSISDN)

	// Get and display contacts
	contacts, err := client.GetContacts()
	if err != nil {
		log.Printf("Warning: failed to get contacts: %v", err)
	} else {
		formattedContacts := display.FormatContacts(contacts)
		display.DisplayContacts(formattedContacts)
	}

	// Get and display SMS
	sms, err := client.GetSMS()
	if err != nil {
		log.Printf("Warning: failed to get SMS: %v", err)
	} else {
		formattedSMS := display.FormatSMS(sms)
		display.DisplaySMS(formattedSMS)
	}

	// Get and display APN
	apn, err := client.GetAPN()
	if err != nil {
		log.Printf("Warning: failed to get APN: %v", err)
	} else {
		formattedAPN := display.FormatAPN(apn)
		display.DisplayAPN(formattedAPN)
	}
}
