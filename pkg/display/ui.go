package display

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

// init initializes the clear/cls commands for the appropriate runtime environment
func init() {
	clear = make(map[string]func()) // Initialize
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	// Add Windows support if needed
	// clear["windows"] = func() {
	//     cmd := exec.Command("cmd", "/c", "cls")
	//     cmd.Stdout = os.Stdout
	//     cmd.Run()
	// }
}

// ClearScreen clears the terminal window
func ClearScreen() error {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
		return nil
	}
	return fmt.Errorf("unsupported platform for screen clearing: %s", runtime.GOOS)
}

// LoadBanner loads and returns the banner text
func LoadBanner(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to load banner: %w", err)
	}
	return string(b), nil
}

// DisplayIdentifiers prints SIM identifiers in a formatted way
func DisplayIdentifiers(imsi, iccid, msisdn string) {
	fmt.Printf("Identifiers:\n\n")
	fmt.Printf("Type    Value\n")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("IMSI:   %v\n", imsi)
	fmt.Printf("ICCID:  %v\n", iccid)
	fmt.Printf("MSISDN: %v\n", msisdn)
	fmt.Println("-------------------------------------------------")
}

// DisplayContacts prints SIM contacts in a formatted way
func DisplayContacts(contacts string) {
	fmt.Printf("\n\nContacts:\n\n")
	fmt.Printf("#  Number        Name\n")
	fmt.Println("-------------------------------------------------")
	fmt.Println(contacts)
	fmt.Println("-------------------------------------------------")
}

// DisplaySMS prints SMS messages in a formatted way
func DisplaySMS(sms string) {
	fmt.Printf("\nStored SMS Messages: \n")
	fmt.Println(sms)
	fmt.Println("-------------------------------------------------")
}

// DisplayAPN prints APN settings in a formatted way
func DisplayAPN(apn string) {
	fmt.Printf("\nAPN Settings: \n")
	fmt.Println(apn)
}
