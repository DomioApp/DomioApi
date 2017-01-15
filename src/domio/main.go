package main

import (
    "log"
    "fmt"
    "domio/components/arguments"
    "os"
    "github.com/fatih/color"
    "domio/components/config"
    "domio/components/logger"
)

var Version string
var Hash string
var Buildstamp string

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    config.AppStatusInfo = config.AppStatus{
        Buildstamp:Buildstamp,
        Hash:Hash,
        Version:Version,
    }
}

func main() {
    printHeader()

    argumentsList, argumentsError := arguments.ProcessArguments()

    log.Print(argumentsList.Command)

    if (argumentsError != nil) {
        fmt.Print(argumentsError)
        logger.Crit(argumentsError.Error())
        os.Exit(1)
    }
}

func printHeader() {

    color.Set(color.FgHiCyan)
    fmt.Println()
    fmt.Println("-----------------------------------------------------")
    fmt.Println("Version:   ", config.AppStatusInfo.Version)
    fmt.Println("Buildstamp:", config.AppStatusInfo.Buildstamp)
    fmt.Println("Build:     ", config.AppStatusInfo.GetBuildDateTime())
    fmt.Println("Hash:      ", config.AppStatusInfo.Hash)
    fmt.Println("-----------------------------------------------------")
    fmt.Println()
    color.Unset()
}