package delete_domain_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "github.com/gorilla/mux"
    "domio/messages"
    "log"
)

func DeleteDomainHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    deleteError := domiodb.DeleteDomain(domainName, userProfile.Email)

    if (deleteError != domioerrors.DomioError{}) {
        log.Println(deleteError)
        responses.ReturnErrorResponse(w, deleteError)
        return
    }
    responses.ReturnObjectResponse(w, messages.DomainDeleted)

    defer req.Body.Close()
}
