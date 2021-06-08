package security

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/data"
)

type KeyUser struct{}

// MiddlewareValidatUser validates the user in the request and calls next if there are no errors
func MiddlewareCreateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		//for now everyone can create a user for testing purposes, otherwise only a logged in user can create another user
		/*//validate the token
		_, err := security.ValidateToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}*/

		//deserializes the user from the request
		var user data.Person
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		//validates the user
		err = user.ValidateUser()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		//hash the password
		hashedP, err := HashPassword(user.Password)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}

		user.Password = string(hashedP)

		var group data.Group
		person := r.Context().Value(KeyUser{}).(data.Person)

		//check if the users group exists
		if person.GroupName != "" {
			foundGroup := data.DB.Where("name = ?", person.GroupName).First(&group)
			if foundGroup.Error != nil {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(foundGroup.Error.Error()))
				return
			}
		}

		//adds the user to the context and calls the next handler if there are no errors
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

type KeyGroup struct{}

// MiddlewareValidatGroup validates the group in the request and calls next if there are no errors
func MiddlewareValidateGroup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		//validate the jwt token
		id, err := ValidateToken(rw, r)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte(err.Error()))
			return
		}

		//check if the user is authorized for this request
		if id != os.Getenv("ADMIN_ID") {
			rw.WriteHeader(http.StatusUnauthorized)
			http.Error(rw, "User not authorized", http.StatusBadRequest)
			return
		}

		//deserializes the group from the request
		var group data.Group
		err = json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		//validates the group
		err = group.ValidateGroup()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		//adds the group to the context and calls the next handler if there are no errors
		ctx := context.WithValue(r.Context(), KeyGroup{}, group)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

type KeyUpdateUser struct{}

// MiddlewareUpdateUser validates the user in the request and calls next if there are no errors
func MiddlewareUpdateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//validate the jwt token
		id, err := ValidateToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		params := mux.Vars(r)
		//check if the user is authorized for this request
		if id != params["id"] {
			if id != os.Getenv("ADMIN_ID") {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "User not authorized", http.StatusBadRequest)
				return
			}
		}

		//deserializes the user from the request
		var upPerson data.Person
		err = json.NewDecoder(r.Body).Decode(&upPerson)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		//check the password
		if len(upPerson.Password) > 0 && len(upPerson.Password) < 4 {
			http.Error(w, "Password too short", http.StatusBadRequest)
			return
		}

		//hash the password
		hashedP, err := HashPassword(upPerson.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		upPerson.Password = string(hashedP)

		//validate the email
		if upPerson.Email != "" && data.ValidateEmail(upPerson.Email) != nil {
			http.Error(w, "Invalid email", http.StatusBadRequest)
			return
		}

		//adds the user to the context and calls the next handler if there are no errors
		ctx := context.WithValue(r.Context(), KeyUpdateUser{}, upPerson)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// MiddlewareDelete validates the authorization of the client and calls next if there are no errors
func MiddlewareDelete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		//validate the jwt token
		id, err := ValidateToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		//check if the user is authorized for this request
		if id != params["id"] {
			if id != os.Getenv("ADMIN_ID") {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "User not authorized", http.StatusBadRequest)
				return
			}
		}

		//if everything is ok call the next handler
		next.ServeHTTP(w, r)
	})
}

// MiddlewareGet validates the authorization of the user calls next if there are no errors
func MiddlewareGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//validate the jwt token
		_, err := ValidateToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		//if everything is ok call the next handler
		next.ServeHTTP(w, r)
	})
}
