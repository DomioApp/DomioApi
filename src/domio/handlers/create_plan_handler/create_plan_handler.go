package create_plan_handler

import (
	"net/http"
	domioerrors  "domio/errors"
	"domio/components/tokens"
	"domio/components/responses"
	"domio/components/requests"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go"
	"log"
)

type NewPlan struct {
	Name     string `json:"name"`
	Interval stripe.PlanInterval `json:"interval"`
	Amount   uint64 `json:"amount"`
	ID       string `json:"id"`
}

func CreatePlanHandler(w http.ResponseWriter, req *http.Request) {
	var newPlan NewPlan

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

	stripePlan, planCreationError := createPlan(&newPlan)

	if (planCreationError != nil) {
		responses.ReturnErrorResponseWithCustomCode(w, planCreationError, http.StatusUnprocessableEntity)
		return
	}

	responses.ReturnObjectResponse(w, stripePlan)

	defer req.Body.Close()
}

func createPlan(newPlan *NewPlan) (stripe.Plan, error) {
	log.Print(newPlan)
	stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

	p, err := plan.New(&stripe.PlanParams{
		Amount: newPlan.Amount,
		Interval: newPlan.Interval,
		Name: newPlan.Name,
		Currency: "usd",
		ID: newPlan.ID,
	})
	return *p, err
}