package r53

import (
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws"
    "log"
    "github.com/fatih/color"
)

func DeleteRecord(zoneId string, domainName string, key string, value string) (*route53.ChangeResourceRecordSetsOutput, error) {

    svc, _ := GetAwsService();

    recordSet := &route53.ResourceRecordSet{// Required
        Name: aws.String(domainName), // Required
        Type: aws.String(key), // Required,
        ResourceRecords: []*route53.ResourceRecord{
            {// Required
                Value: aws.String(value), // Required
            },
        },
        TTL:aws.Int64(3600),
    }

    params := &route53.ChangeResourceRecordSetsInput{
        HostedZoneId: aws.String(zoneId), // Required

        ChangeBatch: &route53.ChangeBatch{// Required
            Comment: aws.String("Sample delete."),
            Changes: []*route53.Change{// Required
                {// Required
                    Action: aws.String("DELETE"), // Required
                    ResourceRecordSet: recordSet,
                },
            },
        },
    }

    resp, err := svc.ChangeResourceRecordSets(params)

    if (err != nil) {
        color.Set(color.FgYellow)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        log.Print(err)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        color.Unset()
    }

    return resp, err
}