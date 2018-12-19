package init_movie

import (
	"day/day13/refactoring/videostore5/distributed"
	"day/day13/refactoring/videostore5/movie"
)

func InitMovie()  {
	movie1:=make([]*movie.AddMovie,10)
	movie1[0]=&movie.AddMovie{
		ID:1,
		Title:"Star Wars",
		Year:"1977",
		Author:"乔治·卢卡斯",
		Star:5,
		PriceCode:0,
	}
	movie1[1]=&movie.AddMovie{
		ID:2,
		Title:"The Godfather Part V",
		Year:"1977",
		Author:"弗朗西斯·福特·科波拉",
		Star:5,
		PriceCode:1,
	}
	movie1[2]=&movie.AddMovie{
		ID:3,
		Title:"Casablanca",
		Year:"1977",
		Author:"迈克尔·柯蒂斯",
		Star:5,
		PriceCode:2,
	}
	distributed.NewDistributed().InitMovieToRedis(movie1)
}
