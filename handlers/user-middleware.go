package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/matejelenc/rest-api/data"
)

type KeyUser struct{}

// MiddlewareValidatUser validates the user in the request and calls next if there are no errors
func MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}

		//deserializes the user from the request
		err := user.UserFromJSON(r.Body)
		if err != nil {
			fmt.Println("[ERROR] deserializing user", err)
			http.Error(rw, "Error reading user", http.StatusBadRequest)
			return
		}

		//validates the user
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

		//adds the user to the context and calls the next handler if there are no errors
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
