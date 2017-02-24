package update_domain_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "UpdateDomain",
        http.MethodPut,
        "/domain/{name}",
        UpdateDomainHandler,
        CheckAccessFunc,
    }
}