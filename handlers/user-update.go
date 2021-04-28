package handlers

import (
	"net/http"
	"strconv"

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
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(KeyUser{}).(data.User)

	err = data.UpdateUser(id, &user)
	if err != nil {
		http.Error(rw, "User not found", http.StatusBadRequest)
	}

}
