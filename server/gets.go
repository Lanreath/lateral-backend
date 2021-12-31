package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var all []data
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := s.db.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cur.All(ctx, &all); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(all)
}

func (s *Server) GetUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := s.db.Collection("users").FindOne(ctx, bson.M{"id": id}).Decode(bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(err)
}

func (s *Server) GetTasks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var all []data
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := s.db.Collection("tasks").Find(ctx, bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cur.All(ctx, &all); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(all)
}

func (s *Server) GetTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := s.db.Collection("tasks").FindOne(ctx, bson.M{"id": id}).Decode(bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(err)
}

func (s *Server) GetCalendars(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var all []data
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := s.db.Collection("calendars").Find(ctx, bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cur.All(ctx, &all); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(res).Encode(all)
}

func (s *Server) GetCalendar(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := s.db.Collection("calendars").FindOne(ctx, bson.M{"id": id}).Decode(bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(err)
}
