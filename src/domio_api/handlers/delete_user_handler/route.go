package delete_user_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "DeleteUser",
        http.MethodDelete,
        "/user",
        DeleteUserHandler,
        CheckAccessFunc,
        nil,
    }
}