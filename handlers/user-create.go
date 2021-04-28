package handlers

import (
	"net/http"

	"github.com/matejelenc/rest-api/data"
)

// swagger:route POST /users users createUser
// Create a new user
//
// responses:
//	200: userResponse
//  422: errorValidation
//  501: errorResponse

//CreateUser creates a user with the requests context
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(KeyUser{}).(data.User)

	data.CreateUser(&user)
}
