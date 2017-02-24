package create_card_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}