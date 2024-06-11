package server

import (
	"context"
	"fmt"

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

func (s *ExplorerServer) LikedYou(ctx context.Context, request *exp.LikedYouRequest) (*exp.LikedYouResponse, error) {
	return nil, nil
}

func (s *ExplorerServer) PutSwipe(ctx context.Context, request *exp.PutSwipeRequest) (*exp.PutSwipeResponse, error) {
	_, err := s.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(SWIPE_TABLE),
		Item: map[string]types.AttributeValue{
			"pk_swipe":                      &types.AttributeValueMemberS{Value: fmt.Sprintf("%d-%d-%d", request.GetActorMarriageProfileId(), request.GetRecipientMarriageProfileId(), request.GetTimestamp())},
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorMarriageProfileId())},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetRecipientMarriageProfileId())},
			"actor_gender":                  &types.AttributeValueMemberS{Value: fmt.Sprintf("%d", request.GetActorGender())},
			"timestamp":                     &types.AttributeValueMemberS{Value: fmt.Sprintf("%d", request.GetTimestamp())},
			"like":                          &types.AttributeValueMemberBOOL{Value: request.GetLike()},
		},
		ConditionExpression: aws.String("attribute_not_exists(pk_swipe)"),
	})

	return &exp.PutSwipeResponse{}, err
}
