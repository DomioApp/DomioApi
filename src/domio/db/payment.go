package domiodb

import (
	_ "github.com/lib/pq"
	domioerrors "domio/errors"
	"github.com/lib/pq"
	"log"
)

type PaymentRequest struct {
	RentalId int `json:"rental_id"`
}

type Payment struct {
	MonthsCount   int `json:"months_count" db:"months_count"`
	AmountInCents int `json:"amount_in_cents" db:"amount_in_cents"`
	PaidBy        string `json:"paid_by" db:"paid_by"`
	Id            int `json:"id" db:"id"`
}

func CreatePayment(paymentRequest PaymentRequest, ownerEmail string, stripeId string) (Payment, *domioerrors.DomioError) {
	var paymentResult Payment
	log.Print("******************************************")
	log.Print(paymentRequest)
	log.Print("******************************************")


	insertErr := Db.QueryRowx("INSERT INTO payments (amount_in_cents, paid_by, stripe_id) VALUES ($1, $2, $3) RETURNING id, amount_in_cents, paid_by", 300, ownerEmail, stripeId).
		StructScan(&paymentResult)

	var pqError *pq.Error

	if (insertErr != nil) {
		pqError = insertErr.(*pq.Error)
	}

	if (pqError != nil) {
		if (pqError.Code.Name() == "foreign_key_violation") {
			return paymentResult, &domioerrors.RentableDomainNotExist
		}
		log.Print(pqError.Code.Name())
		log.Print(pqError.Code.Class())
		return paymentResult, &domioerrors.UnknownError
	}

	return paymentResult, nil
}
