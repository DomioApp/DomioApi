package r53

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/route53"
    "fmt"
    "domio_api/db"
)

var useRealR53 = true

func DeleteDomainZone(domain *domiodb.Domain) error {

    if (useRealR53) {
        return DeleteDomainZoneReal(domain)
    } else {
        return DeleteDomainZoneMock(domain)
    }
}
func CreateDomainZone(domain *domiodb.Domain) (*route53.CreateHostedZoneOutput, error) {

    if (useRealR53) {
        return CreateDomainZoneReal(domain)
    } else {
        return CreateDomainZoneMock(domain)
    }
}

func GetHostedZone(domain *domiodb.Domain) interface{} {
    svc, _ := GetAwsService();

    params := &route53.GetHostedZoneInput{
        Id: aws.String(domain.ZoneId.String), // Required
    }
    resp, err := svc.GetHostedZone(params)

    if err != nil {
        fmt.Println(err)
        return nil
    }

    return resp
}