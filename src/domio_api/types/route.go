package types

import (
    "net/http"
    "domio_api/components/tokens"
)

type Route struct {
    Name            string
    Method          string
    Pattern         string
    HandlerFunc     HandlerFuncWithParams
    CheckAccessFunc CheckAccessFunc
}

type CheckAccessFunc func(userProfile *tokens.UserTokenWithClaims, req *http.Request) bool

type HandlerFuncWithParams func(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessEnabled bool)
