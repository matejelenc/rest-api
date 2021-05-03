package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route PATCH /groups/{id} groups updateGroup
// Update a group
//
// responses:
//	201: groupResponse
//  401: unauthorizedResponse
//  400: badRequestResponse

//UpdateGroup updates a group with requests context
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	//validate the jwt token
	id, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	//check if the user is authorized for this request
	if id != os.Getenv("ADMIN_ID") {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "User not authorized", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)

	//get the group
	var group data.Group
	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	//deserialize the group
	var upGroup data.Group
	err = json.NewDecoder(r.Body).Decode(&upGroup)
	if err != nil {
		http.Error(w, "Could not deserialize requests body", http.StatusBadRequest)
		return
	}

	//update the group
	updatedGroup := data.DB.Model(&group).Updates(upGroup)
	if updatedGroup.Error != nil {
		http.Error(w, "Could not update group", http.StatusInternalServerError)
		return
	}

	//update all the users in the group
	data.DB.Model(&data.Person{}).Where("group_name = ?", upGroup.Name).Update("group_name", upGroup.Name)

	data.DB.First(&group, params["id"])

	json.NewEncoder(w).Encode(group)
}
