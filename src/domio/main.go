package main

import (
    "net/http"
    "log"
    "domio/router"
)

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
    domiorouter := router.NewRouter()
    http.ListenAndServe(":8080", domiorouter)
}