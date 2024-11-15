package datasources

import (
	"database/sql"
	"fmt"
	"gopi/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB(creds *config.Credentials) *sql.DB {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", creds.DbUser, creds.DbPassword, creds.DbHost, creds.DbPort, creds.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDB(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id int NOT NULL AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(256) NOT NULL,
		role INT NOT NULL DEFAULT 0,
		PRIMARY KEY(id)
	);`)

	if err != nil {
		log.Fatal(err)
	}

}
