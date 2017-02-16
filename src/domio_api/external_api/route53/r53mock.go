package route53

import (
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
)

func DeleteDomainZoneMock(domain *domiodb.Domain) error {
    return nil
}

func CreateDomainZoneMock(domain *domiodb.Domain) (*route53.CreateHostedZoneOutput, error) {
    hostedZoneName := "dummy_hosted_zone";
    ns1 := "dummy_hosted_zone";
    ns2 := "dummy_hosted_zone";
    ns3 := "dummy_hosted_zone";
    ns4 := "dummy_hosted_zone";

    delSet := route53.DelegationSet{
        NameServers:[]*string{
            &ns1,
            &ns2,
            &ns3,
            &ns4,
        }}

    hostedZone := route53.CreateHostedZoneOutput{
        HostedZone:&route53.HostedZone{Name:&hostedZoneName},
        DelegationSet:&delSet,
    }

    return &hostedZone, nil

}