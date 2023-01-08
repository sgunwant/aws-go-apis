package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	log.SetFlags(0)
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the EC2 client
	svc := ec2.NewFromConfig(cfg)

	// Build the request with its input parameters
	res, err := svc.DescribeRegions(context.Background(), &ec2.DescribeRegionsInput{})

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	region_name := map[string]string{
		"ap-northeast-1": "Asia Pacific (Tokyo)",
		"ap-northeast-2": "Asia Pacific (Seoul)",
		"ap-northeast-3": "Asia Pacific (Osaka)",
		"ap-southeast-1": "Asia Pacific (Singapore)",
		"ap-southeast-2": "Asia Pacific (Sydney)",
		"ap-south-1":     "Asia Pacific (Mumbai)",
		"eu-central-1":   "EU (Frankfurt)",
		"eu-north-1":     "Europe (Stockholm)",
		"eu-west-1":      "EU (Ireland)",
		"eu-west-2":      "Europe (London)",
		"eu-west-3":      "Europe (Paris)",
		"us-east-1":      "US East (N. Virginia)",
		"us-east-2":      "US East (Ohio)",
		"us-west-1":      "US West (N. California)",
		"us-west-2":      "US West (Oregon)",
		"sa-east-1":      "South America (Sao Paulo)",
		"ca-central-1":   "Canada (Central)",
	}

	log.Print("Regions:")
	for _, region := range res.Regions {
		log.Println("\t* RegionID:", *region.RegionName, "\tRegionName: ", region_name[*region.RegionName], "\tStatus: ", *region.OptInStatus)
	}

}

