

// Go connection Sample Code:
package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	"log"
)

var db *sql.DB
var server = "sqlserver-sample777.database.windows.net"
var port = 1433
var user = "user"
var password = "Passwor"
var database = "azuresqldatabase-sample"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}