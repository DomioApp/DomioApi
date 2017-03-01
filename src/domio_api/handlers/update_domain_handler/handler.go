package update_domain_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
    "github.com/gorilla/mux"
)

func UpdateDomainHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {

    var domainToEdit domiodb.DomainToEdit
    var updatedDomain domiodb.DomainJson

    err := requests.DecodeJsonRequestBody(req, &domainToEdit)

    if err != nil {
        log.Print(err)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    domainUpdateError := domiodb.UpdateDomain(domainName, domainToEdit)

    if (domainUpdateError != nil) {
        log.Print(domainUpdateError)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    responses.ReturnObjectResponse(w, updatedDomain)

    defer req.Body.Close()
}
