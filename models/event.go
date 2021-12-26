package models

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `bson:"id"`
	CreatedOn   primitive.DateTime `bson:"-"`
	UpdatedOn   primitive.DateTime `bson:"-"`
	Description string             `bson:"description" validate:"required"`
	DateTime    primitive.DateTime `bson:"datetime" validate:"required"`
	Status      string             `bson:"status" validate:"progress"`
	Frequency   string             `bson:"frequency" validate:"required,frequency"`
}

func (p *Event) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("progress", validateProgress)
	validate.RegisterValidation("frequency", validateFrequency)

	return validate.Struct(p)
}

func validateProgress(f1 validator.FieldLevel) bool {
	progress := [2]string{"Ongoing", "Done"}
	for _, a := range progress {
		if a == f1.Field().String() {
			return true
		}
	}

	return false
}

func validateFrequency(f1 validator.FieldLevel) bool {
	frequency := [...]string{"Daily", "Weekly", "Monthly", "Annually", "Once"}
	for _, a := range frequency {
		if a == f1.Field().String() {
			return true
		}
	}

	return false
}
