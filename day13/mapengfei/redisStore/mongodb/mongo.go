package mongodb

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type MovieLib struct {
	Key      string
	Id       int
	Title    string
	Year     int
	Author   string
	Star     float64
	PriceCode int
}

func (mv *MovieLib) Db()(db string){
	return "skel"
}
func (mv *MovieLib) Table() (table string){
	return "movie"
}

func (mv *MovieLib) GetC() (c *mgo.Collection) { //模板方法

	db, table := mv.Db(), mv.Table()
//	session := mongo.DBS[db]
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}

func (mv *MovieLib) Add()(err error){
	c := mv.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(mv)
	if err != nil {
		return
	}

	return
}
func Exists(m string) (title string, pri int, err error) {
	var mv MovieLib
	c := mv.GetC()
	defer c.Database.Session.Close()
	var ret *MovieLib
	err = c.Find(bson.M{"key": m}).One(&ret)
	if err != nil {
		if err != mgo.ErrNotFound {
			log.Fatal(err)
		}
		fmt.Println("Not find in redis or mongodb")
	}
	err = nil
	title = ret.Title
	pri = ret.PriceCode

	return
}

func CreatMvMongo(){
	mv := MovieLib{
		Key: "mv6",
		Id:  6,
		Title:    "XNG666",
		Year:     2018,
		Author:   "Li",
		Star:     4.9,
		PriceCode: 0,
	}
	err := mv.Add()
	if err!=nil {
		fmt.Println("Store in mongo failed")
	}
}

