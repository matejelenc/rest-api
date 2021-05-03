package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route DELETE /users/{id} users deleteUser
// Deletes a user
// responses:
//	201: noContent
//  401: unauthorizedResponse
//  400: badRequestResponse

//DeleteUser handles DELETE requests and removes a user from the database
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	var person data.Person

	//validate the jwt token
	id, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	//check if the user is authorized for this request
	if id != params["id"] {
		if id != os.Getenv("ADMIN_ID") {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "User not authorized", http.StatusBadRequest)
			return
		}
	}

	//get the user
	foundPerson := data.DB.First(&person, params["id"])
	if err := foundPerson.Error; err != nil {
		json.NewEncoder(w).Encode(err)
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	//delete the user
	deletedPerson := data.DB.Delete(&person)
	if err := deletedPerson.Error; err != nil {
		json.NewEncoder(w).Encode(err)
		http.Error(w, "User could not be deleted", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("User %v was successfully deleted", person.Name))
}
