package main

import (
	"fmt"

	"leroi-training/vehicles"
)

func main() {
	fmt.Printf("Hello, world!\n")
	fmt.Println(vehicles.Retrieve(2))
	vehicles.List()
}
