package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route PUT/users/{id} users updateUser
// Update a user
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

//UpdateUser updates a user with requests context
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var person data.Person
	var group data.Group

	foundPerson := DB.First(&person, params["id"])
	if foundPerson.Error != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	upPerson := r.Context().Value(KeyUpdateUser{}).(data.Person)

	if upPerson.GroupName != "" {
		foundGroup := DB.First(&group, "Name = ?", upPerson.GroupName)
		if foundGroup.Error != nil {
			http.Error(w, "Group not found", http.StatusBadRequest)
			return
		}
	}

	updatedPerson := DB.Model(&person).Updates(upPerson)
	if updatedPerson.Error != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	DB.First(&person, params["id"])

	json.NewEncoder(w).Encode(person)
}
