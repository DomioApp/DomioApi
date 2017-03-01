package get_user_subscriptions_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetUserSubscriptions",
        http.MethodGet,
        "/subscriptions",
        GetUserSubscriptionsHandler,
        CheckAccessFunc,
        nil,
    }
}