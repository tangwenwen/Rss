package movieHandle

import (
	"RSs/dao/userDao"
	"RSs/models/movieModel"
	"RSs/plugins/etcd"
	"RSs/types"
	usersType "RSs/types/users"
	"encoding/json"
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
	err := movieModel.MovieRate(int(movieId1), int(rating1), userToken)
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, "ok")
	}
}


func (e *Movie) Recommend(req *restful.Request, rsp *restful.Response) {
	defer func() {
		if e := recover(); e != nil {
			types.RspFailRestData(rsp, e.(error).Error())
		}
	}()
	userToken := req.Request.Header.Get("Authorization")
	user := new(usersType.TokenValue)
	data, status, err := etcd.EtcdGet(etcd.ETCDROOT + "/" + userToken)
	if err != nil {
		panic(err)
	}
	if status != 1 {
		json.Unmarshal([]byte(data), user)
	}
	userId, err := userDao.GetUserIdByUserName(user.UserName)
	if err != nil {
		panic(err)
	}
	data1,err := movieModel.Recommend(userId)
	if err != nil {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseFailHttpData(rsp, err.Error())
	} else {
		rsp.Header().Add("Content-Type", "application/json")
		types.ResponseSuccessHttpData(rsp, data1)
	}
}