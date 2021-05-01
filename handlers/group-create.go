package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/matejelenc/rest-api/data"
)

var DB *gorm.DB

// swagger:route POST /groups groups createGroup
// Create a new group
//
// responses:
//	200: groupResponse
//  422: errorValidation
//  501: errorResponse

//CreateGroup creates a group with requests context
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	group := r.Context().Value(KeyGroup{}).(data.Group)

	createdGroup := DB.Create(&group)
	err = createdGroup.Error
	if err != nil {
		http.Error(w, "Could not create a group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&group)
}
