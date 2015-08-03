package vehicles_test

import "testing"

type Vehicle struct {
  Year int
  Make string
  Model string
  Stocknumber int
}

func TestStock(t *testing.T) {
    v := Vehicle{Year: 2002, Make: "Toyota", Model: "Tacoma", Stocknumber: 1}
    if v.Stocknumber != 1 {
      t.Error("Expected 1, got ", v.Stocknumber)
    }
}
