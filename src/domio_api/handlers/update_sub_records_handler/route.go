package update_sub_records_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "UpdateSubscriptionRecords",
        http.MethodPut,
        "/subscription/{id}/records",
        UpdateSubscriptionRecordsHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}