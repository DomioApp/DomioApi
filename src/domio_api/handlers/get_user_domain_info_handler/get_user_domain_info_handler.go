package get_domain_info_handler

import (
    "net/http"
    "domio_api/db"
    "github.com/gorilla/mux"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func GetDomainInfoHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    domainInfo, err := domiodb.GetDomainInfo(domainName)
    if (err != nil) {
        responses.ReturnErrorResponse(w, err)
        return
    }

    //TODO figure out when info from the bottom needed and refactor accordingly
    //domiodb.GetHostedZone(&domainInfo)

    responses.ReturnObjectResponse(w, domainInfo)

}