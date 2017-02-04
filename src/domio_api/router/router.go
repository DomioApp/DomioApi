package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio_api/routes"
    //"path/filepath"
    //"os"
    "domio_api/components/config"
    "log"
)

func NewRouter() *mux.Router {
    conf := config.Config

    router := mux.NewRouter().StrictSlash(true)

    if (conf.ENV == "development") {
        log.Print("Development environment, handling static files by Go...")
        router.PathPrefix("/swagger").Handler(http.FileServer(http.Dir("/usr/local/domio_api/www")))
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