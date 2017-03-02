package router

import (
    "net/http"
    "domio_api/types"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/errors"
    "log"
)

func ManageRoute(route *types.Route) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        var userProfile *tokens.UserTokenWithClaims
        var verifyTokenError *domioerrors.DomioError

        authHeader := req.Header.Get("Authorization")

        log.Print("-----------------------------------------------------------")
        log.Print(req.URL)
        log.Print(authHeader)
        log.Print("-----------------------------------------------------------")

        if (authHeader != "") {
            userProfile, verifyTokenError = tokens.VerifyTokenString(authHeader)

            if (verifyTokenError != nil) {
                log.Print(verifyTokenError)
            }

        } else {
            userProfile = nil
        }

        isAccessGranted := route.CheckAccessFunc(userProfile, req)

        if (isAccessGranted == false) {
            responses.ReturnErrorResponseWithCustomCode(w, domioerrors.AccessIsForbidden, http.StatusUnauthorized)
            return
        }

        data := route.DataGetterFunc(req)

        route.HandlerFunc(w, req, userProfile, isAccessGranted, data)
    }
}