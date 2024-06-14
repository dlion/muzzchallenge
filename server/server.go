package server

import (
	"context"
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

	s.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(SWIPE_TABLE),
		Item: map[string]types.AttributeValue{
			"pk_swipe":                      &types.AttributeValueMemberS{Value: fmt.Sprintf("%d-%d", request.GetActorMarriageProfileId(), request.GetRecipientMarriageProfileId())},
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorMarriageProfileId())},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetRecipientMarriageProfileId())},
			"actor_gender":                  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorGender().Number())},
			"timestamp":                     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetTimestamp())},
			"like":                          &types.AttributeValueMemberBOOL{Value: request.GetLike()},
			"likedBack":                     &types.AttributeValueMemberBOOL{Value: recipientOutput != nil && len(recipientOutput.Attributes) > 0 && recipientOutput.Attributes["like"] != nil && recipientOutput.Attributes["like"].(*types.AttributeValueMemberBOOL).Value},
		},
		ConditionExpression: aws.String("attribute_not_exists(pk_swipe)"),
	})

	return &exp.PutSwipeResponse{}, err
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
	for _, v := range output.Items {
		timestamp, _ := strconv.ParseUint(v["timestamp"].(*types.AttributeValueMemberN).Value, 10, 32)
		actorMarriageProfileID, _ := strconv.ParseUint(v["actor_marriage_profile_id"].(*types.AttributeValueMemberN).Value, 10, 32)

		profiles = append(profiles, &exp.ExploreProfile{
			Timestamp:         uint32(timestamp),
			MarriageProfileId: uint32(actorMarriageProfileID),
		})
	}

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Timestamp > profiles[j].Timestamp
	})

	if request.Limit > 0 && len(profiles) > int(request.Limit) {
		profiles = profiles[:request.Limit]
	}

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
