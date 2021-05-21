package models

import (
	"errors"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []User{
	User{Username: "user1", Password: "pass1"},
	User{Username: "user2", Password: "pass2"},
	User{Username: "user3", Password: "pass3"},
}

func IsUserValid(username string, password string) bool {
	for _, user := range userList {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func RegisterNewUser(username string, password string) (*User, error) {
	if strings.TrimSpace(password) == "" || strings.TrimSpace(username) == "" {
		return nil, errors.New("username and password cannot be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("username is not available")
	} else {
		userList = append(userList, User{Username: username, Password: password})
		return &userList[len(userList)-1], nil
	}
}

func isUsernameAvailable(username string) bool {
	for _, user := range userList {
		if user.Username == username {
			return false
		}
	}
	return true
}
