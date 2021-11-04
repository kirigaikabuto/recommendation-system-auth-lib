package recommendation_system_auth_lib

import (
	"github.com/djumanoff/amqp"
	score "github.com/kirigaikabuto/recommendation-system-score-store"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type AmqpRequests struct {
	clt amqp.Client
}

func NewAmqpRequests(clt amqp.Client) AmqpRequests {
	return AmqpRequests{clt: clt}
}

func (r *AmqpRequests) CreateScore(cmd *CreateScoreCommand) (*score.Score, error) {
	response := &score.Score{}
	err := setdata_common.AmqpCall(r.clt, "score.create", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) ListScore(cmd *ListScoreCommand) ([]score.Score, error) {
	response := []score.Score{}
	err := setdata_common.AmqpCall(r.clt, "score.list", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
