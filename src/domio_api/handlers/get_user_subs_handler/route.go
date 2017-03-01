package get_user_subs_handler

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
        DataGetterFunc,
    }
}