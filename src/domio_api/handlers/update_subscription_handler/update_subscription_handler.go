package update_subscription_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
    "github.com/gorilla/mux"
    r53 "domio_api/external_api/route53"
)

func UpdateSubscriptionHandler(w http.ResponseWriter, req *http.Request) {

    var domainToEdit domiodb.DomainToEdit
    var updatedDomain domiodb.DomainJson

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

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

    result, domainUpdateError := r53.UpdateCNAME(subId, domainToEdit)

    log.Print(result);

    if (domainUpdateError != nil) {
        log.Print(domainUpdateError)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    responses.ReturnObjectResponse(w, updatedDomain)

    defer req.Body.Close()
}
