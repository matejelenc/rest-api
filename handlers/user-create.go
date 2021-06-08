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
//	201: userResponse
//  401: unauthorizedResponse
//  400: badRequestResponse

//CreateUser creates a user with the requests context
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	person := r.Context().Value(security.KeyUser{}).(data.Person)

	//create the user
	createdPerson := data.DB.Create(&person)
	err := createdPerson.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&person)
}
