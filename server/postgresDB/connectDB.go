package postgresDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rospi"
	password = "rpitele"
	dbname   = "rospi"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectDB() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	return db
}