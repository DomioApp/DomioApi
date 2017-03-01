package get_user_domains_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func GetUserDomainsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {

    userEmail := userProfile.Email

    userDomains := domiodb.GetUserDomains(userEmail)

    responses.ReturnObjectResponse(w, userDomains)

    defer req.Body.Close()
}