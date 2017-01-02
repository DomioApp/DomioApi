package domiodb

import (
    domioerrors  "domio/errors"
    "github.com/lib/pq"
    "log"
    "github.com/aws/aws-sdk-go/aws/session"
    "fmt"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/fatih/color"
    "time"
)

type Domain struct {
    Name          string `json:"name" db:"name"`
    Owner         string `json:"owner" db:"owner"`
    PricePerMonth uint64 `json:"price_per_month" db:"price_per_month"`
    IsRented      bool `json:"is_rented" db:"is_rented"`
    ZoneId        string `json:"zone_id" db:"zone_id"`
}

func GetAvailableDomains() []Domain {
    var domains []Domain
    Db.Select(&domains, "SELECT name, price_per_month FROM domains WHERE is_rented=false ORDER BY price_per_month")
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
    Db.MustExec("UPDATE domains SET is_rented=true WHERE name=$1", domainName)
}

func SetDomainZoneId(domain *Domain, id *string) {
    Db.MustExec("UPDATE domains SET zone_id=$1 WHERE name=$2", id, domain.Name)
}

func CreateDomain(domain Domain, ownerEmail string) (Domain, *pq.Error) {
    var domainResult Domain
    insertErr := Db.QueryRowx("INSERT INTO domains (name, price_per_month, owner) VALUES ($1, $2, $3) RETURNING name, price_per_month, owner", domain.Name, domain.PricePerMonth, ownerEmail).StructScan(&domainResult)

    if (insertErr != nil) {
        return domainResult, insertErr.(*pq.Error)
    }

    createDomainZone(&domainResult)

    return domainResult, nil
}

func createDomainZone(domain *Domain) {
    log.Print(domain)
    sess, err := session.NewSession()
    if err != nil {
        fmt.Println("failed to create session,", err)
        return
    }

    svc := route53.New(sess)
    id := time.Now().Format(time.RFC850);

    params := &route53.CreateHostedZoneInput{
        CallerReference: &id,
        Name:            aws.String(domain.Name),
    }
    resp, err := svc.CreateHostedZone(params)

    if err != nil {
        color.Set(color.FgRed)
        log.Println(params.CallerReference)
        log.Println(id)
        log.Println(err)
        color.Unset()
        return
    }
    SetDomainZoneId(domain, resp.HostedZone.Id)
    log.Println(resp)
}

func GetHostedZone(domain *Domain) {
    sess, err := session.NewSession()
    if err != nil {
        fmt.Println("failed to create session,", err)
        return
    }

    svc := route53.New(sess)

    params := &route53.GetHostedZoneInput{
        Id: aws.String(domain.ZoneId), // Required
    }
    resp, err := svc.GetHostedZone(params)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(resp)
}