package get_user_cards_handler

import (
    "net/http"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/external_api/stripe/card"
)

func GetUserCardsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims) {

    userCards, _ := stripe_card_adapter.GetCards(userProfile.Email)

    responses.ReturnObjectResponse(w, userCards)

    defer req.Body.Close()
}