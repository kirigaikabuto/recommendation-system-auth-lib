package recommendation_system_auth_lib

import movies_lib "github.com/kirigaikabuto/recommendation-system-movie-store"

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

type LoginUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *LoginUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).LoginUser(cmd)
}

type GetUserByUsernameAndPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId    string `json:"user_id"`
	AccessKey string `json:"access_key"`
}

type CreateUserCommand struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	Country   string `json:"country"`
}

func (cmd *CreateUserCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).RegisterUser(cmd)
}

type ListMoviesCommand struct {
	Count int64 `json:"count"`
}

func (cmd *ListMoviesCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).ListMovies(cmd)
}

type ListCollaborativeFilteringCommand struct {
	UserId  string `json:"user_id"`
	MovieId int32  `json:"movie_id"`
	Count   int32  `json:"count"`
}

func (cmd *ListCollaborativeFilteringCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).ListCollaborativeFiltering(cmd)
}

type FilteredMovie struct {
	movies_lib.Movie
	PredictedRating float32 `json:"predicted_rating"`
}

type GetMovieById struct {
	Id int64 `json:"id"`
}

type CollaborativeFilteringResponse struct {
	Current           *movies_lib.Movie `json:"current"`
	RecommendedMovies []FilteredMovie   `json:"recommended_movies"`
}

type ListContentBasedFilteringCommand struct {
	MovieId int32 `json:"movie_id"`
	Count   int32 `json:"count"`
}

func (cmd *ListContentBasedFilteringCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(AuthLibService).ListContentBasedFiltering(cmd)
}
