package vehicledb_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "dealerpeak"
	dbPassword = "go@dmin"
	dbName     = "leroi"
)

func TestConnect(t *testing.T) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	fmt.Println("# Inserting values")

	var lastInsertStocknum int
	err = db.QueryRow("INSERT INTO vehicleinfo(year,make,model,stocknumber) VALUES($1,$2,$3,$4) returning id;", "2015", "Nissan", "Frontier", "NSFR123").Scan(&lastInsertStocknum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("last inserted vehicle =", lastInsertStocknum)

	fmt.Println("# Test Query")
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
		fmt.Println(id, year, make, model, stocknumber)
	}
}
