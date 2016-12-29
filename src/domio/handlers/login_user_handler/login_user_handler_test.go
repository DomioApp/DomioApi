package login_user_handler

import (
    "testing"
    . "github.com/franela/goblin"
    "gopkg.in/h2non/gentleman.v1/plugins/body"
)

func TestLoginUser(t *testing.T) {
    g := Goblin(t)
    g.Describe("LoginUser tests:", func() {
        g.It("Logins with correct email and password", func(done Done) {
            go func() {
                var req = cli.Request()
                req.Path("/users/login")
                req.Method("POST")

                data := map[string]string{"email": "jack@gmail.com", "password":"jack@gmail.com"}
                req.Use(body.JSON(data))

                res, _ := req.Send()

                var jsonResp UserTokenResponseBody
                _ = res.JSON(&jsonResp)



                /*
                fmt.Println(body.JSON(data))
                log.Println(res.StatusCode)
                fmt.Printf("Body: %s", res.String())
                log.Print("**************************************")

                */
                /*
                log.Print("**************************************")
                log.Print(res.String())
                log.Print("**************************************")
                log.Print(jsonResp)
                log.Print(jsonError)
                log.Print("**************************************")
                */

                g.Assert(res.StatusCode).Equal(200)

                g.Assert(jsonResp.Email).Equal("jack@gmail.com")
                //g.Assert(jsonResp.Token).Equal("Wrong JSON input")

                done()
            }()

        })

        g.It("Rejects incorrect email and password", func(done Done) {
            go func() {
                var req = cli.Request()
                req.Path("/users/login")
                req.Method("POST")

                data := map[string]string{"email": "abra@kadabra.com", "password":"boom@boom.com"}
                req.Use(body.JSON(data))

                res, _ := req.Send()

                g.Assert(res.StatusCode).Equal(401)

                done()
            }()

        })

        g.It("Rejects empty email and password", func(done Done) {
            go func() {
                var req = cli.Request()
                req.Path("/users/login")
                req.Method("POST")

                res, _ := req.Send()

                var jsonResp ErrorResponseBody
                res.JSON(&jsonResp)

                g.Assert(res.StatusCode).Equal(401)
                g.Assert(jsonResp.Code).Equal(100)
                g.Assert(jsonResp.Message).Equal("Wrong JSON input")
                done()
            }()
        })

    })
}