package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route GET /users users getUsers
// Returns a list of users
// responses:
//	200: usersResponse

//GEtUsers returns all users from the database
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	u := data.GetUsers()

	err := u.UsersToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}

// swagger:route GET /users/{id} users getUser
// Return a single user
// responses:
//	200: userResponse
//	404: errorResponse

//GetUser returns a specified user
func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	u, _, e := data.GetUser(id)
	if e != nil {
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}

	err = u.UserToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}
