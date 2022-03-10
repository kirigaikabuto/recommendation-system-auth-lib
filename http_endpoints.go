package recommendation_system_auth_lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HttpEndpoints interface {
	MakeCreateScoreEndpoint() gin.HandlerFunc
	MakeListScoreEndpoint() gin.HandlerFunc

	MakeRegisterEndpoint() gin.HandlerFunc
	MakeLoginEndpoint() gin.HandlerFunc
	MakeListMovies() gin.HandlerFunc

	MakeListCollaborativeFiltering() gin.HandlerFunc
}

type httpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewHttpEndpoints(ch setdata_common.CommandHandler) HttpEndpoints {
	return &httpEndpoints{ch: ch}
}

func (h *httpEndpoints) MakeCreateScoreEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &CreateScoreCommand{}
		dataBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeListScoreEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &ListScoreCommand{}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeLoginEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &LoginUserCommand{}
		dataBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeRegisterEndpoint() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &CreateUserCommand{}
		dataBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		err = json.Unmarshal(dataBytes, &cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusCreated, response)
	}
}

func (h *httpEndpoints) MakeListMovies() gin.HandlerFunc {
	return func(context *gin.Context) {
		(context.Writer).Header().Set("Access-Control-Allow-Origin", "*")
		(context.Writer).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(context.Writer).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
		cmd := &ListMoviesCommand{}
		countStr := context.Request.URL.Query().Get("count")
		count, err := strconv.ParseInt(countStr, 10, 64)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		cmd.Count = count
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, response)
	}
}

func (h *httpEndpoints) MakeListCollaborativeFiltering() gin.HandlerFunc {
	return func(context *gin.Context) {
		cmd := &ListCollaborativeFilteringCommand{}
		userId, ok := context.Get("user_id")
		if !ok {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(errors.New("no user id")))
			return

		}
		fmt.Println(userId)
		movieIdStr := context.Request.URL.Query().Get("id")
		movieId, err := strconv.ParseInt(movieIdStr, 10, 64)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		cmd.MovieId = int32(movieId)

		countStr := context.Request.URL.Query().Get("count")
		count, err := strconv.ParseInt(countStr, 10, 64)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		cmd.Count = int32(count)
		cmd.UserId = userId.(string)
		response, err := h.ch.ExecCommand(cmd)
		if err != nil {
			respondJSON(context.Writer, http.StatusInternalServerError, setdata_common.ErrToHttpResponse(err))
			return
		}
		respondJSON(context.Writer, http.StatusOK, response)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	setupResponse(&w)
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	w.Write(response)
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
}
