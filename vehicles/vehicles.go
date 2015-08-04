package vehicles

import (
	"fmt"

	"leroi-training/vstruct"
)

var vList []vstruct.Vehicle

func Add() {
	v1 := vstruct.Vehicle{Year: 2003, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
	v2 := vstruct.Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 2}

	vList = append(vList, v1, v2)
}

func List() {
	fmt.Println("Vehicle List:")
	for _, vehicle := range vList {
		fmt.Println(vehicle)
	}
}
