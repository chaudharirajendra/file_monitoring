package file

import (
	"os"
)

func GetFileSize(filename string) (int64, error) {
	fileInfo, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return int64(len(fileInfo)), nil
}
