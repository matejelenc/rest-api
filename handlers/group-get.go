package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route GET /groups groups getGroups
// Returns a list of groups
// responses:
//	200: groupsResponse
//  400: badRequestResponse

//GetGroups returns all the groups in the database
func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	//validate the jwt token
	_, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	//get all groups
	var groups []data.Group
	foundGroups := data.DB.Find(&groups)
	if foundGroups.Error != nil {
		http.Error(w, "No groups exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&groups)
}

// swagger:route GET /groups/{id} groups getGroup
// Returns a group
// responses:
//	200: groupResponse
//  400: badRequestResponse

//GetGroup returns a group
func GetGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	var group data.Group

	//get the group
	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(group)
}

// swagger:route GET /groups/{id}/members groups getMembers
// Returns members of a group
// responses:
//	200: membersResponse
//  400: badRequestResponse
// swagger:parameters updateGroup deleteGroup getGroup
//GetMembers returns members of a group
func GetMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var group data.Group
	var people []data.Person

	//get the group
	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	//get all mmbers in this group
	foundMembers := data.DB.Where("group_name = ?", group.Name).Find(&people)
	if foundMembers.Error != nil {
		http.Error(w, "This group has no members", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&people)

}
