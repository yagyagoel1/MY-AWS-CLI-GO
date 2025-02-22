package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

func ListEC2Instances() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load sdk config , %v", err)
	}
	client := ec2.NewFromConfig(cfg)
	resp, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatalf("falied to descirbe the instances %v", err)
	}
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("Instance ID: %s, Type: %s, State: %s\n",
				*instance.InstanceId, instance.InstanceType, instance.State.Name)
		}
	}
}

var EC2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "List EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		ListEC2Instances()
	},
}
