package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
)

//swagger:model group
type Group struct {
	// the name for this group
	//
	// required: true
	// max length: 50
	Name string `json:"name" validate:"required"`
	// the id for the group
	//
	// required: false
	// min: 0
	ID int `json:"id"`
	//List of all users in this group
	//
	//required: false
	Members []*User `json:"members"`
}

//Groups defines a slice of Group
type Groups []*Group

//GetGroups returns all the groups from the database
func GetGroups() Groups {
	return groupList
}

//error raised when a group is not found
var ErrGroupNotFound = fmt.Errorf("Group not found")

//GetGroup returns a group by specified id
func GetGroup(id int) (*Group, error) {
	for _, g := range groupList {
		if g.ID == id {
			return g, nil
		}
	}

	return nil, ErrGroupNotFound
}

//ValidateGroup validates a group
func (group *Group) ValidateGroup() error {
	validate := validator.New()
	return validate.Struct(group)

}

//UpdateGroup updates a group and returns an error if the group does not exist
func UpdateGroup(id int, g *Group) error {
	for i, group := range groupList {
		if group.ID == id {
			g.ID = id
			groupList[i] = g
			return nil
		}
	}
	return ErrUserNotFound
}

//CreateGroup creates a group and adds to the database
func CreateGroup(g *Group) {
	g.ID = groupList[len(groupList)-1].ID + 1
	groupList = append(groupList, g)
}

//DeleteGroup deletes a group
func DeleteGroup(id int) error {
	groups := Groups{}
	exist := false
	for _, p := range groupList {
		if p.ID != id {
			groups = append(groups, p)
		} else {
			exist = true
		}
	}

	//returns a error if the group does not exist
	groupList = groups
	if !exist {
		return ErrUserNotFound
	}

	exist = false
	return nil

}

//AddToGroup adds a user to a group
func AddToGroup(u *User) {
	for _, g := range groupList {
		if g.ID == u.GroupID {
			g.Members = append(g.Members, u)
			return
		}
	}

}

//RemoveFromGroup removes a user from a group
func RemoveFromGroup(u *User) {
	for _, g := range groupList {
		if g.ID == u.GroupID {
			members := Users{}
			for _, m := range g.Members {
				if m.ID != u.ID {
					members = append(members, m)
					return
				}
			}

			g.Members = members
		}
	}
}

//GroupFromJSON deserializes a group from json string
func (g *Group) GroupFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(g)
}

//test variable simulating a database for now
var groupList = Groups{
	&Group{
		Name: "Slovenija",
		ID:   0,
	},
	&Group{
		Name: "Madžarska",
		ID:   1,
	},
}
