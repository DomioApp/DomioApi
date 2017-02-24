package get_domain_info_handler

import (
    "net/http"
    "domio_api/types"
    "domio_api/handlers/get_user_domains_handler"
)

func GetRoute() *types.Route {
    return &types.Route{
        "CreateUser",
        http.MethodPost,
        "/users",
        GetDomainInfoHandler,
        CheckAccessFunc,
    }
}