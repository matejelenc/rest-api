package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/matejelenc/rest-api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", handlers.GetUsers)
	getRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUser)
	getRouter.HandleFunc("/groups", handlers.GetGroups)
	getRouter.HandleFunc("/groups/{id:[0-9]+}", handlers.GetGroup)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateUser)
	putRouter.HandleFunc("/groups/{id:[0-9]+}", handlers.UpdateGroup)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", handlers.CreateUser)
	postRouter.HandleFunc("/groups", handlers.CreateGroup)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteUser)
	deleteRouter.HandleFunc("/groups/{id:[0-9]+}", handlers.DeleteGroup)

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
