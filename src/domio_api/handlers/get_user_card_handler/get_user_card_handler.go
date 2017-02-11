package get_user_card_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/db"
    "github.com/gorilla/mux"
)

func GetUserCardHandler(w http.ResponseWriter, req *http.Request) {

    requestVars := mux.Vars(req)
    cardId := requestVars["id"]

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    userCard, _ := domiodb.GetCard(userEmail, cardId)

    responses.ReturnObjectResponse(w, userCard)

    defer req.Body.Close()
}