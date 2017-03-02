package get_user_domains_count_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/tokens"
    "domio_api/components/responses"
)

type DomainsCountResponse struct {
    Count int `json:"count"`
}

func GetUserDomainsCountHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    userDomainsCount := domiodb.GetUserDomainsCount(userProfile.Email)

    resp := DomainsCountResponse{Count:userDomainsCount}

    responses.ReturnObjectResponse(w, resp)
}