package delete_user_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/messages"
    "log"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
    "domio_api/components/config"
)

func DeleteUserHandler(w http.ResponseWriter, req *http.Request, data *interface{}) {

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    deletedUser, deleteError := domiodb.DeleteUser(userProfile.Email)

    if (deleteError != nil) {
        log.Println(deleteError)
        responses.ReturnErrorResponse(w, deleteError)
        return
    }

    deleteStripeCustomer(deletedUser.StripeId)

    responses.ReturnObjectResponse(w, messages.UserDeleted)

    defer req.Body.Close()
}

func deleteStripeCustomer(stripe_customer_id string) (*stripe.Customer, error) {

    stripe.Key = config.Config.STRIPE_KEY

    c, err := customer.Del(stripe_customer_id)
    return c, err
}