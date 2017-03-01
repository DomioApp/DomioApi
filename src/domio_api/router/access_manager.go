package router

import (
    "net/http"
    "domio_api/types"
    "domio_api/components/tokens"
    "log"
)

func ManageAccess(handlerFunc types.HandlerFuncWithParams, checkAccessFunc types.CheckAccessFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {

        userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

        log.Print(userProfile)
        log.Print(verifyTokenError)

        result := checkAccessFunc(req)
        log.Print(result)

        handlerFunc(w, req, userProfile)
}
}