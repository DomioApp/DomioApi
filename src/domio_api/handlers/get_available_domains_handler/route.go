package get_available_domains_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetAvailableDomains",
        http.MethodGet,
        "/domains/available",
        GetAvailableDomainsHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}