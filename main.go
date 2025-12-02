package main

import (
	"admin_apis/config"
	"admin_apis/controllers"
	"admin_apis/database"
	"admin_apis/logger"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/test", controllers.Test).Methods("GET")

	err = http.ListenAndServe(fmt.Sprintf("%v:%v", *config.SERVER_HOST, *config.SERVER_PORT), r)
	if err != nil {
		logger.Log("ERROR", fmt.Sprintf("Error in starting the server: %v", err))
		return
	}

	logger.Log("INFO", fmt.Sprintf("Server starting at port: %v", *config.SERVER_PORT))
}
