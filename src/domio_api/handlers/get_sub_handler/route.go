package get_sub_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetSubscription",
        http.MethodGet,
        "/subscriptions/{id}",
        GetSubscriptionHandler,
        CheckAccessFunc,
        nil,
    }
}