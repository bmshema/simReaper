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
	fmt.Println("---------------------------")
	fmt.Printf("IMSI:   %v\n", CleanImsi(GetImsi()))
	fmt.Printf("ICCID: %v\n", CleanIccid(GetIccid()))
	fmt.Printf("MSISDN: %v\n", CleanMsisdn(GetMsisdn()))
	fmt.Println("---------------------------")
	fmt.Printf("\nContacts:\n\n")
	fmt.Printf("#  Number        Name\n")
	fmt.Println("---------------------------")
	fmt.Println(CleanContacts(GetContacts()))
	fmt.Println("---------------------------")
}
