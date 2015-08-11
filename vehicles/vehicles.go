package vehicles

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"leroi-training/vehicledb"
)

// Vehicle is a struct which will help us organize basic vehicle data
type Vehicle struct {
	Year        int                 `json:"year"`
	Make        string              `json:"make"`
	Model       string              `json:"model"`
	Stocknumber string              `json:"stocknumber"`
	Images      map[string][]string `json:"images"`
}

// ImagePaths is a map of stocknumbers to image filepaths
// var ImagePaths = make(map[string][]string)

// vList is the primary data cache (starts with a couple dummy vehicles)
var vList = make([]Vehicle, 0)

// these temporarily hold vehicle data during ImportPhotoData()
var tempVehicle Vehicle
var tempYear int
var tempMake string
var tempModel string
var tempStock string
var tempImg string
var tempPath string

// ImportPhotoData retrieves files from a given directory.
func ImportPhotoData(root string) {
	dir := os.ExpandEnv(root)
	numImported := 0
	var walkpath = func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		pathinfo := strings.Split(f.Name(), "_")
		tempYear, _ = strconv.Atoi(pathinfo[0])
		tempMake = pathinfo[1]
		tempModel = pathinfo[2]
		// in order to trim off the '-[imgview int].jpg' :
		stockfile := pathinfo[3]
		stockExt := filepath.Ext(stockfile)
		tempStock = stockfile[0 : len(stockfile)-(len(stockExt)+2)]
		tempMap := make(map[string][]string)
		tempPath = path
		tempVehicle = Vehicle{Year: tempYear, Make: tempMake, Model: tempModel, Stocknumber: tempStock, Images: tempMap}
		tempVehicle.Images[tempStock] = append(tempVehicle.Images[tempStock], path)

		Add(tempVehicle)
		numImported++

		return nil
	}

	filepath.Walk(dir, walkpath)
	fmt.Printf("Imported %d Files.\nNote: To insert imported data into local database, use vehicles.ImportToDb().\n", numImported)
}

// ImportToDb is the next (optional) step after ImportPhotoData() that loads parsed files into database.
// func ImportToDb() {
// 	for _, item := range vList {
// 		yr := item.Year
// 		mk := item.Make
// 		md := item.Model
// 		st := item.Stocknumber
// 		vehicledb.DbInsert(yr, mk, md, st)
// 		for _, p := range ImagePaths[st] {
// 			vehicledb.DbPhotoInsert(st, p)
// 		}
// 	}
// }

// DbList queries the database to list all vehicles.
func DbList() {
	vehicledb.DbQueryAll()
	vehicledb.DbQueryPhotos()
}

// used to search/filter slices for a given string (e.g. stocknumber)
func contains(sl []Vehicle, st string) bool {
	for _, v := range sl {
		if v.Stocknumber == st {
			return true
		}
	}
	return false
}

// Add a new Vehicle to cached list
func Add(newVehicle Vehicle) {
	if !contains(vList, newVehicle.Stocknumber) {
		vList = append(vList, newVehicle)
		fmt.Println("Vehicle Added!")
	} else {
		for _, item := range vList {
			if item.Stocknumber == newVehicle.Stocknumber {
				item.Images[newVehicle.Stocknumber] = append(item.Images[newVehicle.Stocknumber], tempPath)
			}
		}
	}
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

// ListAsJson turns vList(type []Vehicle) into []byte so can
// be used by HTTP handler (e.g. w.Write(vehicles.ListAsJson()))
func ListAsJson() ([]byte, error) {
	byte, err := json.Marshal(vList)
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
	}

	return byte, err
}
