package domiodb

import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
    "log"
    "github.com/stripe/stripe-go/sub"
)

func GetUserSubscriptions(userId string) (*stripe.SubList, error) {
    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

    customer, err := customer.Get(userId, nil)
    if (err != nil) {
        log.Print(err)
        return nil, err
    }
    return customer.Subs, err

}

func GetUserSubscription(subscriptionId string) *stripe.Sub {
    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"
    s, _ := sub.Get(subscriptionId, nil)
    return s;
}

func DeleteUserSubscription(userId string, subscriptionId string) {
    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

    sub, err := sub.Cancel(
        subscriptionId,
        &stripe.SubParams{Customer: userId},
    )

    log.Print(err)
    log.Print(sub)
}
