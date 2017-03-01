package update_subscription_records_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/db"
    "domio_api/external_api/r53"
    "domio_api/components/requests"
)

type Record struct {
    Key    string `json:"key"`
    Value  string `json:"value"`
    TTL    int64 `json:"ttl"`
    Weight int64 `json:"weight"`
}

func UpdateSubscriptionRecordsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    var record Record

    requestVars := mux.Vars(req)
    subId := requestVars["id"]

    decodeErr := requests.DecodeJsonRequestBody(req, &record)

    log.Print("============================================")
    log.Print(record)
    log.Print(record.Key)
    log.Print("============================================")

    if decodeErr != nil {
        log.Print(decodeErr)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    userEmail := userProfile.Email
    log.Print(userEmail)

    domain, err := domiodb.GetDomainInfoBySubscriptionId(subId)

    r53.UpdateRecord(domain.ZoneId.String, "www." + domain.Name, record.Key, record.Value, record.TTL, record.Weight)

    if (err != nil) {
        log.Print(err)

    }

    log.Print("===========================================================")
    log.Print(domain)
    log.Print("===========================================================")
    responses.ReturnObjectResponse(w, domain)
}