package main

import (
	"encoding/json"
	"net/http"

	"leroi-training/vehicles"
)

func vehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	byte, err := json.Marshal(vehicles.Vlist)

	if err != nil {
		return
	}

	w.Write(byte)
}

func main() {
	vehicles.ImportPhotoData("$HOME/Documents/exercise/photos/")
	http.HandleFunc("/", vehiclesHandler)
	http.ListenAndServe(":8080", nil)
}
