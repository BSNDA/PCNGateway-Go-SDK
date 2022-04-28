package keystore

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/file"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"io/ioutil"
	"os"
	path2 "path"
	"strings"
)

func NewUserCertStore(path string) UserCertStore {

	file.CheckDir(path)

	us := &FileUserCertStore{
		FilePath: path,
	}

	return us
}

type FileUserCertStore struct {
	FilePath string
}

func (f *FileUserCertStore) Load(user *msp.UserData) error {
	key := storeKeyName(user)
	filePath := path2.Join(f.FilePath, key)

	if _, err1 := os.Stat(filePath); os.IsNotExist(err1) {
		return errors.New("user not found")
	}

	bytes, err := ioutil.ReadFile(filePath) // nolint: gas
	if err != nil {
		return err
	}
	if bytes == nil {
		return errors.New("user not found")
	}
	user.EnrollmentCertificate = bytes
	return nil

}

func (f *FileUserCertStore) Store(user *msp.UserData) error {
	key := storeKeyName(user)

	path := path2.Join(f.FilePath, key)

	valueBytes := user.EnrollmentCertificate

	err := os.MkdirAll(path2.Dir(path), 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, valueBytes, 0600)

}

func (f *FileUserCertStore) LoadAll(appCode string) []*msp.UserData {

	var users []*msp.UserData

	//Traverse files under the folder
	files, err := ioutil.ReadDir(f.FilePath)
	if err != nil {
		return users
	}

	for _, file := range files {
		filePath := path2.Join(f.FilePath, file.Name())

		//get the file name
		name := getPemName(file.Name(), appCode)
		if name != "" {
			//get
			user := &msp.UserData{}
			bytes, err := ioutil.ReadFile(filePath) // nolint: gas
			if err == nil && bytes != nil {
				user.EnrollmentCertificate = bytes
				user.UserName = name
				user.AppCode = appCode
				users = append(users, user)
			}
		}
	}

	return users

}

func getPemName(name, appCode string) string {

	ext := "@" + appCode + "-cert.pem"

	i := strings.Index(name, ext)
	if i != -1 {
		return name[:i]
	} else {
		return ""
	}
}

func storeKeyName(user *msp.UserData) string {
	return user.UserName + "@" + user.AppCode + "-cert.pem"
}
