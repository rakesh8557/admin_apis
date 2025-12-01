package database

import (
	"admin_apis/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	// dsn := *config.DB_USER + ":" + *config.DB_PASS + "@tcp(" + *config.DB_HOST + ":" + *config.DB_PORT + ")/" + *config.DB_NAME + "?parseTime=true&charsetutf=8mb4" //username:password@tcp(host:port)/dbname?param=value

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		*config.DB_USER,
		*config.DB_PASS,
		*config.DB_HOST,
		*config.DB_PORT,
		*config.DB_NAME,
	)

	var err error

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// log.Fatal("Failed to open DB: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// log.Fatal("Failed to connect to DB: ", err)
		return nil, err
	}

	// log.Println("Database connected successfully!")
	return db, nil
}
