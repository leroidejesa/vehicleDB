package vehicles_test

import (
	"fmt"

	"leroi-training/vstruct"
)

var v []vstruct.Vehicle

func ExampleAdd() {
	v1 := vstruct.Vehicle{Year: 2003, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
	v2 := vstruct.Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 2}

	v = append(v, v1)
	fmt.Println(v)

	v = append(v, v2)
	fmt.Println(v)
	// Output:
	// [{2003 Toyota Tacoma 1}]
	// [{2003 Toyota Tacoma 1} {2002 Toyota Tacoma 2}]
}
