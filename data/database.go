package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		Printfln("Driver: %v", driver)
	}
}

var insertNewCategory *sql.Stmt
var changeProductCategory *sql.Stmt

func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "products.db")
	if err == nil {
		// 697 Chapter 26 ■ Working with Databases
		Printfln("Opened database")
		// prepare some statements for using later
		insertNewCategory, _ = db.Prepare("INSERT INTO Categories (Name) VALUES (?)")
		changeProductCategory, _ =
			db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")
	}
	return
}
