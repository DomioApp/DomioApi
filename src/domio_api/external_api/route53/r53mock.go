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
    resp := route53.CreateHostedZoneOutput{
        HostedZone:route53.HostedZone{Name:string("dummy_hosted_zone")},
    }
    resp.DelegationSet.NameServers = &[]route53domains.Nameserver{
        *route53domains.Nameserver{Name:"1"},
        *route53domains.Nameserver{Name:"2"},
        *route53domains.Nameserver{Name:"3"},
        *route53domains.Nameserver{Name:"4"},
    }

    return resp, nil

}