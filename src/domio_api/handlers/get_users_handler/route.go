package get_users_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetUsers",
        http.MethodGet,
        "/users",
        GetUsersHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}