package responses

import (
    "net/http"
    "domio_api/messages"
    "encoding/json"
)

func ReturnMessageResponse(w http.ResponseWriter, message messages.DomioMessage) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(message)
}