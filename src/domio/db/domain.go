package domiodb

import (
	_ "github.com/lib/pq"
	"database/sql"
	"errors"
	domioerrors  "domio/errors"
	"github.com/lib/pq"
	"log"
)

type Domain struct {
	Name  string `json:"name" db:"name"`
	Owner string `json:"owner" db:"owner"`
	Plan string `json:"plan" db:"plan"`
}

type AvailableDomain struct {
	Name  string `json:"name" db:"name"`
	Price int `json:"price" db:"price"`
}

type DomainWithRental struct {
	Name        string `json:"name"`
	PeriodRange string `json:"period_range"`
}

type DomainWithoutRental struct {
	Domain
}

var ErrDomainNotFound = errors.New("Domain is is not found in the database.")

func GetAvailableDomains() []AvailableDomain {
	domains := make([]AvailableDomain, 0)
	err := Db.Select(&domains, `SELECT name, price
                             FROM public.domains domain
                             LEFT JOIN rentals rental
                             ON domain.name = rental.domain_name
                             WHERE rental.domain_name IS NULL
                             OR NOW()::date - upper(rental.period_range::daterange) > 0
                         ORDER BY price`)

	if (err != nil) {
		log.Fatalln(err)
	}

	return domains
}

func GetUserDomains(userEmail string) []Domain {
	var domains []Domain
	Db.Select(&domains, "SELECT * FROM domains WHERE owner=$1 ORDER BY price", userEmail)
	return domains
}

func GetDomainInfo(domainName string) (interface{}, *domioerrors.DomioError) {
	var domain Domain
	var domainRental Rental
	var domainWithRentalInfo DomainWithRental

	domainError := Db.QueryRowx("SELECT * FROM domains where name=$1", domainName).StructScan(&domain)
	if domainError != nil {
		return Domain{}, &domioerrors.DomainNotFound
	}

	domainRentalError := Db.QueryRowx(`SELECT domain_name, period_range
                                            FROM rentals WHERE domain=$1
                                            AND NOW()::date - upper(period_range::daterange) <= 0
                                            LIMIT 1`, domain.Name).StructScan(&domainRental)

	if domainRentalError != nil && domainRentalError == sql.ErrNoRows {
		return domain, nil
	}
	domainWithRentalInfo = DomainWithRental{Name:domain.Name, PeriodRange:domainRental.PeriodRange}

	return domainWithRentalInfo, nil
}

func IsDomainRented(domainName string) bool {
	const RentedDomainQuery = `SELECT domain_name, period_range
                                        FROM rentals WHERE domain_name=$1
                                        AND NOW()::date - upper(period_range::daterange) <= 0
                                        LIMIT 1`
	var domainRental Rental

	domainIsNotRentedError := Db.QueryRowx(RentedDomainQuery, domainName).StructScan(&domainRental)
	log.Print(domainIsNotRentedError)

	if domainIsNotRentedError != nil && domainIsNotRentedError == sql.ErrNoRows {
		return false
	}

	return true
}

func GetDomainOwner(domainName string) (string, error) {
	const DomainOwnerQuery = `SELECT * FROM domains where name=$1 LIMIT 1`
	var domainInfo Domain

	domainQueryError := Db.QueryRowx(DomainOwnerQuery, domainName).StructScan(&domainInfo)

	if domainQueryError != nil && domainQueryError == sql.ErrNoRows {
		return "", ErrDomainNotFound
	}
	return domainInfo.Owner, nil
}

func CreateDomain(domain Domain, ownerEmail string) (Domain, *pq.Error) {
	var domainResult Domain
	insertErr := Db.QueryRowx("INSERT INTO domains (name, plan, owner) VALUES ($1, $2, $3) RETURNING name, plan, owner", domain.Name, domain.Plan, ownerEmail).StructScan(&domainResult)

	if (insertErr != nil) {
		return domainResult, insertErr.(*pq.Error)
	}

	return domainResult, nil
}
