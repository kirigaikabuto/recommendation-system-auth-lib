package recommendation_system_auth_lib

import score "github.com/kirigaikabuto/recommendation-system-score-store"

type AuthLibService interface {
	CreateScore(cmd *CreateScoreCommand) (*score.Score, error)
	ListScore(cmd *ListScoreCommand) ([]score.Score, error)
}

type authLibService struct {
	amqpRequest AmqpRequests
}

func NewAuthLibService(a AmqpRequests) AuthLibService {
	return &authLibService{amqpRequest: a}
}

func (a *authLibService) CreateScore(cmd *CreateScoreCommand) (*score.Score, error) {
	return a.amqpRequest.CreateScore(cmd)
}

func (a *authLibService) ListScore(cmd *ListScoreCommand) ([]score.Score, error) {
	return a.amqpRequest.ListScore(cmd)
}
