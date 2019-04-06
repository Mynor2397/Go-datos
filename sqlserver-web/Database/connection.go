package Database

import (
	"database/sql"
	"fmt"

	_"github.com/denisenkom/go-mssqldb"
)


func Dbconn() (db *sql.DB) {

	server := `MAICK\SQLEXPRESS`
	port := 1433
	user := "sa"
	password := "Miprincesa1009"
	database := "employee"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

