package vehicledb

import (
	"database/sql"
	"fmt"
	"log"
	// go-plus is chastising me here for inserting
	// a blank import in a non-main or test package.
	// is there any important reason why?
	_ "github.com/lib/pq"
)

const (
	dbUser     = "dealerpeak"
	dbPassword = "go@dmin"
	dbName     = "leroi"
)

var year int
var make string
var model string
var stocknumber string

// DbInsert loads pre-parsed (see pkg "leroi-training/vehicles") data into postgresql database
func DbInsert(y int, mk string, md string, st string) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	var lastInsertStocknum int
	err = db.QueryRow("INSERT INTO vehicleinfo(year,make,model,stocknumber) VALUES($1,$2,$3,$4) returning id;", y, mk, md, st).Scan(&lastInsertStocknum)
	if err != nil {
		log.Fatal(err)
	}
}

// DbQueryAll lists out all the vehicles in the database
func DbQueryAll() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

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
