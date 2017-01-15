package handlers
/*

import (
    "gopkg.in/h2non/gentleman.v1/plugins/body"
    "gopkg.in/h2non/gentleman.v1"
    "domio_api/db"
)

type VirtualUser struct {
    Email    string `json:"email"`
    Password string `json:"email"`
    Token    string `json:"token"`
}

var AnonymousUser = VirtualUser{}

func (vu *VirtualUser) Login() {
    var req = cli.Request()
    req.Path("/users/login")
    req.Method("POST")

    data := map[string]string{"email": "jack@gmail.com", "password":"jack@gmail.com"}
    req.Use(body.JSON(data))

    res, _ := req.Send()

    var jsonResp UserTokenResponseBody
    _ = res.JSON(&jsonResp)
    vu.Email = jsonResp.Email
    vu.Token = jsonResp.Token
}

func GetAvailableDomainsAs(user *VirtualUser) (*gentleman.Response, error) {
    var req = cli.Request()
    req.Path("/domains/available")
    req.Method("GET")
    req.AddHeader("Authorization", "Bearer " + user.Token)

    res, err := req.Send()
    return res, err
}

func GetAvailableDomainsViaDummyHTTPAs(user *VirtualUser, url string) (*gentleman.Response, error) {
    var req = cli.Request()
    req.Path("/domains/available")
    req.Method("GET")
    req.AddHeader("Authorization", "Bearer " + user.Token)

    res, err := req.Send()
    return res, err
}

func GetUserDomainsAs(user *VirtualUser) (*gentleman.Response, error) {
    var req = cli.Request()
    req.Path("/domains/user")
    req.Method("GET")
    req.AddHeader("Authorization", "Bearer " + user.Token)

    res, err := req.Send()
    return res, err
}

func CreateDomainAs(user *VirtualUser, domain domiodb.Rental) (*gentleman.Response, error) {
    var req = cli.Request()

    if user != (&VirtualUser{}) {
        req.AddHeader("Authorization", "Bearer " + user.Token)
    }

    req.Path("/domains")
    req.Method("POST")

    req.Use(body.JSON(domain))

    res, err := req.Send()
    return res, err
}
*/
