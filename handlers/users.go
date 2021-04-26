package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	u := data.GetUsers()

	encoder := json.NewEncoder(rw)
	err := encoder.Encode(u)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	u, _, e := data.GetUser(id)
	if e != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(rw)
	err = encoder.Encode(u)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

}
