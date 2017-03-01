package get_subscription_handler

import (
    "net/http"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "log"
    "domio_api/external_api/stripe/subscription"
)

func GetSubscriptionHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["id"]

    userEmail := userProfile.Email
    log.Print(userEmail)

    subscription := stripe_subscription_adapter.GetUserSubscription(subscriptionId)

    responses.ReturnObjectResponse(w, subscription)
}