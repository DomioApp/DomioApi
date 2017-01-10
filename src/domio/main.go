package main

import (
    "net/http"
    "log"
    "domio/router"
    "fmt"
    "domio/components/config"
    "domio/components/logger"
)

var Config = config.Configuration{}

func init() {
    log.Print("Main init...")
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    Config = config.LoadConfig()
}

func main() {
    logger.Logger.Info("Domio started")
    domiorouter := router.NewRouter()
    log.Printf("Web server is running on http://localhost:%d", Config.PORT)
    err := http.ListenAndServe(fmt.Sprintf(":%v", Config.PORT), domiorouter)
    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", Config.PORT)
        log.Fatal(msg)
    }
}