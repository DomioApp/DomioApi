package get_domain_info_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetDomainInfo",
        http.MethodGet,
        "/domain/{name}",
        GetDomainInfoHandler,
        CheckAccessFunc,
    }
}