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
//  401: unauthorizedResponse
//  400: badRequestResponse

//DeleteGroup removes a group from the database
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// validate the jwt token
	id, err := security.ValidateToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	//check if the user is authorized for this request
	if id != os.Getenv("ADMIN_ID") {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "User not authorized", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	var group data.Group

	//check if the group exists
	foundGroup := data.DB.First(&group, params["id"])
	if foundGroup.Error != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	//delete the group
	deletedGroup := data.DB.Delete(&group)
	if deletedGroup.Error != nil {
		http.Error(w, "Could not delete group", http.StatusInternalServerError)
		return
	}

	//update all the users in this group that there group has been deleted
	data.DB.Model(&data.Person{}).Where("group_name = ?", group.Name).Update("group_name", group.Name+"-deleted")

	json.NewEncoder(w).Encode(fmt.Sprintf("Group %v was successfully deleted", group.Name))
}
