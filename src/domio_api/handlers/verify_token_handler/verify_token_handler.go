package verify_token_handler

import (
    "net/http"
    "encoding/json"
    "github.com/dgrijalva/jwt-go"
    domioerrors  "domio_api/errors"
    "log"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
)

type Token struct {
    Value string `json:"value"`
}
type Claims struct {
    jwt.StandardClaims
    Email string `json:"email"`
}

func VerifyTokenHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool, data interface{}) {
    var tokenToVerify Token

    err := requests.DecodeJsonRequestBody(req, &tokenToVerify)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.IncorrectJSONInputError)
        return
    }

    var jwtParser = jwt.Parser{UseJSONNumber:false}
    token, jwtParseError := jwtParser.Parse(tokenToVerify.Value, tokens.TokenFunc)

    if jwtParseError != nil {
        log.Print(jwtParseError)
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        json.NewEncoder(w).Encode(claims)
    } else {
        responses.ReturnErrorResponse(w, domioerrors.JwtClaimsError)
    }
}
