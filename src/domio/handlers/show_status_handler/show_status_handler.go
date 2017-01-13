package show_status_handler

import (
    "net/http"
    "domio/components/responses"
    "domio/components/config"
)

func ShowStatusHandler(w http.ResponseWriter, req *http.Request) {
    responses.ReturnObjectResponse(w, config.AppStatusInfo)
}