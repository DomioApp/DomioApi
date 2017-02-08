package domiodb

import (
    _ "github.com/lib/pq"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "github.com/lib/pq"
    "time"
    "log"
    "domio_api/components/logger"
    "github.com/fatih/color"
    domioerrors  "domio_api/errors"
)

type EmailAndPasswordPair struct {
    Email    string  `json:"email"`
    Password string  `json:"password"`
}
type NewCustomer struct {
    Email    string  `json:"email"`
    Password string  `json:"password"`
    Id       string  `json:"id"`
    Role     string  `json:"role"`
}

func (emailAndPasswordPair *EmailAndPasswordPair) IsValid() bool {
    return emailAndPasswordPair.Email != "" && emailAndPasswordPair.Password != ""
}

type UserToken struct {
    Email string `json:"email"`
    Token string `json:"token"`
}

type UserInfo struct {
    Email    string `json:"email" db:"email"`
    Password string `json:"password" db:"password"`
    Id       string `json:"id" db:"id"`
    Role     string `json:"role" db:"role"`
}

func CreateUser(customer NewCustomer) (*jwt.StandardClaims, string, error) {
    encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
    _, creationError := Db.Exec("INSERT INTO users (id, email, password, role) VALUES ($1, $2, $3, $4)", customer.Id, customer.Email, string(encryptedPassword), "user")

    if (creationError != nil) {
        pqErr := creationError.(*pq.Error)
        log.Println(pqErr)
        return nil, "", pqErr
    }
    userToLogin := EmailAndPasswordPair{Email:customer.Email, Password:customer.Password}

    return LoginUser(userToLogin)
}

func LoginUser(user EmailAndPasswordPair) (*jwt.StandardClaims, string, error) {

    userDb := NewCustomer{}

    err := Db.Get(&userDb, "SELECT * FROM users WHERE email=$1", user.Email)

    if err != nil {
        /*if err == sql.ErrNoRows {}*/
        return nil, "", err
    }

    loginError := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))

    newClaims := &jwt.StandardClaims{
        ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
        Subject: userDb.Email,
        Id: userDb.Id,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

    tokenString, _ := token.SignedString([]byte("karamba"))

    if (loginError != nil) {
        log.Println(loginError)
        logger.Logger.Err("User login error: " + user.Email)
    }

    if (newClaims != nil) {
        log.Println(newClaims)
        logger.Logger.Info("User logged in: " + newClaims.Subject)
    }

    return newClaims, tokenString, loginError
}

func DeleteUser(userEmail string) (UserInfo, *domioerrors.DomioError) {
    user, userError := GetUser(userEmail)

    if (userError != nil) {
        log.Print(userError)
        return UserInfo{}, &domioerrors.UserNotFound
    }

    userDomains := GetUserDomains(userEmail);

    if (len(userDomains) != 0) {
        return UserInfo{}, &domioerrors.UserHasDomains
    }

    result := Db.MustExec("DELETE FROM users where email=$1", userEmail)

    rowsAffected, err := result.RowsAffected()

    if (err != nil) {
        color.Set(color.FgHiRed)
        log.Print(err)
        color.Unset()
        return UserInfo{}, &domioerrors.UserNotFound
    }

    color.Set(color.FgHiCyan)
    log.Print(rowsAffected)
    color.Unset()

    if (rowsAffected == 0) {
        return UserInfo{}, &domioerrors.UserNotFound
    }

    return user, nil
}

func GetUser(email string) (UserInfo, *domioerrors.DomioError) {
    user := UserInfo{}
    err := Db.QueryRowx("SELECT * FROM users where email=$1", email).StructScan(&user)

    if (err != nil) {
        log.Println(err)
        return UserInfo{}, &domioerrors.UserNotFound
    }
    return user, nil
}

func GetUsers() []UserInfo {
    users := []UserInfo{}
    Db.Select(&users, "SELECT email FROM users")
    return users
}
