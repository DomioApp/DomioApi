package get_subscription_records_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/db"
)

func GetSubscriptionRecordsHandler(w http.ResponseWriter, req *http.Request) {

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

    //subscription := stripe_subscription_adapter.GetUserSubscription(subscriptionId)
    domain, err := domiodb.GetDomainInfoBySubscriptionId(subscriptionId)

    if (err != nil) {
        log.Print(err)

    }

    log.Print("===========================================================")
    log.Print(domain)
    log.Print("===========================================================")
    responses.ReturnObjectResponse(w, domain)

    defer req.Body.Close()
}