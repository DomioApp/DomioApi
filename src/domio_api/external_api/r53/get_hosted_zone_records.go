package r53

import (
    "fmt"
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
    "domio_api/components/config"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
)

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