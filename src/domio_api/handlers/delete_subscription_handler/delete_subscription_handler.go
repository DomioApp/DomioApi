package delete_subscription_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/messages"
)

func DeleteSubscriptionHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }
    sub := domiodb.GetUserSubscription(subscriptionId)

    domiodb.DeleteUserSubscription(userProfile.Id, subscriptionId)

    domiodb.SetDomainAsAvailable(sub.Meta["domain"], &userProfile)

    responses.ReturnObjectResponse(w, messages.SubscriptionDeleted)

    defer req.Body.Close()
}
