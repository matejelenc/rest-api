package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/matejelenc/rest-api/data"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	u := data.GetUsers()

	err := u.UsersToJSON(rw)
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
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}

	err = u.UserToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(KeyUser{}).(data.User)

	err = data.UpdateUser(id, &user)
	if err != nil {
		http.Error(rw, "User not found", http.StatusBadRequest)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(KeyUser{}).(data.User)

	data.CreateUser(&user)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteUser(id)
	if err != nil {
		http.Error(rw, "User not found", http.StatusBadRequest)
	}
}

type KeyUser struct{}

func MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}

		err := user.UserFromJSON(r.Body)
		if err != nil {
			fmt.Println("[ERROR] deserializing user", err)
			http.Error(rw, "Error reading user", http.StatusBadRequest)
			return
		}

		err = user.ValidateUser()
		if err != nil {
			fmt.Println("[ERROR] validating user", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
