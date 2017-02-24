package types

import "net/http"

type Route struct {
    Name            string
    Method          string
    Pattern         string
    HandlerFunc     HandlerFuncWithParams
    CheckAccessFunc CheckAccessFunc
}

type CheckAccessFunc func(req *http.Request) bool

type HandlerFuncWithParams func(w http.ResponseWriter, req *http.Request, data *interface{})
