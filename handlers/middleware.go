package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/matejelenc/rest-api/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var CLIENT, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Matej:MambaMentality11@user-group.amopm.mongodb.net/user-group?retryWrites=true&w=majority"))
var database = CLIENT.Database("Database")
var userDB = database.Collection("users")
var groupDB = database.Collection("groups")

type KeyUser struct{}

// MiddlewareValidatUser validates the user in the request and calls next if there are no errors
func MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

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

		hashedP, err := HashPassword(user.Password)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}

		user.Password = string(hashedP)

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

		//deserializes the group from the request
		var group data.Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}
		//validates the user
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

// MiddlewareValidatUser validates the user in the request and calls next if there are no errors
func MiddlewareValidateUserUpdate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		//deserializes the user from the request
		var upPerson data.Person
		err = json.NewDecoder(r.Body).Decode(&upPerson)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		if len(upPerson.Password) > 0 && len(upPerson.Password) < 4 {
			http.Error(rw, "Password too short", http.StatusBadRequest)
			return
		}

		hashedP, err := HashPassword(upPerson.Password)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}

		upPerson.Password = string(hashedP)

		if upPerson.Email != "" && data.ValidateEmail(upPerson.Email) != nil {
			http.Error(rw, "Invalid email", http.StatusBadRequest)
			return
		}

		//adds the user to the context and calls the next handler if there are no errors
		ctx := context.WithValue(r.Context(), KeyUpdateUser{}, upPerson)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
