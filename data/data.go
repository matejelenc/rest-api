package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//swagger:model user
type Person struct {
	// creates a unique id, created at, updated at and deleted at fields for this user
	//
	// required: false
	gorm.Model
	// the name for this user
	//
	// required: true
	Name string `json:"name" validate:"required" bson:"name,omitempty"`
	// the email for this user
	//
	// required: true
	Email string `json:"email" validate:"required" bson:"email,omitempty" gorm:"typevarchar(100);unique_index"`
	// the password for the user
	//
	// required: true
	// min: 4
	Password string `json:"password" validate:"required,gte=4" bson:"password,omitempty"`
	// the name of the group a user is in
	//
	// required: false
	GroupName string `json:"group name"`
}

//swagger:model group
type Group struct {
	// creates a unique id, created at, updated at and deleted at fields for this group
	//
	// required: false
	gorm.Model
	// the name for this group
	//
	// required: true
	Name string `json:"name" validate:"required" bson:"name,omitempty" gorm:"typevarchar(100);unique_index"`
}

//Groups defines a slice of Group
type Groups []*Group

//Users defines a slice of User
type Users []*Person

//error raised when a group is not found
var ErrGroupNotFound = fmt.Errorf("Group not found")

//ValidateGroup validates a group
func (group *Group) ValidateGroup() error {
	validate := validator.New()
	return validate.Struct(group)

}

var email_reg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//ValidateUser validates a user
func (user *Person) ValidateUser() error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}

	if !email_reg.MatchString(user.Email) {
		return fmt.Errorf("Invalid email")
	}

	return nil
}

//ValidateEmail validates an email
func ValidateEmail(email string) error {
	if !email_reg.MatchString(email) {
		return fmt.Errorf("Invalid email")
	}

	return nil
}
