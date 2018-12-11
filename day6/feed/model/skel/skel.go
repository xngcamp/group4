/*
Package feed just for demo
*/
package skel

import (
	mgo "github.com/globalsign/mgo"
	"camp/feed/api"
	"camp/feed/mongo"
)

// Feed 定义db对应的类型
type Skel api.Feed

// Db 返回db name
func (skel *Skel) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (skel *Skel) Table() (table string) {
	return "feed"
}

// GetC 返回db col
func (skel *Skel) GetC() (c *mgo.Collection) {
	db, table := skel.Db(), skel.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
