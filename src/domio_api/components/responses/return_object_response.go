package responses

import (
    "net/http"
    "encoding/json"
)

func ReturnObjectResponse(w http.ResponseWriter, obj interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(obj)
}