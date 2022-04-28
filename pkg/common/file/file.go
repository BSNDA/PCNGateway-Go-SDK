package file

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	path2 "path"
	"strings"
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

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的指定后缀的文件,不包含子目录下的文件
func GetAllFiles(dirPth, fix string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathListSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if !fi.IsDir() {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), fix)
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, nil
}

func CheckDir(path string) {
	if _, err1 := os.Stat(path); !os.IsNotExist(err1) {
		return
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logger.Warn("path [%s] create failed : %s", path, err.Error())
	}

}
