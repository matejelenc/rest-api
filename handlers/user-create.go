package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route POST /users users createUser
// Create a new user
//
// responses:
//	200: userResponse
//  422: errorValidation
//  501: errorResponse

//CreateUser creates a user with the requests context
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	_, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	var group data.Group
	person := r.Context().Value(security.KeyUser{}).(data.Person)

	if person.GroupName != "" {
		foundGroup := data.DB.Where("name = ?", person.GroupName).First(&group)
		if foundGroup.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(foundGroup.Error.Error()))
			return
		}
	}

	createdPerson := data.DB.Create(&person)
	err = createdPerson.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&person)
}
