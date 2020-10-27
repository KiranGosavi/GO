package database

import (
	"database/sql"
	"log"
	"time"
)

//DbConn exported
var DbConn *sql.DB

//SetupDatabase to set up db
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:root123@tcp(mysql:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
