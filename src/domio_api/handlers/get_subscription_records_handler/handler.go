package get_subscription_records_handler

import (
    "net/http"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/db"
    "domio_api/external_api/r53"
)


func GetSubscriptionRecordsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["id"]

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