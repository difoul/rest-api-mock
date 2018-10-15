package helpers

import (
	"io/ioutil"
	"os"
)

func IsFileOrDirExists(file_path string) bool {
	if _, err := os.Stat(file_path); err == nil {
		return true
	}
	return false
}

func ReadFile(fpath string) []byte {
	dat, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil
	}
	return dat
}
