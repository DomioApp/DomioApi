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
    "github.com/fatih/color"
)

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
    var user domiodb.EmailAndPasswordPair

    err := requests.DecodeJsonRequestBody(req, &user)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    existingUser := domiodb.GetUser(user.Email)
    if (existingUser != domiodb.UserInfo{}) {
        color.Set(color.FgRed)
        log.Print("===================")
        log.Print("User already exists")
        log.Print("===================")
        color.Unset()
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.UserEmailExists, http.StatusUnprocessableEntity)
        return
    }

    log.Print("Creating stripe user")

    newlyCreatedCustomer, err := createStripeCustomer(user.Email)

    if err != nil {
        log.Print(err)
        responses.ReturnErrorResponse(w, domioerrors.StripeCustomerCreationError)
        return
    }

    new_customer := domiodb.NewCustomer{Id:newlyCreatedCustomer.ID, Email:newlyCreatedCustomer.Email, Password:user.Password}

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