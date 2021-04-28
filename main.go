package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/handlers"
)

func main() {

	//router
	router := mux.NewRouter()

	//subrouter for groups
	gr := router.PathPrefix("/groups").Subrouter()
	gr.HandleFunc("", handlers.GetGroups).Methods(http.MethodGet)
	gr.HandleFunc("/{id:[0-9]+}", handlers.GetGroup).Methods(http.MethodGet)

	//subrouter for groups that handles POST requests
	grPost := gr.Methods(http.MethodPost).Subrouter()
	grPost.HandleFunc("", handlers.CreateGroup)
	grPost.Use(handlers.MiddlewareValidateGroup)

	//subrouter for groups that handles PUT requests
	grPut := gr.Methods(http.MethodPut).Subrouter()
	grPut.HandleFunc("/{id:[0-9]+}", handlers.UpdateGroup)
	grPut.Use(handlers.MiddlewareValidateGroup)

	//subrouter for groups that handles DELETE requests
	grDelete := gr.Methods(http.MethodDelete).Subrouter()
	grDelete.HandleFunc("/{id:[0-9]+}", handlers.DeleteGroup)

	//subrouter for users
	ur := router.PathPrefix("/users").Subrouter()
	ur.HandleFunc("", handlers.GetUsers).Methods(http.MethodGet)
	ur.HandleFunc("/{id:[0-9]+}", handlers.GetUser).Methods(http.MethodGet)

	//subrouter for users that handles POST requests
	urPost := ur.Methods(http.MethodPost).Subrouter()
	urPost.HandleFunc("", handlers.CreateUser)
	urPost.Use(handlers.MiddlewareValidateUser)

	//subrouter for users that handles PUT requests
	urPut := ur.Methods(http.MethodPut).Subrouter()
	urPut.HandleFunc("/{id:[0-9]+}", handlers.UpdateUser)
	urPut.Use(handlers.MiddlewareValidateUser)

	//subrouter for user that handles DELETE requests
	urDelete := ur.Methods(http.MethodDelete).Subrouter()
	urDelete.HandleFunc("/{id:[0-9]+}", handlers.DeleteUser)

	//subrouter that handles docs
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	dr := router.PathPrefix("/docs").Subrouter()
	dr.Handle("", sh)

	//subrouter that serves swagger.yaml file to the server
	yr := router.PathPrefix("/swagger.yaml").Subrouter()
	yr.Handle("", http.FileServer(http.Dir("./")))

	//creates a server with the above router
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//checks for interruptions and errors
	go func() {
		fmt.Println("Starting server...")

		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting the server: %s ", err)
			os.Exit(1)
		}
	}()

	//sends the signal when an interruption occurs
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//receives the signal and shuts down the server
	sig := <-c
	fmt.Println("Got signal: ", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}
