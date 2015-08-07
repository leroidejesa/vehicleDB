package vehicles_test

import (
	"fmt"
	"testing"
)

type vehicle struct {
	Year        int
	Make        string
	Model       string
	Stocknumber int
}

var TestList = []vehicle{
	vehicle{
		Year:        2003,
		Make:        "Nissan",
		Model:       "Frontier",
		Stocknumber: 1,
	},
	vehicle{
		Year:        2002,
		Make:        "Toyota",
		Model:       "Tacoma",
		Stocknumber: 2,
	},
}

func TestAdd(t *testing.T) {
	TestVehicle := vehicle{
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
