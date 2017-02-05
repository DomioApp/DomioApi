package server

import (
    "domio_api/router"
    "log"
    "domio_api/components/config"
    "github.com/gorilla/handlers"
    "net/http"
    "fmt"
)

func StartRouter() {
    log.Print("Starting router...")
    domiorouter := router.NewRouter()
    log.Printf("Web server is running on http://localhost:%d", config.Config.PORT)
    var rt http.Handler

    if (config.Config.ENV == "development") {
        log.Println("CORS is managed by Gorilla...")

        corsObj := handlers.AllowedOrigins([]string{"*"})
        corsObj2 := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Credentials"})
        corsObj3 := handlers.AllowCredentials()

        rt = handlers.CORS(corsObj, corsObj2, corsObj3)(domiorouter)
    } else {
        rt = domiorouter
    }

    err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config.PORT), rt)

    if (err != nil) {
        msg := fmt.Sprintf("Failed to start web server on port %d", config.Config.PORT)
        log.Fatal(msg)
    }

}
func Start() {
    fmt.Print("Starting app...")
    StartRouter()
}
