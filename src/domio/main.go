package main

import (
    "net/http"
    "log"
    "domio/router"
    "fmt"
    "domio/components/config"
)

var Config = config.Configuration{}

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    Config = config.LoadConfig()
}

func main() {
    port := 8080
    domiorouter := router.NewRouter()
    log.Printf("Web server is running on http://localhost:%d", port)
    err := http.ListenAndServe(":8080", domiorouter)
    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", port)
        log.Fatal(msg)
    }
}