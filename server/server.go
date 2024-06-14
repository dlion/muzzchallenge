package server

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dlion/muzzchallenge/explore"
	exp "github.com/dlion/muzzchallenge/explore"
)

const (
	SWIPE_TABLE = "Swipe"
)

type ExplorerServer struct {
	exp.UnimplementedExploreServiceServer
	dbClient *dynamodb.Client
}

func (s *ExplorerServer) PutSwipe(ctx context.Context, request *exp.PutSwipeRequest) (*exp.PutSwipeResponse, error) {
	recipient, err := s.updateRecipient(ctx, request)
	if err != nil {
		return &exp.PutSwipeResponse{}, err
	}

	err = s.addActor(ctx, request, recipient)
	return &exp.PutSwipeResponse{}, err
}

func (s *ExplorerServer) updateRecipient(ctx context.Context, request *exp.PutSwipeRequest) (*dynamodb.UpdateItemOutput, error) {
	recipientOutput, err := s.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(SWIPE_TABLE),
		Key: map[string]types.AttributeValue{
			"pk_swipe": &types.AttributeValueMemberS{Value: fmt.Sprintf("%d-%d", request.GetRecipientMarriageProfileId(), request.GetActorMarriageProfileId())},
		},
		UpdateExpression: aws.String("SET likedBack = :likedBack"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":likedBack": &types.AttributeValueMemberBOOL{Value: request.GetLike()},
		},
		ConditionExpression: aws.String("attribute_exists(pk_swipe)"),
		ReturnValues:        types.ReturnValueAllNew,
	})
	var conditionalCheckFailed *types.ConditionalCheckFailedException
	if err != nil && !errors.As(err, &conditionalCheckFailed) {
		return nil, err
	}

	return recipientOutput, nil
}

func (s *ExplorerServer) addActor(ctx context.Context, request *exp.PutSwipeRequest, recipient *dynamodb.UpdateItemOutput) error {
	_, err := s.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(SWIPE_TABLE),
		Item: map[string]types.AttributeValue{
			"pk_swipe":                      &types.AttributeValueMemberS{Value: fmt.Sprintf("%d-%d", request.GetActorMarriageProfileId(), request.GetRecipientMarriageProfileId())},
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorMarriageProfileId())},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetRecipientMarriageProfileId())},
			"actor_gender":                  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorGender().Number())},
			"timestamp":                     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetTimestamp())},
			"like":                          &types.AttributeValueMemberBOOL{Value: request.GetLike()},
			"likedBack":                     &types.AttributeValueMemberBOOL{Value: getLikedBackFromRecipient(recipient)},
		},
		ConditionExpression: aws.String("attribute_not_exists(pk_swipe)"),
	})
	return err
}

func getLikedBackFromRecipient(recipient *dynamodb.UpdateItemOutput) bool {
	return recipient != nil &&
		len(recipient.Attributes) > 0 &&
		recipient.Attributes["like"] != nil &&
		recipient.Attributes["like"].(*types.AttributeValueMemberBOOL).Value
}

func (s *ExplorerServer) LikedYou(ctx context.Context, request *exp.LikedYouRequest) (*exp.LikedYouResponse, error) {
	profiles, err := s.getProfilesWhoLikedTheProfile(ctx, request)
	if err != nil {
		return nil, err
	}

	return &exp.LikedYouResponse{Profiles: profiles}, nil
}

func (s *ExplorerServer) getProfilesWhoLikedTheProfile(ctx context.Context, request *exp.LikedYouRequest) ([]*exp.ExploreProfile, error) {

	queryInput := &dynamodb.ScanInput{
		TableName:            aws.String(SWIPE_TABLE),
		ProjectionExpression: aws.String("actor_marriage_profile_id, #like, #likedBack, #timestamp"),
		FilterExpression:     aws.String("#actor_gender = :actor_gender AND #like = :like AND #likedBack = :likedBack"),
		ExpressionAttributeNames: map[string]string{
			"#actor_gender": "actor_gender",
			"#like":         "like",
			"#timestamp":    "timestamp",
			"#likedBack":    "likedBack",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":actor_gender": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.Gender.Number())},
			":like":         &types.AttributeValueMemberBOOL{Value: true},
			":likedBack":    &types.AttributeValueMemberBOOL{Value: getFilter(request.GetFilter())},
		},
	}

	output, err := s.dbClient.Scan(ctx, queryInput)
	if err != nil {
		return nil, fmt.Errorf("error scanning profiles who liked the profile: %v", err)
	}

	var profiles []*exp.ExploreProfile
	limit := len(output.Items)
	if request.Limit > 0 && int(request.Limit) < limit {
		limit = int(request.Limit)
	}

	for i := 0; i < limit; i++ {
		timestamp, _ := strconv.ParseUint(output.Items[i]["timestamp"].(*types.AttributeValueMemberN).Value, 10, 32)
		actorMarriageProfileID, _ := strconv.ParseUint(output.Items[i]["actor_marriage_profile_id"].(*types.AttributeValueMemberN).Value, 10, 32)

		profiles = append(profiles, &exp.ExploreProfile{
			Timestamp:         uint32(timestamp),
			MarriageProfileId: uint32(actorMarriageProfileID),
		})
	}

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Timestamp > profiles[j].Timestamp
	})

	return profiles, nil
}

func getFilter(filter explore.LikedYou) bool {
	switch filter {
	case exp.LikedYou_LIKED_YOU_NEW:
		return false
	case exp.LikedYou_LIKED_YOU_SWIPED:
		return true
	default:
		return false
	}
}
