package update_subscription_records_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/db"
    "domio_api/external_api/r53"
)

func UpdateSubscriptionRecordsHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))
    log.Print(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    log.Print(userEmail)

    domain, err := domiodb.GetDomainInfoBySubscriptionId(subscriptionId)
    r53.UpdateCNAME(domain.ZoneId.String, "www." + domain.Name, "heyho100", 3600, 200)

    if (err != nil) {
        log.Print(err)

    }

    log.Print("===========================================================")
    log.Print(domain)
    log.Print("===========================================================")
    responses.ReturnObjectResponse(w, domain)

    defer req.Body.Close()
}