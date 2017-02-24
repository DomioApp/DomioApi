package create_domain_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}