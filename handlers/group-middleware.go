package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/matejelenc/rest-api/data"
)

type KeyGroup struct{}

// MiddlewareValidatGroup validates the group in the request and calls next if there are no errors
func MiddlewareValidateGroup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		group := data.Group{}

		//deserializes the group from the request
		err := group.GroupFromJSON(r.Body)
		if err != nil {
			fmt.Println("[ERROR] deserializing group", err)
			http.Error(rw, "Error reading group", http.StatusBadRequest)
			return
		}

		//validates the group
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

		//adds the group to the context and calls the next handler if there are no errors
		ctx := context.WithValue(r.Context(), KeyGroup{}, group)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
