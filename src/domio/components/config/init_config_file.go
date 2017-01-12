package config

import (
    "fmt"
    "os"
    "log"
    "encoding/json"
    "io/ioutil"
    "path"
)

func InitConfigFile(filenameFlag *string, awsAccessKeyIdFlag *string, awsSecretAccessKeyFlag *string, dbNameFlag *string, dbUserFlag *string, dbPasswordFlag *string, webPortFlag *uint, envFlag *string) error {

    argsErr := false

    if *filenameFlag == "" {
        fmt.Println("Please supply the filename --file option.")
        argsErr = true
    }

    if *awsAccessKeyIdFlag == "" {
        fmt.Println("Please supply the --aws-access-key-id option.")
        argsErr = true
    }

    if *awsSecretAccessKeyFlag == "" {
        fmt.Println("Please supply the --aws-secret-access-key option.")
        argsErr = true
    }

    if *dbNameFlag == "" {
        fmt.Println("Please supply the DB name --db-name option.")
        argsErr = true
    }

    if *dbUserFlag == "" {
        fmt.Println("Please supply the DB user name --db-user option.")
        argsErr = true
    }

    if *dbPasswordFlag == "" {
        fmt.Println("Please supply the DB password --db-password option.")
        argsErr = true
    }

    if argsErr {
        fmt.Println("-----------------------------------------")
        fmt.Println("Please provide required options.")
        os.Exit(1)
    }
    //dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    conf := Configuration{
        AWS_ACCESS_KEY_ID: *awsAccessKeyIdFlag,
        AWS_SECRET_ACCESS_KEY:*awsSecretAccessKeyFlag,
        DOMIO_DB_NAME: *dbNameFlag,
        DOMIO_DB_USER: *dbUserFlag,
        DOMIO_DB_PASSWORD: *dbPasswordFlag,
        PORT: *webPortFlag,
        ENV: *envFlag,
    }

    if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
        log.Print("Creating config folder...")
        os.MkdirAll(ConfigPath, 0660)
    }

    jsonConfig, _ := json.MarshalIndent(conf, "", "    ")
    err := ioutil.WriteFile(path.Join(ConfigPath, *filenameFlag), jsonConfig, 0660)
    if (err != nil) {
        log.Println(err)
        os.Exit(1)
    }
    return nil

}
