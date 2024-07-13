package main

import "fmt"

func main() {

	// Initialize slice integer variable with values
	var vlanSlice = []int{8, 43, 3}

	// Loop through slice and print values
	for _, v := range vlanSlice {
		fmt.Printf("VLAN ID: %d\n", v)
	}

	// Append value to the slice and assign to a new variable
	newVlanSlice := append(vlanSlice, 66)
	fmt.Printf("VLAN slice after append: %d\n", newVlanSlice)
	fmt.Printf("VLAN slice length: %d\n", len(newVlanSlice))

	// Make function to create a slice with a capacity of 50
	// and a length of 2
	fmt.Println("--- Slice Make example ---")
	prealloactedSlice := make([]int, 2, 50)
	fmt.Printf("prealloactedSlice cap is %d, len is %d\n", cap(prealloactedSlice), len(prealloactedSlice))

	// MAPS
	fmt.Println("--- Maps example ----")
	vlanMap := map[string]int{}
	vlanMap["VLAN_100"] = 100
	vlanMap["VLAN_200"] = 200
	vlanMap["VLAN_300"] = 300

	delete(vlanMap, "VLAN_200")

	for k, v := range vlanMap {
		fmt.Printf("Key: %s, Value: %d\n", k, v)

	}

	// Retreive a key that doesn't exist
	fmt.Printf("The VLAN_400 key doesn't exist: %d\n", vlanMap["VLAN_400"])

	// Test whether key exists
	if val, found := vlanMap["VLAN_300"]; found {
		fmt.Printf("Test found vlan value: %d\n", val)

	}
}
