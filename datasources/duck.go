package datasources

import (
	"database/sql"
	"fmt"
	"gopi/config"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

func InitDuckDB(mysql *config.Credentials, db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS logs (
		client_ip TEXT,
		timestamp TIMESTAMP,
		method TEXT,
		path TEXT,
		protocol TEXT,
		status_code INTEGER,
		latency TEXT,
		user_agent TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	);`)

	if err != nil {
		fmt.Printf("Error creating logs table: %v\n", err)
		return
	}

	query := fmt.Sprintf(`INSTALL mysql; LOAD mysql; CREATE SECRET (
		TYPE MYSQL,
		HOST '%s',
		PORT %s,
		DATABASE %s,
		USER '%s',
		PASSWORD '%s'
	); ATTACH '' AS mysql_db (TYPE MYSQL);`, mysql.DbHost, mysql.DbPort, mysql.DbName, mysql.DbUser, mysql.DbPassword)

	_, err = db.Exec(query)
	if err != nil {
		fmt.Printf("Error creating logs table: %v\n", err)
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS mysql_db.logs (
			client_ip TEXT,
			timestamp TIMESTAMP,
			method TEXT,
			path TEXT,
			protocol TEXT,
			status_code INTEGER,
			latency TEXT,
			user_agent TEXT,
			created_at TIMESTAMP,
	);`)

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
