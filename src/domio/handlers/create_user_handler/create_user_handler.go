package create_user_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/responses"
    "github.com/stripe/stripe-go/customer"
    "github.com/stripe/stripe-go"
    "log"
    "domio/components/requests"
    "domio/messages"
)

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
    var user domiodb.EmailAndPasswordPair

    err := requests.DecodeJsonRequestBody(req, &user)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    existingUser := domiodb.GetUser(user.Email)
    log.Print("+++++++++++++++++++++++++++")
    log.Print(existingUser)
    log.Print("+++++++++++++++++++++++++++")
    if (existingUser != domiodb.UserInfo{}) {
        log.Print("====================")
        log.Print("User exists")
        log.Print("====================")
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.UserEmailExists, http.StatusUnprocessableEntity)
        return
    }

    log.Print("Creating stripe user")

    newCustomer, err := createStripeCustomer(user.Email)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.StripeCustomerCreationError)
        return
    }

    new_customer := domiodb.NewCustomer{Id:newCustomer.ID, Email:newCustomer.Email, Password:user.Password}

    log.Print(new_customer)
    newUser, userCreationError := domiodb.CreateUser(new_customer)
    log.Print(newUser)

    if (userCreationError != nil) {
        if (userCreationError.Code.Name() == "unique_violation") {
            responses.ReturnErrorResponseWithCustomCode(w, domioerrors.UserEmailExists, http.StatusUnprocessableEntity)
            return
        }
    }

    responses.ReturnMessageResponse(w, messages.UserCreated)
    defer req.Body.Close()
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