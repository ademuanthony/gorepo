package mongo

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session


func Connect(server string) (err error) {
	session, err = mgo.Dial(server)
	return
}

func NewDb(databaseName string) *mgo.Database {
	newSession := session.Copy()
	db := newSession.DB(databaseName)
	return db
}
