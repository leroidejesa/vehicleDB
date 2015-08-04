package vstruct_test

import (
	"testing"

	"leroi-training/vstruct"
)

func TestStock(t *testing.T) {
	v := vstruct.Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
	if v.Stocknumber != 1 {
		t.Error("Expected 1, got ", v.Stocknumber)
	}
}
