package movieModel

import (
	"RSs/dao/movieDao"
	"RSs/dao/userDao"
	"RSs/plugins/etcd"
	"RSs/plugins/pagination"
	movieType "RSs/types/movie"
	usersType "RSs/types/users"
	"encoding/json"
	"errors"
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


func MovieRate(movieId ,rating int ,token string)error{
	user:=new(usersType.TokenValue)
	data, status, err := etcd.EtcdGet(etcd.ETCDROOT + "/" + token)
	if err != nil {
		return err
	}
	if status != 1 {
		json.Unmarshal([]byte(data), user)
	}
	userId,err:=userDao.GetUserIdByUserName(user.UserName)
	if err!=nil{
		return err
	}
	if has,_ :=movieDao.ExistRating(movieId,userId);has{
		return errors.New("您已评价")
	}
	err =  movieDao.InsertOnRating(movieId,userId,rating)
	if err!=nil{
		return err
	}
	return nil

}
