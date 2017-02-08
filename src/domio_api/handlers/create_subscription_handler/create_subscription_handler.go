package create_subscription_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/sub"
    "domio_api/db"
)

type NewSubscription struct {
    Name       string `json:"name"`
    Domain     string `json:"domain"`
    CustomerId string `json:"customer_id"`
}

func CreateSubscriptionHandler(w http.ResponseWriter, req *http.Request) {
    var newSubscription NewSubscription

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, verifyTokenError)
        return
    }

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

    if (domainInfo != domiodb.Domain{} && domainInfo.IsRented) {
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.DomainIsAlreadyRented, http.StatusUnprocessableEntity)
        return
    }

    newSubscription.CustomerId = userProfile.Id

    stripeSubscription, subscriptionCreationError := createSubscription(&newSubscription, &domainInfo)

    if (subscriptionCreationError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, subscriptionCreationError, http.StatusUnprocessableEntity)
        return
    }

    domiodb.SetDomainAsRented(domainInfo.Name, &userProfile)

    responses.ReturnObjectResponse(w, stripeSubscription)

    defer req.Body.Close()
}

func createSubscription(newSubscription *NewSubscription, domainInfo *domiodb.Domain) (stripe.Sub, error) {
    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"
    subParams := &stripe.SubParams{
        Customer: newSubscription.CustomerId,
        Plan: "month-1",

    }

    subParams.AddMeta("domain", newSubscription.Domain)

    subParams.Quantity = domainInfo.PricePerMonth;
    s, err := sub.New(subParams)
    return *s, err
}