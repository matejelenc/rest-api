package data

import (
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}

type Users []*User

func GetUsers() Users {
	return userList
}

var ErrUserNotFound = fmt.Errorf("User not found")

func GetUser(id int) (*User, int, error) {
	for i, u := range userList {
		if u.ID == id {
			return u, i, nil
		}
	}

	return nil, -1, ErrUserNotFound
}

func CreateUser(u *User) {
	u.ID = len(userList)
	userList = append(userList, u)
}

var userList = Users{
	&User{
		Name:     "Baži",
		Email:    "b@gmail.com",
		Password: "b123",
		ID:       0,
	},
	&User{
		Name:     "Kiki",
		Email:    "k@gmail.com",
		Password: "k123",
		ID:       1,
	},
}
