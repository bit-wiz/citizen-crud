package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Citizen struct {
	ID          string              `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string              `json:"first_name" bson:"first_name"`
	LastName    string              `json:"last_name" bson:"last_name"`
	DateOfBirth time.Time           `json:"date_of_birth" bson:"date_of_birth"`
	Gender      string              `json:"gender" bson:"gender"`
	Address     string              `json:"address" bson:"address"`
	City        string              `json:"city" bson:"city"`
	State       string              `json:"state" bson:"state"`
	Pincode     string              `json:"pincode" bson:"pincode"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.Timestamp `json:"updated_at" bson:"updated_at"`
}
