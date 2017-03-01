package create_card_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
    "github.com/fatih/color"
    "domio_api/external_api/stripe/card"
)

func CreateCardHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {

    var cardRequest stripe_card_adapter.CardRequest

    err := requests.DecodeJsonRequestBody(req, &cardRequest)

    if err != nil {
        color.Set(color.FgRed)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        log.Print(err)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        color.Unset()
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    existingUser, _ := domiodb.GetUser(userProfile.Email)

    newCard, cardCreationError := stripe_card_adapter.CreateCard(&cardRequest, &existingUser)

    if cardCreationError != nil {
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        log.Print(cardCreationError)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        responses.ReturnErrorResponse(w, cardCreationError)
        return
    }

    responses.ReturnObjectResponse(w, newCard)

    defer req.Body.Close()
}
