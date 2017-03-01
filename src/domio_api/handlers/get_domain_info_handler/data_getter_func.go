package get_domain_info_handler

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Data struct {
    Domain string
}

func DataGetterFunc(req *http.Request) interface{} {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    return Data{Domain:domainName}
}