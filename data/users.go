package data

import (
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	GroupID  int    `json:"group id"`
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
	u.ID = userList[len(userList)-1].ID + 1
	userList = append(userList, u)
}

func UpdateUser(id int, u *User) error {
	for _, user := range userList {
		if user.ID == id {
			u.ID = id
			userList[id] = u
			return nil
		}
	}
	return ErrUserNotFound

}

func DeleteUser(id int) error {
	users := Users{}
	exist := false
	for _, p := range userList {
		if p.ID != id {
			users = append(users, p)
		} else {
			exist = true
		}
	}

	userList = users
	if !exist {
		return ErrUserNotFound
	}

	exist = false
	return nil

}

var userList = Users{
	&User{
		Name:     "Baži",
		Email:    "b@gmail.com",
		Password: "b123",
		GroupID:  0,
		ID:       0,
	},
	&User{
		Name:     "Kiki",
		Email:    "k@gmail.com",
		Password: "k123",
		GroupID:  0,
		ID:       1,
	},
}
