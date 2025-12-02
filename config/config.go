package config

import (
	"database/sql"
	"flag"
)

var (
	DB *sql.DB // Global, Exported DB connection

	DB_USER       = flag.String("DB_USER", "root", "Database Username")
	DB_PASS       = flag.String("DB_PASS", "mandeep123", "Database Password")
	DB_HOST       = flag.String("DB_HOST", "127.0.0.1", "Database Host")
	DB_PORT       = flag.String("DB_PORT", "3306", "Database Port")
	DB_NAME       = flag.String("DB_NAME", "test_db", "Database Name")
	ENV           = flag.String("ENV", "dev", "Environment (dev|prod)")
	LOG_FILE_PATH = flag.String("LOG_FILE_PATH", "D:/go_admin_apis_logs/app.log", "Path of the log file")
	SERVER_HOST   = flag.String("SERVER_HOST", "localhost", "Server Host")
	SERVER_PORT   = flag.String("SERVER_PORT", "8080", "Server Port")
)

func InitiateConfig() {
	flag.Parse()
}
