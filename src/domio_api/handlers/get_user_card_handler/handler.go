package get_user_card_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "github.com/gorilla/mux"
    "domio_api/external_api/stripe/card"
)

func GetUserCardHandler(w http.ResponseWriter, req *http.Request, data *interface{}) {

    requestVars := mux.Vars(req)
    cardId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    userCard, _ := stripe_card_adapter.GetCard(userEmail, cardId)

    responses.ReturnObjectResponse(w, userCard)

    defer req.Body.Close()
}