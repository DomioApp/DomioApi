package get_domain_info_handler

import (
    "net/http"
    "domio_api/db"
    "github.com/gorilla/mux"
    "domio_api/components/responses"
    "domio_api/components/tokens"
    domioerrors  "domio_api/errors"
    "fmt"
    "log"
)

type DomainInfo struct {
    Domain     domiodb.Domain
    HostedZone interface{}
}

func GetDomainInfoHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    key := requestVars["key"]
    value := requestVars["value"]



    log.Print("============================================")
    log.Print(key)
    log.Print(value)
    log.Print("============================================")

    var isAuthenticated = false

    defer req.Body.Close()

    _, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError == domioerrors.DomioError{}) {
        isAuthenticated = true
    } else {
        isAuthenticated = false
    }

    fmt.Printf("isAuthenticated: %s", isAuthenticated)

    var domainInfo domiodb.Domain
    var err error

    if (key == "name") {
        domainInfo, err = domiodb.GetDomainInfo(value)
    } else if (key == "sub") {
        domainInfo, err = domiodb.GetDomainInfoBySubscriptionId(value)
    }

    if (err != nil) {
        responses.ReturnErrorResponse(w, err)
        return
    }

    if (isAuthenticated) {

        //hostedZoneInfo := domiodb.GetHostedZone(&domainInfo)
        //fullInfo.HostedZone = hostedZoneInfo


        outDomainInfo := domiodb.DomainJson{
            Name:domainInfo.Name,
            PricePerMonth:domainInfo.PricePerMonth,
            Owner:domainInfo.Owner,
            IsRented:domainInfo.IsRented,
            IsVisible:domainInfo.IsVisible,
        }

        if (domainInfo.RentedBy.Valid) {
            outDomainInfo.RentedBy = domainInfo.RentedBy.String
        }

        if (domainInfo.NS1.Valid) {
            outDomainInfo.NS1 = domainInfo.NS1.String
        }

        if (domainInfo.NS2.Valid) {
            outDomainInfo.NS2 = domainInfo.NS2.String
        }

        if (domainInfo.NS3.Valid) {
            outDomainInfo.NS3 = domainInfo.NS3.String
        }

        if (domainInfo.NS4.Valid) {
            outDomainInfo.NS4 = domainInfo.NS4.String
        }

        if (domainInfo.ZoneId.Valid) {
            outDomainInfo.ZoneId = domainInfo.ZoneId.String
        }

        responses.ReturnObjectResponse(w, outDomainInfo)

    } else {
        domainInfo := domiodb.AvailableDomainJson{
            Name:domainInfo.Name,
            PricePerMonth:domainInfo.PricePerMonth,
        }
        responses.ReturnObjectResponse(w, domainInfo)
    }
}