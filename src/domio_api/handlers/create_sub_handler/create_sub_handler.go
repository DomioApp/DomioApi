package create_subscription_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "domio_api/db"
    "domio_api/external_api/stripe/subscription"
)

func CreateSubscriptionHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {
    var newSubscription stripe_subscription_adapter.NewSubscription

    err := requests.DecodeJsonRequestBody(req, &newSubscription)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    domainInfo, domainError := domiodb.GetDomainInfo(newSubscription.Domain)

    if (domainError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, domainError, http.StatusUnprocessableEntity)
        return
    }

    if (domainInfo != nil && domainInfo.IsRented) {
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.DomainIsAlreadyRented, http.StatusUnprocessableEntity)
        return
    }

    newSubscription.CustomerId = userProfile.Id

    stripeSubscription, subscriptionCreationError := stripe_subscription_adapter.CreateSubscription(&newSubscription, domainInfo)

    if (subscriptionCreationError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, subscriptionCreationError, http.StatusUnprocessableEntity)
        return
    }

    domiodb.SetDomainAsRented(domainInfo.Name, stripeSubscription.ID, userProfile)

    responses.ReturnObjectResponse(w, stripeSubscription)
}
