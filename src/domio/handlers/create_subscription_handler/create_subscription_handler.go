package create_subscription_handler

import (
    "net/http"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "domio/components/requests"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/sub"
    "domio/db"
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

    newSubscription.CustomerId = userProfile.Id

    stripeSubscription, subscriptionCreationError := createSubscription(&newSubscription)

    if (subscriptionCreationError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, subscriptionCreationError, http.StatusUnprocessableEntity)
        return
    }

    responses.ReturnObjectResponse(w, stripeSubscription)

    defer req.Body.Close()
}

func createSubscription(newSubscription *NewSubscription) (stripe.Sub, error) {
    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"
    subParams := &stripe.SubParams{
        Customer: newSubscription.CustomerId,
        Plan: "month-1",

    }

    subParams.AddMeta("Domain", newSubscription.Domain)
    domain, _ := domiodb.GetDomain(newSubscription.Domain)

    subParams.Quantity = domain.PricePerMonth;
    s, err := sub.New(subParams)
    return *s, err
}