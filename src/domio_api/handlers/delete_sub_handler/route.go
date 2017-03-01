package delete_subscription_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "DeleteSubscription",
        http.MethodDelete,
        "/subscription/{subId}",
        DeleteSubscriptionHandler,
        CheckAccessFunc,
        nil,
    }
}