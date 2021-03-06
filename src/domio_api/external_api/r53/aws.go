package r53

import (
    "domio_api/components/config"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/route53"
    "github.com/aws/aws-sdk-go/aws/session"
    "fmt"
)

func GetAwsService() (*route53.Route53, error) {
    creds := credentials.NewStaticCredentials(config.Config.AWS_ACCESS_KEY_ID, config.Config.AWS_SECRET_ACCESS_KEY, "")
    sess, err := session.NewSession(&aws.Config{Credentials: creds})

    if err != nil {
        fmt.Println("failed to create session,", err)
        return nil, err

    }
    svc := route53.New(sess)

    return svc, err
}