package bsonread

import (
	"os"
	"testing"
)

func Test_BsonRead(t *testing.T) {
	osFile, err := os.Open("../testfile/test.bson")
	if err != nil {
		t.Error("err new", err)
		return
	}
	defer osFile.Close()
	bs := NewBsonReader(osFile)

	for {
		var doc interface{}
		err := bs.Next(&doc)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				t.Error("err new", err)
				return
			}
		}
		t.Log(doc)
	}
}
