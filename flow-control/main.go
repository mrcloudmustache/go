package main

import "fmt"

func main() {

	snmpConfigured := true

	if !snmpConfigured {
		fmt.Println("SNMP is configured.")
	} else {
		fmt.Println("SNMP is NOT configured.")
	}

	var vlanID int = 50

	if vlanID < 100 {
		fmt.Println("vlanID is less than 100")
	} else if vlanID > 100 && vlanID < 750 {
		fmt.Println("vlanID is between 100 and 750")
	} else {
		fmt.Println("vlanID is greater than 750.")
	}

}
