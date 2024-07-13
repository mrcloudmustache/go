package main

import "fmt"

func main() {

	// Loop basic exmaple
	for i := 1; i <= 5; i++ {
		fmt.Printf("VLAN %d\n", i)
	}

	// Loop continue example
	fmt.Println("--- Continue example ---")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Printf("VLAN %d\n", i)
	}

	// Loop break example
	fmt.Println("--- Break example ---")
	var vlanID int = 1

	for {
		fmt.Printf("Looking at VLAN %d\n", vlanID)

		if vlanID > 5 {
			break
		}

		vlanID++
	}

}
