package delete_domain_handler

import (
    "net/http"
)

func CheckAccessFunc(w http.ResponseWriter) bool {
    return true
}