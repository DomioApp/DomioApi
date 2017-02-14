package create_user_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/responses"
    "github.com/stripe/stripe-go/customer"
    "github.com/stripe/stripe-go"
    "log"
    "domio_api/components/requests"
    "domio_api/handlers/login_user_handler"
    "domio_api/components/config"
)

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
    //colorRed := color.New(color.FgRed)

    var user domiodb.EmailAndPasswordPair

    err := requests.DecodeJsonRequestBody(req, &user)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    if (user.Email == "" || user.Password == "") {
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.UserEmailOrPasswordEmpty, http.StatusUnprocessableEntity)
        return
    }

    existingUser, _ := domiodb.GetUser(user.Email)

    if (existingUser != domiodb.UserInfo{}) {
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
    newUser, token, userCreationError := domiodb.CreateUser(new_customer)

    if (userCreationError != nil) {
        //TODO manage the error properly
        //if (userCreationError.Code.Name() == "unique_violation") {
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.UserEmailExists, http.StatusUnprocessableEntity)
        return
        //}
    }

    log.Print("--------------------------------------")
    log.Print(newUser)
    log.Print(token)
    log.Print(userCreationError)
    log.Print("--------------------------------------")

    loggedInUser := login_user_handler.UserLoggedinObject{
        Email:newUser.Subject,
        Id:newUser.Id,
        TokenString:token,
    }

    responses.ReturnObjectResponse(w, loggedInUser)
    defer req.Body.Close()
}

func createStripeCustomer(email string) (*stripe.Customer, error) {

    stripe.Key = config.Config.STRIPE_KEY
    customerParams := &stripe.CustomerParams{
        Desc: "Customer for " + email,
        Email: email,
    }
    //customerParams.SetSource("tok_189fl92eZvKYlo2C0sjTBkKA") // obtained with Stripe.js
    c, err := customer.New(customerParams)
    return c, err
}