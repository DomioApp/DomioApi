package arguments

import (
    "flag"
    "os"
    "fmt"
    "domio/components/config"
    "domio/components/server"
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

    startCommand := flag.NewFlagSet("start", flag.ExitOnError)
    //recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
    //messageFlag := sendCommand.String("message", "", "Text message")

    if len(os.Args) == 1 {
        fmt.Println("usage: domio <command> [<args>]")
        fmt.Println("Commands are: ")
        fmt.Println(" init   Init with new config file")
        fmt.Println(" start  Start server")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "init":
        initCommand.Parse(os.Args[2:])

        if initCommand.Parsed() {
            config.InitConfigFile(filenameFlag, awsAccessKeyIdFlag, awsSecretAccessKeyFlag, dbNameFlag, dbUserFlag, dbPasswordFlag, webPortFlag, envFlag)
        }

        return nil
    case "start":
        startCommand.Parse(os.Args[2:])

        config.LoadConfig()

        if startCommand.Parsed() {
            server.Start()
        }

        return nil

    default:
        fmt.Printf("%q is not valid command.\n", os.Args[1])
        os.Exit(2)
    }

    return nil
}
