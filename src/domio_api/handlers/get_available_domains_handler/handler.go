package get_available_domains_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func GetAvailableDomainsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {
    availableDomains := domiodb.GetAvailableDomains()
    responses.ReturnObjectResponse(w, availableDomains)

    defer req.Body.Close()
}