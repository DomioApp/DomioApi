package update_sub_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
    "github.com/gorilla/mux"
    "domio_api/external_api/r53"
)

func UpdateSubscriptionHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    var domainToEdit domiodb.DomainToEdit
    var updatedDomain domiodb.DomainJson

    err := requests.DecodeJsonRequestBody(req, &domainToEdit)

    if err != nil {
        log.Print(err)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    requestVars := mux.Vars(req)
    subId := requestVars["id"]

    log.Print(userProfile)
    log.Print(subId)

    zoneId := "/hostedzone/Z2PJOVV67RBWSS"
    domainName := "www.john.com"
    key := "a"
    value := "heyho"
    var TTL int64 = 3600
    var weight int64 = 100

    result, domainUpdateError := r53.UpdateRecord(zoneId, domainName, key, value, TTL, weight)

    log.Print(result);

    if (domainUpdateError != nil) {
        log.Print(domainUpdateError)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    responses.ReturnObjectResponse(w, updatedDomain)
}
