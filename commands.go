package recommendation_system_auth_lib

type CreateScoreCommand struct {
	UserId  string  `json:"user_id"`
	MovieId int64   `json:"movie_id"`
	Rating  float64 `json:"rating"`
}

func (cmd *CreateScoreCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).CreateScore(cmd)
}

type ListScoreCommand struct {
}

func (cmd *ListScoreCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).ListScore(cmd)
}
