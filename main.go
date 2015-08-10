package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"leroi-training/vehicles"
)

func vehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("Full Vehicle List Retrieved.")
	w.Write(vehicles.ListAsJson())
}

// variable for mux
var stockpath string

func vehicleStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stock := r.URL.Path[1:]
	stockresult, err := json.Marshal(vehicles.Retrieve(stock))

	if err != nil {
		return
	}

	log.Printf("Vehicle %s Successfully Retrieved.", stock)
	w.Write(stockresult)

	stockpath = stock
}

func main() {
	vehicles.ImportPhotoData("$HOME/Documents/exercise/photos/")
	// vehicles.ImportToDb() OPTIONAL. Default: Cache Into Memory

	r := mux.NewRouter()
	r.HandleFunc("/", vehiclesHandler)
	r.HandleFunc("/{stockpath}", vehicleStock)
	http.Handle("/", r)
	http.Handle("/{stockpath}", r)
	log.Println("Now serving on http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}
