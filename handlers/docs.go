//Package classification of User and Group Management API
//
//Documentation for User and Group Management API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package handlers

import (
	"github.com/matejelenc/rest-api/data"
)

// A list of users returns in the response
// swagger:response usersResponse
type usersResponse struct {
	// All current users in the system
	// in: body
	Body data.Users
}

// A single user returns in the response
// swagger:response userResponse
type userResponse struct {
	// A specific user
	// in: body
	Body data.Person
}

// swagger:parameters updateUser deleteUser getUser
type userIDParameter struct {
	// The id of the user to access, update or delete from database
	// in: path
	// required: true
	ID int `json:id`
}

// A list of groups returns in the response
// swagger:response groupsResponse
type groupsResponse struct {
	// All current groups in the system
	// in: body
	Body data.Groups
}

// A single group returns in the response
// swagger:response groupResponse
type groupResponse struct {
	// A specific group
	// in: body
	Body data.Group
}

// Members of a group are returned in the response
// swagger:response membersResponse
type membersResponse struct {
	// Members of a specified group
	// in: body
	Body data.Users
}

type groupIDParameter struct {
	// The id of the group to access, update or delete from database
	// in: path
	// required: true
	ID int `json:id`
}

// swagger:response noContent
type usersNoContent struct {
}

// An unauthorized error is returned
// swagger:response unauthorizedResponse
type unauthorizedResponse struct {
	Message string `json:"message"`
}

// A bad request error is returned
// swagger:response badRequestResponse
type badRequestResponse struct {
	Message string `json:"message"`
}
