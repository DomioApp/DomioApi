package create_domain_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
    "domio_api/external_api/r53"
)

func CreateDomainHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool) {

    var newDomain domiodb.Domain

    err := requests.DecodeJsonRequestBody(req, &newDomain)

    if err != nil {
        log.Print(err)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    //============================================================================================================================

    createdDomain, domainCreationError := domiodb.CreateDomain(newDomain, userProfile.Email)

    if (domainCreationError != nil) {
        if (domainCreationError.Code.Name() == "unique_violation") {
            responses.ReturnErrorResponseWithCustomCode(w, domioerrors.DomainAlreadyExists, http.StatusUnprocessableEntity)
            return
        }

        if (domainCreationError.Code.Name() == "check_violation") {
            responses.ReturnErrorResponseWithCustomCode(w, domioerrors.DomainCheckViolation, http.StatusUnprocessableEntity)
            return
        }

        responses.ReturnErrorResponseWithCustomCode(w, domainCreationError, http.StatusUnprocessableEntity)
        return
    }

    createdDomainZone, _ := r53.CreateDomainZone(createdDomain)

    domiodb.SetDomainZoneId(createdDomain, createdDomainZone.HostedZone.Id)
    domiodb.SetDomainNameServers(createdDomain, createdDomainZone.DelegationSet.NameServers[0], createdDomainZone.DelegationSet.NameServers[1], createdDomainZone.DelegationSet.NameServers[2], createdDomainZone.DelegationSet.NameServers[3])

    responses.ReturnObjectResponse(w, createdDomain)
}
