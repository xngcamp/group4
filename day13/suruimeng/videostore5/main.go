package main

import (
	"day/day13/refactoring/videostore5/customer"
	"day/day13/refactoring/videostore5/distributed"
	"day/day13/refactoring/videostore5/init_movie"
	"day/day13/refactoring/videostore5/movie"
	"day/day13/refactoring/videostore5/rental"
	"fmt"
	. "strconv"
	"time"
)

/*

重构：
1. 终于谈到继承(Inheritance)了
* 对于电影Movie，有不同的类型，看起来应该是继承关系，生成它的子类(subclass), 图5-1
* 我们可以用多态来替代switch-case

2. 带来的问题
* 电影Movie在它的生命周期内，可以改变类型的，比如新上市的电影(New_Release)可以变为普通电影(Regular)
* 但是，一个对象(Object)在它声明周期内却不能改变它的类

3. 状态模式(State pattern [Gang of Four])可以解决这个情况
* 状态模式定义：当一个对象内在状态改变时允许其改变行为，这个对象看起来像改变了其类。
* 使用场景
	行为随状态的改变而改变。
	如果需要使用大量的条件、分支判断。
* 图5-2


4. 通过引入间接层(Pricer)，对价格做子类化
* 如果你熟悉Gang of Four patterns， 你可能疑惑，这个是状态模式(State Pattern)，还是策略模式(Strategy Pattern)？
* 这个模式选择阶段，反映出你如何思考这个结构
* 这个时候，我认为它是一个电影的状态
* 如果以后我倾向于策略模式，只需要把名字改一下(Price/Pricer, or Pricing Strategy)


*/

func main() {
	init_movie.InitMovie()
	//aCustomer := new(customer.Customer)
	var aCustomer customer.Customer
	aCustomer.Init("Tom")
	err:=distributed.NewDistributed().AddUserLogin(aCustomer.GetName(),time.Now().Format("2006-01-02 15:04:05"))
	if err!=nil{
		fmt.Println(err)
	}
	movieSelice:=make([]string,10)
	var priceCode int
	var rental1 rental.Rental
	var movie1 movie.Movie
	err,movieSelice =distributed.NewDistributed().GetMovieFormRedis(1)
	if err!=nil{
		fmt.Println(err)
	}
	priceCode, err = Atoi(movieSelice[1])
	if err!=nil{
		fmt.Println(err)
	}
	movie1.Init(movieSelice[0], priceCode)
	rental1.Init(movie1, 3)
	aCustomer.AddRental(rental1)
	var rental2 rental.Rental
	var movie2 movie.Movie
	err,movieSelice=distributed.NewDistributed().GetMovieFormRedis(2)
	if err!=nil{
		fmt.Println(err)
	}
	priceCode, err = Atoi(movieSelice[1])
	if err!=nil{
		fmt.Println(err)
	}
	movie2.Init(movieSelice[0],priceCode)
	rental2.Init(movie2, 1)
	aCustomer.AddRental(rental2)

	var rental3 rental.Rental
	var movie3 movie.Movie
	err,movieSelice=distributed.NewDistributed().GetMovieFormRedis(3)
	if err!=nil{
		fmt.Println(err)
	}
	priceCode, err = Atoi(movieSelice[1])
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(priceCode)
	movie3.Init(movieSelice[0],priceCode)
	rental3.Init(movie3, 7)
	aCustomer.AddRental(rental3)

	result := aCustomer.Statement()
	fmt.Println(result)
}
