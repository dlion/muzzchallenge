package server

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dlion/muzzchallenge/explore"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("Should swipe as an actor against a recipient giving like to it", func(t *testing.T) {
		dbContainer := createDynamoDBContainer(t)
		defer func() {
			if err := dbContainer.Terminate(context.Background()); err != nil {
				log.Fatalf("failed to terminate container: %s", err)
			}
		}()
		client := createDynamoDBClient(t, dbContainer)
		createDynamoDBTable(t, client, "Swipe", "swipe_table.json")

		srv := ExplorerServer{dbClient: client}
		srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
			ActorMarriageProfileId:     1,
			RecipientMarriageProfileId: 2,
			ActorGender:                explore.Gender_GENDER_MALE,
			Timestamp:                  uint32(time.Now().Unix()),
			Like:                       true,
		})

		items, err := queryItem(t, client, "Swipe", "1", "2")
		if assert.NoError(t, err) && assert.NotEmpty(t, items) {
			assert.Equal(t, "0", items["actor_gender"].(*types.AttributeValueMemberS).Value)
			assert.Equal(t, true, items["like"].(*types.AttributeValueMemberBOOL).Value)
		}
	})
}
