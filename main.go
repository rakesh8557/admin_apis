package main

import (
	"admin_apis/config"
	"admin_apis/database"
	"admin_apis/logger"
	"database/sql"
	"fmt"
	"os"
	"time"
)

var DB *sql.DB

func main() {
	// Set default timezone to IST
	os.Setenv("TZ", "Asia/Kolkata")
	time.Local = time.FixedZone("IST", 5*3600+1800) // +5:30 hours

	config.InitiateConfig()

	logger.InitLogger()

	logger.Log("INFO", "Application started on environment:"+*config.ENV)

	db, err := database.Connect()
	if err != nil {
		msg := fmt.Sprintf("Database Error: %v", err)
		logger.Log("ERROR", msg)
		return
	}

	logger.Log("INFO", "Database connected successfully!")

	config.DB = db // Put DB into config package
	defer config.DB.Close()
}
