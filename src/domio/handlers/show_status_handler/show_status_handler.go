package show_status_handler

import (
    "net/http"
    "domio/components/responses"
    "domio/components/config"
)

type AppStatusInfo struct {
    Buildstamp    string `json:"app_buildstamp"`
    BuildAgo      string `json:"app_buildago"`
    BuildTimeDate string `json:"app_buildtimedate"`
    Hash          string `json:"app_hash"`
    Version       string `json:"app_version"`
}

func ShowStatusHandler(w http.ResponseWriter, req *http.Request) {
    info := AppStatusInfo{
        Buildstamp:config.AppStatusInfo.Buildstamp,
        BuildAgo:config.AppStatusInfo.GetBuildAgoValue(),
        Hash:config.AppStatusInfo.Hash,
        Version:config.AppStatusInfo.Version,
    }

    responses.ReturnObjectResponse(w, info)
}