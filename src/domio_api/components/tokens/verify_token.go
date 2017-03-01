package tokens

import (
    "github.com/dgrijalva/jwt-go"
    "strings"
    "fmt"
    "domio_api/errors"
)

type UserTokenWithClaims struct {
    Email   string
    Expires int64
    Id      string
}

func VerifyTokenString(tokenString string) (*UserTokenWithClaims, *domioerrors.DomioError) {
    var userTokenWithClaims *UserTokenWithClaims
    pureTokenString := strings.TrimPrefix(tokenString, "Bearer ")

    var jwtParser = jwt.Parser{UseJSONNumber:true}
    token, jwtParseError := jwtParser.Parse(pureTokenString, TokenFunc)

    if validationError, ok := jwtParseError.(*jwt.ValidationError); ok {
        if validationError.Errors & jwt.ValidationErrorMalformed != 0 {
            fmt.Println("That's not even a token")
            return nil, &domioerrors.JwtTokenParseError
        } else if validationError.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
            return nil, &domioerrors.JwtTokenExpiredError
        } else {
            fmt.Println("Couldn't handle this token 1:", jwtParseError)
        }
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        stClaims := jwt.StandardClaims{Subject:claims["sub"].(string), Id:claims["jti"].(string)}
        userTokenWithClaims = &UserTokenWithClaims{Email:stClaims.Subject, Expires:stClaims.ExpiresAt, Id:stClaims.Id}
        return userTokenWithClaims, nil
    } else {
        return nil, &domioerrors.JwtTokenParseError
    }
}

func TokenFunc(token *jwt.Token) (interface{}, error) {
    return []byte("karamba"), nil
}