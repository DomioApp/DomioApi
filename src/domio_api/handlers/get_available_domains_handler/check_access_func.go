package get_available_domains_handler

import (
    "net/http"
)

func CheckAccessFunc(w http.ResponseWriter) bool {
    return true
}