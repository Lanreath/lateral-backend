package models

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"id"`
	CreatedOn   primitive.DateTime `bson:"-"`
	UpdatedOn   primitive.DateTime `bson:"-"`
	Title       string             `bson:"title" validate:"required"`
	Description string             `bson:"description,omitempty"`
	DateTime    primitive.DateTime `bson:"datetime" validate:"required"`
	Status      string             `bson:"status" validate:"progress,required"`
	Frequency   string             `bson:"frequency" validate:"required,frequency"`
}

func (p *Task) Validate() error {
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
