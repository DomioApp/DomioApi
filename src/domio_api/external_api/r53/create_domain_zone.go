package r53

import (
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/fatih/color"
    "time"
    "log"
)

func CreateDomainZone(domain *domiodb.Domain) (*route53.CreateHostedZoneOutput, error) {
    svc, _ := GetAwsService();
    id := time.Now().Format(time.RFC850);

    params := &route53.CreateHostedZoneInput{
        CallerReference: &id,
        Name:            aws.String(domain.Name),
    }

    resp, err := svc.CreateHostedZone(params)

    if err != nil {
        color.Set(color.FgRed)
        log.Println(params.CallerReference)
        log.Println(id)
        log.Println(err)
        color.Unset()
        return nil, err
    }

    return resp, nil

}
