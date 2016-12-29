package requests

import (
    "net/http"
    "encoding/json"
)

func DecodeJsonRequestBody(req *http.Request, obj interface{}) (error) {

    decoder := json.NewDecoder(req.Body)

    decodeError := decoder.Decode(&obj)

    return decodeError
}