package router

import (
    "net/http"
    "log"
    "domio_api/types"
)


func ManageAccess(handlerFunc http.HandlerFunc, checkAccessFunc types.CheckAccessFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        log.Print(checkAccessFunc(w))
        handlerFunc(w, req)
    }
}

