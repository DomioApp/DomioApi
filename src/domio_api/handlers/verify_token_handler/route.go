package verify_token_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "VerifyToken",
        http.MethodPost,
        "/tokens/verify",
        VerifyTokenHandler,
        CheckAccessFunc,
        nil,
    }
}