package domiodb

import (
    domioerrors  "domio/errors"
    "github.com/lib/pq"
    "log"
)

type Domain struct {
    Name          string `json:"name" db:"name"`
    Owner         string `json:"owner" db:"owner"`
    PricePerMonth uint64 `json:"price_per_month" db:"price_per_month"`
    IsRented      bool `json:"is_rented" db:"is_rented"`
}

type AvailableDomain struct {
    Name  string `json:"name" db:"name"`
    PricePerMonth int `json:"price_per_month" db:"price_per_month"`
}

type DomainWithRental struct {
    Name        string `json:"name"`
    PeriodRange string `json:"period_range"`
}

type DomainWithoutRental struct {
    Domain
}

func GetAvailableDomains() []AvailableDomain {
    domains := make([]AvailableDomain, 0)
    err := Db.Select(&domains, `SELECT name, price_per_month FROM public.domains domain WHERE is_rented=false ORDER BY price_per_month`)

    if (err != nil) {
        log.Fatalln(err)
    }

    return domains
}

func GetUserDomains(userEmail string) []Domain {
    var domains []Domain
    Db.Select(&domains, "SELECT * FROM domains WHERE owner=$1 ORDER BY price_per_month", userEmail)
    return domains
}

func GetDomain(domainName string) (Domain, *domioerrors.DomioError) {
    var domain Domain

    domainError := Db.QueryRowx("SELECT * FROM domains where name=$1", domainName).StructScan(&domain)

    if domainError != nil {
        log.Print(domainError)
        return Domain{}, &domioerrors.DomainNotFound
    }

    return domain, nil
}

func SetDomainAsRented(domainName string) {
    domainError := Db.MustExec("UPDATE public.domains SET is_rented=true WHERE name=$1", domainName)
    log.Print(domainError)
}

func CreateDomain(domain Domain, ownerEmail string) (Domain, *pq.Error) {
    var domainResult Domain
    insertErr := Db.QueryRowx("INSERT INTO domains (name, price_per_month, owner) VALUES ($1, $2, $3) RETURNING name, price_per_month, owner", domain.Name, domain.PricePerMonth, ownerEmail).StructScan(&domainResult)

    if (insertErr != nil) {
        return domainResult, insertErr.(*pq.Error)
    }

    return domainResult, nil
}
