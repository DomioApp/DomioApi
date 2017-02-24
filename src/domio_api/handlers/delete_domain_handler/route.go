package delete_domain_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "DeleteDomain",
        http.MethodDelete,
        "/domain/{name}",
        DeleteDomainHandler,
        CheckAccessFunc,
    }
}