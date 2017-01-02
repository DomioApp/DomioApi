package create_card_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "domio/components/requests"
    "log"
    "github.com/fatih/color"
)

func CreateCardHandler(w http.ResponseWriter, req *http.Request) {
    var cardRequest domiodb.CardRequest

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))
    log.Print(userProfile)

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, verifyTokenError)
        return
    }

    err := requests.DecodeJsonRequestBody(req, &cardRequest)

    color.Set(color.FgGreen)
    log.Print(cardRequest)
    color.Unset()

    if err != nil {
        color.Set(color.FgRed)
        log.Print(err)
        color.Unset()
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    existingUser := domiodb.GetUser(userProfile.Email)

    newCard, cardCreationError := domiodb.CreateCard(&cardRequest, &existingUser)
    log.Print(newCard)
    log.Print(cardCreationError)

    responses.ReturnObjectResponse(w, newCard)

    defer req.Body.Close()
}
