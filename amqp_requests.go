package recommendation_system_auth_lib

import (
	"github.com/djumanoff/amqp"
	movies "github.com/kirigaikabuto/recommendation-system-movie-store"
	score "github.com/kirigaikabuto/recommendation-system-score-store"
	users "github.com/kirigaikabuto/recommendation-system-users-store"
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

func (r *AmqpRequests) CreateUser(cmd *CreateUserCommand) (*users.User, error) {
	response := &users.User{}
	err := setdata_common.AmqpCall(r.clt, "users.create", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) GetUserByUsername(cmd *GetUserByUsernameAndPassword) (*users.User, error) {
	response := &users.User{}
	err := setdata_common.AmqpCall(r.clt, "users.getByUsernameAndPassword", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *AmqpRequests) ListMovies(cmd *ListMoviesCommand) ([]movies.Movie, error) {
	response := []movies.Movie{}
	err := setdata_common.AmqpCall(r.clt, "movie.lists", cmd, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
