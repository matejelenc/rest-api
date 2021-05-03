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

//GetGroups returns all the groups in the database
func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	_, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	var groups []data.Group
	foundGroups := data.DB.Find(&groups)
	if foundGroups.Error != nil {
		http.Error(w, "No groups exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&groups)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	_, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	params := mux.Vars(r)

	var group data.Group

	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var group data.Group
	var people []data.Person

	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	foundMembers := data.DB.Where("group_name = ?", group.Name).Find(&people)
	if foundMembers.Error != nil {
		http.Error(w, "This group has no members", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&people)

}
