package main

import (
	"fmt"
	"leroi-training/vehicles"
)

func main() {
	v := vehicles.Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
	fmt.Printf("Hello, world! I drive a green %s!\n", v.Model)
}
