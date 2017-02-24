package delete_user_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}