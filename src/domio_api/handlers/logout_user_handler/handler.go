package login_user_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/responses"
    "domio_api/components/tokens"
)

func LogoutUserHandler(w http.ResponseWriter, req *http.Request) {
    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    userDomains := domiodb.GetUserDomains(userEmail)

    responses.ReturnObjectResponse(w, userDomains)

    defer req.Body.Close()

}