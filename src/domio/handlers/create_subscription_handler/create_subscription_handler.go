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
	var newPlan NewSubscription

	userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

	if (verifyTokenError != domioerrors.DomioError{}) {
		responses.ReturnErrorResponse(w, verifyTokenError)
		return
	}

	err := requests.DecodeJsonRequestBody(req, &newPlan)

	if err != nil {
		responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
		return
	}
	log.Print(userProfile)

	stripePlan, planCreationError := createSubscription(&newPlan)

	if (planCreationError != nil) {
		responses.ReturnErrorResponseWithCustomCode(w, planCreationError, http.StatusUnprocessableEntity)
		return
	}

	responses.ReturnObjectResponse(w, stripePlan)

	defer req.Body.Close()
}

func createSubscription(newPlan *NewSubscription) (stripe.Sub, error) {
	log.Print(newPlan)
	stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

	s, err := sub.New(&stripe.SubParams{
		Customer: "cus_9pFLB7Uou1HKav",
		Plan: "month-plan",
	})
	return *s, err
}