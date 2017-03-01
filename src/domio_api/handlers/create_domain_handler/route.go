package create_domain_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "CreateDomain",
        http.MethodPost,
        "/domains",
        CreateDomainHandler,
        CheckAccessFunc,
        nil,
    }
}