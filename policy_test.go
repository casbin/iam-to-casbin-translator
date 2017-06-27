package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

func TestGetPolicy(t *testing.T) {
	policy := &iam.Policy{
		Arn: aws.String("arn:aws:iam::aws:policy/AWSAccountUsageReportAccess"),
		DefaultVersionId: aws.String("v1"),
	}

	document := DecodePolicy(GetPolicyDocument(policy))
	fmt.Print(document)
}