package create_payment_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/tokens"
    "domio/components/responses"
    "domio/components/requests"
    "domio/components/stripe_adapter"
    "log"
)

func CreatePaymentHandler(w http.ResponseWriter, req *http.Request) {
    var paymentRequest domiodb.PaymentRequest

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, verifyTokenError)
        return
    }

    err := requests.DecodeJsonRequestBody(req, &paymentRequest)
	log.Print(paymentRequest)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    charge, err := stripe_adapter.MakePayment()

    if err != nil {
        responses.ReturnErrorResponseWithCustomCode(w, err, http.StatusUnprocessableEntity)
        return
    }

    log.Print(charge.ID)

    newPayment, paymentCreationError := domiodb.CreatePayment(paymentRequest, userProfile.Email, charge.ID)
    if (paymentCreationError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, paymentCreationError, http.StatusUnprocessableEntity)
        return
    }

    responses.ReturnObjectResponse(w, newPayment)

    defer req.Body.Close()
}
