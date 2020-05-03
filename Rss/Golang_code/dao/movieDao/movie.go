package movieDao

import (
	"RSs/plugins/db"
	movieType "RSs/types/movie"
	"strconv"
	"time"
)

func GetAllMovieCount() (int64, error) {
	orm, err := db.GetEngine()
	if err != nil {
		return 0, err
	}
	return orm.Count(&movieType.Movie{})
}

func GetMovieList(limit, offset int) ([]*movieType.MovieList, error) {
	AllList := make([]*movieType.Movie, 0)
	orm, err := db.GetEngine()
	if err != nil {
		return nil, err
	}
	err = orm.Where("status = ?", 0).Limit(limit, offset).Find(&AllList)
	MList := make([]*movieType.MovieList, len(AllList))
	for i := 0; i < len(AllList); i++ {
		MList[i] = new(movieType.MovieList)
		MList[i].Id = strconv.Itoa(AllList[i].MovieId)
		MList[i].Img = ""
		MList[i].Genres = AllList[i].Genres
		MList[i].Title = AllList[i].Title
		MList[i].Nopa = strconv.Itoa(GetNopa(AllList[i].MovieId))
	}
	return MList, nil
}

func GetNopa(movieId int) int {
	orm, _ := db.GetEngine()
	count, _ := orm.Count(&movieType.Ratings{MovieId: movieId})
	return int(count)
}


func ExistRating(movieId,userId int)(bool,error){
	orm, err := db.GetEngine()
	if err != nil {
		return false, err
	}
	rating := new(movieType.Ratings)
	rating.MovieId = movieId
	rating.UserId = userId
	return orm.Exist(rating)
}

func InsertOnRating(movieId,userId,rating int)error{
	orm, err := db.GetEngine()
	if err != nil {
		return  err
	}
	_,err = orm.InsertOne(&movieType.Ratings{
		MovieId:movieId,
		UserId:userId,
		Rating:rating,
		CreatedTime:time.Now(),
		Status: 0,
	})
	return err
}

func GetMovieById(movieId int)(*movieType.Movie,error){
	movie:=&movieType.Movie{MovieId:movieId}
	orm, err := db.GetEngine()
	if err != nil {
		return  nil,err
	}
	_,err = orm.Get(movie)
	if err!=nil{
		return nil,err
	}
	return movie,nil

}