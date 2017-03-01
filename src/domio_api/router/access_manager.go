package router

import (
    "net/http"
    "domio_api/types"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/errors"
)

func ManageRoute(route *types.Route) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {

        userProfile, _ := tokens.VerifyTokenString(req.Header.Get("Authorization"))

        isAccessGranted := route.CheckAccessFunc(userProfile, req)

        if (isAccessGranted == false) {
            responses.ReturnErrorResponseWithCustomCode(w, domioerrors.AccessIsForbidden, http.StatusUnauthorized)
            return
        }

        data := route.DataGetterFunc(req)

        route.HandlerFunc(w, req, userProfile, isAccessGranted, data)
    }
}