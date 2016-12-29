package create_subscription_handler

import (
    "net/http"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "domio/components/requests"
    "github.com/stripe/stripe-go"
    "log"
    "github.com/stripe/stripe-go/sub"
)

type NewSubscription struct {
    Name     string `json:"name"`
    Interval stripe.PlanInterval `json:"interval"`
    Amount   uint64 `json:"amount"`
    ID       string `json:"id"`
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
    log.Print(userProfile)

    stripeSubscription, subscriptionCreationError := createSubscription(&newSubscription)

    if (subscriptionCreationError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, subscriptionCreationError, http.StatusUnprocessableEntity)
        return
    }

    responses.ReturnObjectResponse(w, stripeSubscription)

    defer req.Body.Close()
}

func createSubscription(newSubscription *NewSubscription) (stripe.Sub, error) {
    log.Print(newSubscription)
    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

    s, err := sub.New(&stripe.SubParams{
        Customer: "cus_9pFLB7Uou1HKav",
        Plan: "month-plan",

    })
    return *s, err
}