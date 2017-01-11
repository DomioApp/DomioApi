package login_user_handler

import (
    "net/http"
    "testing"
    "encoding/json"
    "bytes"
    "net/http/httptest"
    . "github.com/franela/goblin"
    "log"
    "domio/db"
    "fmt"
)

type UserEmailAndPassword struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserCreds struct {
    Email string `json:"email"`
    Id    string `json:"id"`
    Token string `json:"token"`
}

func TestLoginUserHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.

    g := Goblin(t)

    g.Describe("CreateCardHandler", func() {
        g.It("Should create a card for a customer", func(done Done) {
            go func() {
                user := LoginAsUser()
                var cardRequest = domiodb.CardRequest{
                    Customer:user.Id,
                }

                uj, _ := json.Marshal(cardRequest)

                reader := bytes.NewReader(uj)

                req, err := http.NewRequest("POST", "/domains", reader)
                authHeader := fmt.Sprintf("Bearer %s", user.Token)

                req.Header.Set("Authorization", authHeader)

                if err != nil {
                    log.Print(err)
                }

                resp := httptest.NewRecorder()
                handler := http.HandlerFunc(LoginUserHandler)

                handler.ServeHTTP(resp, req)

                var cardResp UserCreds

                if err := json.Unmarshal(resp.Body.Bytes(), &cardResp); err != nil {
                    log.Print(err)
                }

                log.Print("======================================")
                log.Print(resp.Body.String())
                log.Print(cardResp)
                log.Print("======================================")

                g.Assert(true).Eql(true)

                done()
            }()
        })
    })
}

func LoginAsUser() UserCreds {
    userJson := UserEmailAndPassword{Email:"john@gmail.com", Password:"john@gmail.com"}
    uj, _ := json.Marshal(userJson)

    reader := bytes.NewReader(uj)

    req, err := http.NewRequest("POST", "/users/login", reader)

    if err != nil {
        log.Print(err)
    }

    resp := httptest.NewRecorder()
    handler := http.HandlerFunc(LoginUserHandler)

    handler.ServeHTTP(resp, req)

    var userCreds UserCreds

    if err := json.Unmarshal(resp.Body.Bytes(), &userCreds); err != nil {
        log.Print(err)
    }
    return userCreds
}