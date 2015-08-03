package vehicles_test

import (
	"leroi-training/vehicles"
	"testing"
)

func TestStock(t *testing.T) {
	v := vehicles.Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
	if v.Stocknumber != 1 {
		t.Error("Expected 1, got ", v.Stocknumber)
	}
}
