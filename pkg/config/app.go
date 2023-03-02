package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Basically, this is wherre the connection to the database is made.
//dbDriver is the driver type; I'm using mysql
//dbUser is the username of your mysql local server
//dbPass is the password for your mysql local server

var DB_NAME = "" //This was created so that you won't need to manually change the database name in your SQL queries

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := ""
	dbPass := "!"
	dbName := ""
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
