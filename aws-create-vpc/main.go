package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	// Load the Shared AWS Configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an EC2 client
	svc := ec2.NewFromConfig(cfg)

	// Create the VPC
	createVpcInput := &ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeVpc,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("MyVPC"),
					},
				},
			},
		},
	}

	createVpcOutput, err := svc.CreateVpc(context.TODO(), createVpcInput)
	if err != nil {
		log.Fatalf("failed to create VPC, %v", err)
	}

	fmt.Printf("Successfully created VPC with ID %s\n", *createVpcOutput.Vpc.VpcId)
}
