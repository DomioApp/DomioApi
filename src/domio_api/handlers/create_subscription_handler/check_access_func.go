package create_subscription_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}