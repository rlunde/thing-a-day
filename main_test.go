package main

import (
	"testing"
)

func Test_randRecords(t *testing.T) {
	type args struct {
		collection string
		numRecords int
	}
	test1 := args{collection: "predictions", numRecords: 5}
	gotRecords, err := randRecords(test1.collection, test1.numRecords)
	if err != nil {
		t.Errorf("randRecords() unexpected error = %v", err)
		return
	}
	if len(gotRecords) != 5 {
		t.Errorf("randRecords() returned %d records, expected %d", len(gotRecords), test1.numRecords)
	}

}
