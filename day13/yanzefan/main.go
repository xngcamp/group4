package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"xng.com/camp/chap7/customer"
	"xng.com/camp/chap7/movie"
	redis1 "xng.com/camp/chap7/redis"
	"xng.com/camp/chap7/rental"
)

/*
目标：
给一个录影带商店编写一个计算和打印顾客租赁录影带的收费和积分详单


需求：
1. 计算顾客租赁的电影清单和每部电影的租赁费用
2. 计算总的租赁费用
3. 根据电影类型，计算用户积分


有三种类型的电影：普通电影(类型0)，儿童电影(类型2)和新发行的电影(类型1)
收费规则：
1.普通电影：前2天共2元，超过2天后，每天1.5元
2.新发行的电影：每天3元
3.儿童电影：前3天共1.5元，超过3天后，每天1.5元

积分规则：
每租一部电影得1积分
新上市电影：租期大于1天，额外得1积分

结果示例：

Rental Record for Tom:
	1. The Godfather	6.50
	2. Fast & Furious	3.00
Amount owned is 9.50
Tom earned 2 frequent renter points.

备注：
videostore1.pos  使用 https://www.processon.com 浏览

*/

const key  = "movieList"

func main() {
	InitMovie()
	aCustomer := addRental()
	result := aCustomer.Statement()
	fmt.Println(result)
}

func InitMovie() {
	conn := redis1.GetRedisConn()
	defer conn.Close()

	movie1 := new(movie.Movie)
	movie2 := new(movie.Movie)
	movie1.Init(1,"The Godfather",movie.REGULAR,"2016","StevenJack",1000)
	movie2.Init(2,"Fast & Furious",movie.NEW_RELEASE,"2009","yanzefan",503)

	value1, _ := json.Marshal(movie1)
	value2, _ := json.Marshal(movie2)

	if _,err := conn.Do("SET", movie1.Id, string(value1)) ; err != nil {
		log.Fatal(err)
	}
	if _,err := conn.Do("SET", movie2.Id, string(value2)) ; err != nil {
		log.Fatal(err)
	}
}

func addRental() *customer.Customer{
	conn := redis1.GetRedisConn()
	defer conn.Close()

	aCustomer := new(customer.Customer)
	aCustomer.Init("Tom")

	objStr1,_ := redis.String(conn.Do("GET", 1))
	objStr2,_ := redis.String(conn.Do("GET", 2))

	temp1 := getMovie(objStr1)
	temp2 := getMovie(objStr2)

	rental1 := rental.Rental{AMovie: movie.Movie{Title: temp1.Title, PriceCode: temp1.PriceCode,}, DaysRented: 5,}
	aCustomer.AddRental(rental1)

	rental2 := rental.Rental{AMovie: movie.Movie{Title: temp2.Title, PriceCode: temp2.PriceCode,}, DaysRented: 1,}
	aCustomer.AddRental(rental2)
	return aCustomer
}

func getMovie(objStr string) *movie.Movie{
	b := []byte(objStr)
	temp := &movie.Movie{}
	if err := json.Unmarshal(b, temp) ;err != nil {
		log.Fatal(err)
	}
	return temp
}


