package show_status_handler

import (
    "net/http"
    "domio_api/components/responses"
    "domio_api/components/config"
    "domio_api/components/tokens"
)

type AppStatusInfo struct {
    Version       string `json:"app_version"`
    BuildAgo      string `json:"app_buildago"`
    Buildstamp    string `json:"app_buildstamp"`
    BuildTimeDate string `json:"app_builddatetime"`
    Hash          string `json:"app_hash"`
}

func ShowStatusHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {
    info := AppStatusInfo{
        Buildstamp:config.AppStatusInfo.Buildstamp,
        BuildTimeDate:config.AppStatusInfo.GetBuildDateTime(),
        BuildAgo:config.AppStatusInfo.GetBuildAgoValue(),
        Hash:config.AppStatusInfo.Hash,
        Version:config.AppStatusInfo.Version,
    }

    responses.ReturnObjectResponse(w, info)
}