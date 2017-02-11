package get_subscription_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/db"
    "github.com/gorilla/mux"
    "log"
)

func GetSubscriptionHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))
    log.Print(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    log.Print(userEmail)

    subscription := domiodb.GetUserSubscription(subscriptionId)

    log.Print("===========================================================")
    log.Print(subscription)
    log.Print("===========================================================")
    responses.ReturnObjectResponse(w, subscription)

    defer req.Body.Close()
}