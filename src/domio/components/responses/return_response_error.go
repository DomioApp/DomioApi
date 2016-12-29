package responses

import (
    "net/http"
    "encoding/json"
    "log"
)

func ReturnErrorResponse(w http.ResponseWriter, messageError interface{}) {
    log.Print(messageError)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(messageError)
}
func ReturnErrorResponseWithCustomCode(w http.ResponseWriter, messageError interface{}, httpErrorCode int) {
    log.Print(messageError)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpErrorCode)
    json.NewEncoder(w).Encode(messageError)
}