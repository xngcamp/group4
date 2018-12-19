package movie

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func MovieToRedis(id string, title string, year int, author string, star float64, priceCode int) {
	redisAddr := "127.0.0.1:6379"
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		fmt.Println("connect to redis error:", err)
		return
	}
	defer c.Close()

	key := id
	data := map[string]interface{}{"title": title, "year": year, "author": author, "star": star, "priceCode": priceCode}
	value, _ := json.Marshal(data)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	if n == int64(1) {
		fmt.Println("set success")
	}
}
