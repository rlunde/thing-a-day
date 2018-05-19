package main

import (
	"gopkg.in/mgo.v2"
)

func init() {

}

//UserDailyRecordsKey -- we store the data for a field/user/day using this key
//TODO: decide if we need to know the number of records here -- if we didn't have
//the same number when we stored it as we want to return, it will overcomplicate
//things...I think
type UserDailyRecordsKey struct {
	Field string
	User  string
	Day   string // yyyy-mm-dd
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
	key := UserDailyRecordsKey{
		Field: fieldname,
		User:  user,
		Day:   day,
	}
	query := c.Find(key)
	result := &records
	err = query.Limit(1).One(result)

	return records, err
}
