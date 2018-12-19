package main

import (
	"camp/day13/redisStore/mongodb"
	"camp/day13/redisStore/redis"
	"fmt"
	"camp/day13/redisStore/customer"
	//"gothinking/refactoring/videostore5/customer"
	//"gothinking/refactoring/videostore5/movie"
	//"gothinking/refactoring/videostore5/rental"
	"camp/day13/redisStore/movie"
	"camp/day13/redisStore/rental"
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
	//aCustomer := new(customer.Customer)
	var aCustomer customer.Customer
	aCustomer.Init("Tom")

	//	创建电影库  将数据写入redis(mv1-mv5),
	//	若查询redis返回nil,则去mongoDB中查找(mv6)
	redis.CreatMvLib()
	mongodb.CreatMvMongo()


	var rental1 rental.Rental
	var movie1 movie.Movie
	movie1.Init("mv1")                        //redis
	rental1.Init(movie1, 3)
	aCustomer.AddRental(rental1)

	var rental2 rental.Rental
	var movie2 movie.Movie
	movie2.Init("mv2")                       //redis
	rental2.Init(movie2, 1)
	aCustomer.AddRental(rental2)

	var rental3 rental.Rental
	var movie3 movie.Movie
	movie3.Init("mv6")                      //mongoDB
	rental3.Init(movie3, 7)
	aCustomer.AddRental(rental3)

	result := aCustomer.Statement()
	fmt.Println(result)
}
