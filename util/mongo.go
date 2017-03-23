package util

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	mgoUri    string = "mongodb"
	mgoDbName string = "user-svc"
)

type myMongo mgo.Session

var mgoCollection *mgo.Collection

func GetMgoCollection(mgoCollName string) *mgo.Collection {
	if mgoCollection == nil {
		mgoSession, err := mgo.Dial(mgoUri)
		if err != nil {
			panic(err)
		}
		session := mgoSession.Clone()
		mgoCollection = session.DB(mgoDbName).C(mgoCollName)
		setUserUnique()
	}
	return mgoCollection
}

func setUserUnique() {
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := mgoCollection.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
}
