package redis

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	redisAddress string
	pool *redis.Pool
)



func init(){
	flag.StringVar(&redisAddress, "s","127.0.0.1:6379","redis server address")
	flag.Parse()

	poolSize := 20
	pool = &redis.Pool{
		MaxIdle: poolSize,
		IdleTimeout: time.Minute,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", redisAddress)
			if err != nil{
				return  nil, err
			}
			return conn, nil
		},
	}
}

func GetRedisConn() redis.Conn {
	return  pool.Get()
}

func CreatMvLib(){
	conn := GetRedisConn()
	//fmt.Printf("conn=%p, %v\n", conn, conn)
	//记得销毁本次链连接
	defer conn.Close()

	//	向redis中写入每条电影记录并限制redis中List长度
	_, err := conn.Do("RPUSH", "mv1", "id", 1, "Title", "XNG111", "Year", 2014,
			"Author", "Ding", "Star", 4.5, "PriceCode", 0)
	if err != nil {
		fmt.Println("mv1 stored failed")
	}
	_, err = conn.Do("LTRIM", "mv1", 0, 11)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("RPUSH", "mv2", "id", 2, "Title", "XNG222", "Year", 2015,
		"Author", "Lin", "Star", 4.6, "PriceCode", 2)
	if err != nil {
		fmt.Println("mv2 stored failed")
	}
	_, err = conn.Do("LTRIM", "mv2", 0, 11)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("RPUSH", "mv3", "id", 3, "Title", "XNG333", "Year", 2016,
		"Author", "Wang", "Star", 4.7, "PriceCode", 2)
	if err != nil {
		fmt.Println("mv3 stored failed")
	}
	_, err = conn.Do("LTRIM", "mv3", 0, 11)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("RPUSH", "mv4", "id", 4, "Title", "XNG444", "Year", 2017,
		"Author", "Yang", "Star", 4.8, "PriceCode", 1)
	if err != nil {
		fmt.Println("mv4 stored failed")
	}
	_, err = conn.Do("LTRIM", "mv4", 0, 11)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("RPUSH", "mv5", "id", 5, "Title", "XNG555", "Year", 2018,
		"Author", "Zhao", "Star", 4.9, "PriceCode", 1)
	if err != nil {
		fmt.Println("mv5 stored failed")
	}
	_, err = conn.Do("LTRIM", "mv5", 0, 11)
	if err != nil {
		log.Fatal(err)
	}




}
