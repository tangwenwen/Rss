package movieModel

import (
	"RSs/dao/movieDao"
	"RSs/dao/userDao"
	"RSs/plugins/etcd"
	"RSs/plugins/logs"
	"RSs/plugins/pagination"
	pb "RSs/rpc"
	movieType "RSs/types/movie"
	usersType "RSs/types/users"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func GetMovieList(pageNo, pageSize int) (*movieType.MoviePageResp, error) {
	movieList := make([]*movieType.MovieList, 0)
	pagg := new(movieType.Pagination)
	total, err := movieDao.GetAllMovieCount()
	if err != nil {
		return nil, err
	}
	limit, offset, err := pagination.GetORMLimitOffset(pageNo, pageSize, int(total))
	if err != nil {
		return nil, err
	}
	pagg.Pages = pagination.GetPageCnt(limit, int(total))
	pagg.PageSize = limit
	pagg.Total = int(total)

	movieList, err = movieDao.GetMovieList(limit, offset)
	if err != nil {
		return nil, err
	}
	moviePageResp := &movieType.MoviePageResp{
		MList: movieList,
		Page:  pagg,
	}
	return moviePageResp, nil
}

func MovieRate(movieId, rating int, token string) error {
	user := new(usersType.TokenValue)
	data, status, err := etcd.EtcdGet(etcd.ETCDROOT + "/" + token)
	if err != nil {
		return err
	}
	if status != 1 {
		json.Unmarshal([]byte(data), user)
	}
	userId, err := userDao.GetUserIdByUserName(user.UserName)
	if err != nil {
		return err
	}
	if has, _ := movieDao.ExistRating(movieId, userId); has {
		return errors.New("您已评价")
	}
	err = movieDao.InsertOnRating(movieId, userId, rating)
	if err != nil {
		return err
	}
	return nil

}


func Recommend(userId int)([]*movieType.MovieList, error) {
	Reply:=make([]*pb.RecommendItem,0)
	Reply,err:= Rpc(int32(userId))
	if err!=nil{
		return nil,err
	}
	movieList :=make([]*movieType.MovieList,len(Reply))
	for i:=0;i<len(Reply);i++{
		movieList[i] = new(movieType.MovieList)
		movieObj,_:= movieDao.GetMovieById(int(Reply[i].MoiveId))
		movieList[i].Id = strconv.Itoa(movieObj.MovieId)
		movieList[i].Nopa = strconv.Itoa(movieDao.GetNopa(int(Reply[i].MoiveId)))
		movieList[i].Img = ""
		movieList[i].Title = movieObj.Title
		movieList[i].Genres = movieObj.Genres
	}
	return movieList,nil

}


func Rpc(userId int32) ([]*pb.RecommendItem,error){
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		logs.Error("did not connect: %v", err)
		return nil,err
	}
	defer conn.Close()
	c := pb.NewRecommendClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*210)
	defer cancel()
	r, err := c.Recommend(ctx, &pb.RecommendReq{UserId: userId})
	if err!=nil{
		logs.Error("rpc err: ", err)
		return nil,err
	}
	return r.RecommendItem,err
}
