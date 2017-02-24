package get_user_subscriptions_handler

import (
    "net/http"
)

func CheckAccessFunc(w http.ResponseWriter) bool {
    return true
}