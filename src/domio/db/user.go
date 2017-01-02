package domiodb

import (
    _ "github.com/lib/pq"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "github.com/lib/pq"
    "database/sql"
    "time"
    "log"
)

type EmailAndPasswordPair struct {
    Email    string  `json:"email"`
    Password string  `json:"password"`
}
type NewCustomer struct {
    Email    string  `json:"email"`
    Password string  `json:"password"`
    Id       string  `json:"id"`
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
}

func CreateUser(customer NewCustomer) (sql.Result, *pq.Error) {
    encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
    result, creationError := Db.Exec("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", customer.Id, customer.Email, string(encryptedPassword))

    if (creationError != nil) {
        pqErr := creationError.(*pq.Error)
        return result, pqErr
    }
    return result, nil

}

func LoginUser(user EmailAndPasswordPair) (error, *jwt.StandardClaims, string) {

    userDb := NewCustomer{}

    err := Db.Get(&userDb, "SELECT * FROM users WHERE email=$1", user.Email)

    if err != nil {
        /*if err == sql.ErrNoRows {}*/
        return err, nil, ""
    }

    log.Print(userDb)

    loginError := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))

    newClaims := &jwt.StandardClaims{
        ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
        Subject: userDb.Email,
        Id: userDb.Id,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

    tokenString, _ := token.SignedString([]byte("karamba"))

    return loginError, newClaims, tokenString
}

func GetUser(email string) UserInfo {
    user := UserInfo{}
    err := Db.QueryRowx("SELECT * FROM users where email=$1", email).StructScan(&user)
    if (err != nil) {
        log.Println(err)
    }
    return user
}

func GetUsers() []UserInfo {
    users := []UserInfo{}
    Db.Select(&users, "SELECT email FROM users")
    return users
}
