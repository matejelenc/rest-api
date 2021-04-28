package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

// swagger:route DELETE /groups/{id} groups deleteGroup
// Deletes a group
// responses:
//	201: noContent

//DeleteGroup removes a group from the database
func DeleteGroup(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteGroup(id)
	if err != nil {
		http.Error(rw, "Group not found", http.StatusBadRequest)
	}
}
