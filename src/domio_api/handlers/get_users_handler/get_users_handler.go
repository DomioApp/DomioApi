package get_users_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func GetUsersHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {
    users := domiodb.GetUsers()
    responses.ReturnObjectResponse(w, users)
}