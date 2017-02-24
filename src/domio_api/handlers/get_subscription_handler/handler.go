package get_subscription_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/external_api/stripe/subscription"
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

    subscription := stripe_subscription_adapter.GetUserSubscription(subscriptionId)

    //log.Print("===========================================================")
    //log.Print(subscription)
    //log.Print("===========================================================")

    responses.ReturnObjectResponse(w, subscription)

    defer req.Body.Close()
}