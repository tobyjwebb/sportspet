package user_service

import "fmt"

type UserServiceMock struct {
	LoginFn func(nick string) (sessionID string, err error)
}

func (usm *UserServiceMock) Login(nick string) (sessionID string, err error) {
	if usm.LoginFn != nil {
		return usm.LoginFn(nick)
	}
	return "", fmt.Errorf("mock login not implemented")
}
