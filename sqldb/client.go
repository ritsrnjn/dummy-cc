package sqldb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// define a global variable to hold the database connection
var db *sql.DB

// make a connection with the dummy-cc-db database
func ConnectWithDb() error {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "codemysql"
	dbName := "ecommerce"

	localDb, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return err
	}
	db = localDb
	return db.Ping()
}

func Execute(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}
