package main

import (
	API "forms-api/api"
	CONFIG "forms-api/config"
	DB "forms-api/database"
	"os"
	"strings"

	_ "forms-api/docs"
)

// @title Brankas Task - Form API
// @version 1.0
// @description This is an api to submit form
func main() {

	CONFIG.LoadConfig()
	DB.ConnectDatabase()

	//if not prod, setup database tables
	if !strings.EqualFold(os.Getenv("prod"), "true") {
		DB.SetupDB()
	}

	API.StartServer()
}
