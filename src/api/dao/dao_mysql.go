package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func StartMySQL(user, pass string) {

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// See "Important settings" section.
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO users VALUES ( ?, ? )", user, pass)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}