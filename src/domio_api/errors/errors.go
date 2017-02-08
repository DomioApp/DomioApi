package errors

type DomioError struct {
    Code    int `json:"code"`
    Message string  `json:"message"`
}

func (e *DomioError) Error() string {
    return e.Message
}

var IncorrectJSONInputError = DomioError{1000, "Wrong JSON input"}
var WrongEmailOrPassword = DomioError{1010, "Wrong email or password"}

var JsonDecodeError = DomioError{1020, "Request Json Decoding Error"}
var PayloadValidationError = DomioError{1030, "Request Json Decoding Error"}
var JwtTokenParseError = DomioError{1040, "JWT token parsing Error"}
var JwtTokenExpiredError = DomioError{1050, "JWT token Expired"}
var JwtClaimsError = DomioError{1060, "JWT claims Error"}

var DomainInRent = DomioError{2010, "Domain is already in rent and cannot be rented until the current rent period ends"}

var DomainIsOwnedByUser = DomioError{2020, "Domain is owned by the user and cannot be rented by the owner"}
var DomainIsAlreadyRented = DomioError{2030, "Domain is already in rent and cannot be rented until the current rent period ends"}
var DomainAlreadyExists = DomioError{2040, "Domain with given name is already created and exists in the database and cannot be added"}
var DomainCheckViolation = DomioError{2050, "Domain fields check violation. Check correctness of price and name fields"}
var RentableDomainNotExist = DomioError{2060, "Domain requested for rent doesn't exist in the database"}
var DomainNotFound = DomioError{2070, "Domain not found"}

var UserEmailExists = DomioError{2080, "User email already exists and can't be used for a new registration"}
var UserEmailOrPasswordEmpty = DomioError{2090, "User email or password for registration is empty"}
var UserNotFound = DomioError{2091, "User not found"}
var UserHasDomains = DomioError{2092, "User has domains which must be deleted before user can be deleted"}

var StripeCustomerCreationError = DomioError{3010, "Stripe customer creation error"}

var UnknownError = DomioError{2990, "Unknown error"}