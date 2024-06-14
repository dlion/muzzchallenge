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
		t.Run("Should swipe as an actor against a recipient giving like to it, the recipient doesnt exist yet", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")

			srv := ExplorerServer{dbClient: client}
			_, err := srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
				ActorMarriageProfileId:     1,
				RecipientMarriageProfileId: 2,
				ActorGender:                explore.Gender_GENDER_MALE,
				Timestamp:                  uint32(time.Now().Unix()),
				Like:                       true,
			})
			assert.NoError(t, err)

			items, err := queryItem(t, client, "Swipe", "1", "2")
			if assert.NoError(t, err) && assert.NotEmpty(t, items) {
				assert.Equal(t, "0", items["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, true, items["like"].(*types.AttributeValueMemberBOOL).Value)
				assert.Equal(t, false, items["likedBack"].(*types.AttributeValueMemberBOOL).Value)
			}
		})

		t.Run("Should swipe as an actor against a recipient giving like to it, the recipient exist and didn't give a like back, the recipient should get a likedBack update, the actor shouldnt have a likedBack false", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")
			err := addSwipeToTable(client, "Swipe", fmt.Sprintf("%d", time.Now().Unix()), "2", "1", explore.Gender_GENDER_FEMALE, false, false)
			assert.NoError(t, err)

			srv := ExplorerServer{dbClient: client}
			_, err = srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
				ActorMarriageProfileId:     1,
				RecipientMarriageProfileId: 2,
				ActorGender:                explore.Gender_GENDER_MALE,
				Timestamp:                  uint32(time.Now().Unix()),
				Like:                       true,
			})
			assert.NoError(t, err)

			items, err := queryItem(t, client, "Swipe", "1", "2")
			if assert.NoError(t, err) && assert.NotEmpty(t, items) {
				assert.Equal(t, "0", items["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, true, items["like"].(*types.AttributeValueMemberBOOL).Value)
				assert.Equal(t, false, items["likedBack"].(*types.AttributeValueMemberBOOL).Value)
			}

			items, err = queryItem(t, client, "Swipe", "2", "1")
			if assert.NoError(t, err) && assert.NotEmpty(t, items) {
				assert.Equal(t, "1", items["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, false, items["like"].(*types.AttributeValueMemberBOOL).Value)
				assert.Equal(t, true, items["likedBack"].(*types.AttributeValueMemberBOOL).Value)
			}
		})

		t.Run("Should swipe as an actor against a recipient giving like to it, the recipient exist and gave a like back, the recipient should get a likedback update, the actor should have a likedBack true", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")
			err := addSwipeToTable(client, "Swipe", fmt.Sprintf("%d", time.Now().Unix()), "2", "1", explore.Gender_GENDER_FEMALE, true, false)
			assert.NoError(t, err)

			srv := ExplorerServer{dbClient: client}
			_, err = srv.PutSwipe(context.Background(), &explore.PutSwipeRequest{
				ActorMarriageProfileId:     1,
				RecipientMarriageProfileId: 2,
				ActorGender:                explore.Gender_GENDER_MALE,
				Timestamp:                  uint32(time.Now().Unix()),
				Like:                       true,
			})
			assert.NoError(t, err)

			items, err := queryItem(t, client, "Swipe", "1", "2")
			if assert.NoError(t, err) && assert.NotEmpty(t, items) {
				assert.Equal(t, "0", items["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, true, items["like"].(*types.AttributeValueMemberBOOL).Value)
				assert.Equal(t, true, items["likedBack"].(*types.AttributeValueMemberBOOL).Value)
			}

			items, err = queryItem(t, client, "Swipe", "2", "1")
			if assert.NoError(t, err) && assert.NotEmpty(t, items) {
				assert.Equal(t, "1", items["actor_gender"].(*types.AttributeValueMemberN).Value)
				assert.Equal(t, true, items["like"].(*types.AttributeValueMemberBOOL).Value)
				assert.Equal(t, true, items["likedBack"].(*types.AttributeValueMemberBOOL).Value)
			}
		})
	})

	t.Run("LikedYou", func(t *testing.T) {
		t.Run("Using the filter LikedYou_LIKED_YOU_NEW, it should get back all the people who liked the specific profile, but that didn't get the like back", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")
			timestamp1 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp1, "1", "2", explore.Gender_GENDER_FEMALE, true, false)
			timestamp2 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp2, "3", "2", explore.Gender_GENDER_FEMALE, true, false)
			timestamp4 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp4, "5", "2", explore.Gender_GENDER_FEMALE, true, false)
			timestamp5 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp5, "4", "2", explore.Gender_GENDER_FEMALE, true, true)

			server := ExplorerServer{dbClient: client}
			response, err := server.LikedYou(context.Background(), &explore.LikedYouRequest{
				MarriageProfileId: 2,
				Gender:            explore.Gender_GENDER_FEMALE,
				Filter:            explore.LikedYou_LIKED_YOU_NEW,
			}) // Give me all the female profiles who liked 2 but that didn't get a like back from 2
			assert.NoError(t, err)
			profiles := response.GetProfiles()

			assert.Equal(t, 3, len(profiles))
			assert.Equal(t, timestamp1, fmt.Sprintf("%d", profiles[0].Timestamp))
			assert.Equal(t, timestamp2, fmt.Sprintf("%d", profiles[1].Timestamp))
			assert.Equal(t, timestamp4, fmt.Sprintf("%d", profiles[2].Timestamp))
		})

		t.Run("Using the filter LikedYou_LIKED_YOU_SWIPED, it should get back all the people who liked the specific profile and that got the like back", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")
			timestamp1 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp1, "1", "2", explore.Gender_GENDER_FEMALE, true, true)
			timestamp2 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp2, "3", "2", explore.Gender_GENDER_FEMALE, true, true)
			timestamp4 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp4, "5", "2", explore.Gender_GENDER_FEMALE, true, false)
			timestamp5 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp5, "4", "2", explore.Gender_GENDER_FEMALE, true, true)

			server := ExplorerServer{dbClient: client}
			response, err := server.LikedYou(context.Background(), &explore.LikedYouRequest{
				MarriageProfileId: 2,
				Gender:            explore.Gender_GENDER_FEMALE,
				Filter:            explore.LikedYou_LIKED_YOU_SWIPED,
			})
			assert.NoError(t, err)
			profiles := response.GetProfiles()

			assert.Equal(t, 3, len(profiles))
			assert.Equal(t, timestamp1, fmt.Sprintf("%d", profiles[0].Timestamp))
			assert.Equal(t, timestamp2, fmt.Sprintf("%d", profiles[1].Timestamp))
			assert.Equal(t, timestamp5, fmt.Sprintf("%d", profiles[2].Timestamp))
		})

		t.Run("Using the filter LikedYou_LIKED_YOU_SWIPED, it should get back all the people who liked the specific profile and that got the like back with a limit of 2 people got back", func(t *testing.T) {
			dbContainer := createDynamoDBContainer(t)
			defer func() {
				if err := dbContainer.Terminate(context.Background()); err != nil {
					log.Fatalf("failed to terminate container: %s", err)
				}
			}()
			client := createDynamoDBClient(t, dbContainer)
			createDynamoDBTable(t, client, "Swipe", "swipe_table.json")
			timestamp1 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp1, "1", "2", explore.Gender_GENDER_FEMALE, true, true)
			timestamp2 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp2, "3", "2", explore.Gender_GENDER_FEMALE, true, true)
			timestamp4 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp4, "5", "2", explore.Gender_GENDER_FEMALE, true, false)
			timestamp5 := fmt.Sprintf("%d", time.Now().Unix())
			addSwipeToTable(client, "Swipe", timestamp5, "4", "2", explore.Gender_GENDER_FEMALE, true, true)

			server := ExplorerServer{dbClient: client}
			response, err := server.LikedYou(context.Background(), &explore.LikedYouRequest{
				MarriageProfileId: 2,
				Gender:            explore.Gender_GENDER_FEMALE,
				Filter:            explore.LikedYou_LIKED_YOU_SWIPED,
				Limit:             2,
			})
			assert.NoError(t, err)
			profiles := response.GetProfiles()

			assert.Equal(t, 2, len(profiles))
			assert.Equal(t, timestamp1, fmt.Sprintf("%d", profiles[0].Timestamp))
			assert.Equal(t, timestamp2, fmt.Sprintf("%d", profiles[1].Timestamp))
		})
	})
}
