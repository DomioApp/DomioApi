package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio/routes"
    "path/filepath"
    "os"
    "domio/components/config"
    "log"
)

func NewRouter() *mux.Router {
    conf := config.Config
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    router := mux.NewRouter().StrictSlash(true)

    if (conf.ENV == "development") {
        log.Println("Running development environment.")
        router.PathPrefix("/swagger").Handler(http.FileServer(http.Dir(dir)))
    }

    for _, route := range routes.RoutesList {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
        Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }

    return router
}