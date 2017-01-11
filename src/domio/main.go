package main

import (
    "net/http"
    "log"
    "domio/router"
    "fmt"
    "domio/components/config"
    "domio/components/arguments"
)

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
    result := arguments.ProcessArguments()
    log.Print(result)

    //startRouter()
}

func startRouter() {
    domiorouter := router.NewRouter()
    log.Printf("Web server is running on http://localhost:%d", config.Config.PORT)
    err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config.PORT), domiorouter)
    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", config.Config.PORT)
        log.Fatal(msg)
    }

}