package arguments

import (
    "flag"
    "os"
    "fmt"
    "domio/components/config"
    "log"
    "encoding/json"
    "io/ioutil"
    "path"
)

func ProcessArguments() error {
    initCommand := flag.NewFlagSet("init", flag.ExitOnError)
    filenameFlag := initCommand.String("file", "config.json", "config file absolute path")
    awsAccessKeyIdFlag := initCommand.String("aws-access-key-id", "", "AWS Access Key ID")
    awsSecretAccessKeyFlag := initCommand.String("aws-secret-access-key", "", "AWS Secret Access Key")
    dbNameFlag := initCommand.String("db-name", "", "DB name")
    dbUserFlag := initCommand.String("db-user", "", "DB user name")
    dbPasswordFlag := initCommand.String("db-password", "", "DB password")
    webPortFlag := initCommand.Uint("port", 8080, "Port for the HTTP server to run on")
    envFlag := initCommand.String("env", "development", "Environment name: development, testing, production")

    sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
    //recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
    //messageFlag := sendCommand.String("message", "", "Text message")

    if len(os.Args) == 1 {
        fmt.Println("usage: siri <command> [<args>]")
        fmt.Println("The most commonly used git commands are: ")
        fmt.Println(" ask   Ask questions")
        fmt.Println(" send  Send messages to your contacts")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "init":
        initCommand.Parse(os.Args[2:])

        if initCommand.Parsed() {
            initConfig(filenameFlag, awsAccessKeyIdFlag, awsSecretAccessKeyFlag, dbNameFlag, dbUserFlag, dbPasswordFlag, webPortFlag, envFlag)
        }

        return nil
    case "send":
        sendCommand.Parse(os.Args[2:])
    default:
        fmt.Printf("%q is not valid command.\n", os.Args[1])
        os.Exit(2)
    }

    return nil
}

func initConfig(filenameFlag *string, awsAccessKeyIdFlag *string, awsSecretAccessKeyFlag *string, dbNameFlag *string, dbUserFlag *string, dbPasswordFlag *string, webPortFlag *uint, envFlag *string) error {

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

    conf := config.Configuration{
        AWS_ACCESS_KEY_ID: *awsAccessKeyIdFlag,
        AWS_SECRET_ACCESS_KEY:*awsSecretAccessKeyFlag,
        DOMIO_DB_NAME: *dbNameFlag,
        DOMIO_DB_USER: *dbUserFlag,
        DOMIO_DB_PASSWORD: *dbPasswordFlag,
        PORT: *webPortFlag,
        ENV: *envFlag,
    }

    if _, err := os.Stat(config.ConfigPath); os.IsNotExist(err) {
        log.Print("Creating config folder...")
        os.MkdirAll(config.ConfigPath, 0755)
    }

    jsonConfig, _ := json.MarshalIndent(conf, "", "    ")
    err := ioutil.WriteFile(path.Join(config.ConfigPath, *filenameFlag), jsonConfig, 0755)
    if (err != nil) {
        log.Println(err)
        os.Exit(1)
    }
    return nil

}