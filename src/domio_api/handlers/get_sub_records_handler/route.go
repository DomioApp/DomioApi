package get_sub_records_handler

import (
    "net/http"
    "domio_api/types"
)

func GetRoute() *types.Route {
    return &types.Route{
        "GetSubscriptionRecords",
        http.MethodGet,
        "/subscriptions/{id}/records",
        GetSubscriptionRecordsHandler,
        CheckAccessFunc,
        DataGetterFunc,
    }
}