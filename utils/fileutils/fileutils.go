package fileutils

import (
	"io/ioutil"
	"os"
	"urlmonitor/utils/errorutils"
)

func OpenFile(filepath string, flag int, perm os.FileMode) *os.File {
	file, err := os.OpenFile(filepath, flag, perm)
	errorutils.CheckError(err)

	return file
}

func ReadFile(filepath string) []byte {
	bytes, err := ioutil.ReadFile(filepath)
	errorutils.CheckError(err)

	return bytes
}

func CloseFile(file *os.File) {
	err := file.Close()
	errorutils.CheckError(err)
}

func WriteStringToFile(file *os.File, str string) {
	_, err := file.WriteString(str)
	errorutils.CheckError(err)
}
