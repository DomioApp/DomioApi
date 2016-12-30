package handlers
/*

import (
    "gopkg.in/h2non/gentleman.v1"
    "domio/db"
)

var cli = gentleman.New()

type ErrorResponseBody struct {
    Code    int `json:"code"`
    Message string `json:"message"`
}
type UserTokenResponseBody struct {
    Email string `json:"email"`
    Token string `json:"token"`
}

func init() {
    cli.URL("http://localhost:8080")
    //resetDb()
}

func resetDb() {
    domiodb.Db.Query("DELETE FROM domains;")
    domiodb.Db.Query("DELETE FROM users;")

    jack := domiodb.EmailAndPasswordPair{Email:"jack@gmail.com", Password:"jack@gmail.com"}
    domiodb.CreateUser(jack)

    john := domiodb.EmailAndPasswordPair{Email:"john@gmail.com", Password:"john@gmail.com"}
    domiodb.CreateUser(john)

    jane := domiodb.EmailAndPasswordPair{Email:"jane@gmail.com", Password:"jane@gmail.com"}
    domiodb.CreateUser(jane)

    domiodb.Db.Query("INSERT INTO domains (name, price, owner) VALUES ('jack200.com', 200, 'jack@gmail.com');")
    domiodb.Db.Query("INSERT INTO domains (name, price, owner) VALUES ('jack500.com', 500, 'jack@gmail.com');")
    domiodb.Db.Query("INSERT INTO domains (name, price, owner) VALUES ('jack3000.com', 3000, 'jack@gmail.com');")
    domiodb.Db.Query("INSERT INTO domains (name, price, owner) VALUES ('john100.com', 100, 'john@gmail.com');")

}*/
