package create_user_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "CreateUser",
        http.MethodPost,
        "/users",
        CreateUserHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}