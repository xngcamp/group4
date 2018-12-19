package redis

import (
	"day/day13/refactoring/videostore5/movie"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

var (
	redisAddress string
	pool         *redis.Pool
)

func init() {
	flag.StringVar(&redisAddress, "s", "127.0.0.1:6379", "redis server address")
	flag.Parse()

	poolSize := 20
	pool = &redis.Pool{
		MaxIdle:     poolSize,
		IdleTimeout: time.Minute,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}

func AddUserLogin(key string,value string) (err error){
	conn:=GetRedisConn()
	defer conn.Close()
	value=key+"login:"+value
	_, err = conn.Do("LPUSH", key,value )
	if err != nil {
		return
	}
	return
}
func InitMovieToRedis(movie1 []*movie.AddMovie) (err error) {
	conn:=GetRedisConn()
	defer conn.Close()
	addMovieToRedis(movie1[0],conn)
	addMovieToRedis(movie1[1],conn)
	addMovieToRedis(movie1[2],conn)
	return
}

func addMovieToRedis(movie1 *movie.AddMovie,conn redis.Conn) (err error) {

	if err!=nil{
		return
	}
	_,err=conn.Do("DEL","movie"+strconv.Itoa(movie1.ID))
	if err!=nil {
		return
	}
	_,err=conn.Do("SET","movie"+strconv.Itoa(movie1.ID),movie1.Title+"&&&"+strconv.Itoa(movie1.PriceCode))
	if err!=nil {

		fmt.Println(err)
		return
	}
	return
}

func GetMovieFormRedis(movieID int) (err error,value []string ) {
	conn:=GetRedisConn()
	defer conn.Close()
	value1, err := redis.String(conn.Do("GET", "movie"+strconv.Itoa(movieID)))
	value=strings.Split(value1,"&&&")
	fmt.Printf("value:%v \n",value)
	if err!=nil {
		return
	}
	return
}
