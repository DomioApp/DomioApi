package delete_card_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/messages"
    "log"
    "domio_api/external_api/stripe/card"
)

func DeleteCardHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    cardId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    log.Print("==============================================")
    log.Print(userProfile.Id)
    log.Print("==============================================")
    stripe_card_adapter.DeleteCard(userProfile.Id, cardId)

    responses.ReturnObjectResponse(w, messages.CardDeleted)

    defer req.Body.Close()
}
