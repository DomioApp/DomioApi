package r53

import (
    "domio_api/components/config"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
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
    conf := config.Config
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

    if err != nil {
        fmt.Println("failed to create session,", err)
        return nil
    }

    svc := route53.New(sess)

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
func GetHostedZoneRecords(domain *domiodb.Domain) *route53.ListResourceRecordSetsOutput {
    conf := config.Config
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

    if err != nil {
        fmt.Println("failed to create session,", err)
        return nil
    }

    svc := route53.New(sess)

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