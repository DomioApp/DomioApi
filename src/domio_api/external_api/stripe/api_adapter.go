package stripe_adapter

import (
    "github.com/stripe/stripe-go"
    "domio_api/components/config"
    "github.com/stripe/stripe-go/sub"
    "log"
    "github.com/stripe/stripe-go/customer"
    "domio_api/db"
)

type NewSubscription struct {
    Name       string `json:"name"`
    Domain     string `json:"domain"`
    CustomerId string `json:"customer_id"`
}

func DeleteUserSubscription(userId string, subscriptionId string) {
    stripe.Key = config.Config.STRIPE_KEY

    userSub, err := sub.Cancel(
        subscriptionId,
        &stripe.SubParams{Customer: userId},
    )

    if (err != nil) {
        log.Print(userSub)
        log.Print(nil)
    }
}

func GetUserSubscription(subscriptionId string) *stripe.Sub {
    stripe.Key = config.Config.STRIPE_KEY
    s, _ := sub.Get(subscriptionId, nil)
    return s;
}

func GetUserSubscriptions(userId string) ([]*stripe.Sub, error) {
    stripe.Key = config.Config.STRIPE_KEY

    cstmr, err := customer.Get(userId, nil)
    if (err != nil) {
        log.Print(err)
        return nil, err
    }

    return cstmr.Subs.Values, err

}
func CreateSubscription(newSubscription *NewSubscription, domainInfo *domiodb.Domain) (stripe.Sub, error) {
    stripe.Key = config.Config.STRIPE_KEY

    subParams := &stripe.SubParams{
        Customer: newSubscription.CustomerId,
        Plan: "month-1",
        Quantity : domainInfo.PricePerMonth,
    }

    subParams.AddMeta("domain", newSubscription.Domain)

    s, err := sub.New(subParams)
    return *s, err
}