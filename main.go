package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Lanreath/lateral-backend/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	connection := &server.Server{
		Database: client.Database("MyFirstDatabase"),
		Router:   mux.NewRouter(),
	}

	connection.Router.HandleFunc("/users", connection.GetUser).Methods("GET")
	connection.Router.HandleFunc("/user/{id:[0-9]+}", connection.GetUser).Methods("GET")
	connection.Router.HandleFunc("/tasks", connection.GetTask).Methods("GET")
	connection.Router.HandleFunc("/task/{id:[0-9]+}", connection.GetTask).Methods("GET")
	connection.Router.HandleFunc("/calendars", connection.GetCalendar).Methods("GET")
	connection.Router.HandleFunc("/calendar/{id:[0-9]+}", connection.GetCalendars).Methods("GET")

	connection.Router.HandleFunc("/user", connection.GetUser).Methods("POST")
	connection.Router.HandleFunc("/task", connection.GetTask).Methods("POST")
	connection.Router.HandleFunc("/calendar", connection.GetCalendar).Methods("POST")

	connection.Router.HandleFunc("/user/{id:[0-9]}", connection.UpdateUser).Methods("PUT")
	connection.Router.HandleFunc("/task/{id:[0-9]}", connection.UpdateTask).Methods("PUT")
	connection.Router.HandleFunc("/calendar/{id:[0-9]}", connection.UpdateCalendar).Methods("PUT")

	connection.Router.HandleFunc("/user/{id:[0-9]}", connection.DeleteUser).Methods("DELETE")
	connection.Router.HandleFunc("/task/{id:[0-9]}", connection.DeleteTask).Methods("DELETE")
	connection.Router.HandleFunc("/calendar/{id:[0-9]}", connection.DeleteCalendar).Methods("DELETE")

	http.ListenAndServe(":8080", connection.Router)

}
