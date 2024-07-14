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

	Switch01 := Switch{
		hostname: "Switch01",
		vlans:    []vlan{myVLAN},
	}

	Router01 := Router{
		hostname: "Router01",
		vrfs:     []string{"wireless", "dmz"},
	}

	// Set values after instantiation
	// myVLAN.name = "VLAN_200"
	// myVLAN.id = 200

	// fmt.Println(myVLAN)
	// fmt.Println(Switch01)

	// Switch01.printHostname()
	// Switch01.setHostname("switch100")
	// Switch01.printHostname()

	printHostname(Router01)

	printHostname(Switch01)

}

// Struct definition definition
// Must initialize the struct after its creation
type vlan struct {
	id   uint
	name string
}

type Switch struct {
	hostname string
	vlans    []vlan
}

type Router struct {
	hostname string
	vrfs     []string
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

func (r Router) printHostname() {
	fmt.Println(r.hostname)
}

func (d *Router) setHostname(hostname string) {
	if len(d.hostname) > 10 {
		hostname = hostname[:10]
	}

	d.hostname = hostname
}

type HostNamer interface {
	GetHostname() string
}

func (r Router) GetHostname() string {
	return r.hostname
}

func (s Switch) GetHostname() string {
	return s.hostname + " " + "I'm as switch"
}

func printHostname(device HostNamer) {
	fmt.Printf("The hostname is %s\n", device.GetHostname())
}
