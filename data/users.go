package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
)

//swagger:model user
type User struct {
	// the id for the user
	//
	// required: false
	// min: 0
	ID int `json:"id"`
	// the name for this user
	//
	// required: true
	// max length: 50
	Name string `json:"name" validate:"required"`
	// the email for this user
	//
	// required: true
	Email string `json:"email" validate:"required,email"`
	// the password for the user
	//
	// required: true
	// min: 4
	Password string `json:"password" validate:"required,gte=4"`
	//whether a user is in a group or not
	//
	//required: false
	InGroup bool `json:"in group"`
	//users group id
	//
	//required: false
	GroupID int `json:"group id"`
}

//Users defines a slice of User
type Users []*User

//GetUsers returns all users from the database
func GetUsers() Users {
	return userList
}

//ValidateUser validates a user
func (user *User) ValidateUser() error {
	validate := validator.New()
	return validate.Struct(user)

}

//error which is raised when a user does not exist
var ErrUserNotFound = fmt.Errorf("User not found")

//GetUser returns a user and an error if the user does not exist
func GetUser(id int) (*User, int, error) {
	for i, u := range userList {
		if u.ID == id {
			return u, i, nil
		}
	}

	return nil, -1, ErrUserNotFound
}

//CreateUser creates a user
func CreateUser(u *User) {
	u.ID = userList[len(userList)-1].ID + 1
	userList = append(userList, u)

	//adds the user to a group if it is specified
	if u.InGroup {
		AddToGroup(u)
	}
}

//UpdateUser updates a user and returns an error if the user does not exist
func UpdateUser(id int, u *User) error {

	for i, user := range userList {
		if user.ID == id {
			u.ID = id
			userList[i] = u

			//removes the outdated user from a group
			if user.InGroup {
				RemoveFromGroup(user)
			}

			//adds the updated user to a group
			if u.InGroup {
				AddToGroup(u)
			}

			return nil
		}
	}

	return ErrUserNotFound

}

//DeleteUser deletes a user
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

	//error is raised if the user does not exist
	userList = users
	if !exist {
		return ErrUserNotFound
	}

	exist = false
	return nil

}

//UsersFromJSON deserializes a user from json string
func (u *User) UserFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

//UsersFromJSON serializes a user to json string
func (u *User) UserToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

//UsersFromJSON deserializes users from json string
func (u *Users) UsersFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

//UsersFromJSON serializes users to json string
func (u *Users) UsersToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

//test variable simulating a database for now
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
