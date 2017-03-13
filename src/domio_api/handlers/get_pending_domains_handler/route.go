package get_pending_domains_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetPendingDomains",
        http.MethodGet,
        "/domains/pending",
        GetPendingDomainsHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}