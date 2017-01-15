package get_available_domains_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/responses"
)

func GetAvailableDomainsHandler(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()
    availableDomains := domiodb.GetAvailableDomains()
    responses.ReturnObjectResponse(w, availableDomains)
}