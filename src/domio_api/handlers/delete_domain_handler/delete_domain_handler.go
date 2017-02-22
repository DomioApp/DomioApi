package delete_domain_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/messages"
    "log"
    "domio_api/external_api/r53"
    "github.com/fatih/color"
)

func DeleteDomainHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    domain, deleteError := domiodb.DeleteDomain(domainName, userProfile.Email)

    r53.DeleteDomainZone(domain)
    if (deleteError != nil) {
        color.Set(color.FgRed)
        log.Println(deleteError)
        color.Unset()
        responses.ReturnErrorResponse(w, deleteError)
        return
    }
    responses.ReturnObjectResponse(w, messages.DomainDeleted)

    defer req.Body.Close()
}
