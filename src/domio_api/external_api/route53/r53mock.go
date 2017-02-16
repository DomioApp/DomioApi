package route53

import (
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
)

func DeleteDomainZoneMock(domain *domiodb.Domain) error {
    return nil
}

func CreateDomainZoneMock(domain *domiodb.Domain) (*route53.CreateHostedZoneOutput, error) {
    hostedZoneId := "dummy_hosted_zone";
    ns1 := "ns1";
    ns2 := "ns2";
    ns3 := "ns3";
    ns4 := "ns4";

    delSet := route53.DelegationSet{
        NameServers:[]*string{
            &ns1,
            &ns2,
            &ns3,
            &ns4,
        }}

    hostedZone := route53.CreateHostedZoneOutput{
        HostedZone:&route53.HostedZone{Id:&hostedZoneId},
        DelegationSet:&delSet,
    }

    return &hostedZone, nil

}