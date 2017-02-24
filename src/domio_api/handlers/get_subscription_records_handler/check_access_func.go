package get_subscription_records_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}