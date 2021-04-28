package handlers

import (
	"net/http"

	"github.com/matejelenc/rest-api/data"
)

// swagger:route POST /groups groups createGroup
// Create a new group
//
// responses:
//	200: groupResponse
//  422: errorValidation
//  501: errorResponse

//CreateGroup creates a group with requests context
func CreateGroup(rw http.ResponseWriter, r *http.Request) {
	g := r.Context().Value(KeyGroup{}).(data.Group)

	data.CreateGroup(&g)
}
