package server

import (
    "domio/router"
    "log"
    "domio/components/config"
    "net/http"
    "fmt"
)

func StartRouter() {
    log.Print("Starting router...")
    log.Print("-------------------------------------------")
    log.Print(config.Config)
    domiorouter := router.NewRouter()
    log.Printf("Web server is running on http://localhost:%d", config.Config.PORT)
    err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config.PORT), domiorouter)
    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", config.Config.PORT)
        log.Fatal(msg)
    }

}
func Start() {
    fmt.Print("Starting app...")
    StartRouter()
}
