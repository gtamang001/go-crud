package initializers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbServer := os.Getenv("DATABASE_SERVER")
	dbPort := os.Getenv("DATABASE_PORT")
	log.Println("Database host: " + dbServer)
	log.Println("Database port: " + dbPort)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbServer+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
