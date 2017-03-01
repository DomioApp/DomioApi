package router

import (
    "net/http"
    "domio_api/types"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    domioerrors "domio_api/errors"
)

func ManageAccess(handlerFunc types.HandlerFuncWithParams, checkAccessFunc types.CheckAccessFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {

        userProfile, _ := tokens.VerifyTokenString(req.Header.Get("Authorization"))

        isAccessGranted := checkAccessFunc(userProfile, req)

        if (isAccessGranted == false) {
            responses.ReturnErrorResponseWithCustomCode(w, domioerrors.AccessIsForbidden, http.StatusUnauthorized)
            return
        }

        handlerFunc(w, req, userProfile, isAccessGranted)
    }
}