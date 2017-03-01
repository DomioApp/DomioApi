package get_domain_info_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/responses"
    "domio_api/components/tokens"
    "log"
)

type DomainInfo struct {
    Domain     domiodb.Domain
    HostedZone interface{}
}

func GetDomainInfoHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    dataStruct := data.(Data)

    log.Print(dataStruct)

    domainInfo, domainInfoError := domiodb.GetDomainInfo(dataStruct.Domain)

    if (domainInfoError != nil) {

        log.Print(domainInfoError)
        responses.ReturnErrorResponseWithCustomCode(w, domainInfoError, http.StatusNotFound)
        return
    }

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
}