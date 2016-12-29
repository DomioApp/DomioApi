package router

import (
    "net/http"
    "github.com/gorilla/mux"
    "domio/routes"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.PathPrefix("/swagger").Handler(http.FileServer(http.Dir("./")))

    for _, route := range routes.RoutesList {
        var handler http.Handler
        handler = route.HandlerFunc
        //handler = Logger.Logger(handler, route.Name)

        router.
        Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    //log.Println("Listening on 8080")
    return router
}