package vehicles

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Vehicle is a struct which will help us organize basic vehicle data
type Vehicle struct {
	Year        int    `json:"year"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	Stocknumber string `json:"stocknumber"`
}

// Vlist is the primary data cache
var Vlist = []Vehicle{
	Vehicle{
		Year:        2003,
		Make:        "Nissan",
		Model:       "Frontier",
		Stocknumber: "test",
	},
	Vehicle{
		Year:        2002,
		Make:        "Toyota",
		Model:       "Tacoma",
		Stocknumber: "test",
	},
}

// Vslices holds pre-parsed slices of vehicle data
var Vslices [][]string

// push places photo files into a multidimensional slice (Vslices) for func parseFiles()
func push(p []string) {
	Vslices = append(Vslices, p)
}

// ImportPhotoData is an exported function that
// retrieves files from a given directory.
// Possibly consolidate/refactor into one function?
func ImportPhotoData(root string) {
	dir := os.ExpandEnv(root)
	numImported := 0
	var walkpath = func(path string, f os.FileInfo, err error) error {
		// Considering using ioutil.ReadDir() instead of filepath.Walk. Tabled for now.
		if f.IsDir() {
			return nil
		}
		pathinfo := strings.Split(f.Name(), "_")
		push(pathinfo)
		numImported++
		return nil
	}
	defer parseFiles()
	filepath.Walk(dir, walkpath)
	fmt.Printf("Imported %d Files\n", numImported)
}

// organizes import results and adds vehicles to memory in proper format.
func parseFiles() {
	for _, file := range Vslices {
		year, _ := strconv.Atoi(file[0])
		make := file[1]
		model := file[2]
		// in order to trim off the '-[imgview int].jpg' :
		stockfile := file[3]
		stockExt := filepath.Ext(stockfile)
		stock := stockfile[0 : len(stockfile)-(len(stockExt)+2)]
		Add(Vehicle{Year: year, Make: make, Model: model, Stocknumber: stock})
	}
}

// Add a new Vehicle to the list
func Add(newVehicle Vehicle) {
	Vlist = append(Vlist, newVehicle)
}

// Retrieve a Vehicle by stocknumber.
// Example output:
// {2003 Nissan Frontier 1}
func Retrieve(s string) []Vehicle {
	var result Vehicle
	for _, vehicle := range Vlist {
		if vehicle.Stocknumber == s {
			result = vehicle
		}
	}

	var resultSlice []Vehicle
	resultSlice = append(resultSlice, result)
	return resultSlice
}

// List returns complete vehicle list.
func List() {
	fmt.Println("Vehicle List:")
	for _, vehicle := range Vlist {
		fmt.Println(vehicle)
	}
}
