package show_status_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}