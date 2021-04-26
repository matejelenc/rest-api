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
	for i, p := range userList {
		if p.ID == id {
			userList = append(userList[:i], userList[i+1])
			return nil
		}
	}

	return ErrUserNotFound
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
