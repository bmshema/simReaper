package display

import (
	"regexp"
	"strings"
)

// FormatIMSI formats the IMSI response
func FormatIMSI(response string) string {
	// Extract digits only from the response
	re := regexp.MustCompile(`\d{15,16}`)
	matches := re.FindStringSubmatch(response)
	if len(matches) > 0 {
		return matches[0]
	}

	// Fallback to string replacer if regex doesn't match
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "\n", "", "AT+CIMI", "")
	return strings.TrimSpace(replacer.Replace(response))
}

// FormatICCID formats the ICCID response
func FormatICCID(response string) string {
	// First try to extract ICCID using regex
	re := regexp.MustCompile(`\+ICCID:\s*([0-9A-F]+)`)
	matches := re.FindStringSubmatch(response)
	if len(matches) > 1 {
		return matches[1]
	}

	// Fallback to string replacer if regex doesn't match
	replacer := strings.NewReplacer("+ICCID:", "", "OK", "", "\t", "", "\r", "", "\n", "")
	return strings.TrimSpace(replacer.Replace(response))
}

// FormatContacts formats the contacts response
func FormatContacts(response string) string {
	if response == "" || strings.Contains(response, "ERROR") {
		return "No contacts found"
	}

	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "+CPBR: ", "", "\"", "", ",", "  ", "129", "", "0", "")
	return strings.TrimSpace(replacer.Replace(response))
}

// FormatMSISDN formats the MSISDN response
func FormatMSISDN(response string) string {
	// Try to extract phone number using regex
	re := regexp.MustCompile(`\+CNUM:\s*,"([^"]+)"`)
	matches := re.FindStringSubmatch(response)
	if len(matches) > 1 {
		return matches[1]
	}

	// Fallback to string replacer if regex doesn't match
	replacer := strings.NewReplacer("\"", "", ",", "", "129", "", "+CNUM:", "", "OK", "", "\r", "", "\n", "")
	return strings.TrimSpace(replacer.Replace(response))
}

// FormatAPN formats the APN response
func FormatAPN(response string) string {
	if response == "" || strings.Contains(response, "ERROR") {
		return "No APN settings found"
	}

	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "")
	return strings.TrimSpace(replacer.Replace(response))
}

// FormatSMS formats the SMS response
func FormatSMS(response string) string {
	if response == "" || strings.Contains(response, "ERROR") {
		return "No messages stored"
	}

	if !strings.Contains(response, "+CMGL:") {
		return "No messages stored"
	}

	replacer := strings.NewReplacer("\t", "", "\r", "")
	return strings.TrimSpace(replacer.Replace(response))
}
