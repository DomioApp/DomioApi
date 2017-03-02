package get_user_domains_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/tokens"
    "domio_api/components/responses"
)

func GetUserDomainsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    userDomains := domiodb.GetUserDomains(userProfile.Email)

    responses.ReturnObjectResponse(w, formatDomainsAsJson(userDomains))
}

func formatDomainsAsJson(domains []domiodb.Domain) []domiodb.DomainJson {

    var domainsJson []domiodb.DomainJson = make([]domiodb.DomainJson, 0)

    for i := range domains {
        currentDomain := domains[i]

        domain := domiodb.DomainJson{
            Name:currentDomain.Name,
            PricePerMonth:currentDomain.PricePerMonth,
        }

        domainsJson = append(domainsJson, domain)
    }
    return domainsJson
}