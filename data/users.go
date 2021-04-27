package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
)

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4"`
	InGroup  bool   `json:"in group"`
	GroupID  int    `json:"group id"`
	ID       int    `json:"-"`
}

type Users []*User

func GetUsers() Users {
	return userList
}

func (user *User) ValidateUser() error {
	validate := validator.New()
	return validate.Struct(user)

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

	if u.InGroup {
		AddToGroup(u)
	}
}

func UpdateUser(id int, u *User) error {

	for i, user := range userList {
		if user.ID == id {
			u.ID = id
			userList[i] = u

			if user.InGroup {
				RemoveFromGroup(user)
			}

			if u.InGroup {
				AddToGroup(u)
			}

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
			RemoveFromGroup(p)
		}
	}

	userList = users
	if !exist {
		return ErrUserNotFound
	}

	exist = false
	return nil

}

func (u *User) UserFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) UserToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Users) UsersFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *Users) UsersToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
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
