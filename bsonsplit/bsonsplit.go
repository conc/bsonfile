package bsonsplit

import (
	"io"
	"os"
)

//return offset list
func BsonSplit(filePath string, splitNum int64) ([]int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	var (
		res           []int64
		expectStep    int64 = fileSize / int64(splitNum)
		currentOffset int64 = 0
		usedOffset    int64 = 0
	)

	for {
		dataLen := make([]byte, 4)
		_, err := io.ReadFull(f, dataLen)
		if err != nil {
			return nil, err
		}
		dataSize := int64(dataLen[0]) + int64(dataLen[1])*256 + int64(dataLen[2])*256*256 + int64(dataLen[3])*256*256*256
		usedOffset += dataSize
		currentOffset += dataSize
		if usedOffset >= expectStep {
			res = append(res, currentOffset)
			usedOffset = 0
		}
		if currentOffset >= fileSize {
			break
		}

		_, err = f.Seek(dataSize-4, 1)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
