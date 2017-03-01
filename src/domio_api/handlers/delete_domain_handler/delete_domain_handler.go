package delete_domain_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/messages"
    "domio_api/external_api/r53"
    "domio_api/utils"
)

func DeleteDomainHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    domain, deleteError := domiodb.DeleteDomain(domainName, userProfile.Email)

    if (deleteError != nil) {

        utils.ShowError(deleteError)

        responses.ReturnErrorResponse(w, deleteError)
        return
    }

    r53.DeleteDomainZone(domain)
    responses.ReturnObjectResponse(w, messages.DomainDeleted)
}
