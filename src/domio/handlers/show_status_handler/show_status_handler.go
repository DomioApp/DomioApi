package show_status_handler

import (
    "net/http"
    "domio/components/responses"
)

type AppStatus struct {
    AppVersion string `json:"app_version"`
}

func ShowStatusHandler(w http.ResponseWriter, req *http.Request) {
    appStatus := AppStatus{AppVersion:"1"}
    responses.ReturnObjectResponse(w, appStatus)
}