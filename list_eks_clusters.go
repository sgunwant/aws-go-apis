package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func main() {
	log.SetFlags(0)
	region := "ap-northeast-1"
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the eks client
	svc := eks.NewFromConfig(cfg)

	// Build the request with its input parameters
	res, err := svc.ListClusters(context.Background(), &eks.ListClustersInput{})

	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	log.Print("EKS clusters:")
	for _, cluster := range res.Clusters {
		cluster_info, _ := svc.DescribeCluster(context.Background(), &eks.DescribeClusterInput{Name: &cluster})

		log.Println("\tName: ", *cluster_info.Cluster.Name, "\tPlatformVersion: ", *cluster_info.Cluster.PlatformVersion,
			"\tStatus: ", *&cluster_info.Cluster.Status, "\tKubeVersion: ", *cluster_info.Cluster.Version,
			"\tResources: ", *&cluster_info.Cluster.ResourcesVpcConfig.SubnetIds)
	}

}

