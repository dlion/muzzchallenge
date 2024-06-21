package testutils

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	dynamodblocal "github.com/abhirockzz/dynamodb-local-testcontainers-go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dlion/muzzchallenge/explore"
	"gotest.tools/v3/assert"
)

//go:embed swipe_table.json
var tableDefinitionFile []byte

func CreateDynamoDBContainer(t *testing.T) *dynamodblocal.DynamodbLocalContainer {
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

func CreateDynamoDBClient(t *testing.T, dbContainer *dynamodblocal.DynamodbLocalContainer) *dynamodb.Client {
	t.Helper()

	client, err := dbContainer.GetDynamoDBClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CreateDynamoDBTable(t *testing.T, client *dynamodb.Client, tableName string) {
	t.Helper()

	if err := createTableFromFile(client, tableDefinitionFile); err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	result, err := client.ListTables(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, tableName, result.TableNames[0])

	log.Printf("%s Table created successfully\n", tableName)
}

func createTableFromFile(client *dynamodb.Client, tableDefinitionFile []byte) error {
	var createTableInput dynamodb.CreateTableInput
	if err := json.NewDecoder(bytes.NewReader(tableDefinitionFile)).Decode(&createTableInput); err != nil {
		return fmt.Errorf("failed to decode table definition JSON: %w", err)
	}
	_, err := client.CreateTable(context.Background(), &createTableInput)
	return err
}

func AddSwipeToTable(
	client *dynamodb.Client,
	tablename,
	timestamp,
	actorId,
	recipientId string,
	gender explore.Gender,
	like bool,
	likedBack bool) error {

	_, err := client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tablename),
		Item: map[string]types.AttributeValue{
			"pk_swipe":                      &types.AttributeValueMemberS{Value: fmt.Sprintf("%s-%s", actorId, recipientId)},
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: actorId},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: recipientId},
			"actor_gender":                  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", gender.Number())},
			"like":                          &types.AttributeValueMemberBOOL{Value: like},
			"timestamp":                     &types.AttributeValueMemberN{Value: timestamp},
			"likedBack":                     &types.AttributeValueMemberBOOL{Value: likedBack},
		},
	})

	return err
}

func QueryItem(t *testing.T, client *dynamodb.Client, tableName, actorId, recipientId string) (map[string]types.AttributeValue, error) {
	t.Helper()

	output, err := client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"pk_swipe": &types.AttributeValueMemberS{
				Value: fmt.Sprintf("%s-%s", actorId, recipientId),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return output.Item, nil
}
