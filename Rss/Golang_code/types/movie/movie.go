package movieType

import "time"

type MoviePageResp struct {
	MList []*MovieList `json:"movieList"`
	Page  *Pagination  `json:"pagination"`
}

type MovieList struct {
	Id     string `json:"id"`
	Img    string `json:"img"`
	Title  string `json:"title"`
	Genres string `json:"genres"`
	Nopa   string `json:"nopa"`
}
type Pagination struct {
	PageSize int `json:"pageSize"`
	Pages    int `json:"pages"`
	Total    int `json:"total"`
}

type Movie struct {
	Id          int    `xorm:"id"`
	MovieId     int `xorm:"movie_id"`
	Title       string `xorm:"title"`
	Genres      string `xorm:"genres"`
	CreatedTime string `xorm:"created_time"`
	Status      int `xorm:"status"`
}

type Ratings struct {
	Id          int    `xorm:"id"`
	UserId      int `xorm:"user_id"`
	MovieId     int `xorm:"movie_id"`
	Rating      int `xorm:"rating"`
	CreatedTime time.Time `xorm:"created_time"`
	Status      int `xorm:"status"`
}
