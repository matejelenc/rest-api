package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

func GetGroups(rw http.ResponseWriter, r *http.Request) {
	g := data.GetGroups()

	encoder := json.NewEncoder(rw)
	err := encoder.Encode(g)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}

func GetGroup(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	g, err := data.GetGroup(id)
	if err != nil {
		http.Error(rw, "Group not found", http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(rw)
	err = encoder.Encode(g)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
