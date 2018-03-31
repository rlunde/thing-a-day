package main

import (
	"testing"
)

func Test_randRecords(t *testing.T) {
	type args struct {
		collection string
		fieldName  string
		numRecords int
	}
	StartSession()
	test1 := args{collection: "predictions", fieldName: "prediction", numRecords: 5}
	gotRecords, err := GetRandRecords(test1.collection, test1.fieldName, test1.numRecords)
	if err != nil {
		t.Errorf("randRecords() unexpected error = %v", err)
		return
	}
	if len(gotRecords) != 5 {
		t.Errorf("randRecords() returned %d records, expected %d", len(gotRecords), test1.numRecords)
	}

}
