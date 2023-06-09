package main

import (
	"fmt"
)

func main() {
	CallClear()
	fmt.Println(Banner())
	StartCommands()
	fmt.Printf("Identifiers:\n\n")
	fmt.Printf("Type    Value\n")
	fmt.Println("-------------------------------------------------")
	fmt.Printf("IMSI:   %v\n", CleanImsi(GetImsi()))
	fmt.Printf("ICCID: %v\n", CleanIccid(GetIccid()))
	fmt.Printf("MSISDN: %v\n", CleanMsisdn(GetMsisdn()))
	fmt.Println("-------------------------------------------------")
	fmt.Printf("\n\nContacts:\n\n")
	fmt.Printf("#  Number        Name\n")
	fmt.Println("-------------------------------------------------")
	fmt.Println(CleanContacts(GetContacts()))
	fmt.Println("-------------------------------------------------")
	fmt.Printf("\nStored SMS Messages: \n")
	fmt.Println(GetSms())
	fmt.Println("-------------------------------------------------")
	fmt.Printf("\nAPN Settings: \n")
	fmt.Println(CleanApn(GetApn()))
}
