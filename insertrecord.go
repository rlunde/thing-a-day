package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

}

//InsertRecord -- insert a new record into a collection and return the generated ID
func InsertRecord(collection string, fieldName string, record string) (id string, err error) {
	session := GetSession()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("thing-a-day")
	c := db.C(collection)
	objectID := bson.NewObjectId()
	err = c.Insert(bson.M{"_id": objectID, fieldName: record})
	if err != nil {
		return "", err
	}

	return objectID.Hex(), nil
}
