package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route PUT /groups/{id} groups updateGroup
// Update a group
//
// responses:
//	201: groupResponse
//  404: errorResponse
//  422: errorValidation

//UpdateGroup updates a group with requests context
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	var group data.Group
	foundGroup := DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	var upGroup data.Group
	err := json.NewDecoder(r.Body).Decode(&upGroup)
	if err != nil {
		http.Error(w, "Could not deserialize requests body", http.StatusBadRequest)
		return
	}

	updatedGroup := DB.Model(&group).Updates(upGroup)
	if updatedGroup.Error != nil {
		http.Error(w, "Could not update group", http.StatusInternalServerError)
		return
	}

	DB.Model(&data.Person{}).Where("group_name = ?", upGroup.Name).Update("group_name", upGroup.Name)

	DB.First(&group, params["id"])

	json.NewEncoder(w).Encode(group)
}
