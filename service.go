package recommendation_system_auth_lib

import (
	"github.com/google/uuid"
	score "github.com/kirigaikabuto/recommendation-system-score-store"
	users "github.com/kirigaikabuto/recommendation-system-users-store"
	redis "github.com/kirigaikabuto/setdata-common/redis_connect"
	"time"
)

type AuthLibService interface {
	CreateScore(cmd *CreateScoreCommand) (*score.Score, error)
	ListScore(cmd *ListScoreCommand) ([]score.Score, error)

	RegisterUser(cmd *CreateUserCommand) (*users.User, error)
	LoginUser(cmd *LoginUserCommand) (*LoginResponse, error)
}

type authLibService struct {
	amqpRequest    AmqpRequests
	userTokenStore redis.RedisConnectStore
}

func NewAuthLibService(a AmqpRequests, u redis.RedisConnectStore) AuthLibService {
	return &authLibService{amqpRequest: a, userTokenStore: u}
}

func (a *authLibService) CreateScore(cmd *CreateScoreCommand) (*score.Score, error) {
	return a.amqpRequest.CreateScore(cmd)
}

func (a *authLibService) ListScore(cmd *ListScoreCommand) ([]score.Score, error) {
	return a.amqpRequest.ListScore(cmd)
}

func (a *authLibService) RegisterUser(cmd *CreateUserCommand) (*users.User, error) {
	return a.amqpRequest.CreateUser(cmd)
}

func (a *authLibService) LoginUser(cmd *LoginUserCommand) (*LoginResponse, error) {
	user, err := a.amqpRequest.GetUserByUsername(&GetUserByUsernameAndPassword{Username: cmd.Username, Password: cmd.Password})
	if err != nil {
		return nil, err
	}
	uuId := uuid.New()
	accessKey := uuId.String()
	err = a.userTokenStore.Save(accessKey, user.Id, 5*time.Minute)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		UserId:    user.Id,
		AccessKey: accessKey,
	}, nil
}
