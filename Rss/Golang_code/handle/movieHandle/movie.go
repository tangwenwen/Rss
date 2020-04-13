package movieHandle

import (
	"RSs/models/movieModel"
	"RSs/types"
	"github.com/emicklei/go-restful"
	"strconv"
)

type Movie struct {
}

func (e *Movie) GetAllMovieList(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	pageNo := req.Request.URL.Query().Get("pageNo")
	pageSize := req.Request.URL.Query().Get("pageSize")
	pageno, _ := strconv.ParseInt(pageNo, 10, 64)
	pagesize, _ := strconv.ParseInt(pageSize, 10, 64)
	data, err := movieModel.GetMovieList(int(pageno), int(pagesize))
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, data)
	}
}


func (e *Movie) MovieRate(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	userToken := req.Request.Header.Get("Authorization")

	movieId := req.Request.URL.Query().Get("movieId")
	rating := req.Request.URL.Query().Get("rating")
	movieId1, _ := strconv.ParseInt(movieId, 10, 64)
	rating1, _ := strconv.ParseInt(rating, 10, 64)
	err := movieModel.MovieRate(int(movieId1), int(rating1),userToken)
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, "ok")
	}
}