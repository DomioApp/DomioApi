package get_available_domains_handler

import (
    "testing"
    . "github.com/franela/goblin"
    "domio/db"
)

func TestGetAvailableDomainsHandler(t *testing.T) {
    g := Goblin(t)
    g.Describe("GetAvailableDomainsHandler tests", func() {
        g.It("Should login and get available domains", func(done Done) {
            go func() {
                var virtualUser = VirtualUser{Email:"jack@gmail.com", Password:"jack@gmail.com"}
                virtualUser.Login()
                domainsResponse, _ := GetAvailableDomainsAs(&virtualUser)

                var domainsList []domiodb.Rental
                domainsResponse.JSON(&domainsList)

                g.Assert(domainsResponse.StatusCode).Equal(200)
                g.Assert(len(domainsList)).Equal(5)

                g.Assert(domainsList[0]).Equal(domiodb.Rental{Name:"john100.com", Owner:"john@gmail.com", Price:100})

                done()
            }()
        })
    })
}