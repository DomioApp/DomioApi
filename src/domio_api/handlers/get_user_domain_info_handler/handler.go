package get_domain_info_handler

import (
    "net/http"
    "domio_api/db"
    "github.com/gorilla/mux"
    "domio_api/components/responses"
)

func GetDomainInfoHandler(w http.ResponseWriter, req *http.Request, data *interface{}) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    defer req.Body.Close()

    /*
    _, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }
    */

    domainInfo, err := domiodb.GetDomainInfo(domainName)
    if (err != nil) {
        responses.ReturnErrorResponse(w, err)
        return
    }
    //TODO figure out when info from the bottom needed and refactor accordingly
    //domiodb.GetHostedZone(&domainInfo)

    responses.ReturnObjectResponse(w, domainInfo)
}