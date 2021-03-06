package domiodb

import (
    "github.com/jmoiron/sqlx"
    "log"
    "fmt"
    "domio_api/components/config"
    "domio_api/components/logger"
)

var Db *sqlx.DB

func InitDb() {
    appconfig := config.Config
    var err error
    var dbusername = appconfig.DOMIO_DB_USER
    var dbpassword = appconfig.DOMIO_DB_PASSWORD
    var dbhost = appconfig.DOMIO_DB_HOST
    var dbname = appconfig.DOMIO_DB_NAME
    var aws_access_key_id = appconfig.AWS_ACCESS_KEY_ID
    var aws_secret_access_key = appconfig.AWS_SECRET_ACCESS_KEY

    if (dbusername == "") {
        log.Fatalln("DOMIO_DB_USER not set")
    }

    if (dbpassword == "") {
        log.Fatalln("DOMIO_DB_PASSWORD not set")
    }

    if (dbhost == "") {
        log.Fatalln("DOMIO_DB_HOST not set")
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

    var dbconfig = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbusername, dbpassword, dbname)

    Db, err = sqlx.Connect("postgres", dbconfig)
    if err != nil {
        log.Fatalln("Couldn't connect to Database:\n", err)
    }

    logger.Logger.Info("Database Initialized")
}
