package get_user_domains_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetUserDomains",
        http.MethodGet,
        "/user/domains",
        GetUserDomainsHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}