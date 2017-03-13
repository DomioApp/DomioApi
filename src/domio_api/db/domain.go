package domiodb

import (
    "github.com/lib/pq"
    "log"
    "github.com/fatih/color"
    "domio_api/components/tokens"
    "database/sql"
    "strings"
    "domio_api/errors"
)

type Domain struct {
    Name           string `json:"name" db:"name"`
    Owner          string `json:"owner" db:"owner"`
    PricePerMonth  uint64 `json:"price_per_month" db:"price_per_month"`
    IsApproved     bool `json:"is_approved" db:"is_approved"`
    IsVisible      bool `json:"is_visible" db:"is_visible"`
    IsRented       bool `json:"is_rented" db:"is_rented"`
    RentedBy       sql.NullString `json:"rented_by" db:"rented_by"`
    SubscriptionId sql.NullString `json:"subscription_id" db:"subscription_id"`
    ZoneId         sql.NullString `json:"zone_id" db:"zone_id"`
    NS1            sql.NullString `json:"ns1" db:"ns1"`
    NS2            sql.NullString `json:"ns2" db:"ns2"`
    NS3            sql.NullString `json:"ns3" db:"ns3"`
    NS4            sql.NullString `json:"ns4" db:"ns4"`
}

type AvailableDomainJson struct {
    Name          string `json:"name" db:"name"`
    PricePerMonth uint64 `json:"price_per_month" db:"price_per_month"`
}

type DomainToEdit struct {
    IsVisible     bool `json:"is_visible" db:"is_visible"`
    PricePerMonth uint64 `json:"price_per_month" db:"price_per_month"`
}

type DomainJson struct {
    Name          string `json:"name" db:"name"`
    Owner         string `json:"owner" db:"owner"`
    PricePerMonth uint64 `json:"price_per_month" db:"price_per_month"`
    IsApproved     bool `json:"is_approved" db:"is_approved"`
    IsRented      bool   `json:"is_rented" db:"is_rented"`
    IsVisible     bool `json:"is_visible" db:"is_visible"`
    RentedBy      string `json:"rented_by,omitempty" db:"rented_by"`
    ZoneId        string `json:"zone_id,omitempty" db:"zone_id"`
    NS1           string `json:"ns1" db:"ns1"`
    NS2           string `json:"ns2" db:"ns2"`
    NS3           string `json:"ns3" db:"ns3"`
    NS4           string `json:"ns4" db:"ns4"`
}

func GetAvailableDomains() []AvailableDomainJson {
    var domains = make([]AvailableDomainJson, 0)
    Db.Select(&domains, "SELECT name, price_per_month FROM domains WHERE is_rented=false AND is_visible=true ORDER BY price_per_month")

    return domains
}

func GetPendingDomains() []DomainJson {
    var domains = make([]DomainJson, 0)
    err := Db.Select(&domains, "SELECT name, price_per_month FROM domains WHERE is_approved!=$1", true)

    if (err != nil) {
        log.Print(err)
    }

    return domains
}

func GetUserDomains(userEmail string) []Domain {
    var domains []Domain = make([]Domain, 0)

    selectErr := Db.Select(&domains, "SELECT * FROM domains WHERE owner=$1 ORDER BY price_per_month", userEmail)
    if (selectErr != nil) {
        log.Print(selectErr)
    }
    return domains
}

func GetUserDomainsCount(userEmail string) int {
    domains := GetUserDomains(userEmail);
    return len(domains)
}

func GetDomainInfo(domainName string) (*Domain, *domioerrors.DomioError) {
    var domain Domain

    domainError := Db.QueryRowx("SELECT * FROM domains where name=$1", domainName).StructScan(&domain)

    if domainError != nil {

        color.Set(color.FgRed)
        log.Print(domainError)
        color.Unset()

        return nil, &domioerrors.DomainNotFound
    }

    return &domain, nil
}

func GetDomainInfoBySubscriptionId(subId string) (Domain, *domioerrors.DomioError) {
    var domain Domain

    domainError := Db.QueryRowx("SELECT * FROM domains where subscription_id=$1", subId).StructScan(&domain)

    if domainError != nil {
        log.Print(domainError)
        return Domain{}, &domioerrors.DomainNotFound
    }

    return domain, nil
}

func DeleteDomain(domainName string, ownerEmail string) (*Domain, *domioerrors.DomioError) {
    domain, domainError := GetDomainInfo(domainName)

    if (domainError != nil) {
        color.Set(color.FgHiRed)
        log.Print(domainError)
        color.Unset()

        return nil, &domioerrors.DomainNotFound
    }

    if (domain.IsRented) {
        return nil, &domioerrors.DomainInRent
    }

    result := Db.MustExec("DELETE FROM domains where name=$1 AND owner=$2 AND is_rented=false", domainName, ownerEmail)

    rowsAffected, err := result.RowsAffected()
    if (err != nil) {
        return nil, &domioerrors.DomainNotFound
    }

    color.Set(color.FgHiCyan)
    log.Print(rowsAffected)
    color.Unset()

    if (rowsAffected == 0) {
        return nil, &domioerrors.DomainNotFound
    }

    log.Print("Domain removed from local database")
    return domain, nil
}

func SetDomainAsRented(domainName string, subId string, userProfile *tokens.UserTokenWithClaims) {
    Db.MustExec("UPDATE domains SET is_rented=true, rented_by=$2, subscription_id=$3 WHERE name=$1", domainName, userProfile.Email, subId)
}

func SetDomainAsAvailable(domainName string, userProfile *tokens.UserTokenWithClaims) {
    Db.MustExec("UPDATE domains SET is_rented=false, rented_by=$3, subscription_id=$3 WHERE rented_by=$1 AND name=$2", userProfile.Email, domainName, nil)
}

func SetDomainZoneId(domain *Domain, hostedZoneId *string) {
    Db.MustExec("UPDATE domains SET zone_id=$2 WHERE name=$1", domain.Name, hostedZoneId)
}

func SetDomainNameServers(domain *Domain, ns1 *string, ns2 *string, ns3 *string, ns4 *string) {
    Db.MustExec("UPDATE domains SET ns1=$2, ns2=$3, ns3=$4, ns4=$5 WHERE name=$1", domain.Name, ns1, ns2, ns3, ns4, )
}

func CreateDomain(domain Domain, ownerEmail string) (*Domain, *pq.Error) {
    var domainResultDb Domain

    insertErr := Db.QueryRowx("INSERT INTO domains (name, price_per_month, owner) VALUES ($1, $2, $3) RETURNING name, price_per_month, owner",
        strings.ToLower(domain.Name), domain.PricePerMonth, ownerEmail).
        StructScan(&domainResultDb)

    if (insertErr != nil) {
        log.Println(insertErr)
        return nil, insertErr.(*pq.Error)
    }

    domainInfo, getDomainError := GetDomainInfo(domainResultDb.Name)

    if (getDomainError != nil) {
        log.Println(getDomainError)
    }

    return domainInfo, nil
}

func UpdateDomain(domainName string, domainToEdit DomainToEdit) error {
    log.Print(domainToEdit.IsVisible)

    _, err := Db.Exec("UPDATE domains SET price_per_month=$2, is_visible=$3 WHERE name=$1", domainName, domainToEdit.PricePerMonth, domainToEdit.IsVisible)

    if (err != nil) {
        log.Print(err)
    }

    return nil
}
