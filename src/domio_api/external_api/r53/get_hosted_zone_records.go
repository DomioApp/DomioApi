package r53

import (
    "fmt"
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws"
)

func GetHostedZoneRecords(domain *domiodb.Domain) *route53.ListResourceRecordSetsOutput {
    svc, _ := GetAwsService();

    listParams := &route53.ListResourceRecordSetsInput{
        HostedZoneId: aws.String(domain.ZoneId.String), // Required
    }
    resp, err := svc.ListResourceRecordSets(listParams)

    if err != nil {
        fmt.Println(err)
        return nil
    }

    return resp
}