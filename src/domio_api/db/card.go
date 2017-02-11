package domiodb

import (
    _ "github.com/lib/pq"
    domioerrors "domio_api/errors"
    "log"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/card"
    "strconv"
)

type CardRequest struct {
    Customer string
    Currency string

    Name     string `json:"name_on_the_card"`
    Month    uint64 `json:"expiry_month"`
    Year     uint64 `json:"expiry_year"`
    CVC      uint64 `json:"cvc"`
    Number   uint64 `json:"card_number"`
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
        Month: strconv.FormatUint(cardRequest.Month, 10),
        Year: strconv.FormatUint(cardRequest.Year, 10),
        CVC: strconv.FormatUint(cardRequest.CVC, 10),
        Number: strconv.FormatUint(cardRequest.Number, 10),
    }

    c, err := card.New(&cardParams)

    log.Print("**********************************************")
    log.Print(c)
    log.Print(err)
    log.Print("**********************************************")

    return cardResult, nil
}

func GetCards(userEmail string) ([]stripe.Card, *domioerrors.DomioError) {
    userInfo, _ := GetUser(userEmail);

    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"
    cards := card.List(&stripe.CardListParams{Customer: userInfo.Id})

    cardsList := make([]stripe.Card, 0)

    for cards.Next() {
        c := *cards.Card()
        cardsList = append(cardsList, c)
    }

    return cardsList, nil
}
