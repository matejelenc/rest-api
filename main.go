package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/matejelenc/rest-api/data"
	"github.com/matejelenc/rest-api/handlers"
	"github.com/matejelenc/rest-api/security"
)

func main() {

	//declaring environment variables used for connecting to the database
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//connecting to the database
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	var err error
	conn, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}
	data.DB = conn

	//the database will close once the server stops running
	defer data.DB.Close()

	//data models for groups and users are created in the database
	data.DB.AutoMigrate(&data.Group{})
	data.DB.AutoMigrate(&data.Person{})

	//router
	router := mux.NewRouter()

	//handling user login
	router.HandleFunc("/login", handlers.Login)

	//subrouter for groups
	gr := router.PathPrefix("/groups").Subrouter()
	gr.HandleFunc("", handlers.GetGroups).Methods(http.MethodGet)
	gr.HandleFunc("/{id}", handlers.GetGroup).Methods(http.MethodGet)
	gr.HandleFunc("/{id}/members", handlers.GetMembers).Methods(http.MethodGet)

	//subrouter for groups that handles POST requests
	grPost := gr.Methods(http.MethodPost).Subrouter()
	grPost.HandleFunc("", handlers.CreateGroup)
	grPost.Use(security.MiddlewareValidateGroup)

	//subrouter for groups that handles PUT requests
	grPatch := gr.Methods(http.MethodPatch).Subrouter()
	grPatch.HandleFunc("/{id}", handlers.UpdateGroup)

	//subrouter for groups that handles DELETE requests
	grDelete := gr.Methods(http.MethodDelete).Subrouter()
	grDelete.HandleFunc("/{id}", handlers.DeleteGroup)

	//subrouter for users
	ur := router.PathPrefix("/users").Subrouter()
	ur.HandleFunc("", handlers.GetPeople).Methods(http.MethodGet)
	ur.HandleFunc("/{id}", handlers.GetPerson).Methods(http.MethodGet)

	//subrouter for users that handles POST requests
	urPost := ur.Methods(http.MethodPost).Subrouter()
	urPost.HandleFunc("", handlers.CreatePerson)
	urPost.Use(security.MiddlewareCreateUser)

	//subrouter for users that handles PUT requests
	urPatch := ur.Methods(http.MethodPatch).Subrouter()
	urPatch.HandleFunc("/{id}", handlers.UpdatePerson)
	urPatch.Use(security.MiddlewareUpdateUser)

	//subrouter for user that handles DELETE requests
	urDelete := ur.Methods(http.MethodDelete).Subrouter()
	urDelete.HandleFunc("/{id}", handlers.DeletePerson)

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
