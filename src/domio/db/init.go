package domiodb

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"fmt"
)

var Db *sqlx.DB

func init() {
	var err error
	var username = os.Getenv("DOMIO_DB_USER")
	var password = os.Getenv("DOMIO_DB_PASSWORD")
	var dbname = os.Getenv("DOMIO_DB_NAME")

	if (username == "") {
		log.Fatalln("DOMIO_DB_USER not set")
	}

	if (password == "") {
		log.Fatalln("DOMIO_DB_PASSWORD not set")
	}

	if (dbname == "") {
		log.Fatalln("DOMIO_DB_NAME not set")
	}

	var dbconfig = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

	Db, err = sqlx.Connect("postgres", dbconfig)
	if err != nil {
		log.Fatalln("Couldn't connect to Database:\n", err)
	}
	log.Println("Database Initialized.")
}
