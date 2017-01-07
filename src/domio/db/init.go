package domiodb

import (
    "github.com/jmoiron/sqlx"
    "log"
    "fmt"
    "domio/components/config"
)

var Db *sqlx.DB

func init() {
    appconfig := config.LoadConfig()
    var err error
    var username = appconfig.DOMIO_DB_USER
    var password = appconfig.DOMIO_DB_PASSWORD
    var dbname = appconfig.DOMIO_DB_NAME
    var aws_access_key_id = appconfig.AWS_ACCESS_KEY_ID
    var aws_secret_access_key = appconfig.AWS_SECRET_ACCESS_KEY

    if (username == "") {
        log.Fatalln("DOMIO_DB_USER not set")
    }

    if (password == "") {
        log.Fatalln("DOMIO_DB_PASSWORD not set")
    }

    if (dbname == "") {
        log.Fatalln("DOMIO_DB_NAME not set")
    }

    if (aws_access_key_id == "") {
        log.Fatalln("AWS_ACCESS_KEY_ID not set")
    }

    if (aws_secret_access_key == "") {
        log.Fatalln("AWS_SECRET_ACCESS_KEY not set")
    }

    var dbconfig = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

    Db, err = sqlx.Connect("postgres", dbconfig)
    if err != nil {
        log.Fatalln("Couldn't connect to Database:\n", err)
    }
    log.Println("Database Initialized.")
}
