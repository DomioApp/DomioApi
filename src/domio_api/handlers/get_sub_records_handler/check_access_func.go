package get_sub_records_handler

import (
    "net/http"
    "domio_api/components/tokens"
)

func CheckAccessFunc(userProfile *tokens.UserTokenWithClaims, req *http.Request) bool {
    return true
}