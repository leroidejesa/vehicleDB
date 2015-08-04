package vehicles

import (
	"fmt"

	"leroi-training/vstruct"
)

var vList = []vstruct.Vehicle{
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

// Retrieve returns a vehicle by stocknumber.
// Example output:
// {2003 Nissan Frontier 1}
func Retrieve(s int) vstruct.Vehicle {
	var result vstruct.Vehicle
	for _, vehicle := range vList {
		if vehicle.Stocknumber == s {
			result = vehicle
		}
	}
	return result
}

// List returns complete vehicle list.
func List() {
	fmt.Println("Vehicle List:")
	for _, vehicle := range vList {
		fmt.Println(vehicle)
	}
}
