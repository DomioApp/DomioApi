package get_user_domains_count_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetUserDomainsCount",
        http.MethodGet,
        "/user/domains/count",
        GetUserDomainsCountHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}