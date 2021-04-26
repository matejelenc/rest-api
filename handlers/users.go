package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/matejelenc/rest-api/data"

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

type KeyUser struct{}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	u := data.User{}
	err = decoder.Decode(&u)
	if err != nil {
		fmt.Println("[ERROR] deserializing user", err)
		http.Error(rw, "Error reading user", http.StatusBadRequest)
		return
	}

	err = data.UpdateUser(id, &u)
	if err != nil {
		http.Error(rw, "User not found", http.StatusBadRequest)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	u := data.User{}
	err := decoder.Decode(&u)
	if err != nil {
		fmt.Println("[ERROR] deserializing user", err)
		http.Error(rw, "Error reading user", http.StatusBadRequest)
		return
	}

	data.CreateUser(&u)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

}
