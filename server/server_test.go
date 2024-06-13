package server

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dlion/muzzchallenge/explore"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("PutSwipe", func(t *testing.T) {
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
			timestamp := time.Now().Unix()
			srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
				ActorMarriageProfileId:     1,
				RecipientMarriageProfileId: 2,
				ActorGender:                explore.Gender_GENDER_MALE,
				Timestamp:                  uint32(timestamp),
				Like:                       true,
			})

			items, err := queryItem(t, client, "Swipe", "1", "2", fmt.Sprintf("%d", timestamp))
			if assert.NoError(t, err) && assert.NotEmpty(t, items) {
				assert.Equal(t, "0", items["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, true, items["like"].(*types.AttributeValueMemberBOOL).Value)
			}
		})

		t.Run("Should be idempotent", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")

			srv := ExplorerServer{dbClient: client}
			timestamp := time.Now().Unix()
			srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
				ActorMarriageProfileId:     1,
				RecipientMarriageProfileId: 2,
				ActorGender:                explore.Gender_GENDER_MALE,
				Timestamp:                  uint32(timestamp),
				Like:                       true,
			})

			srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
				ActorMarriageProfileId:     1,
				RecipientMarriageProfileId: 2,
				ActorGender:                explore.Gender_GENDER_MALE,
				Timestamp:                  uint32(timestamp),
				Like:                       false,
			})

			item, err := queryItem(t, client, "Swipe", "1", "2", fmt.Sprintf("%d", timestamp))
			if assert.NoError(t, err) && assert.NotEmpty(t, item) {
				assert.Equal(t, "0", item["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, true, item["like"].(*types.AttributeValueMemberBOOL).Value)
			}
		})
	})

	t.Run("LikedYou", func(t *testing.T) {
		t.Run("Should returns all profiles who liked a specific user but that didn't get a like back from that specific user in descending order", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")

			//Actor 1 likes recipient 2
			timestamp1 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp1, "1", "2", explore.Gender_GENDER_FEMALE, true)
			//Actor 3 likes recipient 2
			timestamp2 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp2, "3", "2", explore.Gender_GENDER_FEMALE, true)
			// Actor 4 likes recipient 2
			timestamp3 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp3, "4", "2", explore.Gender_GENDER_FEMALE, true)
			// Actor 2 likes recipient 1
			timestamp4 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp4, "2", "1", explore.Gender_GENDER_FEMALE, true)
			// Actor 2 likes recipient 4
			timestamp5 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp5, "2", "4", explore.Gender_GENDER_FEMALE, true)
			// Actor 2 doesnt like recipient 3
			timestamp6 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp6, "2", "3", explore.Gender_GENDER_FEMALE, false)
			// Actor 2 doesnt like recipient 3
			timestamp7 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp7, "5", "2", explore.Gender_GENDER_FEMALE, true)

			server := ExplorerServer{dbClient: client}
			response, err := server.LikedYou(context.Background(), &explore.LikedYouRequest{
				MarriageProfileId: 2,
				Gender:            explore.Gender_GENDER_FEMALE,
				Filter:            explore.LikedYou_LIKED_YOU_NEW,
			})
			if err != nil {
				log.Fatalf("error getting back the response %v", err)
			}
			profiles := response.GetProfiles()

			assert.Equal(t, 2, len(profiles))
			assert.Equal(t, timestamp2, fmt.Sprintf("%d", profiles[0].Timestamp))
			assert.Equal(t, timestamp7, fmt.Sprintf("%d", profiles[1].Timestamp))
		})

		t.Run("Should returns all profiles who liked a specific user and that got a like back from that specific user in descending order", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")

			// Actor 1 likes Recipient 2
			timestamp1 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp1, "1", "2", explore.Gender_GENDER_FEMALE, true)
			// Actor 3 likes Recipient 2
			timestamp2 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp2, "3", "2", explore.Gender_GENDER_FEMALE, true)
			// Actror 4 likes recipient 2
			timestamp3 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp3, "4", "2", explore.Gender_GENDER_FEMALE, true)
			// Actor 2 likes recipient 1
			timestamp4 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp4, "2", "1", explore.Gender_GENDER_FEMALE, true)
			// Actor 2 likes recipient 4
			timestamp5 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp5, "2", "4", explore.Gender_GENDER_FEMALE, true)
			// Actor 2 didnt like recipient 3
			timestamp6 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp6, "2", "3", explore.Gender_GENDER_FEMALE, false)
			// Actor 2 like recipient 5
			timestamp7 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp7, "2", "5", explore.Gender_GENDER_FEMALE, true)
			// Actor 5 like recipient 2
			timestamp8 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp8, "5", "2", explore.Gender_GENDER_FEMALE, true)

			// 1 -> 2, 2 -> 1 = match
			//3 -> 2, 2 -> 3 = NO
			//4 -> 2, 2 -> 4 = match
			//2 -> 5, 5 -> 2 = match

			server := ExplorerServer{dbClient: client}
			response, err := server.LikedYou(context.Background(), &explore.LikedYouRequest{
				MarriageProfileId: 2,
				Gender:            explore.Gender_GENDER_FEMALE,
				Filter:            explore.LikedYou_LIKED_YOU_SWIPED,
				Limit:             2,
			})
			if err != nil {
				log.Fatalf("error getting back the response %v", err)
			}
			profiles := response.GetProfiles()

			assert.Equal(t, 2, len(profiles))
			assert.Equal(t, timestamp1, fmt.Sprintf("%d", profiles[0].Timestamp))
			assert.Equal(t, timestamp3, fmt.Sprintf("%d", profiles[1].Timestamp))
		})
	})
}
