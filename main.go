package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest1RegionID),
	}))

	svc := ec2.New(sess)

	params := &ec2.CopyImageInput{
		Name:          aws.String("TestAMI2"),     // Required
		SourceImageId: aws.String("ami-41e7a357"), // Required
		SourceRegion:  aws.String("us-east-1"),    // Required
		//ClientToken:   aws.String("String"),
		Description: aws.String("Copy of TestAMI"),
		DryRun:      aws.Bool(false),
		Encrypted:   aws.Bool(false),
		//KmsKeyId:      aws.String("String"),
	}
	resp, err := svc.CopyImage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

	/*
		srcAmi := flag.String("srcAmi", "", "Source AMI")
		srcRegion := flag.String("srcRegion", "", "Source Region")
		flag.Parse()
		fmt.Printf("srcAmi: %s, srcRegion: %s\n", *srcAmi, *srcRegion)

			var srcAmi, srcRegion string
			var err error
			fmt.Print("Enter Source AMI:  ")
			_, err = fmt.Scanln(&srcAmi)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Print("Enter source region to copy from:  ")
			_, err = fmt.Scanln(&srcRegion)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("Copying", srcAmi, "in region", srcRegion)
	*/
}
