package main

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

}

//PreviousUserDailyRecord -- store the day and ID of previously used records
type PreviousUserDailyRecord struct {
	Field   string    `bson:"field"`
	User    string    `bson:"user"`
	Date    time.Time `bson:"date"`
	Records []string  `bson:"records"`
}

/*GetRandRecordsForUser -- get all previously selected records from a collection for a user,
 * and then order them by day. Delete all the history records that are older than N days.
 * Then use the remaining history records to filter out random new records until we have nr
 * new ones. Then save new history records and return the nr records.
 */
func GetRandRecordsForUser(collection, fieldname, user, day string, nr int) (records []string, err error) {
	session := GetSession()
	session.SetMode(mgo.Monotonic, true)
	// for every collection "foo" we have a "foo.history" collection
	historyCollection := collection + ".history"

	/*
	 * If we need NR new unique records per day, and we have R total records in the collection,
	 * then we will run out if (numHistoryDays + 1) * NR > R. Let's be conservative and say that
	 * we want to have 50% of them to re-use if possible. If not possible, then at least two days
	 * worth.
	 */

	db := session.DB("thing-a-day")
	col := db.C(collection)
	colhist := db.C(historyCollection)
	// r is the count of records in the collection
	// this assumes that the field we want is in every record of the collection
	r, err := col.Find(bson.M{}).Count()
	want := (r / 2)
	minNeeded := nr * 2
	if minNeeded > want {
		want = minNeeded
	}
	if want > r {
		msg := fmt.Sprintf("Need at least %d records of type %s, but only have %d", want, fieldname, r)
		err = errors.New(msg)
		return
	}

	//historyDays is all the records in history for that field, for that user
	historyDays := colhist.Find([]bson.M{{"field": fieldname, "user": user, "$sort": bson.M{"date": 1}}})
	numHistoryDays, err := historyDays.Count()
	if r-(numHistoryDays*nr) > want {
		toDelete := ((r - want) / nr) + 1
		// have to delete the oldest
		// TODO: sort by age and delete (want / nr) + 1 days
		// sorting: https://docs.mongodb.com/manual/reference/operator/aggregation/sort/
		// date type: https://play.golang.org/p/A0n6DGBAqt
		//            https://github.com/go-mgo/mgo/blob/v2-unstable/bson/encode.go#L43-L57
		// it looks like if I just put a time.Time struct in the mgo key or value, it will be
		// converted into a mongodb time. So maybe something like:
		// delete the earliest N records
		iter := historyDays.Iter()
		deleted := 0
		var result PreviousUserDailyRecord
		for {
			if deleted >= toDelete {
				break
			}
			deleted++
			if iter.Next(&result) {
				//TODO: delete this from history rather than print it out
				fmt.Printf("Result: %v\n", result.Records)
			} else { // oops! ran out unexpectedly
				break
			}

		}
		if err := iter.Close(); err != nil {
			return nil, err
		}
		// TODO
	}
	// initialize random collection
	// TODO: loop, getting random items, compare for uniqueness, and make sure they're not in history

	return records, err
}
