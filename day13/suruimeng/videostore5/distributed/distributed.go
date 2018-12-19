package distributed

import (
	"day/day13/refactoring/videostore5/movie"
	"day/day13/refactoring/videostore5/redis"
)

type distributed struct {

}

func NewDistributed() *distributed   {
	return &distributed{}
}

func (*distributed) AddUserLogin(key string,value string) (err error) {
	err =redis.AddUserLogin(key,value)
	if err!=nil{
		return
	}
	return
}

func (*distributed) InitMovieToRedis(movie1 []*movie.AddMovie) (err error) {
	 err =redis.InitMovieToRedis(movie1)
	if err!=nil{
		return
	}
	return
}

func (*distributed) GetMovieFormRedis(movieID int) (err error,value []string) {
	err ,value =redis.GetMovieFormRedis(movieID)
	if err!=nil{
		return
	}
	return
}

