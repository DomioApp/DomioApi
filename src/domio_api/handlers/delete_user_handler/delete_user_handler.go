package delete_user_handler

import (
    "net/http"
    "domio_api/db"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/messages"
    "log"
    "domio_api/external_api/stripe/customer"
)

func DeleteUserHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {

    deletedUser, deleteError := domiodb.DeleteUser(userProfile.Email)

    if (deleteError != nil) {
        log.Println(deleteError)
        responses.ReturnErrorResponse(w, deleteError)
        return
    }

    stripe_user_adapter.DeleteStripeCustomer(deletedUser.StripeId)

    responses.ReturnObjectResponse(w, messages.UserDeleted)
}