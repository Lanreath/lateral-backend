package server

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	db     *mongo.Database
	router *mux.Router
}
