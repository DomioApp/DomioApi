package get_user_subscriptions_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/db"
    "log"
    "github.com/fatih/color"
)

func GetUserSubscriptionsHandler(w http.ResponseWriter, req *http.Request) {
    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userSubscriptions, _ := domiodb.GetUserSubscriptions(userProfile.Id)

    color.Set(color.FgHiCyan)
    log.Print(userSubscriptions)
    color.Unset()

    responses.ReturnObjectResponse(w, userSubscriptions)

    defer req.Body.Close()
}