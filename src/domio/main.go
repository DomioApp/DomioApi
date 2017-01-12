package main

import (
    "log"
    "fmt"
    "domio/components/arguments"
    "os"
    "github.com/fatih/color"
)

var Version string
var Buildstamp string

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
    printHeader()
    result := arguments.ProcessArguments()
    if (result != nil) {
        fmt.Print(result)
        os.Exit(1)
    }

    //server.StartRouter()
}

func printHeader() {
    color.Set(color.FgHiCyan)
    fmt.Println()
    fmt.Println("------------------------------------------------------")
    fmt.Println("Buildstamp: ", Buildstamp)
    fmt.Println("Hash:       ", Version)
    fmt.Println("------------------------------------------------------")
    fmt.Println()
    color.Unset()

}