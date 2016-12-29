package stripe_adapter

import "github.com/stripe/stripe-go"

func MakePayment() (*stripe.Charge, error) {
	stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"
	/*
	params := stripe.CustomerParams{
		Email:  "george.costanza@mail.com",
		Desc:   "short, bald",
		Card:   &stripe.CardParams{
			Name     : "George Costanza",
			Number   : "4242424242424242",
			ExpYear  : 2012,
			ExpMonth : 5,
			CVC      : "26726",
		},
	}
	customer, err := stripe.Customers.Create(&params)
	*/

	/*
	params := stripe.ChargeParams{
		Desc:     "Calzone",
		Amount:   400,
		Currency: "usd",
		Card:     &stripe.CardParams{
			Name     : "George Costanza",
			Number   : "4242424242424242",
			ExpYear  : 2017,
			ExpMonth : 5,
			CVC      : "197",
		},
	}
	*/

	/*
	charge, err := stripe.Charges.Create(&params)
	return charge, err
	*/
	return nil, nil
}
