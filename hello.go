package main

import (
	"fmt"

	"leroi-training/vehicles"
	"leroi-training/vstruct"
)

func main() {
	v := vstruct.Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
	fmt.Printf("Hello, world! I drive a green %s!\n", v.Model)
	vehicles.Add()
	vehicles.List()
}
