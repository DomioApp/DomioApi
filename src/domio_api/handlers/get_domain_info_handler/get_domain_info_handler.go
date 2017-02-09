package get_domain_info_handler

import (
    "net/http"
    "domio_api/db"
    "github.com/gorilla/mux"
    "domio_api/components/responses"
    "domio_api/components/tokens"
    domioerrors  "domio_api/errors"
    "fmt"
)

type DomainInfo struct {
    Domain     domiodb.Domain
    HostedZone interface{}
}

func GetDomainInfoHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    var isAuthenticated = false

    defer req.Body.Close()

    _, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError == domioerrors.DomioError{}) {
        isAuthenticated = true
    } else {
        isAuthenticated = false
    }

    fmt.Printf("isAuthenticated: %s", isAuthenticated)

    domainInfo, err := domiodb.GetDomainInfo(domainName)
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