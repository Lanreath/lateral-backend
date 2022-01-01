package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Lanreath/lateral-backend/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var user models.User
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	json.NewDecoder(req.Body).Decode(&user)
	ret, err := s.db.Collection("users").UpdateOne(ctx, bson.M{"id": id}, bson.D{{"$set", user}})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(ret)
}

func (s *Server) UpdateTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var task models.Task
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	json.NewDecoder(req.Body).Decode(&task)
	ret, err := s.db.Collection("tasks").UpdateOne(ctx, bson.M{"id": id}, bson.D{{"$set", task}})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(ret)
}

func (s *Server) UpdateCalendar(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var calendar models.Calendar
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	json.NewDecoder(req.Body).Decode(&calendar)
	ret, err := s.db.Collection("calendars").UpdateOne(ctx, bson.M{"id": id}, bson.D{{"$set", calendar}})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(ret)
}
