package create_domain_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "domio/components/requests"
)

func CreateDomainHandler(w http.ResponseWriter, req *http.Request) {

    var domain domiodb.Domain
    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    err := requests.DecodeJsonRequestBody(req, &domain)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    newDomain, domainCreationError := domiodb.CreateDomain(domain, userProfile.Email)
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
    responses.ReturnObjectResponse(w, newDomain)

    defer req.Body.Close()
}
