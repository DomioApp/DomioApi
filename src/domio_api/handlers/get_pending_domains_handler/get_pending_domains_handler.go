package get_pending_domains_handler

import (
    "net/http"
    "log"
    "domio_api/db"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func GetPendingDomainsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {
    pendingDomains := domiodb.GetPendingDomains()
    log.Print(pendingDomains)

    responses.ReturnObjectResponse(w, pendingDomains)
}