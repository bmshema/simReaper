package main

import (
	"fmt"
)

func main() {
	StartCommands()

	fmt.Printf("IMSI:   %s\n", CleanImsi(ImsiCommand()))

	fmt.Printf("ICCID: %v\n\n", CleanIccid(IccidCommand()))

	fmt.Printf("Contacts:\n %v\n", CleanContacts(GetContacts()))

}
