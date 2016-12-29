package domiodb

import (
	_ "github.com/lib/pq"
	domioerrors  "domio/errors"
	"github.com/lib/pq"
	"log"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go"
)

type Rental struct {
	Id          int `json:"id" db:"id"`
	Domain      string `json:"domain_name" db:"domain_name"`
	PeriodRange string `json:"period_range" db:"period_range"`
	Payment     int `json:"payment" db:"payment"`
}

func CreateRental(rentalRequest Rental, ownerEmail string) (Rental, *domioerrors.DomioError) {
	var rentalResult Rental

	domainOwner, _ := GetDomainOwner(rentalRequest.Domain)

	if domainOwner == ownerEmail {
		return rentalResult, &domioerrors.DomainIsOwnedByUser
	}

	isDomainRented := IsDomainRented(rentalRequest.Domain)

	if (isDomainRented == true) {
		return rentalResult, &domioerrors.DomainIsAlreadyRented
	}

	log.Print("************************************************")
	log.Print(rentalRequest)
	insertErr := Db.QueryRowx("INSERT INTO rentals (domain_name, renter_id) VALUES ($1, $2) RETURNING domain_name, id", rentalRequest.Domain, ownerEmail).
		StructScan(&rentalResult)

	var pqError *pq.Error

	if (insertErr != nil) {
		pqError = insertErr.(*pq.Error)
	}

	if (pqError != nil) {
		log.Print(pqError)
		if (pqError.Code.Name() == "foreign_key_violation") {
			log.Print(pqError)
			return rentalResult, &domioerrors.RentableDomainNotExist
		}
		log.Print(pqError.Code.Name())
		return rentalResult, &domioerrors.UnknownError

	}
	makePayment()
	return rentalResult, nil
}

func makePayment() {
	stripe.Key = "sk_test_BQokikJOvBiI2HlWgH4olfQ2"

	chargeParams := &stripe.ChargeParams{
		Amount: 2000,
		Currency: "usd",
		Desc: "Charge for jacob.davis@example.com",
	}
	chargeParams.SetSource("tok_189fl92eZvKYlo2C0sjTBkKA")
	ch, err := charge.New(chargeParams)
	log.Print(ch)
	log.Print(err)
}