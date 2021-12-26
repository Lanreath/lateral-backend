package handlers

import (
	"github.com/Lanreath/lateral-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type data struct {
	models.User
	models.Calendar
	models.Event
}

type Collections struct {
	users     *mongo.Collection
	calendars *mongo.Collection
	events    *mongo.Collection
}
