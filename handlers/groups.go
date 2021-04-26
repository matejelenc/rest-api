package handlers

import (
	"encoding/json"
	"fmt"
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

func UpdateGroup(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	g := data.Group{}
	err = decoder.Decode(&g)
	if err != nil {
		fmt.Println("[ERROR] deserializing group", err)
		http.Error(rw, "Error reading group", http.StatusBadRequest)
		return
	}

	err = data.UpdateGroup(id, &g)
	if err != nil {
		http.Error(rw, "Group not found", http.StatusBadRequest)
	}
}

func CreateGroup(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	g := data.Group{}
	err := decoder.Decode(&g)
	if err != nil {
		fmt.Println("[ERROR] deserializing user", err)
		http.Error(rw, "Error reading user", http.StatusBadRequest)
		return
	}

	data.CreateGroup(&g)
}

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
