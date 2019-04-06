package Database
import (
	"database/sql"
	"fmt"

	_"github.com/denisenkom/go-mssqldb"
)


func Dbconn() (db *sql.DB) {

	server := `yOUR SERVER / LOCALHOST`
	port := 1433
	user := "sa" //USER SYSTEM ADMINISTRATOR
	password := "YOUR PASSWORD"
	database := "YOUR DATABASE"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

