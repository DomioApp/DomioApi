package show_status_handler

import (
    "net/http"
    "domio/components/responses"
    "domio/components/config"
)

type AppStatusInfo struct {
    Version       string `json:"app_version"`
    BuildAgo      string `json:"app_buildago"`
    Buildstamp    string `json:"app_buildstamp"`
    BuildTimeDate string `json:"app_buildtimedate"`
    Hash          string `json:"app_hash"`
}

func ShowStatusHandler(w http.ResponseWriter, req *http.Request) {
    info := AppStatusInfo{
        Buildstamp:config.AppStatusInfo.Buildstamp,
        BuildTimeDate:config.AppStatusInfo.GetBuildDateTime(),
        BuildAgo:config.AppStatusInfo.GetBuildAgoValue(),
        Hash:config.AppStatusInfo.Hash,
        Version:config.AppStatusInfo.Version,
    }

    responses.ReturnObjectResponse(w, info)
}