package movie

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func GetMovieInfo(id string) (movie Movie, err error){
	redisAddr := "127.0.0.1:6379"
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		fmt.Println("connect to redis error:", err)
		return
	}
	defer c.Close()

	key := id

	var dataGet = make(map[string]interface{})
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Printf("redispt get %s failed:%s\n", key, err)
	}

	errShal := json.Unmarshal(valueGet, &dataGet)
	if errShal != nil {
		fmt.Println(errShal)
	}

	movie.title = dataGet["title"].(string)
	movie.author = dataGet["author"].(string)
	movie.star = dataGet["star"].(float64)
	movie.year = int(dataGet["year"].(float64))
	code :=  int(dataGet["priceCode"].(float64))
	movie.SetPriceCode(code)

	return
}
