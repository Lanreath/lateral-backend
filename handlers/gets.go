package handlers

import (
	"context"
	"net/http"

	"github.com/Lanreath/lateral-backend/models"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var users []models.User
	ctx, _ := context.WithTimeout()
}
