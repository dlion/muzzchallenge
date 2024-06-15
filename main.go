package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dlion/muzzchallenge/explore"
	"github.com/dlion/muzzchallenge/server"
	"google.golang.org/grpc"
)

var (
	port                  = flag.Int("p", 37857, "The service port")
	dynamoEndpoint        = flag.String("dynamoEndpoint", "http://localhost:8000", "DynamoDB Endpoint")
	dynamoAccessKey       = flag.String("dynamoAccessKey", "dummy", "DynamoDB Accesskey")
	dynamoSecretAccessKey = flag.String("dynamoSecretAccessKey", "dummy", "DynamoDB SecretAccessKey")
	dynamoRegion          = flag.String("dynamoRegion", "local", "DynamoDB Region")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dbClient, err := getDynamoDBClient()
	if err != nil {
		log.Fatalf("Can't create a DynamoDB client: %v", err)
	}

	s := grpc.NewServer()
	explore.RegisterExploreServiceServer(s, &server.ExplorerServer{DbClient: dbClient})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func getDynamoDBClient() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(*dynamoRegion),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: *dynamoEndpoint}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     *dynamoAccessKey,
				SecretAccessKey: *dynamoSecretAccessKey,
				SessionToken:    "",
				Source:          "Credentials from the parameters",
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
