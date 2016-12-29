package get_user_domains_handler

import (
	"net/http"
	domioerrors  "domio/errors"
	"domio/components/tokens"
	"domio/components/responses"
	"domio/db"
)

func GetUserDomainsHandler(w http.ResponseWriter, req *http.Request) {
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