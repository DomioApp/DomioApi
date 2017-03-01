package delete_record_handler

import (
    "net/http"
    "log"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/external_api/r53"
    "domio_api/db"
    "domio_api/components/requests"
    "domio_api/errors"
)

type Record struct {
    Key    string `json:"key"`
    Value  string `json:"value"`
    TTL    int64 `json:"ttl"`
    Weight int64 `json:"weight"`
}

func DeleteRecordHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    var record Record

    requestVars := mux.Vars(req)
    subId := requestVars["subId"]

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

    domain, domainErr := domiodb.GetDomainInfoBySubscriptionId(subId)

    if domainErr != nil {
        log.Print(domainErr)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    r53.DeleteRecord(domain.ZoneId.String, "www." + domain.Name, record.Key, record.Value)

    responses.ReturnObjectResponse(w, domain)
}