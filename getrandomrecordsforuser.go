package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

}

//PreviousUserDailyRecordsKey -- store the day and ID of previously used records
type PreviousUserDailyRecordsKey struct {
	Field string `bson:"field"`
	User  string `bson:"user"`
}

/*GetRandRecordsForUser -- get all previously selected records from a collection for a user,
 * and then order them by day. Delete all the history records that are older than N days.
 * Then use the remaining history records to filter out random new records until we have nr
 * new ones. Then save new history records and return the nr records.
 */
func GetRandRecordsForUser(collection, fieldname, user, day string, nr int) (records []string, err error) {
	session := GetSession()
	session.SetMode(mgo.Monotonic, true)
	historyCollection := collection + ".history"

	/*
	 * If we need NR new unique records per day, and we have R total records in the collection,
	 * then we will run out if (numHistoryDays + 1) * NR > R. Let's be conservative and say that
	 * we want to have 50% of them to re-use if possible. If not possible, then at least two days
	 * worth.
	 */

	db := session.DB("thing-a-day")
	col := db.C(collection)
	c := db.C(historyCollection)
	// count the number of records in the collection
	r, err := col.Find(bson.M{}).Count()
	want := (r / 2)
	minNeeded := nr * 2
	if minNeeded > want {
		want = minNeeded
	}
	key := UserDailyRecordsKey{
		Field: fieldname,
		User:  user,
	}
	historyDays := c.Find(key)
	numHistoryDays, err := historyDays.Count()
	if (numHistoryDays * nr) > want {
		// have to delete the oldest
		// TODO: sort by age and delete (want / nr) + 1 days
	}
	// initialize random collection
	// TODO: loop, getting random items, compare for uniqueness, and make sure they're not in history

	return records, err
}
