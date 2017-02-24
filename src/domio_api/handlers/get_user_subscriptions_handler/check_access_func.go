package get_user_subscriptions_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}