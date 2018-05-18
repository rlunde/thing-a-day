package main

import (
	"gopkg.in/mgo.v2"
)

func init() {

}

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

	db := session.DB("thing-a-day")
	c := db.C(collection)
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
