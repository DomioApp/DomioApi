package create_domain_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
)

func CreateDomainHandler(w http.ResponseWriter, req *http.Request) {

    var newDomain domiodb.Domain
    var createdDomain domiodb.DomainJson

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

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

    responses.ReturnObjectResponse(w, createdDomain)

    defer req.Body.Close()
}
