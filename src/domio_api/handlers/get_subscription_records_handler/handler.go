package get_subscription_records_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/db"
    "domio_api/external_api/r53"
)


func GetSubscriptionRecordsHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))
    log.Print(req.Header.Get("Authorization"))

    if (verifyTokenError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    log.Print(userEmail)

    //subscription := stripe_subscription_adapter.GetUserSubscription(subscriptionId)
    domainInfo, err := domiodb.GetDomainInfoBySubscriptionId(subscriptionId)

    if (err != nil) {
        log.Print(err)
    }

    records := r53.GetHostedZoneRecords(&domainInfo)

    log.Print(records)

    responses.ReturnObjectResponse(w, records.ResourceRecordSets)

    defer req.Body.Close()
}