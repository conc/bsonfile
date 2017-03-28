package bsonsplit

import (
	"testing"
)

func Test_BsonSplit(t *testing.T) {
	res, err := BsonSplit("../testfile/test.bson", 10)
	if err != nil {
		t.Error("err new", err)
		return
	}
	t.Log(res)
}
