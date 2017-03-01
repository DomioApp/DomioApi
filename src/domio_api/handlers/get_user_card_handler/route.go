package get_user_card_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetUserCard",
        http.MethodGet,
        "/cards/{id}",
        GetUserCardHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}