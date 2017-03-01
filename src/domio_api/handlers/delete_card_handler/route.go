package delete_card_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "DeleteUserCard",
        http.MethodDelete,
        "/cards/{id}",
        DeleteCardHandler,
        CheckAccessFunc,
        nil,
    }
}