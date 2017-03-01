package get_user_domains_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func GetUserDomainsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    userDomains := domiodb.GetUserDomains(userProfile.Email)

    responses.ReturnObjectResponse(w, userDomains)
}