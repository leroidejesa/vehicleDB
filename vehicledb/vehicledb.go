package vehicledb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	dbLock sync.Mutex
)

func getDbConnection() (*sql.DB, error) {
	dbLock.Lock()
	defer dbLock.Unlock()

	if db != nil {
		return db, nil
	}
	connPool, err := sql.Open("postgres", "user=dealerpeak password=go@dmin dbname=leroi sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db = connPool

	return db, nil
}

// DbInsert loads pre-parsed (see pkg "leroi-training/vehicles") data into postgresql database
func DbInsert(y int, mk string, md string, st string) (err error) {
	db, err := getDbConnection()
	if err != nil {
		return errors.New("Something went wrong:" + err.Error())
	}
	_, err = db.Exec("INSERT INTO vehicleinfo(year,make,model,stocknumber) VALUES($1,$2,$3,$4);", y, mk, md, st)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// DbPhotoInsert loads image filepaths into database
func DbPhotoInsert(stock string, p string) (err error) {
	db, err := getDbConnection()
	if err != nil {
		return errors.New("Something went wrong:" + err.Error())
	}
	_, err = db.Exec("INSERT INTO vehiclephotos(stocknumber,path) VALUES($1,$2);", stock, p)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// DbQueryAll lists out all the vehicles in the database
func DbQueryAll() {
	fmt.Println("# Query")
	rows, err := db.Query("SELECT * FROM vehicleinfo")

	for rows.Next() {
		var id int
		var year int
		var make string
		var model string
		var stocknumber string
		err = rows.Scan(&id, &year, &make, &model, &stocknumber)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(year, make, model, stocknumber)
	}
}

// DbQueryPhotos lists out all the photo paths in the database
func DbQueryPhotos() {
	fmt.Println("# Photo Query")
	rows, err := db.Query("SELECT * FROM vehiclephotos")

	for rows.Next() {
		var id int
		var stocknumber string
		var filepath string
		err = rows.Scan(&id, &stocknumber, &filepath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, stocknumber, filepath)
	}
}
