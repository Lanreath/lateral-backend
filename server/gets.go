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

func (s *Server) GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var users []models.User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := s.Database.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cur.All(ctx, &users); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(users)
}

func (s *Server) GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var user models.User
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := s.Database.Collection("users").FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(err)
}

func (s *Server) GetTasks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var tasks []models.Task
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := s.Database.Collection("tasks").Find(ctx, bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cur.All(ctx, &tasks); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(tasks)
}

func (s *Server) GetTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := s.Database.Collection("tasks").FindOne(ctx, bson.M{"id": id}).Decode(bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(err)
}

func (s *Server) GetCalendars(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var calendars []models.Calendar
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := s.Database.Collection("calendars").Find(ctx, bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cur.All(ctx, &calendars); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(calendars)
}

func (s *Server) GetCalendar(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := s.Database.Collection("calendars").FindOne(ctx, bson.M{"id": id}).Decode(bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(err)
}
