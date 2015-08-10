package vehicles

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"leroi-training/vehicledb"
)

// Vehicle is a struct which will help us organize basic vehicle data
type Vehicle struct {
	Year        int    `json:"year"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	Stocknumber string `json:"stocknumber"`
}

// vList is the primary data cache
var vList = []Vehicle{
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

// vSlices holds pre-parsed slices of vehicle data
var vSlices [][]string

// push places photo files into a multidimensional slice (vSlices) for func parseAndCache()
func push(p []string) {
	vSlices = append(vSlices, p)
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
	defer parseAndCache()
	filepath.Walk(dir, walkpath)
	fmt.Printf("Imported %d Files.\nNote: To insert imported data into local database, use vehicles.ImportToDb().\n", numImported)
}

// parseAndCache() organizes import results and adds vehicles to memory in proper format.
func parseAndCache() {
	for _, file := range vSlices {
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

//make these two DRY!

// ImportToDb is the next (optional) step after ImportPhotoData() that loads parsed files into database.
func ImportToDb() {
	for _, file := range vSlices {
		year, _ := strconv.Atoi(file[0])
		make := file[1]
		model := file[2]
		// in order to trim off the '-[imgview int].jpg' :
		stockfile := file[3]
		stockExt := filepath.Ext(stockfile)
		stock := stockfile[0 : len(stockfile)-(len(stockExt)+2)]
		vehicledb.DbInsert(year, make, model, stock)
	}
}

// DbList queries the database to list all vehicles.
func DbList() {
	vehicledb.DbQueryAll()
}

// Add a new Vehicle to cached list
func Add(newVehicle Vehicle) {
	vList = append(vList, newVehicle)
}

// Retrieve a (cached) Vehicle by stocknumber.
// Example output:
// {2003 Nissan Frontier NF2301}
func Retrieve(stockRequest string) []Vehicle {
	for _, v := range vList {
		if v.Stocknumber == stockRequest {
			return []Vehicle{v}
		}
	}

	return nil
}

// List returns complete (cached) vehicle list.
func List() {
	fmt.Println("Vehicle List:")
	for _, vehicle := range vList {
		fmt.Println(vehicle)
	}
}
