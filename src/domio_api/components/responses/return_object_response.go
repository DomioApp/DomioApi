package responses

import (
    "net/http"
    "encoding/json"
    "log"
)

func ReturnObjectResponse(w http.ResponseWriter, obj interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    log.Print(obj)
    json.NewEncoder(w).Encode(obj)
}