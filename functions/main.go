package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// doSomething()
	// printMessage("I'm printing a message bro!")
	// fmt.Println(totalIPv4Addresses(24))
	createVLAN(45)
}

func doSomething() {
	fmt.Println("I'm doing something bro!")
}

func printMessage(msg string) {
	fmt.Println(msg)
}

func totalIPv4Addresses(prefixlen int) int {
	x := 32 - prefixlen

	addrCount := math.Pow(2.0, float64(x))
	return int(addrCount)
}

func createVLAN(id uint) error {

	if id > 4096 {
		return errors.New("VLAN ID must be <= 4096")
	}

	fmt.Printf("Creating VLAN with ID of %d\n", id)
	return nil

}
