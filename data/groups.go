package data

import "fmt"

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
