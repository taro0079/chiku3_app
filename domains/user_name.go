package domains

import (
	"time"
)

type UserName struct {
	Name string
}

func NewUserName(name string) (*UserName, error) {
	if len(name) < 3 {
		return nil, UserNameError{time.Now(), "error"}
	}

	username := new(UserName)
	username.Name = name

	return username, nil
}
