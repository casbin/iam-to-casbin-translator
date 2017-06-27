package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func WriteFile(path string, text string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	w.WriteString(text)
	w.Flush()
	f.Close()
	return nil
}

func WritePolicyToFile(policy *iam.Policy) {
	document := DecodePolicy(GetPolicyDocument(policy))
	fmt.Print(document)
	path := "policy/" + *policy.PolicyName + ".json"
	WriteFile(path, document)
}

func DecodePolicy(text string) string {
	text, _ = url.QueryUnescape(text)
	return text
}

func ListPolicies() []*iam.Policy {
	sess := session.Must(session.NewSession())

	svc := iam.New(sess)

	params := &iam.ListPoliciesInput{
		MaxItems:     aws.Int64(1000),
		OnlyAttached: aws.Bool(false),
		Scope:        aws.String("All"),
	}
	resp, err := svc.ListPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	// Pretty-print the response data.
	fmt.Println(resp)
	return resp.Policies
}

func GetPolicyDocument(policy *iam.Policy) string {
	sess := session.Must(session.NewSession())

	svc := iam.New(sess)

	params := &iam.GetPolicyVersionInput{
		PolicyArn: aws.String(*policy.Arn),             // Required
		VersionId: aws.String(*policy.DefaultVersionId), // Required
	}
	resp, err := svc.GetPolicyVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return ""
	}

	// Pretty-print the response data.
	//fmt.Println(resp)
	return *resp.PolicyVersion.Document
}

func main() {
	// policies := ListPolicies()
}