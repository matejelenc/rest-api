package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"rest-api/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", handlers.GetUsers)
	getRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUser)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}/update", handlers.UpdateUser)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/new", handlers.CreateUser)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}/delete", handlers.GetUsers)

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
