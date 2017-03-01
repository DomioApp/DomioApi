package create_card_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "CreateCard",
        http.MethodPost,
        "/cards",
        CreateCardHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}