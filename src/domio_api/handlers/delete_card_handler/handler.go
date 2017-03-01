package delete_card_handler

import (
    "net/http"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/messages"
    "domio_api/external_api/stripe/card"
)

func DeleteCardHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {

    requestVars := mux.Vars(req)
    cardId := requestVars["id"]

    stripe_card_adapter.DeleteCard(userProfile.Id, cardId)

    responses.ReturnObjectResponse(w, messages.CardDeleted)

    defer req.Body.Close()
}
