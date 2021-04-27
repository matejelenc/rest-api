package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/matejelenc/rest-api/handlers"
)

func main() {

	router := mux.NewRouter()

	gr := router.PathPrefix("/groups").Subrouter()
	gr.HandleFunc("", handlers.GetGroups).Methods(http.MethodGet)
	gr.HandleFunc("/{id:[0-9]+}", handlers.GetGroup).Methods(http.MethodGet)

	grPost := gr.Methods(http.MethodPost).Subrouter()
	grPost.HandleFunc("", handlers.CreateGroup)
	grPost.Use(handlers.MiddlewareValidateGroup)

	grPut := gr.Methods(http.MethodPut).Subrouter()
	grPut.HandleFunc("/{id:[0-9]+}", handlers.UpdateGroup)
	grPut.Use(handlers.MiddlewareValidateGroup)

	grDelete := gr.Methods(http.MethodDelete).Subrouter()
	grDelete.HandleFunc("/{id:[0-9]+}", handlers.DeleteGroup)

	ur := router.PathPrefix("/users").Subrouter()
	ur.HandleFunc("", handlers.GetUsers).Methods(http.MethodGet)
	ur.HandleFunc("/{id:[0-9]+}", handlers.GetUser).Methods(http.MethodGet)

	urPost := ur.Methods(http.MethodPost).Subrouter()
	urPost.HandleFunc("", handlers.CreateUser)
	urPost.Use(handlers.MiddlewareValidateUser)

	urPut := ur.Methods(http.MethodPut).Subrouter()
	urPut.HandleFunc("/{id:[0-9]+}", handlers.UpdateUser)
	urPut.Use(handlers.MiddlewareValidateUser)

	urDelete := ur.Methods(http.MethodDelete).Subrouter()
	urDelete.HandleFunc("/{id:[0-9]+}", handlers.DeleteUser)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		fmt.Println("Starting server...")

		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting the server: %s ", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	fmt.Println("Got signal: ", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}
