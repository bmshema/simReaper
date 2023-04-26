package main

import (
	"fmt"
)

func main() {
	StartCommands()

	fmt.Printf("IMSI:   %s\n", CleanImsi(ImsiCommand()))

	fmt.Printf("ICCID: %v\n", CleanIccid(IccidCommand()))

}
