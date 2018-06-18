package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

}

/*GetUserDailyRecords -- get previously selected records from a collection for a user for a day
 * see if the given user already has a set of records for this day, and return if so
 */
func GetUserDailyRecords(collection, fieldname, user, day string) (records []string, err error) {
	session := GetSession()
	session.SetMode(mgo.Monotonic, true)
	historyCollection := collection + ".history"

	db := session.DB("thing-a-day")
	c := db.C(historyCollection)

	query := c.Find([]bson.M{{"field": fieldname, "user": user}})
	result := bson.M{}
	err = query.Limit(1).One(&result)
	//TODO: figure out how to get records from result
	thing := fmt.Sprintf("result: %#v", result)
	// I think I need to use bson.Unmarshal ? https://godoc.org/labix.org/v2/mgo/bson#Unmarshal
	records = make([]string, 1)
	records[0] = thing
	return records, err
}
