package get_domain_info_handler

import (
    "net/http"
)

func CheckAccessFunc(w http.ResponseWriter) bool {
    return true
}