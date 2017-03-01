package create_subscription_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "CreateSubscription",
        http.MethodPost,
        "/subscriptions",
        CreateSubscriptionHandler,
        CheckAccessFunc,
        nil,
    }
}