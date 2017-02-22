package r53

import (
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws"
    "domio_api/components/config"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "fmt"
    "log"
    "strings"
    "github.com/fatih/color"
)

func UpdateRecord(zoneId string, domainName string, key string, value string, TTL int64, weight int64) (*route53.ChangeResourceRecordSetsOutput, error) {

    conf := config.Config
    token := ""
    creds := credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, token)
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

    if err != nil {
        fmt.Println("failed to create session,", err)
        return nil, err
    }

    svc := route53.New(sess)

    recordSet := &route53.ResourceRecordSet{// Required
        Name: aws.String(domainName), // Required
        Type: aws.String(strings.ToUpper(key)), // Required
        ResourceRecords: []*route53.ResourceRecord{
            {// Required
                Value: aws.String(value), // Required
            },
        },
        TTL:            aws.Int64(TTL),
        Weight:         aws.Int64(weight),
        SetIdentifier:  aws.String("Arbitrary Id describing this change set"),
    }

    params := &route53.ChangeResourceRecordSetsInput{
        HostedZoneId: aws.String(zoneId), // Required

        ChangeBatch: &route53.ChangeBatch{// Required
            Comment: aws.String("Sample update."),
            Changes: []*route53.Change{// Required
                {// Required
                    Action: aws.String("UPSERT"), // Required
                    ResourceRecordSet: recordSet,
                },
            },
        },
    }

    resp, err := svc.ChangeResourceRecordSets(params)

    if (err != nil) {
        color.Set(color.FgRed)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        log.Print(err)
        log.Print("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
        color.Unset()
    }

    return resp, err
}