package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var thingSession *mgo.Session

func init() {

}

//StartSession -- connect to mongo running in a container:
// docker run -d -p 37017:27017 --rm --name tadmongo -v $MONGO_DATA_PATH:/data/db mongo
func StartSession() error {
	var err error
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"localhost:37017"},
		Direct:    true,
		Timeout:   10 * time.Second,
		Service:   "localhost:37017",
		PoolLimit: 0,
	}
	thingSession, err = mgo.DialWithInfo(dialInfo)
	// thingSession, err = mgo.Dial("172.17.0.2:27017")
	if err != nil {
		log.Fatalf("mgo.Dial returned error %s", err)
	}
	return err
}

//EndSession -- disconnect from mongo
func EndSession() {
	thingSession.Close()
	thingSession = nil
}

//GetSession -- return the mongo session
func GetSession() *mgo.Session {
	return thingSession
}

//GetRandRecords -- get N randomly selected records from a collection
func GetRandRecords(collection string, fieldName string, numRecords int) (records []string, err error) {
	session := GetSession()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("thing-a-day")
	c := db.C(collection)
	pipe := c.Pipe([]bson.M{{"$sample": bson.M{"size": numRecords}}})
	resp := []bson.M{}
	err = pipe.All(&resp)
	var rval []string
	rval = make([]string, 0, numRecords)
	if err != nil {
		log.Printf("error sampling %s: %s", collection, err.Error())
	} else {
		for _, r := range resp {
			//id := r["_id"].(bson.ObjectId)
			field := r[fieldName].(string)
			//fmt.Printf("id: %s, %s: %s\n", id.Hex(), fieldName, field)
			rval = append(rval, field)
		}
	}

	return rval, nil
}
