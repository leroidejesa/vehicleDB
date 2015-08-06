package main

import (
	"fmt"
	"leroi-training/vehicles"
)

func main() {
	vehicles.ImportPhotoData("$HOME/Documents/exercise/photos/")
	fmt.Println(vehicles.Retrieve("PW7165"))
}
