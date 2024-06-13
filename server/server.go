package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	_, err := s.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(SWIPE_TABLE),
		Item: map[string]types.AttributeValue{
			"pk_swipe":                      &types.AttributeValueMemberS{Value: fmt.Sprintf("%d-%d-%d", request.GetActorMarriageProfileId(), request.GetRecipientMarriageProfileId(), request.GetTimestamp())},
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorMarriageProfileId())},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetRecipientMarriageProfileId())},
			"actor_gender":                  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorGender().Number())},
			"timestamp":                     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetTimestamp())},
			"like":                          &types.AttributeValueMemberBOOL{Value: request.GetLike()},
		},
		ConditionExpression: aws.String("attribute_not_exists(pk_swipe)"),
	})

	return &exp.PutSwipeResponse{}, err
}

func (s *ExplorerServer) LikedYou(ctx context.Context, request *exp.LikedYouRequest) (*exp.LikedYouResponse, error) {
	profiles, err := s.getProfilesWhoLikedRecipient(ctx, request)
	if err != nil {
		return nil, err
	}

	filteredProfiles, err := s.filterProfilesByLikeBackStatus(ctx, request, profiles)
	if err != nil {
		return nil, err
	}

	return &exp.LikedYouResponse{Profiles: filteredProfiles}, nil
}

func (s *ExplorerServer) getProfilesWhoLikedRecipient(ctx context.Context, request *exp.LikedYouRequest) ([]*exp.ExploreProfile, error) {
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(SWIPE_TABLE),
		IndexName:              aws.String("RecipientTimestampIndex"),
		KeyConditionExpression: aws.String("recipient_marriage_profile_id = :recipientID"),
		ProjectionExpression:   aws.String("actor_marriage_profile_id, #timestamp"),
		FilterExpression:       aws.String("#actor_gender = :actor_gender AND #like = :like"),
		ExpressionAttributeNames: map[string]string{
			"#actor_gender": "actor_gender",
			"#like":         "like",
			"#timestamp":    "timestamp",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":recipientID":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetMarriageProfileId())},
			":actor_gender": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.Gender.Number())},
			":like":         &types.AttributeValueMemberBOOL{Value: true},
		},
		ScanIndexForward: aws.Bool(false),
	}

	if request.Limit > 0 {
		queryInput.Limit = aws.Int32(int32(request.Limit))
	}

	output, err := s.dbClient.Query(ctx, queryInput)
	if err != nil {
		return nil, fmt.Errorf("error querying profiles who liked recipient: %v", err)
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

	return profiles, nil
}

func (s *ExplorerServer) filterProfilesByLikeBackStatus(ctx context.Context, request *exp.LikedYouRequest, profiles []*exp.ExploreProfile) ([]*exp.ExploreProfile, error) {
	var filteredProfiles []*exp.ExploreProfile
	for _, profile := range profiles {
		likeBackQueryInput := &dynamodb.QueryInput{
			TableName:              aws.String(SWIPE_TABLE),
			KeyConditionExpression: aws.String("pk_swipe = :likeBackKey"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":likeBackKey": &types.AttributeValueMemberS{Value: fmt.Sprintf("%d-%d-%d", request.GetMarriageProfileId(), profile.MarriageProfileId, profile.Timestamp)},
			},
		}

		likeBackOutput, err := s.dbClient.Query(ctx, likeBackQueryInput)
		if err != nil {
			return nil, fmt.Errorf("error checking like back status: %v", err)
		}

		if shouldIncludeProfile(request.Filter, likeBackOutput.Items) {
			filteredProfiles = append(filteredProfiles, profile)
		}
	}

	return filteredProfiles, nil
}

func shouldIncludeProfile(filter exp.LikedYou, likeBackItems []map[string]types.AttributeValue) bool {
	switch filter {
	case exp.LikedYou_LIKED_YOU_NEW:
		return len(likeBackItems) == 0 || !likeBackItems[0]["like"].(*types.AttributeValueMemberBOOL).Value
	case exp.LikedYou_LIKED_YOU_SWIPED:
		return len(likeBackItems) > 0 && likeBackItems[0]["like"].(*types.AttributeValueMemberBOOL).Value
	default:
		return false
	}
}
