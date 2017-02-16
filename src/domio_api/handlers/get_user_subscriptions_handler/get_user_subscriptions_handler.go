package get_user_subscriptions_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "log"
    "github.com/fatih/color"
    "domio_api/external_api/stripe"
)

func GetUserSubscriptionsHandler(w http.ResponseWriter, req *http.Request) {
    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userSubscriptions, _ := stripe_adapter.GetUserSubscriptions(userProfile.Id)

    color.Set(color.FgHiCyan)
    log.Print(userSubscriptions)
    color.Unset()

    responses.ReturnObjectResponse(w, userSubscriptions)

    defer req.Body.Close()
}