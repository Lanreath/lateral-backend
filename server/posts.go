package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Lanreath/lateral-backend/models"
)

func (s *Server) CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ret, err := s.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(ret)
}

func (s *Server) CreateTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var task models.Task
	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ret, err := s.db.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(ret)
}

func (s *Server) CreateCalendar(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var calendar models.Calendar
	if err := json.NewDecoder(req.Body).Decode(&calendar); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ret, err := s.db.Collection("calendars").InsertOne(ctx, calendar)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(ret)
}
