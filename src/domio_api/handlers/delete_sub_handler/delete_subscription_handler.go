package delete_sub_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/messages"
    "domio_api/external_api/stripe/subscription"
)

func DeleteSubscriptionHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    requestVars := mux.Vars(req)
    subscriptionId := requestVars["subId"]

    sub := stripe_subscription_adapter.GetUserSubscription(subscriptionId)

    stripe_subscription_adapter.DeleteUserSubscription(userProfile.Id, subscriptionId)

    domiodb.SetDomainAsAvailable(sub.Meta["domain"], userProfile)

    responses.ReturnObjectResponse(w, messages.SubscriptionDeleted)
}
