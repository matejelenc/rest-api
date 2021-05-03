package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

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

	id, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	if id != os.Getenv("ADMIN_ID") {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "User not authorized", http.StatusBadRequest)
		return
	}

	group := r.Context().Value(security.KeyGroup{}).(data.Group)

	createdGroup := data.DB.Create(&group)
	err = createdGroup.Error
	if err != nil {
		http.Error(w, "Could not create a group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&group)
}
