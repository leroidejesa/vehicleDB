package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"leroi-training/vehicles"
)

// var MyServerName = "http://127.0.0.1:8080/"

// GET all vehicles JSON
func vehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	byte, err := vehicles.ListAsJson()
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Response sent.")
	w.Write(byte)
}

// GET individual vehicle JSON
func vehicleStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	stockpath := vars["stockpath"]
	stockjson, err := json.Marshal(vehicles.Retrieve(stockpath))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Vehicle %s Successfully Retrieved.", stockpath)
	w.Write(stockjson)
}

func main() {
	vehicles.ImportPhotoData("photos/")
	r := mux.NewRouter()
	r.HandleFunc("/api/vehicles/", vehiclesHandler)
	r.HandleFunc("/api/vehicles/{stockpath}", vehicleStock)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/home/leroi/dpnext/go/src/leroi-training/")))
	fmt.Println(http.ListenAndServe(":8080", r))
}
