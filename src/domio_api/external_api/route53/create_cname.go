package route53

import (
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws"
)

func CreateCNAME(svc *route53.Route53, zoneId string, name string, target string, TTL int64, weight int64) (*route53.ChangeResourceRecordSetsOutput, error) {
    recordSet := &route53.ResourceRecordSet{// Required
        Name: aws.String(name), // Required
        Type: aws.String("CNAME"), // Required
        ResourceRecords: []*route53.ResourceRecord{
            {// Required
                Value: aws.String(target), // Required
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

    return resp, err
}