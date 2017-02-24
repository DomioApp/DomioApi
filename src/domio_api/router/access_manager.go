package router

import (
    "net/http"
    "log"
    "domio_api/routes"
)


func ManageAccess(handlerFunc http.HandlerFunc, checkAccessFunc routes.CheckAccessFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        log.Print(checkAccessFunc(w))
        handlerFunc(w, req)
    }
}

