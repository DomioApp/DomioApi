package create_rental_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "domio/components/requests"
)

func CreateRentalHandler(w http.ResponseWriter, req *http.Request) {
    var rental domiodb.Rental

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, verifyTokenError)
        return
    }

    err := requests.DecodeJsonRequestBody(req, &rental)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    newRental, rentalCreationError := domiodb.CreateRental(rental, userProfile.Email)
    if (rentalCreationError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, rentalCreationError, http.StatusUnprocessableEntity)
        return
    }

    responses.ReturnObjectResponse(w, newRental)

    defer req.Body.Close()
}
