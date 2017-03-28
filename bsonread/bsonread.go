package bsonread

import (
	"io"

	"gopkg.in/mgo.v2/bson"
)

type BsonReader struct {
	ioReader io.Reader
}

func NewBsonReader(ioReader io.Reader) *BsonReader {
	return &BsonReader{
		ioReader: ioReader,
	}
}

func (bs *BsonReader) Next(doc interface{}) error {
	dataLen := make([]byte, 4)
	_, err := io.ReadFull(bs.ioReader, dataLen)
	if err != nil {
		return err
	}
	size := int32(dataLen[0]) + int32(dataLen[1])*256 + int32(dataLen[2])*256*256 + int32(dataLen[3])*256*256*256

	fileContent := make([]byte, size)
	_, err = io.ReadFull(bs.ioReader, fileContent[4:])
	if err != nil {
		return err
	}
	copy(fileContent, dataLen)

	if err := bson.Unmarshal(fileContent, doc); err != nil {
		return err
	}
	return nil
}
