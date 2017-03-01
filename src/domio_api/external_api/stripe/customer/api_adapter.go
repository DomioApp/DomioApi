package stripe_user_adapter

import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
    "domio_api/components/config"
)

func DeleteStripeCustomer(stripe_customer_id string) (*stripe.Customer, error) {

    stripe.Key = config.Config.STRIPE_KEY

    c, err := customer.Del(stripe_customer_id)
    return c, err
}