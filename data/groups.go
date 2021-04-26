package data

import (
	"fmt"
)

type Group struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	Members []int  `json:"members"`
}

type Groups []*Group

func GetGroups() Groups {
	return groupList
}

var ErrGroupNotFound = fmt.Errorf("Group not found")

func GetGroup(id int) (*Group, error) {
	for _, g := range groupList {
		if g.ID == id {
			return g, nil
		}
	}

	return nil, ErrGroupNotFound
}

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

func CreateGroup(g *Group) {
	g.ID = groupList[len(groupList)-1].ID + 1
	groupList = append(groupList, g)
}

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

	groupList = groups
	if !exist {
		return ErrUserNotFound
	}

	exist = false
	return nil

}

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
