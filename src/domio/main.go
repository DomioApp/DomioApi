package main

import (
    "log"
    "fmt"
    "domio/components/arguments"
    "os"
    "github.com/fatih/color"
    "domio/components/config"
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
        os.Exit(1)
    }
}

func printHeader() {
    color.Set(color.FgHiCyan)
    fmt.Println()
    fmt.Println("------------------------------------------------------")
    fmt.Println("Buildstamp: ", Buildstamp)
    fmt.Println("Hash:       ", Hash)
    fmt.Println("Version:    ", Version)
    fmt.Println("------------------------------------------------------")
    fmt.Println()
    color.Unset()

}