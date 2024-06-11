package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	dynamodblocal "github.com/abhirockzz/dynamodb-local-testcontainers-go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"gotest.tools/v3/assert"
)

func createDynamoDBContainer(t *testing.T) *dynamodblocal.DynamodbLocalContainer {
	t.Helper()

	ctx := context.Background()

	dynamodbLocalContainer, err := dynamodblocal.RunContainer(
		ctx,
		dynamodblocal.WithTelemetryDisabled(),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	return dynamodbLocalContainer
}

func createDynamoDBClient(t *testing.T, dbContainer *dynamodblocal.DynamodbLocalContainer) *dynamodb.Client {
	t.Helper()

	client, err := dbContainer.GetDynamoDBClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func createDynamoDBTable(t *testing.T, client *dynamodb.Client, tableName, tableDefinitionFilename string) {
	t.Helper()

	relativePath := filepath.Join("..", "docker", "dynamodb", tableDefinitionFilename)
	if err := createTableFromFile(client, relativePath); err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	result, err := client.ListTables(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, tableName, result.TableNames[0])

	log.Printf("%s Table created successfully\n", tableName)
}

func createTableFromFile(client *dynamodb.Client, tableDefinitionFile string) error {
	file, err := os.Open(tableDefinitionFile)
	if err != nil {
		return fmt.Errorf("failed to open table definition file: %w", err)
	}
	defer file.Close()

	var createTableInput dynamodb.CreateTableInput
	if err := json.NewDecoder(file).Decode(&createTableInput); err != nil {
		return fmt.Errorf("failed to decode table definition JSON: %w", err)
	}
	_, err = client.CreateTable(context.Background(), &createTableInput)
	return err
}

func addDataToTable(client *dynamodb.Client, tablename, val string) error {

	_, err := client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tablename),
		Item: map[string]types.AttributeValue{
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: "1"},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: "2"},
			"actor_gender":                  &types.AttributeValueMemberS{Value: val},
		},
	})

	return err
}

func queryItem(t *testing.T, client *dynamodb.Client, tableName, actorId, recipientId string) (map[string]types.AttributeValue, error) {
	t.Helper()

	output, err := client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: actorId},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: recipientId},
		},
	})

	if err != nil {
		return nil, err
	}

	return output.Item, nil
}
