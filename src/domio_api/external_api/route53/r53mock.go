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
    a := "dummy_hosted_zone";
    ns1 := "dummy_hosted_zone";
    ns2 := "dummy_hosted_zone";
    ns3 := "dummy_hosted_zone";
    ns4 := "dummy_hosted_zone";

    resp := route53.CreateHostedZoneOutput{
        HostedZone:&route53.HostedZone{Name:&a},
    }

    var delSet=
    resp.DelegationSet.NameServers = []route53domains.Nameserver{
        {Name:&ns1},
        {Name:&ns2},
        {Name:&ns3},
        {Name:&ns4},
    }

    return &resp, nil

}