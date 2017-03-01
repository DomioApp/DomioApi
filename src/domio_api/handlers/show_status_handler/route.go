package show_status_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "ShowStatus",
        http.MethodGet,
        "/",
        ShowStatusHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}