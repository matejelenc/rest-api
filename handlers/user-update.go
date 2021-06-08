package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route PATCH/users/{id} users updateUser
// Update a user
//
// responses:
//	200: userResponse
//  401: unauthorizedResponse
//  400: badRequestResponse

//UpdateUser updates a user with requests context
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var person data.Person
	var group data.Group

	//get the user
	foundPerson := data.DB.First(&person, params["id"])
	if foundPerson.Error != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	upPerson := r.Context().Value(security.KeyUpdateUser{}).(data.Person)

	//check if the users group exists
	if upPerson.GroupName != "" {
		foundGroup := data.DB.First(&group, "Name = ?", upPerson.GroupName)
		if foundGroup.Error != nil {
			http.Error(w, "Group not found", http.StatusBadRequest)
			return
		}
	}

	//update the user
	updatedPerson := data.DB.Model(&person).Updates(upPerson)
	if updatedPerson.Error != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	data.DB.First(&person, params["id"])

	json.NewEncoder(w).Encode(person)
}
