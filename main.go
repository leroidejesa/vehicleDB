package main

import (
	"fmt"

	"leroi-training/vehicles"
)

func main() {
	fmt.Printf("Hello, world!\n")

	newV := vehicles.Vehicle{
		Year:        1986,
		Make:        "Ferrari",
		Model:       "328",
		Stocknumber: 3,
	}

	vehicles.Add(newV)
	fmt.Println(vehicles.Retrieve(3))
	vehicles.List()
}
