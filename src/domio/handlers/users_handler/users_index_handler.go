package handlers

import (
    "net/http"
    "domio/db"
    "encoding/json"
)

func UsersIndex(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()
    users := domiodb.GetUsers()

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
}
