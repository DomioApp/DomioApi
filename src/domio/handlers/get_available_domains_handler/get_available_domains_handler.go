package get_available_domains_handler

import (
    "net/http"
    "domio/db"
    "domio/components/responses"
)

func GetAvailableDomainsHandler(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()
    availableDomains := domiodb.GetAvailableDomains()
    responses.ReturnObjectResponse(w, availableDomains)
}