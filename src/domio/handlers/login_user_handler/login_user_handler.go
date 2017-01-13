package login_user_handler

import (
    "net/http"
    "domio/db"
    domioerrors  "domio/errors"
    "domio/components/responses"
    "domio/components/requests"
    //"log"
)

type UserLoggedinObject struct {
    Email       string  `json:"email"`
    Id          string  `json:"id"`
    TokenString string  `json:"token"`
}

func LoginUserHandler(w http.ResponseWriter, req *http.Request) {
    var emailAndPasswordPair domiodb.EmailAndPasswordPair

    err := requests.DecodeJsonRequestBody(req, &emailAndPasswordPair)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.IncorrectJSONInputError)
        return
    }

    if (emailAndPasswordPair.IsValid() != true) {
        responses.ReturnErrorResponse(w, domioerrors.PayloadValidationError)
        return
    }

    loginError, userClaims, tokenString := domiodb.LoginUser(emailAndPasswordPair)

    if (loginError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.WrongEmailOrPassword)
        return
    }

    responses.ReturnObjectResponse(w, UserLoggedinObject{Email:userClaims.Subject, Id:userClaims.Id, TokenString:tokenString})

    defer req.Body.Close()
}