package vehicles_test

import (
	"fmt"
	"testing"

	"leroi-training/vstruct"
)

var TestList = []vstruct.Vehicle{
	vstruct.Vehicle{
		Year:        2003,
		Make:        "Nissan",
		Model:       "Frontier",
		Stocknumber: 1,
	},
	vstruct.Vehicle{
		Year:        2002,
		Make:        "Toyota",
		Model:       "Tacoma",
		Stocknumber: 2,
	},
}

func TestAdd(t *testing.T) {
	TestVehicle := vstruct.Vehicle{
		Year:        1996,
		Make:        "Honda",
		Model:       "Passport",
		Stocknumber: 3,
	}
	TestList = append(TestList, TestVehicle)

	fmt.Println("Vehicle List:")
	for _, vehicle := range TestList {
		fmt.Println(vehicle)
	}
}
