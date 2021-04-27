package handlers

import (
	"context"
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

	g := r.Context().Value(KeyGroup{}).(data.Group)

	err = data.UpdateGroup(id, &g)
	if err != nil {
		http.Error(rw, "Group not found", http.StatusBadRequest)
	}
}

func CreateGroup(rw http.ResponseWriter, r *http.Request) {
	g := r.Context().Value(KeyGroup{}).(data.Group)

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

type KeyGroup struct{}

func MiddlewareValidateGroup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		group := data.Group{}

		err := group.GroupFromJSON(r.Body)
		if err != nil {
			fmt.Println("[ERROR] deserializing group", err.Error())
			http.Error(rw, "Error reading group", http.StatusBadRequest)
			return
		}

		err = group.ValidateGroup()
		if err != nil {
			fmt.Println("[ERROR] validating group", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating group: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyGroup{}, group)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
