package datasources

import (
	"database/sql"
	"fmt"
	"gopi/config"
	"log"
	"strings"

	_ "github.com/marcboeker/go-duckdb"
	"github.com/qustavo/dotsql"
)

func IsInitialized(db *sql.DB) bool {
	row := db.QueryRow(`
        SELECT config
        FROM settings
        WHERE id = 'app'
    `)

	var configJSON string
	err := row.Scan(&configJSON)
	if err != nil {
		// add better checks if the error is err == sql.ErrNoRows or table doesn't exist VS unexpected error
		return false
	}

	return strings.Contains(configJSON, `"database_initialized": true`)
}

func InitDuckDB(mysql *config.Credentials, db *sql.DB) {
	dot, err := dotsql.LoadFromFile("datasources/sql/duckdb-init.sql")
	if err != nil {
		log.Fatal("Failed to open duckdb-init.sql")
	}

	_, err = dot.Exec(db, "duckdb-tables")
	if err != nil {
		fmt.Print(err)
		log.Fatal("Failedd to initialize duckdb")
	}

	query := fmt.Sprintf(`
        INSTALL mysql;
        LOAD mysql;
        CREATE SECRET (
            TYPE MYSQL,
            HOST '%s',
            PORT '%s',
            DATABASE '%s',
            USER '%s',
            PASSWORD '%s'
        );
        ATTACH '' AS mysql_db (TYPE MYSQL);
    `,
		mysql.DbHost,
		mysql.DbPort,
		mysql.DbName,
		mysql.DbUser,
		mysql.DbPassword,
	)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	_, err = dot.Exec(db, "mysql-table")
	if err != nil {
		fmt.Print(err)
		log.Fatal("Failed to create mysql table on attached instance")
	}

	_, err = dot.Exec(db, "finish-init")
	if err != nil {
		fmt.Print(err)
		log.Fatal("Failed to finish database initialization")
	}
}

func GetDuckDB() *sql.DB {
	db, err := sql.Open("duckdb", "datasources/db/gopi.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
