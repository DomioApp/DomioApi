package domiodb

import (
    _ "github.com/lib/pq"
    domioerrors "domio_api/errors"
    "log"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/card"
)

type CardRequest struct {
    Customer string
    Currency string

    Name     string `json:"customer_name"`
    Month    string `json:"month"`
    Year     string `json:"year"`
    CVC      string `json:"cvc"`
    Number   string `json:"number"`
}

type Card struct {
    Name string `json:"name" db:"name"`
}

func CreateCard(cardRequest *CardRequest, user *UserInfo) (Card, *domioerrors.DomioError) {
    var cardResult Card
    log.Print("******************************************")
    log.Print(cardRequest)
    log.Print("------------------------------------------")
    log.Print(user)
    log.Print("******************************************")

    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

    cardParams := stripe.CardParams{
        Customer: user.Id,
        Currency: "USD",

        Name: cardRequest.Name,
        Month: cardRequest.Month,
        Year: cardRequest.Year,
        CVC: cardRequest.CVC,
        Number: cardRequest.Number,
    }

    c, err := card.New(&cardParams)

    log.Print("**********************************************")
    log.Print(c)
    log.Print(err)
    log.Print("**********************************************")

    return cardResult, nil
}
