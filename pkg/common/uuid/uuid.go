package uuid

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetUUID() string {

	u1 := uuid.NewV4()
	return strings.Replace(u1.String(), "-", "", -1)

}
