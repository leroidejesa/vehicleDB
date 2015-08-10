package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"leroi-training/vehicles"
)

// var MyServerName = "http://127.0.0.1:8080/"

func vehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	byte, err := vehicles.ListAsJson()
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
	}

	log.Println("Full Vehicle List Retrieved.")
	w.Write(byte)
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
	// // command line flags
	// // port := flag.Int("port", 80, "port to serve on")/
	// dir := flag.String("directory", "web/", "directory of web files")
	// flag.Parse()
	//
	// // handle all requests by serving a file of the same name
	// fs := http.Dir(*dir)
	// fileHandler := http.FileServer(fs)

	vehicles.ImportPhotoData("$HOME/Documents/exercise/photos/")
	// OPTIONAL. Default: Cache Into Memory
	vehicles.ImportToDb()

	r := mux.NewRouter()
	// r.Handle("/", http.RedirectHandler("/static/", 302))
	r.HandleFunc("/", vehiclesHandler)
	r.HandleFunc("/{stockpath}", vehicleStock)
	http.Handle("/", r)

	http.Handle("/{stockpath}", r)
	log.Println("Now serving on http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", nil)
}
