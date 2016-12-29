package create_user_handler

import (
	"net/http"
	"encoding/json"
	"domio/db"
	"fmt"
	domioerrors  "domio/errors"
	"domio/messages"
	"domio/components/responses"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go"
	"log"
)

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user domiodb.EmailAndPasswordPair
	err := decoder.Decode(&user)

	defer req.Body.Close()

	if err != nil {
		fmt.Fprintln(w, "{\"error\":\"wrong json input\"")
		responses.ReturnErrorResponse(w, domioerrors.IncorrectJSONInputError)
		return
	}
	_, creationError := domiodb.CreateUser(user)

	if (creationError != nil) {
		if (creationError.Code.Name() == "unique_violation") {
			responses.ReturnErrorResponseWithCustomCode(w, domioerrors.UserEmailExists, http.StatusUnprocessableEntity)
			return
		}
	}
	c, err := createStripeCustomer(user.Email)
	log.Print(c)
	log.Print(err)

	responses.ReturnMessageResponse(w, messages.UserCreated)
}
func createStripeCustomer(email string) (*stripe.Customer, error) {

	stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"
	customerParams := &stripe.CustomerParams{
		Desc: "Customer for " + email,
		Email: email,
	}
	//customerParams.SetSource("tok_189fl92eZvKYlo2C0sjTBkKA") // obtained with Stripe.js
	c, err := customer.New(customerParams)
	return c, err
}