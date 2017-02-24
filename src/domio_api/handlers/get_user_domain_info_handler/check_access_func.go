package get_domain_info_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}