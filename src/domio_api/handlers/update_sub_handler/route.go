package update_sub_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "UpdateSubscription",
        http.MethodPut,
        "/subscriptions/{id}",
        UpdateSubscriptionHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}