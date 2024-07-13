package main

import (
	"errors"
	"fmt"
)

func main() {

	myVLAN, err := NewVLAN(500, "VLAN_500")
	if err != nil {
		fmt.Println(err)
		return
	}

	Switch01 := device{
		hostname: "switch01",
		vlans:    []vlan{myVLAN},
	}

	// Set values after instantiation
	// myVLAN.name = "VLAN_200"
	// myVLAN.id = 200

	fmt.Println(myVLAN)
	fmt.Println(Switch01)

	Switch01.printHostname()

}

// Struct definition definition
// Must initialize the struct after its creation
type vlan struct {
	id   uint
	name string
}

type device struct {
	hostname string
	vlans    []vlan
}

func NewVLAN(id uint, name string) (vlan, error) {

	if id > 4096 {
		return vlan{}, errors.New("VLAN ID must be <= 4096")
	}

	return vlan{
		id:   id,
		name: name,
	}, nil
}

func (d device) printHostname() {
	fmt.Println(d.hostname)
}
