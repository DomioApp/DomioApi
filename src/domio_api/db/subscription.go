package domiodb

import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
    "log"
    "github.com/stripe/stripe-go/sub"
    "domio_api/components/config"
)

func GetUserSubscriptions(userId string) ([]*stripe.Sub, error) {
    stripe.Key = config.Config.STRIPE_KEY

    cstmr, err := customer.Get(userId, nil)
    if (err != nil) {
        log.Print(err)
        return nil, err
    }

    return cstmr.Subs.Values, err

}

func GetUserSubscription(subscriptionId string) *stripe.Sub {
    stripe.Key = config.Config.STRIPE_KEY
    s, _ := sub.Get(subscriptionId, nil)
    return s;
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
