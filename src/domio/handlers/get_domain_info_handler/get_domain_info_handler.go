package get_domain_info_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "github.com/gorilla/mux"
    "domio/components/tokens"
    "domio/components/responses"
)

func GetDomainInfoHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    defer req.Body.Close()

    _, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }
    domainInfo, err := domiodb.GetDomain(domainName)
    if (err != nil) {
        responses.ReturnErrorResponse(w, err)
        return
    }
    responses.ReturnObjectResponse(w, domainInfo)
}