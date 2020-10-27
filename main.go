package main

import (
	"log"
	"net/http"

	"github.com/KiranGosavi/webservice/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/KiranGosavi/webservice/product"
)

const apiBasePath = "/api"

func main() {

	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
