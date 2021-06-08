package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
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
