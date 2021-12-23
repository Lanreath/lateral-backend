package models

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Calendar struct {
	ID        primitive.ObjectID `bson:"id" validate:"required"`
	CreatedBy User               `bson:"createdby" validate:"required"`
	CreatedOn primitive.DateTime `bson:"-"`
	UpdatedOn primitive.DateTime `bson:"-"`
	Members   []*User            `bson:"members"`
}

func (p *Calendar) Validate() error {
	validate := validator.New()

	return validate.Struct(p)
}
