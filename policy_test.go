package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func TestGetPolicy(t *testing.T) {
	policy := &iam.Policy{
		PolicyName: aws.String("AWSAccountUsageReportAccess"),
		Arn: aws.String("arn:aws:iam::aws:policy/AWSAccountUsageReportAccess"),
		DefaultVersionId: aws.String("v1"),
	}

	WritePolicyToFile(policy)
}