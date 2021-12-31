package server

import (
	"github.com/Lanreath/lateral-backend/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type data struct {
	models.User
	models.Calendar
	models.Event
}

type Collections struct {
	users     *mongo.Collection
	tasks     *mongo.Collection
	calendars *mongo.Collection
}

type Server struct {
	server *mongo.Client
	router *mux.Router
	db     *mongo.Database
}
