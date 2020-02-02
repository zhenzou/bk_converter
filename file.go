package bk_converter

import (
	"gopkg.in/errgo.v2/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func curFile(addLevel int) string {
	var filename string
	var level = -1
	for i := 0; i < 20; i++ {
		_, filename, _, _ = runtime.Caller(i)
		if strings.HasSuffix(filename, "file.go") {
			level = i + 1
			break
		}
	}

	_, filename, _, _ = runtime.Caller(level + addLevel)
	return filename
}

func Open(name string) (*os.File, error) {
	file, err := os.Open(name)
	if os.IsNotExist(err) {
		curDir := curFile(1)
		println(curDir)
		file, err = os.Open(filepath.Join(filepath.Dir(curDir), name))
	}
	return file, err
}

func ReadAll(name string) ([]byte, error) {
	file, err := os.Open(name)
	if os.IsNotExist(err) {
		curDir := curFile(1)
		println(curDir)
		file, err = os.Open(filepath.Join(filepath.Dir(curDir), name))
	}
	if err != nil {
		return nil, errors.Wrap(err)
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
