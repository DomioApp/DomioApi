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
    "domio_api/external_api/stripe/customer"
)

func DeleteUserHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool) {

    deletedUser, deleteError := domiodb.DeleteUser(userProfile.Email)

    if (deleteError != nil) {
        log.Println(deleteError)
        responses.ReturnErrorResponse(w, deleteError)
        return
    }

    stripe_user_adapter.DeleteStripeCustomer(deletedUser.StripeId)

    responses.ReturnObjectResponse(w, messages.UserDeleted)
}