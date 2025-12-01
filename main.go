package main

import (
	"admin_apis/config"
	"admin_apis/database"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func main() {
	fmt.Println("Admin APIs Code")

	config.InitiateConfig()

	fmt.Println("Environment:", *config.ENV)

	db, err := database.Connect()
	if err != nil {
		fmt.Println("Database error:", err)
		return
	}

	fmt.Println("Database connected successfully!")

	config.DB = db // Put DB into config package
	defer config.DB.Close()
}
