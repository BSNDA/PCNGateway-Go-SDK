package file

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"io/ioutil"
	"os"
	path2 "path"
)

func ReadFile(path string) ([]byte, error) {
	if _, err1 := os.Stat(path); os.IsNotExist(err1) {
		return nil, errors.New("file not found")
	}

	bytes, err := ioutil.ReadFile(path) // nolint: gas
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, errors.New("file not found")
	}
	return bytes, nil

}

// WriteFile cover true 覆盖文件，false 检查是否已存在
func WriteFile(data []byte, path string, cover bool) error {

	if !cover {
		if _, err1 := os.Stat(path); !os.IsNotExist(err1) {
			return errors.New("file is exist")
		}
	}

	err := os.MkdirAll(path2.Dir(path), 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0600)

}
