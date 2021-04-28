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

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of users returns in the response
// swagger:response usersResponse
type usersResponse struct {
	// All current users in the system
	// in: body
	Body []data.User
}

// A single user returns in the response
// swagger:response userResponse
type userResponse struct {
	// A specific user
	// in: body
	Body data.User
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
	Body []data.Group
}

// A single group returns in the response
// swagger:response groupResponse
type groupResponse struct {
	// A specific group
	// in: body
	Body data.Group
}

// swagger:parameters updateGroup deleteGroup getGroup
type groupIDParameter struct {
	// The id of the group to access, update or delete from database
	// in: path
	// required: true
	ID int `json:id`
}

// swagger:response noContent
type usersNoContent struct {
}
