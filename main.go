package main

import (
	"fmt"
)

func main() {
	StartCommands()
	fmt.Printf("\nIdentifiers:\n\n")
	fmt.Printf("Type    Value\n")
	fmt.Println("---------------------------")
	fmt.Printf("IMSI:   %s\n", CleanImsi(ImsiCommand()))
	fmt.Printf("ICCID: %v\n", CleanIccid(IccidCommand()))
	fmt.Printf("MSISDN: \n\n")
	fmt.Printf("Contacts:\n %v\n", CleanContacts(GetContacts()))

}
