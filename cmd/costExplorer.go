package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/spf13/cobra"
)

func getAllCostings() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load Sdk config %v", err)
	}
	client := costexplorer.NewFromConfig(cfg)

	end := time.Now().Format("2006-01-02")
	start := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: &start,
			End:   &end,
		},
		Granularity: "DAILY",
		Metrics:     []string{"BlendedCost"},
	}
	resp, err := client.GetCostAndUsage(ctx, input)
	if err != nil {
		log.Fatalf("failed to fetch cost report, %v", err)
	}

	fmt.Println("AWS Cost Report:")
	for _, result := range resp.ResultsByTime {
		fmt.Printf("Date: %s, Cost: $%s\n", *result.TimePeriod.Start, *result.Total["BlendedCost"].Amount)
	}
}

var CostCmd = &cobra.Command{
	Use:   "cost",
	Short: "Get AWS Cost Reports",
	Run: func(cmd *cobra.Command, args []string) {
		getAllCostings()
	},
}
