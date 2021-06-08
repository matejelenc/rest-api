package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route POST /groups groups createGroup
// Create a new group
//
// responses:
//	201: groupResponse
//  401: unauthorizedResponse
//  400: badRequestResponse

//CreateGroup creates a group with requests context
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	group := r.Context().Value(security.KeyGroup{}).(data.Group)

	//create the group
	createdGroup := data.DB.Create(&group)
	err := createdGroup.Error
	if err != nil {
		http.Error(w, "Could not create a group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&group)
}
