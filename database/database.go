package database

import (
	"fmt"
	"hash/fnv"
	"log"

	"gopkg.in/mgo.v2"
)

var databaseName string
var connectionString string

var session *mgo.Session
var err error

//Init creates a connection to the database
func Init(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring + dbName

	session, err = mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
}

//SetCollection sets the collection
func SetCollection(collection string) *mgo.Collection {
	return session.DB(databaseName).C(collection)
}

//Insert allows users to add generic objects to a collection in the database
func Insert(collection string, object interface{}) bool {
	c := SetCollection(collection)
	err := c.Insert(object)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Update allows users to add generic objects to a collection in the database
func Update(collection string, selector interface{}, object interface{}) bool {
	c := SetCollection(collection)
	err := c.Update(selector, object)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//GetAll returns an array of all objects in a collection
func GetAll(collection string) []map[string]string {
	c := SetCollection(collection)
	list := make([]map[string]string, 1000)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//RemoveAll will empty a collection
func RemoveAll(collection string) bool {
	c := SetCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//Hash hashes a string
func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}
