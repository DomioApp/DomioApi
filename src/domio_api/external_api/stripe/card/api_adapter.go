package stripe_card_adapter

import (
    "github.com/stripe/stripe-go"
    "domio_api/components/config"
    "github.com/stripe/stripe-go/card"
    "log"
    "strconv"
    "domio_api/db"
    domioerrors  "domio_api/errors"
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

func CreateCard(cardRequest *CardRequest, user *domiodb.UserInfo) (Card, *domioerrors.DomioError) {
    var cardResult Card
    log.Print("******************************************")
    log.Print(cardRequest)
    log.Print("------------------------------------------")
    log.Print(user)
    log.Print("******************************************")

    stripe.Key = config.Config.STRIPE_KEY

    cardParams := stripe.CardParams{
        Customer: user.StripeId,
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
    userInfo, _ := domiodb.GetUser(userEmail);

    stripe.Key = config.Config.STRIPE_KEY
    cards := card.List(&stripe.CardListParams{Customer: userInfo.StripeId})

    cardsList := make([]stripe.Card, 0)

    for cards.Next() {
        c := *cards.Card()
        cardsList = append(cardsList, c)
    }

    return cardsList, nil
}
func GetCard(userEmail string, id string) (*stripe.Card, *domioerrors.DomioError) {
    userInfo, _ := domiodb.GetUser(userEmail);

    stripe.Key = config.Config.STRIPE_KEY

    userCard, cardError := card.Get(id, &stripe.CardParams{Customer: userInfo.StripeId})

    if (cardError != nil) {
        log.Print(cardError)
    }

    return userCard, nil
}

func DeleteCard(userId string, id string) (*stripe.Card, *domioerrors.DomioError) {

    stripe.Key = config.Config.STRIPE_KEY

    userCard, cardError := card.Del(id, &stripe.CardParams{Customer: userId})

    if (cardError != nil) {
        log.Print(cardError)
    }

    return userCard, nil
}
