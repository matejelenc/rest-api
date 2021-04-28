package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route PUT /groups/{id} groups updateGroup
// Update a group
//
// responses:
//	201: groupResponse
//  404: errorResponse
//  422: errorValidation

//UpdateGroup updates a group with requests context
func UpdateGroup(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	g := r.Context().Value(KeyGroup{}).(data.Group)

	err = data.UpdateGroup(id, &g)
	if err != nil {
		http.Error(rw, "Group not found", http.StatusBadRequest)
	}
}
