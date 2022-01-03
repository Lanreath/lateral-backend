package models

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserValidation(t *testing.T) {
	u := &User{
		LastName: "James",
		Age:      151,
		Email:    "notanemail",
	}

	err := u.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestTaskValidation(t *testing.T) {
	s := &Task{
		Title:       "Sleep",
		Description: "This is the most important of them all",
		DateTime:    primitive.DateTime(time.Now().Day()),
		Status:      "Ongoing",
		Frequency:   "Weekly",
	}

	err := s.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestCalendarValidation(t *testing.T) {
	c := &Calendar{}

	err := c.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
