package get_user_cards_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetUserCards",
        http.MethodGet,
        "/cards",
        GetUserCardsHandler,
        CheckAccessFunc,
    }
}