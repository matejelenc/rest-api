package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/security"
)

// swagger:route DELETE /groups/{id} groups deleteGroup
// Deletes a group
// responses:
//	201: noContent

//DeleteGroup removes a group from the database
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	id, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(id)
	if id != os.Getenv("ADMIN_ID") {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "User not authorized", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	var group data.Group

	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	deletedGroup := data.DB.Delete(&group)
	if deletedGroup.Error != nil {
		http.Error(w, "Could not delete group", http.StatusInternalServerError)
		return
	}

	data.DB.Model(&data.Person{}).Where("group_name = ?", group.Name).Update("group_name", group.Name+"-deleted")

	json.NewEncoder(w).Encode(&group)
}
