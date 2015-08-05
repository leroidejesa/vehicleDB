package vehicles

import (
	"fmt"
)

type Vehicle struct {
	Year        int
	Make        string
	Model       string
	Stocknumber int
}

var vList = []Vehicle{
	Vehicle{
		Year:        2003,
		Make:        "Nissan",
		Model:       "Frontier",
		Stocknumber: 1,
	},
	Vehicle{
		Year:        2002,
		Make:        "Toyota",
		Model:       "Tacoma",
		Stocknumber: 2,
	},
}

// Add a new Vehicle to the list
func Add(newVehicle Vehicle) {
	vList = append(vList, newVehicle)
}

// Retrieve a Vehicle by stocknumber.
// Example output:
// {2003 Nissan Frontier 1}
func Retrieve(s int) Vehicle {
	var result Vehicle
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
