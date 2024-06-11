package server

import (
	"context"
	"fmt"
	"time"

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
		//TODO: Idempotent key
		Item: map[string]types.AttributeValue{
			"actor_marriage_profile_id":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetActorMarriageProfileId())},
			"recipient_marriage_profile_id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", request.GetRecipientMarriageProfileId())},
			"actor_gender":                  &types.AttributeValueMemberS{Value: fmt.Sprintf("%d", request.GetActorGender())},
			"timestamp":                     &types.AttributeValueMemberS{Value: fmt.Sprintf("%d", time.Now().Unix())},
			"like":                          &types.AttributeValueMemberBOOL{Value: request.GetLike()},
		},
	})

	return &exp.PutSwipeResponse{}, err
}
