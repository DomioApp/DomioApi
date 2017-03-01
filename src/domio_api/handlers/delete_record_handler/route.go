package delete_record_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "DeleteRecord",
        http.MethodDelete,
        "/subscription/{subId}/records",
        DeleteRecordHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}