package datasources

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

func InitDuckDB() {
	db := GetDuckDB()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS logs (
			client_ip TEXT,
			time_stamp TIMESTAMP,
			method TEXT,
			path TEXT,
			protocol TEXT,
			status_code INTEGER,
			latency TEXT,
			user_agent TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP
	)`)

	if err != nil {
		fmt.Printf("Error creating logs table: %v\n", err)
		return
	}
}

func GetDuckDB() *sql.DB {
	db, err := sql.Open("duckdb", "datasources/db/gopi.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
