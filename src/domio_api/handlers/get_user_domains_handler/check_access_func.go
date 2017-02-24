package get_user_domains_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}