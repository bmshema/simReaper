package display

import (
	"strings"
)

// FormatIMSI formats the IMSI response
func FormatIMSI(response string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "\n", "", "AT+CIMI", "")
	return replacer.Replace(response)
}

// FormatICCID formats the ICCID response
func FormatICCID(response string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "\n", "", "+ICCID:", "")
	return replacer.Replace(response)
}

// FormatContacts formats the contacts response
func FormatContacts(response string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "", "+CPBR: ", "", "\"", "", ",", "  ", "129", "", "0", "")
	return replacer.Replace(response)
}

// FormatMSISDN formats the MSISDN response
func FormatMSISDN(response string) string {
	replacer := strings.NewReplacer("\"", "", ",", "", "129", "", "+CNUM: ", "")
	return replacer.Replace(response)
}

// FormatAPN formats the APN response
func FormatAPN(response string) string {
	replacer := strings.NewReplacer("OK", "", "\t", "", "\r", "")
	return replacer.Replace(response)
}

// FormatSMS formats the SMS response
func FormatSMS(response string) string {
	replacer := strings.NewReplacer("\t", "", "\r", "", "\n", "")
	formatted := replacer.Replace(response)

	if strings.HasPrefix(formatted, "") {
		return "\nNo messages stored...\n"
	}
	return formatted
}
