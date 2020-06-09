package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "dbuser:dbpass@tcp(127.0.0.1:3306)/xgproj")

	if err != nil {
		fmt.Printf("error opening db: %v", err)
		panic(err.Error)
	}

	fmt.Printf("closing db ...")
	defer db.Close()

	results, err := db.Query("SELECT uid, name FROM customers")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var uid, name string
		err = results.Scan(&uid, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("\tuid: %s, name: %s\n", uid, name)
	}
}
