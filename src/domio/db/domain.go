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
    "domio/components/tokens"
    "database/sql"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "domio/components/config"
)

type Domain struct {
    Name          string `json:"name" db:"name"`
    Owner         string `json:"owner" db:"owner"`
    PricePerMonth uint64 `json:"price_per_month" db:"price_per_month"`
    IsRented      bool `json:"is_rented" db:"is_rented"`
    RentedBy      sql.NullString `json:"rented_by" db:"rented_by"`
    ZoneId        string `json:"zone_id" db:"zone_id"`
    NS1           sql.NullString `json:"ns1" db:"ns1"`
    NS2           sql.NullString `json:"ns2" db:"ns2"`
    NS3           sql.NullString `json:"ns3" db:"ns3"`
    NS4           sql.NullString `json:"ns4" db:"ns4"`
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

func DeleteDomain(domainName string, ownerEmail string) domioerrors.DomioError {
    domain, domainError := GetDomain(domainName)
    if (domainError != nil) {
        log.Print(domainError)
        return domioerrors.DomainNotFound
    }

    result := Db.MustExec("DELETE FROM domains where name=$1 AND owner=$2 AND is_rented=false", domainName, ownerEmail)

    rowsAffected, err := result.RowsAffected()
    if (err != nil) {
        color.Set(color.FgHiRed)
        log.Print(err)
        color.Unset()
        return domioerrors.DomainNotFound
    }

    color.Set(color.FgHiCyan)
    log.Print(rowsAffected)
    color.Unset()

    if (rowsAffected == 0) {
        return domioerrors.DomainNotFound
    }

    deleteDomainZone(&domain)

    log.Print("Domain removed from local database")
    return domioerrors.DomioError{}
}

func SetDomainAsRented(domainName string, userProfile *tokens.UserTokenWithClaims) {
    Db.MustExec("UPDATE domains SET is_rented=true, rented_by=$1 WHERE name=$2", userProfile.Email, domainName)
}

func SetDomainAsAvailable(domainName string, userProfile *tokens.UserTokenWithClaims) {
    Db.MustExec("UPDATE domains SET is_rented=false, rented_by=NULL, ns1=NULL, ns2=NULL, ns3=NUL, ns4=NULL WHERE rented_by=$1 AND name=$2", userProfile.Email, domainName)
}

func SetDomainZoneId(domain *Domain, id *string) {
    Db.MustExec("UPDATE domains SET zone_id=$1 WHERE name=$2", id, domain.Name)
}

func SetDomainNameServers(domain *Domain, ns1 *string, ns2 *string, ns3 *string, ns4 *string) {
    Db.MustExec("UPDATE domains SET ns1=$2, ns2=$3, ns3=$4, ns4=$5 WHERE name=$1", domain.Name, ns1, ns2, ns3, ns4, )
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
    conf := config.LoadConfig()
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})
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
    SetDomainNameServers(domain, resp.DelegationSet.NameServers[0], resp.DelegationSet.NameServers[1], resp.DelegationSet.NameServers[2], resp.DelegationSet.NameServers[3])

    log.Println(resp)
}
func deleteDomainZone(domain *Domain) {
    conf := config.LoadConfig()
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

    if err != nil {
        fmt.Println("failed to create session,", err)
        return
    }

    svc := route53.New(sess)

    params := &route53.DeleteHostedZoneInput{
        Id: &domain.ZoneId,
    }
    //resp, err := svc.DeleteHostedZone(params)
    resp, err := svc.DeleteHostedZone(params)

    if err != nil {
        color.Set(color.FgRed)
        log.Println(err)
        color.Unset()
        return
    }
    log.Println(resp)
    log.Print("Domain zone removed from Route 53")
}

func GetHostedZone(domain *Domain) {
    conf := config.LoadConfig()
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

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