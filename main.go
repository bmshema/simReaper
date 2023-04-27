package main

import (
	"fmt"
)

func main() {
	StartCommands()
	fmt.Printf("\nIdentifiers:\n\n")
	fmt.Printf("Type    Value\n")
	fmt.Println("---------------------------")
	fmt.Printf("IMSI:   %v\n", CleanImsi(GetImsi()))
	fmt.Printf("ICCID: %v\n", CleanIccid(GetIccid()))
	fmt.Printf("MSISDN: %v\n\n", CleanMsisdn(GetMsisdn()))
	fmt.Printf("Contacts:\n\n")
	fmt.Printf("#  Number        Name\n")
	fmt.Println("---------------------------")
	fmt.Println(CleanContacts(GetContacts()))
}
