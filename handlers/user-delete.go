package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route DELETE /users/{id} users deleteUser
// Deletes a user
// responses:
//	201: noContent

//DeleteUser handles DELETE requests and removes a user from the database
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteUser(id)
	if err != nil {
		http.Error(rw, "User not found", http.StatusBadRequest)
		return
	}

}
