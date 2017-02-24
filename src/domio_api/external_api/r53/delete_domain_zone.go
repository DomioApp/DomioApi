package r53

import (
    "domio_api/db"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/fatih/color"
    "log"
)

func DeleteDomainZone(domain *domiodb.Domain) error {

    svc, _ := GetAwsService();

    params := &route53.DeleteHostedZoneInput{
        Id: &domain.ZoneId.String,
    }

    resp, err := svc.DeleteHostedZone(params)
    if err != nil {
        color.Set(color.FgRed)
        log.Println(err)
        color.Unset()
        return err
    }
    log.Println(resp)
    log.Print("Domain zone removed from Route 53")
    return nil
}
