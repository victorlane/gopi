package datasources

import (
	"database/sql"
	"fmt"
	"gopi/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(creds *config.Credentials) *sql.DB {
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
