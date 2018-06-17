#!/bin/bash
# drop the collection when loading the first set of names (firstnames)
mongoimport -v -d thing-a-day -c names --port 27017 --drop --type csv --fields firstname --file firstnames.csv 
# don't drop the collection now
mongoimport -v -d thing-a-day -c names --port 27017 --type csv --fields lastname --file lastnames.csv 
