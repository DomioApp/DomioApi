package route53

import (
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/service/route53domains"
)

func DeleteDomainZoneMock(domain *domiodb.Domain) error {
    return nil
}

func CreateDomainZoneMock(domain *domiodb.Domain) (*route53.CreateHostedZoneOutput, error) {
    resp := route53.CreateHostedZoneOutput{}

    resp.HostedZone.Id = "hosted_zone_id"
    resp.DelegationSet.NameServers = &[]route53domains.Nameserver{
        *route53domains.Nameserver{Name:"1"},
        *route53domains.Nameserver{Name:"2"},
        *route53domains.Nameserver{Name:"3"},
        *route53domains.Nameserver{Name:"4"},
    }

    return resp, nil

}