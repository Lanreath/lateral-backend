package models

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"id"`
	FirstName  string             `bson:"first_name" validate:"required"`
	LastName   string             `bson:"last_name" validate:"required"`
	Age        int                `bson:"age,omitempty" validate:"gt=0,lte=150"`
	Occupation string             `bson:"occupation,omit_empty"`
	Email      string             `validate:"required, email"`
	Password   string             `bson:"password" validate:"required"`
	CreatedOn  primitive.DateTime `bson:"="`
	Events     []*Event           `bson:"events,omitempty"`
	Calendars  []*Calendar        `bson:"calendars,omitempty"`
}

func (p *User) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
