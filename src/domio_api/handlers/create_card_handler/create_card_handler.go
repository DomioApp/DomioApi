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

    if err != nil {
        color.Set(color.FgRed)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        log.Print(err)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        color.Unset()
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    existingUser := domiodb.GetUser(userProfile.Email)

    newCard, cardCreationError := domiodb.CreateCard(&cardRequest, &existingUser)

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
